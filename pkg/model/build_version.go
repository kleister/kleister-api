package model

import (
	"time"
)

// BuildVersion represents a build version model definition.
type BuildVersion struct {
	BuildID   string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Build     *Build
	VersionID string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Version   *Version
	CreatedAt time.Time
	UpdatedAt time.Time
}
