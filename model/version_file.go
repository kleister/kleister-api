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

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	ID          int              `json:"id" gorm:"primary_key"`
	VersionID   int              `json:"-"`
	Version     *Pack            `json:"-"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5"`
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
		absolutePath, err := u.AbsolutePath()

		if err != nil {
			return fmt.Errorf("Missing storage path for logo")
		}

		errDir := os.MkdirAll(
			filepath.Dir(
				absolutePath,
			),
			os.ModePerm,
		)

		if errDir != nil {
			return fmt.Errorf("Failed to create file directory")
		}

		file, errCreate := os.Create(
			absolutePath,
		)

		if errCreate != nil {
			return fmt.Errorf("Failed to open version at %s", absolutePath)
		}

		_, errWrite := u.Upload.WriteTo(
			file,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write version at %s", absolutePath)
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *VersionFile) Validate(db *gorm.DB) {
	if u.Upload == nil {
		db.AddError(fmt.Errorf("A file is required"))
	}

	if isInvalidVersionFileType(u.Upload.MediaType.String()) {
		db.AddError(fmt.Errorf("Invalid file media type"))
	}
}

// Path generates the absolute path to the file.
func (u *VersionFile) AbsolutePath() (string, error) {
	if u.Path == "" {
		return "", fmt.Errorf("Missing storage path for file")
	}

	return path.Join(
		u.Path,
		"version",
		u.MD5,
	), nil
}

func isInvalidVersionFileType(mediaType string) bool {
	logrus.Debugf("Got %s version file media type", mediaType)
	return false
}
