package store

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/kleister/kleister-api/pkg/authz"
	"github.com/kleister/kleister-api/pkg/model"
)

// permTiers lists the relation names that make up an object's permission
// hierarchy, highest tier first.
var permTiers = []string{model.OwnerPerm, model.AdminPerm, model.UserPerm}

func tierRank(perm string) int {
	switch perm {
	case model.OwnerPerm:
		return 3
	case model.AdminPerm:
		return 2
	case model.UserPerm:
		return 1
	default:
		return 0
	}
}

func authzObject(kind, id string) string {
	return fmt.Sprintf("%s:%s", kind, id)
}

func authzUserSubject(id string) string {
	return fmt.Sprintf("user:%s", id)
}

func authzGroupSubject(id string) string {
	return fmt.Sprintf("group:%s#user", id)
}

func userIDFromSubject(subject string) string {
	return strings.TrimPrefix(subject, "user:")
}

func groupIDFromSubject(subject string) string {
	return strings.TrimSuffix(strings.TrimPrefix(subject, "group:"), "#user")
}

// searchMatches reports whether any of the given fields contains the search
// term as a case-insensitive substring.
func searchMatches(search string, fields ...string) bool {
	term := strings.ToLower(search)

	for _, field := range fields {
		if strings.Contains(strings.ToLower(field), term) {
			return true
		}
	}

	return false
}

// authzAssignment is the direct relationship tuple a subject holds on an
// object, mirroring one row of the former mapping tables.
type authzAssignment struct {
	Subject   string
	Perm      string
	CreatedAt time.Time
}

// authzDirectAssignments returns, per subject, the direct assignment held on
// the given object across the permission tiers.
func authzDirectAssignments(ctx context.Context, checker authz.Checker, object string) (map[string]authzAssignment, error) {
	tuples, err := checker.Read(ctx, object, permTiers...)

	if err != nil {
		return nil, err
	}

	out := make(map[string]authzAssignment, len(tuples))

	for _, t := range tuples {
		if existing, ok := out[t.Subject]; ok && tierRank(existing.Perm) >= tierRank(t.Relation) {
			continue
		}

		out[t.Subject] = authzAssignment{
			Subject:   t.Subject,
			Perm:      t.Relation,
			CreatedAt: t.CreatedAt,
		}
	}

	return out, nil
}

// authzGroupAssignments returns, per object ID of the given type, the direct
// assignment held by the group (as "group:<id>#user") across that type —
// the reverse direction of authzDirectAssignments, used to list the mods or
// packs a group has been granted access to.
func authzGroupAssignments(ctx context.Context, checker authz.Checker, objectType, groupID string) (map[string]authzAssignment, error) {
	return authzSubjectAssignments(ctx, checker, objectType, authzGroupSubject(groupID))
}

// authzUserAssignments returns, per object ID of the given type, the direct
// assignment held by the user — the reverse direction of
// authzDirectAssignments, used to list the groups, mods or packs a user has
// been granted access to.
func authzUserAssignments(ctx context.Context, checker authz.Checker, objectType, userID string) (map[string]authzAssignment, error) {
	return authzSubjectAssignments(ctx, checker, objectType, authzUserSubject(userID))
}

func authzSubjectAssignments(ctx context.Context, checker authz.Checker, objectType, subject string) (map[string]authzAssignment, error) {
	tuples, err := checker.ReadByUser(ctx, subject, objectType, permTiers...)

	if err != nil {
		return nil, err
	}

	out := make(map[string]authzAssignment, len(tuples))

	for _, t := range tuples {
		_, id, _ := strings.Cut(t.Object, ":")

		if existing, ok := out[id]; ok && tierRank(existing.Perm) >= tierRank(t.Relation) {
			continue
		}

		out[id] = authzAssignment{
			Subject:   t.Object,
			Perm:      t.Relation,
			CreatedAt: t.CreatedAt,
		}
	}

	return out, nil
}

func authzAssignmentFor(ctx context.Context, checker authz.Checker, object, subject string) (authzAssignment, bool, error) {
	assignments, err := authzDirectAssignments(ctx, checker, object)

	if err != nil {
		return authzAssignment{}, false, err
	}

	a, ok := assignments[subject]

	return a, ok, nil
}

// authzWriteAssignment grants the subject the given tier on the object,
// replacing any existing tier the subject already held.
func authzWriteAssignment(ctx context.Context, checker authz.Checker, object, subject, perm string) error {
	existing, ok, err := authzAssignmentFor(ctx, checker, object, subject)

	if err != nil {
		return err
	}

	if ok && existing.Perm != perm {
		if err := checker.Delete(ctx, authz.TupleKey{Object: object, Relation: existing.Perm, Subject: subject}); err != nil {
			return err
		}
	}

	if ok && existing.Perm == perm {
		return nil
	}

	return checker.Write(ctx, authz.TupleKey{Object: object, Relation: perm, Subject: subject})
}

// authzRemoveAssignment revokes whatever tier the subject currently holds on
// the object, if any.
func authzRemoveAssignment(ctx context.Context, checker authz.Checker, object, subject string) error {
	existing, ok, err := authzAssignmentFor(ctx, checker, object, subject)

	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	return checker.Delete(ctx, authz.TupleKey{Object: object, Relation: existing.Perm, Subject: subject})
}

// authzAssemble loads one item per assignment via load (skipping assignments
// for which load reports ok=false, e.g. a stale tuple pointing at a deleted
// record), filters by search, sorts via less and paginates the result. It
// replaces the SQL-level search/sort/limit/offset bun queries used by the
// former mapping tables, since assignments now live as OpenFGA tuples rather
// than rows in a joinable table.
func authzAssemble[T any](
	assignments map[string]authzAssignment,
	load func(subject string, a authzAssignment) (T, bool, error),
	search string,
	matches func(item T, search string) bool,
	less func(a, b T) bool,
	limit, offset int64,
) ([]T, int64, error) {
	items := make([]T, 0, len(assignments))

	for subject, a := range assignments {
		item, ok, err := load(subject, a)

		if err != nil {
			return nil, 0, err
		}

		if !ok {
			continue
		}

		if search != "" && !matches(item, search) {
			continue
		}

		items = append(items, item)
	}

	if less != nil {
		sort.Slice(items, func(i, j int) bool {
			return less(items[i], items[j])
		})
	}

	counter := int64(len(items))

	if offset > 0 {
		if offset >= counter {
			return []T{}, counter, nil
		}

		items = items[offset:]
	}

	if limit > 0 && limit < int64(len(items)) {
		items = items[:limit]
	}

	return items, counter, nil
}

// descending reports whether an order parameter requests descending sort.
func descending(order string) bool {
	return sortOrder(order) == "DESC"
}

func compareStrings(a, b string, desc bool) bool {
	if desc {
		return a > b
	}

	return a < b
}

func compareTimes(a, b time.Time, desc bool) bool {
	if desc {
		return a.After(b)
	}

	return a.Before(b)
}

// authzDeleteObject removes every direct assignment on the given object.
func authzDeleteObject(ctx context.Context, checker authz.Checker, object string) error {
	assignments, err := authzDirectAssignments(ctx, checker, object)

	if err != nil {
		return err
	}

	tuples := make([]authz.TupleKey, 0, len(assignments))

	for _, a := range assignments {
		tuples = append(tuples, authz.TupleKey{Object: object, Relation: a.Perm, Subject: a.Subject})
	}

	return checker.Delete(ctx, tuples...)
}

// authzDeleteSubject removes every direct assignment held by the given
// subject across the given object types, the reverse direction of
// authzDeleteObject — used to strip a user's own tuples when the user itself
// is being deleted.
func authzDeleteSubject(ctx context.Context, checker authz.Checker, subject string, objectTypes ...string) error {
	tuples := make([]authz.TupleKey, 0)

	for _, objectType := range objectTypes {
		assignments, err := authzSubjectAssignments(ctx, checker, objectType, subject)

		if err != nil {
			return err
		}

		for id, a := range assignments {
			tuples = append(tuples, authz.TupleKey{Object: authzObject(objectType, id), Relation: a.Perm, Subject: subject})
		}
	}

	return checker.Delete(ctx, tuples...)
}
