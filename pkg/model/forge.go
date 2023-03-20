package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Forge within Kleister.
type Forge struct {
	ID        string `gorm:"primaryKey;length:36"`
	Name      string `gorm:"unique;length:255"`
	Minecraft string `gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}

// BeforeSave provides a hook for the database layer.
func (f *Forge) BeforeSave(_ *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}

	return nil
}
