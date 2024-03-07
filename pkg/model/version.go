package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Version within Kleister.
type Version struct {
	ID        string `gorm:"primaryKey;length:36"`
	Mod       *Mod
	ModID     string `gorm:"index;length:36"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	File      *VersionFile
	Public    bool `default:"true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*BuildVersion
}

// BeforeSave defines the hook executed before every save.
func (m *Version) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
