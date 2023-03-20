package repository

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// NewGormRepository initializes a new repository for GormDB.
func NewGormRepository(
	handle *gorm.DB,
) *GormRepository {
	return &GormRepository{
		handle: handle,
	}
}

// GormRepository implements the ModsRepository interface.
type GormRepository struct {
	handle *gorm.DB
}

// List implements the ModsRepository interface.
func (r *GormRepository) List(ctx context.Context, _ string) ([]*model.Mod, error) {
	records := make([]*model.Mod, 0)

	// TODO: use search if given

	if err := r.query(ctx).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Create implements the ModsRepository interface.
func (r *GormRepository) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if mod.Slug == "" {
		mod.Slug = store.Slugify(
			tx.Model(&model.Mod{}),
			mod.Name,
			"",
		)
	}

	mod.ID = uuid.New().String()

	if err := r.validate(ctx, mod, false); err != nil {
		return nil, err
	}

	if err := tx.Create(mod).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return mod, nil
}

// Update implements the ModsRepository interface.
func (r *GormRepository) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if mod.Slug == "" {
		mod.Slug = store.Slugify(
			tx.Model(&model.Mod{}),
			mod.Name,
			mod.ID,
		)
	}

	if err := r.validate(ctx, mod, true); err != nil {
		return nil, err
	}

	if err := tx.Save(mod).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return mod, nil
}

// Show implements the ModsRepository interface.
func (r *GormRepository) Show(ctx context.Context, id string) (*model.Mod, error) {
	record := &model.Mod{}

	err := r.query(ctx).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrModNotFound
	}

	return record, err
}

// Delete implements the ModsRepository interface.
func (r *GormRepository) Delete(ctx context.Context, id string) error {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Delete(
		&model.Mod{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the ModsRepository interface.
func (r *GormRepository) Exists(ctx context.Context, id string) (bool, string, error) {
	record := &model.Mod{}

	res := r.query(ctx).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Find(
		record,
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, "", nil
	}

	if res.Error != nil {
		return false, "", res.Error
	}

	return res.RowsAffected > 0, record.ID, nil
}

// ListUsers implements the ModsRepository interface.
func (r *GormRepository) ListUsers(ctx context.Context, modID, _ string) ([]*model.UserMod, error) {
	records := make([]*model.UserMod, 0)

	// TODO: use search if given

	if err := r.handle.WithContext(
		ctx,
	).Where(
		"mod_id = ?",
		modID,
	).Order(
		"User.username ASC",
	).Model(
		&model.UserMod{},
	).Preload(
		"User",
	).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// AttachUser implements the ModsRepository interface.
func (r *GormRepository) AttachUser(_ context.Context, _, _ string) error {
	return errors.New("not implemented")
}

// DropUser implements the ModsRepository interface.
func (r *GormRepository) DropUser(_ context.Context, _, _ string) error {
	return errors.New("not implemented")
}

// ListTeams implements the ModsRepository interface.
func (r *GormRepository) ListTeams(ctx context.Context, modID, _ string) ([]*model.TeamMod, error) {
	records := make([]*model.TeamMod, 0)

	// TODO: use search if given

	if err := r.handle.WithContext(
		ctx,
	).Where(
		"mod_id = ?",
		modID,
	).Order(
		"Team.name ASC",
	).Model(
		&model.TeamMod{},
	).Preload(
		"Team",
	).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// AttachTeam implements the ModsRepository interface.
func (r *GormRepository) AttachTeam(_ context.Context, _, _ string) error {
	return errors.New("not implemented")
}

// DropTeam implements the ModsRepository interface.
func (r *GormRepository) DropTeam(_ context.Context, _, _ string) error {
	return errors.New("not implemented")
}

func (r *GormRepository) validate(ctx context.Context, record *model.Mod, existing bool) error {
	errs := validate.Errors{}

	if existing {
		if err := validation.Validate(
			record.ID,
			validation.Required,
			is.UUIDv4,
			validation.By(r.uniqueValueIsPresent(ctx, "id", record.ID)),
		); err != nil {
			errs.Errors = append(errs.Errors, validate.Error{
				Field: "id",
				Error: err,
			})
		}
	}

	if err := validation.Validate(
		record.Slug,
		validation.Required,
		validation.Length(3, 255),
		validation.By(r.uniqueValueIsPresent(ctx, "slug", record.ID)),
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
		validation.By(r.uniqueValueIsPresent(ctx, "name", record.ID)),
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

func (r *GormRepository) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		res := r.handle.WithContext(ctx).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Not(
			"id = ?",
			id,
		).Find(
			&model.Mod{},
		)

		if res.RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (r *GormRepository) query(ctx context.Context) *gorm.DB {
	return r.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Mod{},
	).Preload(
		"Icon",
	).Preload(
		"Logo",
	).Preload(
		"Back",
	)
}
