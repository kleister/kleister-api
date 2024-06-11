package model

import (
	"time"
)

// BuildVersion within Kleister.
type BuildVersion struct {
	BuildID   string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Build     *Build
	VersionID string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Version   *Version
	CreatedAt time.Time
	UpdatedAt time.Time
}
