package model

import (
	"time"

	"github.com/vincent-petithory/dataurl"
)

// VersionFile represents a version file model definition.
type VersionFile struct {
	ID          string `storm:"id" gorm:"primaryKey;length:36"`
	Version     *Version
	VersionID   string `storm:"index" gorm:"index;length:36"`
	Slug        string `storm:"unique" gorm:"unique;length:255"`
	ContentType string
	MD5         string           `gorm:"column:md5"`
	Path        string           `json:"-" gorm:"-"`
	URL         string           `json:"-" gorm:"-"`
	Upload      *dataurl.DataURL `json:"-" gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
