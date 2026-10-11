// Package authz provides the authorization abstraction used to replace the
// former flat-role permission mapping tables with OpenFGA relationship
// tuples.
package authz

//go:generate mockgen -source authz.go -destination mock.go -package authz

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/config"
)

// TupleKey defines a relationship tuple between a subject and an object.
type TupleKey struct {
	// Object is the target of the relation, e.g. "pack:01ABC".
	Object string

	// Relation is the relation name, e.g. "owner", "admin", "user" or
	// "parent".
	Relation string

	// Subject is the principal holding the relation, e.g. "user:01ABC" or
	// "group:01ABC#user".
	Subject string

	// CreatedAt is populated by Read and reflects when the tuple was
	// written. It is ignored by Write and Delete.
	CreatedAt time.Time
}

// Checker provides the authorization operations used across the API and
// store layers. It is implemented both by an embedded, in-process OpenFGA
// server (for single-binary deployments) and by a client talking to an
// external OpenFGA service (for larger deployments).
type Checker interface {
	// Check reports whether the subject has the given relation on the object.
	Check(ctx context.Context, subject, relation, object string) (bool, error)

	// Write creates the given tuples. It is idempotent: writing a tuple that
	// already exists is not an error.
	Write(ctx context.Context, tuples ...TupleKey) error

	// Delete removes the given tuples. It is idempotent: deleting a tuple
	// that doesn't exist is not an error.
	Delete(ctx context.Context, tuples ...TupleKey) error

	// Read returns the direct tuples assigned to the given object, optionally
	// filtered to the given relations.
	Read(ctx context.Context, object string, relations ...string) ([]TupleKey, error)

	// ListObjects returns the objects of the given type on which the subject
	// holds the given relation, directly or through group membership.
	ListObjects(ctx context.Context, subject, relation, objectType string) ([]string, error)

	// ReadByUser returns the direct tuples held by the given subject across
	// all objects of the given type, optionally filtered to the given
	// relations. It is the reverse of Read (which is keyed by object).
	ReadByUser(ctx context.Context, subject, objectType string, relations ...string) ([]TupleKey, error)

	// ReadAll returns every tuple in the store, unfiltered by object, subject
	// or relation. It exists to support one-off migration between backends
	// (see the `authz migrate` command) and isn't used on any request path.
	ReadAll(ctx context.Context) ([]TupleKey, error)

	// Ready reports whether the checker is ready to serve requests.
	Ready(ctx context.Context) error

	// Close releases any resources held by the checker.
	Close() error
}

// New creates a Checker based on the given configuration. Driver selects the
// embedded database engine (sqlite3, mysql, mariadb, postgres, postgresql,
// the same names config.Database accepts) or "external" to talk to a
// standalone OpenFGA service instead.
func New(cfg config.Authz) (Checker, error) {
	if cfg.Driver == "external" {
		return newExternal(cfg)
	}

	return newEmbedded(cfg)
}
