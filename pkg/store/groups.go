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
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
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

	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().
			Model(record).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := tx.NewInsert().
			Model(&model.UserGroup{
				GroupID: record.ID,
				UserID:  s.client.principal.ID,
				Perm:    model.UserGroupAdminPerm,
			}).
			Exec(ctx); err != nil {
			return err
		}

		return nil
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

	return nil
}

// ListUsers implements the listing of all users for a group.
func (s *Groups) ListUsers(ctx context.Context, params model.UserGroupParams) ([]*model.UserGroup, int64, error) {
	records := make([]*model.UserGroup, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("User").
		Relation("Group").
		Where("group_id = ?", params.GroupID)

	if val, ok := s.client.Users.ValidSort(params.Sort); ok {
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
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// AttachUser implements the attachment of a group to an user.
func (s *Groups) AttachUser(ctx context.Context, params model.UserGroupParams) error {
	group, err := s.Show(ctx, params.GroupID)

	if err != nil {
		return err
	}

	user, err := s.client.Users.Show(ctx, params.UserID)

	if err != nil {
		return err
	}

	assigned, err := s.isUserAssigned(ctx, group.ID, user.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	record := &model.UserGroup{
		GroupID: group.ID,
		UserID:  user.ID,
		Perm:    params.Perm,
	}

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isUserUnassigned(ctx, group.ID, user.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewUpdate().
		Model((*model.UserGroup)(nil)).
		Set("perm = ?", params.Perm).
		Where("group_id = ? AND user_id = ?", group.ID, user.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isUserUnassigned(ctx, group.ID, user.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.UserGroup)(nil)).
		Where("group_id = ? AND user_id = ?", group.ID, user.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Groups) isUserAssigned(ctx context.Context, groupID, userID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.UserGroup)(nil)).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Groups) isUserUnassigned(ctx context.Context, groupID, userID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.UserGroup)(nil)).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}

// ListMods implements the listing of all mods for a group.
func (s *Groups) ListMods(ctx context.Context, params model.GroupModParams) ([]*model.GroupMod, int64, error) {
	records := make([]*model.GroupMod, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Group").
		Relation("Mod").
		Where("group_id = ?", params.GroupID)

	if val, ok := s.client.Mods.ValidSort(params.Sort); ok {
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
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
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

	assigned, err := s.isModAssigned(ctx, group.ID, mod.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	record := &model.GroupMod{
		GroupID: group.ID,
		ModID:   mod.ID,
		Perm:    params.Perm,
	}

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isModUnassigned(ctx, group.ID, mod.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewUpdate().
		Model((*model.GroupMod)(nil)).
		Set("perm = ?", params.Perm).
		Where("group_id = ? AND mod_id = ?", group.ID, mod.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isModUnassigned(ctx, group.ID, mod.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.GroupMod)(nil)).
		Where("group_id = ? AND mod_id = ?", group.ID, mod.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Groups) isModAssigned(ctx context.Context, groupID, modID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.GroupMod)(nil)).
		Where("group_id = ? AND mod_id = ?", groupID, modID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Groups) isModUnassigned(ctx context.Context, groupID, modID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.GroupMod)(nil)).
		Where("group_id = ? AND mod_id = ?", groupID, modID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}

// ListPacks implements the listing of all packs for a group.
func (s *Groups) ListPacks(ctx context.Context, params model.GroupPackParams) ([]*model.GroupPack, int64, error) {
	records := make([]*model.GroupPack, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Group").
		Relation("Pack").
		Where("group_id = ?", params.GroupID)

	if val, ok := s.client.Packs.ValidSort(params.Sort); ok {
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
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
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

	assigned, err := s.isPackAssigned(ctx, group.ID, pack.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	record := &model.GroupPack{
		GroupID: group.ID,
		PackID:  pack.ID,
		Perm:    params.Perm,
	}

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isPackUnassigned(ctx, group.ID, pack.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewUpdate().
		Model((*model.GroupPack)(nil)).
		Set("perm = ?", params.Perm).
		Where("group_id = ? AND pack_id = ?", group.ID, pack.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
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

	unassigned, err := s.isPackUnassigned(ctx, group.ID, pack.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.GroupPack)(nil)).
		Where("group_id = ? AND pack_id = ?", group.ID, pack.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Groups) isPackAssigned(ctx context.Context, groupID, packID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.GroupPack)(nil)).
		Where("group_id = ? AND pack_id = ?", groupID, packID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Groups) isPackUnassigned(ctx context.Context, groupID, packID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.GroupPack)(nil)).
		Where("group_id = ? AND pack_id = ?", groupID, packID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
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
