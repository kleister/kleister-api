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

// BeforeCreate defines the hook executed before every create.
func (m *Minecraft) BeforeCreate(_ *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
