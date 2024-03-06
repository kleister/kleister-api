package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Minecraft within Kleister.
type Minecraft struct {
	ID        string `storm:"id" gorm:"primaryKey;length:20"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	Type      string `storm:"index" gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeSave defines the hook executed before every save.
func (m *Minecraft) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
