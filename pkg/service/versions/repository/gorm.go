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

// GormRepository implements the VersionsRepository interface.
type GormRepository struct {
	handle *gorm.DB
}

// List implements the VersionsRepository interface.
func (r *GormRepository) List(ctx context.Context, modID string) ([]*model.Version, error) {
	records := make([]*model.Version, 0)

	if err := r.query(ctx).Where(
		"mod_id = ?",
		modID,
	).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Create implements the VersionsRepository interface.
func (r *GormRepository) Create(ctx context.Context, build *model.Version) (*model.Version, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if build.Slug == "" {
		build.Slug = store.Slugify(
			tx.Model(&model.Version{}),
			build.Name,
			"",
		)
	}

	build.ID = uuid.New().String()

	if err := r.validate(ctx, build, false); err != nil {
		return nil, err
	}

	if err := tx.Create(build).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return build, nil
}

// Update implements the VersionsRepository interface.
func (r *GormRepository) Update(ctx context.Context, build *model.Version) (*model.Version, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if build.Slug == "" {
		build.Slug = store.Slugify(
			tx.Model(&model.Version{}),
			build.Name,
			build.ID,
		)
	}

	if err := r.validate(ctx, build, true); err != nil {
		return nil, err
	}

	if err := tx.Save(build).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return build, nil
}

// Show implements the VersionsRepository interface.
func (r *GormRepository) Show(ctx context.Context, _ string, name string) (*model.Version, error) {
	record := &model.Version{}

	// TODO: also add modID to the query
	err := r.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrVersionNotFound
	}

	return record, err
}

// Delete implements the VersionsRepository interface.
func (r *GormRepository) Delete(ctx context.Context, _ string, name string) error {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	// TODO: also add modID to the query
	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Delete(
		&model.Version{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the VersionsRepository interface.
func (r *GormRepository) Exists(ctx context.Context, modID, id string) (bool, string, error) {
	return false, "", errors.New("not implemented")
}

// ListBuilds implements the VersionsRepository interface.
func (r *GormRepository) ListBuilds(ctx context.Context, modID, versionID, search string) ([]*model.Build, error) {
	records := make([]*model.Build, 0)
	return records, errors.New("not implemented")
}

// AttachBuild implements the VersionsRepository interface.
func (r *GormRepository) AttachBuild(ctx context.Context, modID, versionID, packID, buildID string) error {
	return errors.New("not implemented")
}

// DropBuild implements the VersionsRepository interface.
func (r *GormRepository) DropBuild(ctx context.Context, modID, versionID, packID, buildID string) error {
	return errors.New("not implemented")
}

func (r *GormRepository) validate(ctx context.Context, record *model.Version, existing bool) error {
	errs := validate.Errors{}

	if existing {
		if err := validation.Validate(
			record.ID,
			validation.Required,
			is.UUIDv4,
			validation.By(r.uniqueValueIsPresent(ctx, "id", record.ID, "")),
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
		validation.By(r.uniqueValueIsPresent(ctx, "slug", record.ID, record.ModID)),
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
		validation.By(r.uniqueValueIsPresent(ctx, "name", record.ID, record.ModID)),
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

func (r *GormRepository) uniqueValueIsPresent(ctx context.Context, key, id, _ string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		// TODO: also add modID to the query
		res := r.handle.WithContext(ctx).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Not(
			"id = ?",
			id,
		).Find(
			&model.Version{},
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
		&model.Version{},
	)
}
