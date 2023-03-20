package model

import (
	"time"
)

// PackBack within Kleister.
type PackBack struct {
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
