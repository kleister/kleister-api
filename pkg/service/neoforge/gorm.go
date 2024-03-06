package neoforge

import (
	"context"

	neoforgeClient "github.com/kleister/kleister-api/pkg/internal/neoforge"
	"github.com/kleister/kleister-api/pkg/model"
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

// Search implements the Store interface for database persistence.
func (s *GormService) Search(ctx context.Context, search string) ([]*model.Neoforge, error) {
	records := make([]*model.Neoforge, 0)
	q := s.query(ctx)

	if search != "" {
		q = q.Or(
			"name LIKE ?",
			"%"+search+"%",
		)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Sync implements the Store interface for database persistence.
func (s *GormService) Sync(ctx context.Context, versions neoforgeClient.Versions) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	for _, v := range versions {
		if err := tx.Where(
			&model.Neoforge{
				Name: v.Value,
			},
		).FirstOrCreate(
			&model.Neoforge{},
		).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Neoforge{},
	)
}
