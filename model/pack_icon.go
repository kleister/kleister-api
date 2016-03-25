package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/vincent-petithory/dataurl"
)

// PackIcons is simply a collection of pack icon structs.
type PackIcons []*PackIcon

// PackIcon represents a pack icon model definition.
type PackIcon struct {
	ID          int              `json:"id" gorm:"primary_key"`
	PackID      int              `json:"-"`
	Pack        *Pack            `json:"-"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5"`
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
		absolutePath, err := u.AbsolutePath()

		if err != nil {
			return fmt.Errorf("Missing storage path for icon")
		}

		errDir := os.MkdirAll(
			filepath.Dir(
				absolutePath,
			),
			os.ModePerm,
		)

		if errDir != nil {
			return fmt.Errorf("Failed to create icon directory")
		}

		file, errCreate := os.Create(
			absolutePath,
		)

		if errCreate != nil {
			return fmt.Errorf("Failed to open icon at %s", absolutePath)
		}

		_, errWrite := u.Upload.WriteTo(
			file,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write icon at %s", absolutePath)
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *PackIcon) Validate(db *gorm.DB) {
	if u.Upload == nil {
		db.AddError(fmt.Errorf("An icon is required"))
	}

	if isInvalidPackIconType(u.Upload.MediaType.String()) {
		db.AddError(fmt.Errorf("Invalid icon media type"))
	}
}

// Path generates the absolute path to the icon.
func (u *PackIcon) AbsolutePath() (string, error) {
	if u.Path == "" {
		return "", fmt.Errorf("Missing storage path for icon")
	}

	return path.Join(
		u.Path,
		"icon",
		u.MD5,
	), nil
}

func isInvalidPackIconType(mediaType string) bool {
	logrus.Debugf("Got %s pack icon media type", mediaType)
	return false
}
