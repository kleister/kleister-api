package builds

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
func (s *GormService) List(ctx context.Context, packID string) ([]*model.Build, error) {
	records := make([]*model.Build, 0)

	if err := s.query(ctx, packID).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, packID, name string) (*model.Build, error) {
	record := &model.Build{}

	err := s.query(ctx, packID).Where(
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
func (s *GormService) Create(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	build.PackID = packID

	if err := s.validate(ctx, build, false); err != nil {
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

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	build.PackID = packID

	if err := s.validate(ctx, build, true); err != nil {
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

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, packID string, name string) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"pack_id = ?",
		packID,
	).Where(
		"id = ?",
		name,
		name,
	).Delete(
		&model.Build{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, packID string, name string) (bool, error) {
	res := s.query(ctx, packID).Where(
		"id = ?",
		name,
		name,
	).Find(
		&model.Build{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

// Column implements the Service interface for database persistence.
func (s *GormService) Column(ctx context.Context, packID, id, col string, val any) error {
	err := s.handle.WithContext(
		ctx,
	).Table(
		"builds",
	).Where(
		"pack_id = ?",
		packID,
	).Where(
		"id = ?",
		id,
	).Update(
		col,
		val,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	return err
}

func (s *GormService) validate(ctx context.Context, record *model.Build, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID, record.PackID)),
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

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id, packID string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.handle.WithContext(
			ctx,
		).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Where(
			"pack_id = ?",
			packID,
		)

		if id != "" {
			q = q.Not(
				"id = ?",
				id,
			)
		}

		if q.Find(
			&model.Build{},
		).RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *GormService) query(ctx context.Context, packID string) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Build{},
	).Where(
		"pack_id = ?",
		packID,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).Preload(
		"Neoforge",
	).Preload(
		"Quilt",
	).Preload(
		"Fabric",
	)
}
