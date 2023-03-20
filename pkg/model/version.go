package model

import (
	"time"
)

// Version within Kleister.
type Version struct {
	ID        string `gorm:"primaryKey;length:36"`
	Mod       *Mod
	ModID     string `gorm:"index;length:36"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	File      *VersionFile
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*BuildVersion
}
