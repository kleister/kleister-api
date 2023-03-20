package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Minecraft within Kleister.
type Minecraft struct {
	ID        string `gorm:"primaryKey;length:36"`
	Name      string `gorm:"unique;length:255"`
	Type      string `gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}

// BeforeSave provides a hook for the database layer.
func (m *Minecraft) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
