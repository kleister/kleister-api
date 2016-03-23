package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vincent-petithory/dataurl"
)

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	VersionID   int              `json:"-" gorm:"primary_key"`
	Version     *Pack            `json:"-"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5"`
	Content     string           `json:"-" gorm:"type:longtext"`
	Upload      *dataurl.DataURL `json:"upload,omitempty" sql:"-"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *VersionFile) BeforeSave(db *gorm.DB) error {
	if u.Upload != nil {
		check := md5.Sum(
			u.Upload.Data,
		)

		hash := hex.EncodeToString(
			check[:],
		)

		u.MD5 = hash
		u.ContentType = u.Upload.MediaType.String()
	}

	return nil
}

// AfterSave invokes required actions after persisting.
func (u *VersionFile) AfterSave(db *gorm.DB) error {
	if u.Upload != nil {
		if u.Path == "" {
			return fmt.Errorf("Missing storage path for logo")
		}

		file, errCreate := os.Create(
			u.Path,
		)

		if errCreate != nil {
			return fmt.Errorf("Failed to open version at %s", u.Path)
		}

		_, errWrite := u.Upload.WriteTo(
			file,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write version at %s", u.Path)
		}
	}

	return nil
}
