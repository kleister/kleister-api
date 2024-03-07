package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PackLogo within Kleister.
type PackLogo struct {
	ID          string `gorm:"primaryKey;length:36"`
	Pack        *Pack
	PackID      string `gorm:"index;length:36"`
	Slug        string `gorm:"unique;length:255"`
	ContentType string
	MD5         string `gorm:"column:md5"`
	Path        string `gorm:"-"`
	URL         string `gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeSave defines the hook executed before every save.
func (m *PackLogo) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
