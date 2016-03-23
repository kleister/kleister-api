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

// PackIcons is simply a collection of pack icon structs.
type PackIcons []*PackIcon

// PackIcon represents a pack icon model definition.
type PackIcon struct {
	PackID      int              `json:"-" gorm:"primary_key"`
	Pack        *Pack            `json:"-"`
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
func (u *PackIcon) BeforeSave(db *gorm.DB) error {
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
func (u *PackIcon) AfterSave(db *gorm.DB) error {
	if u.Upload != nil {
		if u.Path == "" {
			return fmt.Errorf("Missing storage path for logo")
		}

		file, errCreate := os.Create(
			u.Path,
		)

		if errCreate != nil {
			return fmt.Errorf("Failed to open icon at %s", u.Path)
		}

		_, errWrite := u.Upload.WriteTo(
			file,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write icon at %s", u.Path)
		}
	}

	return nil
}
