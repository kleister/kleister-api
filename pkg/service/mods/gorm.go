package mods

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(handle *gorm.DB) *GormService {
	return &GormService{
		handle: handle,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context) ([]*model.Mod, error) {
	records := make([]*model.Mod, 0)

	if err := s.query(ctx).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, name string) (*model.Mod, error) {
	record := &model.Mod{}

	err := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrNotFound
	}

	return record, err
}

// Create implements the Service interface for database persistence.
func (s *GormService) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if mod.Slug == "" {
		mod.Slug = store.Slugify(
			tx.Model(&model.Mod{}),
			mod.Name,
			"",
		)
	}

	if err := s.validate(ctx, mod, false); err != nil {
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

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if mod.Slug == "" {
		mod.Slug = store.Slugify(
			tx.Model(&model.Mod{}),
			mod.Name,
			mod.ID,
		)
	}

	if err := s.validate(ctx, mod, true); err != nil {
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

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, name string) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Delete(
		&model.Mod{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, name string) (bool, error) {
	res := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Find(
		&model.Mod{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func (s *GormService) validate(ctx context.Context, record *model.Mod, _ bool) error {
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

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		res := s.handle.WithContext(
			ctx,
		).Where(
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

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Mod{},
	).Preload(
		"Users",
	).Preload(
		"Users.User",
	)
}
