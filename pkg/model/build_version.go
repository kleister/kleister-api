package model

import (
	"time"
)

// BuildVersion within Kleister.
type BuildVersion struct {
	BuildID   string `gorm:"index:idx_id,unique;length:20"`
	Build     *Build
	VersionID string `gorm:"index:idx_id,unique;length:20"`
	Version   *Version
	CreatedAt time.Time
	UpdatedAt time.Time
}
