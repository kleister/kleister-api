package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/authz"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/uptrace/bun"
)

// Groups provides all database operations related to groups.
type Groups struct {
	client *Store
}

// List implements the listing of all users.
func (s *Groups) List(ctx context.Context, params model.ListParams) ([]*model.Group, int64, error) {
	records := make([]*model.Group, 0)

	q := s.client.handle.NewSelect().
		Model(&records)

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
func (s *Groups) Show(ctx context.Context, name string) (*model.Group, error) {
	record := &model.Group{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("id = ? OR slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrGroupNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new group.
func (s *Groups) Create(ctx context.Context, record *model.Group) error {
	if record.Slug == "" {
		record.Slug = s.slugify(
			ctx,
			"slug",
			record.Name,
			"",
		)
	}

	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().
			Model(record).
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return s.client.authz.Write(ctx, authz.TupleKey{
		Object:   authzObject("group", record.ID),
		Relation: model.UserGroupAdminPerm,
		Subject:  authzUserSubject(s.client.principal.ID),
	})
}

// Update implements the update of an existing group.
func (s *Groups) Update(ctx context.Context, record *model.Group) error {
	if record.Slug == "" {
		record.Slug = s.slugify(
			ctx,
			"slug",
			record.Name,
			record.ID,
		)
	}

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

// Delete implements the deletion of a group.
func (s *Groups) Delete(ctx context.Context, name string) error {
	record, err := s.Show(ctx, name)

	if err != nil {
		return err
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.Group)(nil)).
		Where("id = ?", record.ID).
		Exec(ctx); err != nil {
		return err
	}

	if err := authzDeleteObject(ctx, s.client.authz, authzObject("group", record.ID)); err != nil {
		return err
	}

	return authzDeleteSubject(ctx, s.client.authz, authzGroupSubject(record.ID), "mod", "pack")
}

// ListUsers implements the listing of all users for a group.
func (s *Groups) ListUsers(ctx context.Context, params model.UserGroupParams) ([]*model.UserGroup, int64, error) {
	assignments, err := authzDirectAssignments(ctx, s.client.authz, authzObject("group", params.GroupID))

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(subject string, a authzAssignment) (*model.UserGroup, bool, error) {
			user, err := s.client.Users.Show(ctx, userIDFromSubject(subject))

			if err != nil {
				if errors.Is(err, ErrUserNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserGroup{
				GroupID:   params.GroupID,
				UserID:    user.ID,
				User:      user,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserGroup, search string) bool {
			return searchMatches(search, item.User.Username, item.User.Email, item.User.Fullname)
		},
		func(a, b *model.UserGroup) bool {
			switch strings.ToLower(params.Sort) {
			case "email":
				return compareStrings(a.User.Email, b.User.Email, desc)
			case "fullname":
				return compareStrings(a.User.Fullname, b.User.Fullname, desc)
			case "created":
				return compareTimes(a.CreatedAt, b.CreatedAt, desc)
			case "updated":
				return compareTimes(a.UpdatedAt, b.UpdatedAt, desc)
			default:
				return compareStrings(a.User.Username, b.User.Username, desc)
			}
		},
		params.Limit,
		params.Offset,
	)
}

// AttachUser implements the attachment of a group to a user.
func (s *Groups) AttachUser(ctx context.Context, params model.UserGroupParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

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

// PermitUser implements the permission update for a user on a group.
func (s *Groups) PermitUser(ctx context.Context, params model.UserGroupParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

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

// DropUser implements the removal of a group from an user.
func (s *Groups) DropUser(ctx context.Context, params model.UserGroupParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

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

// ListMods implements the listing of all mods for a group.
func (s *Groups) ListMods(ctx context.Context, params model.GroupModParams) ([]*model.GroupMod, int64, error) {
	assignments, err := authzGroupAssignments(ctx, s.client.authz, "mod", params.GroupID)

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(modID string, a authzAssignment) (*model.GroupMod, bool, error) {
			mod, err := s.client.Mods.Show(ctx, modID)

			if err != nil {
				if errors.Is(err, ErrModNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.GroupMod{
				GroupID:   params.GroupID,
				ModID:     mod.ID,
				Mod:       mod,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.GroupMod, search string) bool {
			return searchMatches(search, item.Mod.Name, item.Mod.Slug)
		},
		func(a, b *model.GroupMod) bool {
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

// AttachMod implements the attachment of a group to a mod.
func (s *Groups) AttachMod(ctx context.Context, params model.GroupModParams) error {
	group, err := s.Show(ctx, params.GroupID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID), params.Perm)
}

// PermitMod implements the permission update for a mod on a group.
func (s *Groups) PermitMod(ctx context.Context, params model.GroupModParams) error {
	group, err := s.Show(ctx, params.GroupID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID), params.Perm)
}

// DropMod implements the removal of a group from a mod.
func (s *Groups) DropMod(ctx context.Context, params model.GroupModParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzRemoveAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))
}

// ListPacks implements the listing of all packs for a group.
func (s *Groups) ListPacks(ctx context.Context, params model.GroupPackParams) ([]*model.GroupPack, int64, error) {
	assignments, err := authzGroupAssignments(ctx, s.client.authz, "pack", params.GroupID)

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(packID string, a authzAssignment) (*model.GroupPack, bool, error) {
			pack, err := s.client.Packs.Show(ctx, packID)

			if err != nil {
				if errors.Is(err, ErrPackNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.GroupPack{
				GroupID:   params.GroupID,
				PackID:    pack.ID,
				Pack:      pack,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.GroupPack, search string) bool {
			return searchMatches(search, item.Pack.Name, item.Pack.Slug)
		},
		func(a, b *model.GroupPack) bool {
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

// AttachPack implements the attachment of a group to a pack.
func (s *Groups) AttachPack(ctx context.Context, params model.GroupPackParams) error {
	group, err := s.Show(ctx, params.GroupID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID), params.Perm)
}

// PermitPack implements the permission update for a pack on a group.
func (s *Groups) PermitPack(ctx context.Context, params model.GroupPackParams) error {
	group, err := s.Show(ctx, params.GroupID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID), params.Perm)
}

// DropPack implements the removal of a group from a pack.
func (s *Groups) DropPack(ctx context.Context, params model.GroupPackParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzRemoveAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))
}

func (s *Groups) validatePerm(perm string) error {
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

func (s *Groups) validate(ctx context.Context, record *model.Group, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Slug,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (s *Groups) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.Group)(nil)).
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

func (s *Groups) slugify(ctx context.Context, column, value, id string) string {
	var (
		slug string
	)

	for i := 0; true; i++ {
		if i == 0 {
			slug = slugify.Slugify(value)
		} else {
			slug = slugify.Slugify(
				fmt.Sprintf("%s-%s", value, uniuri.NewLen(6)),
			)
		}

		query := s.client.handle.NewSelect().
			Model((*model.Group)(nil)).
			Where("? = ?", bun.Ident(column), slug)

		if id != "" {
			query = query.Where(
				"id != ?",
				id,
			)
		}

		if count, err := query.Count(
			ctx,
		); err == nil && count == 0 {
			break
		}
	}

	return slug
}

// ValidSort validates the given sorting column.
func (s *Groups) ValidSort(val string) (string, bool) {
	if val == "" {
		return "group.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug":    "group.slug",
		"name":    "group.name",
		"created": "group.created_at",
		"updated": "group.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "group.name", true
}
