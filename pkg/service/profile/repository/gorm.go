package repository

import (
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

// GormRepository implements the ProfileRepository interface.
type GormRepository struct {
	handle *gorm.DB
}
