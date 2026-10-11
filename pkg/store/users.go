package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/uptrace/bun"
)

// Users provides all database operations related to users.
type Users struct {
	client *Store
}

// List implements the listing of all users.
func (s *Users) List(ctx context.Context, params model.ListParams) ([]*model.User, int64, error) {
	records := make([]*model.User, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Auths")

	if val, ok := s.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// Show implements the details for a specific user.
func (s *Users) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Relation("Auths").
		Where("id = ? OR username = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrUserNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new user.
func (s *Users) Create(ctx context.Context, record *model.User) error {
	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Update implements the update of an existing user.
func (s *Users) Update(ctx context.Context, record *model.User) error {
	if err := s.validate(ctx, record, true); err != nil {
		return err
	}

	if _, err := s.client.handle.NewUpdate().
		Model(record).
		Where("id = ?", record.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Delete implements the deletion of an user.
func (s *Users) Delete(ctx context.Context, name string) error {
	record, err := s.Show(ctx, name)

	if err != nil {
		return err
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.User)(nil)).
		Where("id = ?", record.ID).
		Exec(ctx); err != nil {
		return err
	}

	return authzDeleteSubject(ctx, s.client.authz, authzUserSubject(record.ID), "group", "mod", "pack")
}

// ShowRedirectToken implements the details for a specific redirect token.
func (s *Users) ShowRedirectToken(ctx context.Context, token string) (*model.UserToken, error) {
	record := &model.UserToken{}
	expired := time.Now().UTC().Add(-5 * time.Minute)

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("token = ? AND kind = ? AND created_at > ?", token, model.UserTokenKindRedirect, expired).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrTokenNotFound
		}

		return record, err
	}

	return record, nil
}

// CreateRedirectToken implements the create of a new redirect token.
func (s *Users) CreateRedirectToken(ctx context.Context, username string) (*model.UserToken, error) {
	user, err := s.Show(ctx, username)

	if err != nil {
		return nil, err
	}

	record := &model.UserToken{
		UserID: user.ID,
		Kind:   model.UserTokenKindRedirect,
		Token:  secret.Generate(32),
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return nil, err
	}

	return record, nil
}

// DeleteRedirectToken implements the deletion of a redirect token.
func (s *Users) DeleteRedirectToken(ctx context.Context, token string) error {
	if _, err := s.client.handle.NewDelete().
		Model((*model.UserToken)(nil)).
		Where("token = ? AND kind = ?", token, model.UserTokenKindRedirect).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// CleanupRedirectTokens implements the cleanup of expired redirect tokens.
func (s *Users) CleanupRedirectTokens(ctx context.Context) error {
	expired := time.Now().UTC().Add(-5 * time.Minute)

	if _, err := s.client.handle.NewDelete().
		Model((*model.UserToken)(nil)).
		Where("kind = ? AND created_at < ?", model.UserTokenKindRedirect, expired).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// ListGroups implements the listing of all groups for a user.
func (s *Users) ListGroups(ctx context.Context, params model.UserGroupParams) ([]*model.UserGroup, int64, error) {
	assignments, err := authzUserAssignments(ctx, s.client.authz, "group", params.UserID)

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(groupID string, a authzAssignment) (*model.UserGroup, bool, error) {
			group, err := s.client.Groups.Show(ctx, groupID)

			if err != nil {
				if errors.Is(err, ErrGroupNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserGroup{
				UserID:    params.UserID,
				GroupID:   group.ID,
				Group:     group,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserGroup, search string) bool {
			return searchMatches(search, item.Group.Name, item.Group.Slug)
		},
		func(a, b *model.UserGroup) bool {
			switch strings.ToLower(params.Sort) {
			case "slug":
				return compareStrings(a.Group.Slug, b.Group.Slug, desc)
			case "created":
				return compareTimes(a.CreatedAt, b.CreatedAt, desc)
			case "updated":
				return compareTimes(a.UpdatedAt, b.UpdatedAt, desc)
			default:
				return compareStrings(a.Group.Name, b.Group.Name, desc)
			}
		},
		params.Limit,
		params.Offset,
	)
}

// AttachGroup implements the attachment of a user to a group.
func (s *Users) AttachGroup(ctx context.Context, params model.UserGroupParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	group, err := s.client.Groups.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID), params.Perm)
}

// PermitGroup implements the permission update for a group on a user.
func (s *Users) PermitGroup(ctx context.Context, params model.UserGroupParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	group, err := s.client.Groups.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID), params.Perm)
}

// DropGroup implements the removal of a user from a group.
func (s *Users) DropGroup(ctx context.Context, params model.UserGroupParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	group, err := s.client.Groups.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzRemoveAssignment(ctx, s.client.authz, authzObject("group", group.ID), authzUserSubject(user.ID))
}

// ListMods implements the listing of all mods for a user.
func (s *Users) ListMods(ctx context.Context, params model.UserModParams) ([]*model.UserMod, int64, error) {
	assignments, err := authzUserAssignments(ctx, s.client.authz, "mod", params.UserID)

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(modID string, a authzAssignment) (*model.UserMod, bool, error) {
			mod, err := s.client.Mods.Show(ctx, modID)

			if err != nil {
				if errors.Is(err, ErrModNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserMod{
				UserID:    params.UserID,
				ModID:     mod.ID,
				Mod:       mod,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserMod, search string) bool {
			return searchMatches(search, item.Mod.Name, item.Mod.Slug)
		},
		func(a, b *model.UserMod) bool {
			switch strings.ToLower(params.Sort) {
			case "slug":
				return compareStrings(a.Mod.Slug, b.Mod.Slug, desc)
			case "created":
				return compareTimes(a.CreatedAt, b.CreatedAt, desc)
			case "updated":
				return compareTimes(a.UpdatedAt, b.UpdatedAt, desc)
			default:
				return compareStrings(a.Mod.Name, b.Mod.Name, desc)
			}
		},
		params.Limit,
		params.Offset,
	)
}

// AttachMod implements the attachment of a user to a mod.
func (s *Users) AttachMod(ctx context.Context, params model.UserModParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID), params.Perm)
}

// PermitMod implements the permission update for a mod on a user.
func (s *Users) PermitMod(ctx context.Context, params model.UserModParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID), params.Perm)
}

// DropMod implements the removal of a user from a mod.
func (s *Users) DropMod(ctx context.Context, params model.UserModParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzRemoveAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))
}

// ListPacks implements the listing of all packs for a user.
func (s *Users) ListPacks(ctx context.Context, params model.UserPackParams) ([]*model.UserPack, int64, error) {
	assignments, err := authzUserAssignments(ctx, s.client.authz, "pack", params.UserID)

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(packID string, a authzAssignment) (*model.UserPack, bool, error) {
			pack, err := s.client.Packs.Show(ctx, packID)

			if err != nil {
				if errors.Is(err, ErrPackNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserPack{
				UserID:    params.UserID,
				PackID:    pack.ID,
				Pack:      pack,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserPack, search string) bool {
			return searchMatches(search, item.Pack.Name, item.Pack.Slug)
		},
		func(a, b *model.UserPack) bool {
			switch strings.ToLower(params.Sort) {
			case "slug":
				return compareStrings(a.Pack.Slug, b.Pack.Slug, desc)
			case "created":
				return compareTimes(a.CreatedAt, b.CreatedAt, desc)
			case "updated":
				return compareTimes(a.UpdatedAt, b.UpdatedAt, desc)
			default:
				return compareStrings(a.Pack.Name, b.Pack.Name, desc)
			}
		},
		params.Limit,
		params.Offset,
	)
}

// AttachPack implements the attachment of a user to a pack.
func (s *Users) AttachPack(ctx context.Context, params model.UserPackParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID), params.Perm)
}

// PermitPack implements the permission update for a pack on a user.
func (s *Users) PermitPack(ctx context.Context, params model.UserPackParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	if err := s.validatePerm(params.Perm); err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID), params.Perm)
}

// DropPack implements the removal of a user from a pack.
func (s *Users) DropPack(ctx context.Context, params model.UserPackParams) error {
	user, err := s.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzRemoveAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))
}

func (s *Users) validatePerm(perm string) error {
	if err := validation.Validate(
		perm,
		validation.In("user", "admin"),
	); err != nil {
		return validate.Errors{
			Errors: []validate.Error{
				{
					Field: "perm",
					Error: fmt.Errorf("invalid permission value"),
				},
			},
		}
	}

	return nil
}

func (s *Users) validate(ctx context.Context, record *model.User, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Username,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "username", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (s *Users) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.User)(nil)).
			Where("? = ?", bun.Ident(key), val)

		if id != "" {
			q = q.Where(
				"id != ?",
				id,
			)
		}

		exists, err := q.Exists(ctx)

		if err != nil {
			return err
		}

		if exists {
			return errors.New("is already taken")
		}

		return nil
	}
}

// ValidSort validates the given sorting column.
func (s *Users) ValidSort(val string) (string, bool) {
	if val == "" {
		return "user.username", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"username": "user.username",
		"email":    "user.email",
		"fullname": "user.fullname",
		"admin":    "user.admin",
		"active":   "user.active",
		"created":  "user.created_at",
		"updated":  "user.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "user.username", true
}
