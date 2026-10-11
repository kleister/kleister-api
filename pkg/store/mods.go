package store

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	"github.com/gabriel-vasile/mimetype"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/authz"
	"github.com/kleister/kleister-api/pkg/identicon"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/uptrace/bun"
)

// Mods provides all database operations related to mods.
type Mods struct {
	client *Store
}

// List implements the listing of all users.
func (s *Mods) List(ctx context.Context, params model.ListParams) ([]*model.Mod, int64, error) {
	records := make([]*model.Mod, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Avatar")

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
func (s *Mods) Show(ctx context.Context, name string) (*model.Mod, error) {
	record := &model.Mod{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Relation("Avatar").
		Where("mod.id = ? OR mod.slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrModNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new mod.
func (s *Mods) Create(ctx context.Context, record *model.Mod) error {
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

	avatar, err := identicon.New(
		"mods",
		record.Name,
	)

	if err != nil {
		return err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().
			Model(record).
			Exec(ctx); err != nil {
			return err
		}

		record.Avatar = &model.ModAvatar{
			ModID: record.ID,
			Slug: strings.Join(
				[]string{
					secret.Generate(64),
					"png",
				},
				".",
			),
		}

		if _, err := tx.NewInsert().
			Model(record.Avatar).
			Exec(ctx); err != nil {
			return err
		}

		return s.client.upload.Upload(
			ctx,
			filepath.Join(
				"avatars",
				record.Avatar.Slug,
			),
			avatar,
		)
	}); err != nil {
		return err
	}

	return s.client.authz.Write(ctx, authz.TupleKey{
		Object:   authzObject("mod", record.ID),
		Relation: model.UserModAdminPerm,
		Subject:  authzUserSubject(s.client.principal.ID),
	})
}

// Update implements the update of an existing mod.
func (s *Mods) Update(ctx context.Context, record *model.Mod) error {
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

// Delete implements the deletion of a mod.
func (s *Mods) Delete(ctx context.Context, name string) error {
	record, err := s.Show(ctx, name)

	if err != nil {
		return err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewDelete().
			Model((*model.Mod)(nil)).
			Where("id = ?", record.ID).
			Exec(ctx); err != nil {
			return err
		}

		if record.Avatar != nil {
			if err := s.client.upload.Delete(
				ctx,
				filepath.Join(
					"avatars",
					record.Avatar.Slug,
				),
				false,
			); err != nil {
				return err
			}
		}

		return s.client.upload.Delete(
			ctx,
			filepath.Join(
				"versions",
				record.ID,
			),
			true,
		)
	}); err != nil {
		return err
	}

	return authzDeleteObject(ctx, s.client.authz, authzObject("mod", record.ID))
}

// CreateAvatar implements the upload and storage of a mod avatar.
func (s *Mods) CreateAvatar(ctx context.Context, name string, content *bytes.Buffer) (*model.ModAvatar, error) {
	record, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		return s.replaceAvatar(ctx, record, tx, content)
	}); err != nil {
		return nil, err
	}

	record, err = s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return record.Avatar, nil
}

// DeleteAvatar implements the deletion of a mod avatar.
func (s *Mods) DeleteAvatar(ctx context.Context, name string) (*model.ModAvatar, error) {
	record, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	content, err := identicon.New(
		"mods",
		record.Name,
	)

	if err != nil {
		return nil, err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		return s.replaceAvatar(ctx, record, tx, content)
	}); err != nil {
		return nil, err
	}

	record, err = s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return record.Avatar, nil
}

func (s *Mods) replaceAvatar(ctx context.Context, record *model.Mod, tx bun.Tx, content *bytes.Buffer) error {
	reader := bytes.NewReader(
		content.Bytes(),
	)

	mtype, err := mimetype.DetectReader(
		reader,
	)

	if err != nil {
		return err
	}

	if record.Avatar != nil {
		if err := s.client.upload.Delete(
			ctx,
			filepath.Join(
				"avatars",
				record.Avatar.Slug,
			),
			false,
		); err != nil {
			return err
		}

		if _, err := tx.NewDelete().
			Model((*model.ModAvatar)(nil)).
			Where("mod_id = ?", record.ID).
			Exec(ctx); err != nil {
			return err
		}
	}

	record.Avatar = &model.ModAvatar{
		ModID: record.ID,
		Slug: strings.Join(
			[]string{
				secret.Generate(64),
				mtype.Extension(),
			},
			"",
		),
	}

	if _, err := tx.NewInsert().
		Model(record.Avatar).
		Exec(ctx); err != nil {
		return err
	}

	return s.client.upload.Upload(
		ctx,
		filepath.Join(
			"avatars",
			record.Avatar.Slug,
		),
		content,
	)
}

// ListGroups implements the listing of all groups for a mod.
func (s *Mods) ListGroups(ctx context.Context, params model.GroupModParams) ([]*model.GroupMod, int64, error) {
	assignments, err := authzDirectAssignments(ctx, s.client.authz, authzObject("mod", params.ModID))

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(subject string, a authzAssignment) (*model.GroupMod, bool, error) {
			group, err := s.client.Groups.Show(ctx, groupIDFromSubject(subject))

			if err != nil {
				if errors.Is(err, ErrGroupNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.GroupMod{
				ModID:     params.ModID,
				GroupID:   group.ID,
				Group:     group,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.GroupMod, search string) bool {
			return searchMatches(search, item.Group.Name, item.Group.Slug)
		},
		func(a, b *model.GroupMod) bool {
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

// AttachGroup implements the attachment of a mod to a group.
func (s *Mods) AttachGroup(ctx context.Context, params model.GroupModParams) error {
	mod, err := s.Show(ctx, params.ModID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID), params.Perm)
}

// PermitGroup implements the permission update for a group on a mod.
func (s *Mods) PermitGroup(ctx context.Context, params model.GroupModParams) error {
	mod, err := s.Show(ctx, params.ModID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzGroupSubject(group.ID), params.Perm)
}

// DropGroup implements the removal of a mod from a group.
func (s *Mods) DropGroup(ctx context.Context, params model.GroupModParams) error {
	mod, err := s.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	group, err := s.client.Groups.Show(ctx, params.GroupID)

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

// ListUsers implements the listing of all users for a mod.
func (s *Mods) ListUsers(ctx context.Context, params model.UserModParams) ([]*model.UserMod, int64, error) {
	assignments, err := authzDirectAssignments(ctx, s.client.authz, authzObject("mod", params.ModID))

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(subject string, a authzAssignment) (*model.UserMod, bool, error) {
			user, err := s.client.Users.Show(ctx, userIDFromSubject(subject))

			if err != nil {
				if errors.Is(err, ErrUserNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserMod{
				ModID:     params.ModID,
				UserID:    user.ID,
				User:      user,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserMod, search string) bool {
			return searchMatches(search, item.User.Username, item.User.Email, item.User.Fullname)
		},
		func(a, b *model.UserMod) bool {
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

// AttachUser implements the attachment of a mod to a user.
func (s *Mods) AttachUser(ctx context.Context, params model.UserModParams) error {
	mod, err := s.Show(ctx, params.ModID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID), params.Perm)
}

// PermitUser implements the permission update for a user on a mod.
func (s *Mods) PermitUser(ctx context.Context, params model.UserModParams) error {
	mod, err := s.Show(ctx, params.ModID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("mod", mod.ID), authzUserSubject(user.ID), params.Perm)
}

// DropUser implements the removal of a mod from a user.
func (s *Mods) DropUser(ctx context.Context, params model.UserModParams) error {
	mod, err := s.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

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

func (s *Mods) validatePerm(perm string) error {
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

func (s *Mods) validate(ctx context.Context, record *model.Mod, _ bool) error {
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

func (s *Mods) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.Mod)(nil)).
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

func (s *Mods) slugify(ctx context.Context, column, value, id string) string {
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
			Model((*model.Mod)(nil)).
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
func (s *Mods) ValidSort(val string) (string, bool) {
	if val == "" {
		return "mod.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug":        "mod.slug",
		"name":        "mod.name",
		"side":        "mod.side",
		"description": "mod.description",
		"author":      "mod.author",
		"website":     "mod.website",
		"public":      "mod.public",
		"created":     "mod.created_at",
		"updated":     "mod.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "mod.name", true
}
