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

// PackBackgrounds is simply a collection of pack background structs.
type PackBackgrounds []*PackBackground

// PackBackground represents a pack background model definition.
type PackBackground struct {
	ID          int              `json:"-" gorm:"primary_key"`
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
func (u *PackBackground) BeforeSave(db *gorm.DB) error {
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
func (u *PackBackground) AfterSave(db *gorm.DB) error {
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
			return fmt.Errorf("Failed to create background directory")
		}

		errWrite := ioutil.WriteFile(
			absolutePath,
			u.Upload.Data,
			0644,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write background at %s", absolutePath)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *PackBackground) AfterDelete(tx *gorm.DB) error {
	absolutePath, err := u.AbsolutePath()

	if err != nil {
		return err
	}

	errRemove := os.Remove(
		absolutePath,
	)

	if errRemove != nil {
		return fmt.Errorf("Failed to remove background")
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *PackBackground) Validate(db *gorm.DB) {
	if u.Upload != nil {
		if isInvalidPackBackgroundType(u.Upload.MediaType.String()) {
			db.AddError(fmt.Errorf("Invalid background media type"))
		}
	}
}

// AbsolutePath generates the absolute path to the background.
func (u *PackBackground) AbsolutePath() (string, error) {
	if config.Server.Storage == "" {
		return "", fmt.Errorf("Missing storage path for background")
	}

	return path.Join(
		config.Server.Storage,
		"background",
		strconv.Itoa(u.PackID),
	), nil
}

func isInvalidPackBackgroundType(mediaType string) bool {
	logrus.Debugf("Got %s pack background media type", mediaType)
	return false
}
