package model

import (
	"time"
)

// VersionFile within Kleister.
type VersionFile struct {
	ID          string `gorm:"primaryKey;length:36"`
	Version     *Version
	VersionID   string `gorm:"index;length:36"`
	Slug        string `gorm:"unique;length:255"`
	ContentType string
	MD5         string `gorm:"column:md5"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
