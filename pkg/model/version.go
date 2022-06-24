package model

import (
	"time"
)

// Version represents a version model definition.
type Version struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	File      *VersionFile
	Mod       *Mod
	ModID     string `storm:"index" gorm:"index;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*BuildVersion
}
