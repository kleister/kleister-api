package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/config"
	"github.com/vincent-petithory/dataurl"
)

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	ID          int              `json:"-" gorm:"primary_key"`
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
			return err
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

		errWrite := ioutil.WriteFile(
			absolutePath,
			u.Upload.Data,
			0644,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write file at %s", absolutePath)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *VersionFile) AfterDelete(tx *gorm.DB) error {
	absolutePath, err := u.AbsolutePath()

	if err != nil {
		return err
	}

	errRemove := os.Remove(
		absolutePath,
	)

	if errRemove != nil {
		return fmt.Errorf("Failed to remove file")
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *VersionFile) Validate(db *gorm.DB) {
	if u.Upload != nil {
		if isInvalidVersionFileType(u.Upload.MediaType.String()) {
			db.AddError(fmt.Errorf("Invalid file media type"))
		}
	}
}

// AbsolutePath generates the absolute path to the file.
func (u *VersionFile) AbsolutePath() (string, error) {
	if config.Server.Storage == "" {
		return "", fmt.Errorf("Missing storage path for file")
	}

	return path.Join(
		config.Server.Storage,
		"version",
		strconv.Itoa(u.VersionID),
	), nil
}

func isInvalidVersionFileType(mediaType string) bool {
	logrus.Debugf("Got %s version file media type", mediaType)
	return false
}
