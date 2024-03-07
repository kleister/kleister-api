package versions

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
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
func (s *GormService) List(ctx context.Context, modID string) ([]*model.Version, error) {
	records := make([]*model.Version, 0)

	if err := s.query(ctx, modID).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, modID, name string) (*model.Version, error) {
	record := &model.Version{}

	err := s.query(ctx, modID).Where(
		"id = ?",
		name,
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
func (s *GormService) Create(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	version.ModID = modID

	if err := s.validate(ctx, version, false); err != nil {
		return nil, err
	}

	if err := tx.Create(version).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return version, nil
}

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	version.ModID = modID

	if err := s.validate(ctx, version, true); err != nil {
		return nil, err
	}

	if err := tx.Save(version).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return version, nil
}

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, modID string, name string) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"mod_id = ?",
		modID,
	).Where(
		"id = ?",
		name,
		name,
	).Delete(
		&model.Version{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, modID string, name string) (bool, error) {
	res := s.query(ctx, modID).Where(
		"id = ?",
		name,
		name,
	).Find(
		&model.Version{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func (s *GormService) validate(ctx context.Context, record *model.Version, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID, record.ModID)),
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

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id, modID string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.handle.WithContext(
			ctx,
		).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Where(
			"mod_id = ?",
			modID,
		)

		if id != "" {
			q = q.Not(
				"id = ?",
				id,
			)
		}

		if q.Find(
			&model.Version{},
		).RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *GormService) query(ctx context.Context, modID string) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Version{},
	).Where(
		"mod_id = ?",
		modID,
	).Preload(
		"Mod",
	)
}
