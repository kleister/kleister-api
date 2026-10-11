package authz_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/kleister/kleister-api/pkg/authz"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testAuthzConfig returns a fresh on-disk sqlite config per test. A true
// ":memory:" DSN can't be reused here since the migrate step and the
// datastore each open their own connection, and SQLite's in-memory
// databases are only shared across connections in shared-cache mode, which
// is fragile to rely on once connections are pooled/closed in between.
//
// journal_mode is forced to DELETE (rather than the WAL default) since WAL
// leaves "-wal"/"-shm" sidecar files whose final checkpoint can race with
// t.TempDir()'s cleanup, intermittently failing it with "directory not
// empty".
func testAuthzConfig(t *testing.T) config.Authz {
	t.Helper()

	return config.Authz{
		Driver:  "sqlite3",
		Name:    filepath.Join(t.TempDir(), "authz.sqlite3"),
		Options: map[string]string{"_pragma": "journal_mode(DELETE)"},
	}
}

func newTestChecker(t *testing.T) authz.Checker {
	t.Helper()

	checker, err := authz.New(testAuthzConfig(t))

	require.NoError(t, err)

	t.Cleanup(func() {
		_ = checker.Close()
	})

	return checker
}

func TestEmbeddedBootstrapIdempotent(t *testing.T) {
	cfg := testAuthzConfig(t)

	first, err := authz.New(cfg)
	require.NoError(t, err)
	defer func() { _ = first.Close() }()

	ctx := context.Background()

	require.NoError(t, first.Write(ctx, authz.TupleKey{
		Object:   "pack:test",
		Relation: "admin",
		Subject:  "user:alice",
	}))

	allowed, err := first.Check(ctx, "user:alice", "admin", "pack:test")
	require.NoError(t, err)
	assert.True(t, allowed)

	// Reopening against the same DSN must reuse the existing store/model
	// rather than erroring or creating a duplicate, and must still see the
	// tuple written through the first instance.
	second, err := authz.New(cfg)
	require.NoError(t, err)
	defer func() { _ = second.Close() }()

	allowed, err = second.Check(ctx, "user:alice", "admin", "pack:test")
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestHierarchyOwnerAdminUser(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx, authz.TupleKey{
		Object:   "pack:p1",
		Relation: "owner",
		Subject:  "user:owner",
	}))

	for _, relation := range []string{"owner", "admin", "user"} {
		allowed, err := checker.Check(ctx, "user:owner", relation, "pack:p1")
		require.NoError(t, err)
		assert.Truef(t, allowed, "owner should satisfy %s", relation)
	}

	require.NoError(t, checker.Write(ctx, authz.TupleKey{
		Object:   "pack:p1",
		Relation: "admin",
		Subject:  "user:admin",
	}))

	allowed, err := checker.Check(ctx, "user:admin", "owner", "pack:p1")
	require.NoError(t, err)
	assert.False(t, allowed, "admin should not satisfy owner")

	allowed, err = checker.Check(ctx, "user:admin", "user", "pack:p1")
	require.NoError(t, err)
	assert.True(t, allowed, "admin should satisfy user")
}

func TestGroupMembershipGrantsPackAccess(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx,
		authz.TupleKey{Object: "group:g1", Relation: "user", Subject: "user:bob"},
		authz.TupleKey{Object: "pack:p1", Relation: "user", Subject: "group:g1#user"},
	))

	allowed, err := checker.Check(ctx, "user:bob", "user", "pack:p1")
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestBuildInheritsFromParentPack(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx,
		authz.TupleKey{Object: "pack:p1", Relation: "admin", Subject: "user:carol"},
		authz.TupleKey{Object: "build:b1", Relation: "parent", Subject: "pack:p1"},
	))

	allowed, err := checker.Check(ctx, "user:carol", "admin", "build:b1")
	require.NoError(t, err)
	assert.True(t, allowed)

	allowed, err = checker.Check(ctx, "user:carol", "user", "build:b1")
	require.NoError(t, err)
	assert.True(t, allowed)

	allowed, err = checker.Check(ctx, "user:dave", "user", "build:b1")
	require.NoError(t, err)
	assert.False(t, allowed)
}

func TestReadReturnsDirectTuplesOnly(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx,
		authz.TupleKey{Object: "pack:p1", Relation: "admin", Subject: "user:alice"},
		authz.TupleKey{Object: "pack:p1", Relation: "user", Subject: "user:bob"},
	))

	tuples, err := checker.Read(ctx, "pack:p1", "owner", "admin", "user")
	require.NoError(t, err)
	assert.Len(t, tuples, 2)

	require.NoError(t, checker.Delete(ctx, authz.TupleKey{
		Object:   "pack:p1",
		Relation: "admin",
		Subject:  "user:alice",
	}))

	tuples, err = checker.Read(ctx, "pack:p1", "owner", "admin", "user")
	require.NoError(t, err)
	assert.Len(t, tuples, 1)
}

func TestReadByUserReturnsObjectsAcrossType(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx,
		authz.TupleKey{Object: "pack:p1", Relation: "admin", Subject: "user:alice"},
		authz.TupleKey{Object: "pack:p2", Relation: "user", Subject: "user:alice"},
		authz.TupleKey{Object: "pack:p3", Relation: "user", Subject: "user:bob"},
	))

	tuples, err := checker.ReadByUser(ctx, "user:alice", "pack", "owner", "admin", "user")
	require.NoError(t, err)
	assert.Len(t, tuples, 2)
}

func TestReadAllReturnsEveryTuple(t *testing.T) {
	checker := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, checker.Write(ctx,
		authz.TupleKey{Object: "pack:p1", Relation: "admin", Subject: "user:alice"},
		authz.TupleKey{Object: "mod:m1", Relation: "user", Subject: "user:bob"},
		authz.TupleKey{Object: "group:g1", Relation: "user", Subject: "user:carol"},
		authz.TupleKey{Object: "build:b1", Relation: "parent", Subject: "pack:p1"},
	))

	tuples, err := checker.ReadAll(ctx)
	require.NoError(t, err)
	assert.Len(t, tuples, 4)
}

func TestReadAllRoundTripsIntoAnotherBackend(t *testing.T) {
	source := newTestChecker(t)
	destination := newTestChecker(t)
	ctx := context.Background()

	require.NoError(t, source.Write(ctx,
		authz.TupleKey{Object: "pack:p1", Relation: "admin", Subject: "user:alice"},
		authz.TupleKey{Object: "group:g1", Relation: "user", Subject: "user:bob"},
	))

	tuples, err := source.ReadAll(ctx)
	require.NoError(t, err)
	require.NoError(t, destination.Write(ctx, tuples...))

	allowed, err := destination.Check(ctx, "user:alice", "admin", "pack:p1")
	require.NoError(t, err)
	assert.True(t, allowed)

	copied, err := destination.ReadAll(ctx)
	require.NoError(t, err)
	assert.Len(t, copied, 2)
}
