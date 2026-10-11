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

// Packs provides all database operations related to packs.
type Packs struct {
	client *Store
}

// List implements the listing of all packs.
func (s *Packs) List(ctx context.Context, params model.ListParams) ([]*model.Pack, int64, error) {
	records := make([]*model.Pack, 0)

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
func (s *Packs) Show(ctx context.Context, name string) (*model.Pack, error) {
	record := &model.Pack{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Relation("Avatar").
		Where("pack.id = ? OR pack.slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrPackNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new pack.
func (s *Packs) Create(ctx context.Context, record *model.Pack) error {
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
		"packs",
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

		record.Avatar = &model.PackAvatar{
			PackID: record.ID,
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
		Object:   authzObject("pack", record.ID),
		Relation: model.UserPackAdminPerm,
		Subject:  authzUserSubject(s.client.principal.ID),
	})
}

// Update implements the update of an existing pack.
func (s *Packs) Update(ctx context.Context, record *model.Pack) error {
	if record.Slug == "" {
		record.Slug = s.slugify(
			ctx,
			"slug",
			record.Name,
			"",
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

// Delete implements the deletion of a pack.
func (s *Packs) Delete(ctx context.Context, name string) error {
	record, err := s.Show(ctx, name)

	if err != nil {
		return err
	}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewDelete().
			Model((*model.Pack)(nil)).
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

		return nil
	}); err != nil {
		return err
	}

	return authzDeleteObject(ctx, s.client.authz, authzObject("pack", record.ID))
}

// CreateAvatar implements the upload and storage of a pack avatar.
func (s *Packs) CreateAvatar(ctx context.Context, name string, content *bytes.Buffer) (*model.PackAvatar, error) {
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
func (s *Packs) DeleteAvatar(ctx context.Context, name string) (*model.PackAvatar, error) {
	record, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	content, err := identicon.New(
		"packs",
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

func (s *Packs) replaceAvatar(ctx context.Context, record *model.Pack, tx bun.Tx, content *bytes.Buffer) error {
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
			Model((*model.PackAvatar)(nil)).
			Where("pack_id = ?", record.ID).
			Exec(ctx); err != nil {
			return err
		}
	}

	record.Avatar = &model.PackAvatar{
		PackID: record.ID,
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

// ListGroups implements the listing of all groups for a pack.
func (s *Packs) ListGroups(ctx context.Context, params model.GroupPackParams) ([]*model.GroupPack, int64, error) {
	assignments, err := authzDirectAssignments(ctx, s.client.authz, authzObject("pack", params.PackID))

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(subject string, a authzAssignment) (*model.GroupPack, bool, error) {
			group, err := s.client.Groups.Show(ctx, groupIDFromSubject(subject))

			if err != nil {
				if errors.Is(err, ErrGroupNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.GroupPack{
				PackID:    params.PackID,
				GroupID:   group.ID,
				Group:     group,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.GroupPack, search string) bool {
			return searchMatches(search, item.Group.Name, item.Group.Slug)
		},
		func(a, b *model.GroupPack) bool {
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

// AttachGroup implements the attachment of a pack to a group.
func (s *Packs) AttachGroup(ctx context.Context, params model.GroupPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID), params.Perm)
}

// PermitGroup implements the permission update for a group on a pack.
func (s *Packs) PermitGroup(ctx context.Context, params model.GroupPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzGroupSubject(group.ID), params.Perm)
}

// DropGroup implements the removal of a pack from a group.
func (s *Packs) DropGroup(ctx context.Context, params model.GroupPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	group, err := s.client.Groups.Show(ctx, params.GroupID)

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

// ListUsers implements the listing of all users for a pack.
func (s *Packs) ListUsers(ctx context.Context, params model.UserPackParams) ([]*model.UserPack, int64, error) {
	assignments, err := authzDirectAssignments(ctx, s.client.authz, authzObject("pack", params.PackID))

	if err != nil {
		return nil, 0, err
	}

	desc := descending(params.Order)

	return authzAssemble(
		assignments,
		func(subject string, a authzAssignment) (*model.UserPack, bool, error) {
			user, err := s.client.Users.Show(ctx, userIDFromSubject(subject))

			if err != nil {
				if errors.Is(err, ErrUserNotFound) {
					return nil, false, nil
				}

				return nil, false, err
			}

			return &model.UserPack{
				PackID:    params.PackID,
				UserID:    user.ID,
				User:      user,
				Perm:      a.Perm,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.CreatedAt,
			}, true, nil
		},
		params.Search,
		func(item *model.UserPack, search string) bool {
			return searchMatches(search, item.User.Username, item.User.Email, item.User.Fullname)
		},
		func(a, b *model.UserPack) bool {
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

// AttachUser implements the attachment of a pack to a user.
func (s *Packs) AttachUser(ctx context.Context, params model.UserPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID), params.Perm)
}

// PermitUser implements the permission update for a user on a pack.
func (s *Packs) PermitUser(ctx context.Context, params model.UserPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

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

	_, assigned, err := authzAssignmentFor(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID))

	if err != nil {
		return err
	}

	if !assigned {
		return ErrNotAssigned
	}

	return authzWriteAssignment(ctx, s.client.authz, authzObject("pack", pack.ID), authzUserSubject(user.ID), params.Perm)
}

// DropUser implements the removal of a pack from a user.
func (s *Packs) DropUser(ctx context.Context, params model.UserPackParams) error {
	pack, err := s.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

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

func (s *Packs) validatePerm(perm string) error {
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

func (s *Packs) validate(ctx context.Context, record *model.Pack, _ bool) error {
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

func (s *Packs) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.Pack)(nil)).
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

func (s *Packs) slugify(ctx context.Context, column, value, id string) string {
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
			Model((*model.Pack)(nil)).
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
func (s *Packs) ValidSort(val string) (string, bool) {
	if val == "" {
		return "pack.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug":    "pack.slug",
		"name":    "pack.name",
		"website": "pack.website",
		"public":  "pack.public",
		"created": "pack.created_at",
		"updated": "pack.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "pack.name", true
}
