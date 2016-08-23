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
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/shared/s3client"
	"github.com/vincent-petithory/dataurl"
)

// PackLogos is simply a collection of pack logo structs.
type PackLogos []*PackLogo

// PackLogo represents a pack logo model definition.
type PackLogo struct {
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
func (u *PackLogo) BeforeSave(db *gorm.DB) error {
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
func (u *PackLogo) AfterSave(db *gorm.DB) error {
	if u.Upload != nil {
		if config.S3.Enabled {
			_, err := s3client.New().Upload(
				u.RelativePath(),
				u.ContentType,
				u.Upload.Data,
			)

			if err != nil {
				return fmt.Errorf("Failed to upload logo to S3. %s", err)
			}
		} else {
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
				return fmt.Errorf("Failed to create logo directory")
			}

			errWrite := ioutil.WriteFile(
				absolutePath,
				u.Upload.Data,
				0644,
			)

			if errWrite != nil {
				return fmt.Errorf("Failed to write logo at %s", absolutePath)
			}
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *PackLogo) AfterDelete(tx *gorm.DB) error {
	if config.S3.Enabled {
		_, err := s3client.New().Delete(
			u.RelativePath(),
		)

		return fmt.Errorf("Failed to remove logo from S3. %s", err)
	} else {
		absolutePath, err := u.AbsolutePath()

		if err != nil {
			return err
		}

		errRemove := os.Remove(
			absolutePath,
		)

		if errRemove != nil {
			return fmt.Errorf("Failed to remove logo")
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *PackLogo) Validate(db *gorm.DB) {
	if u.Upload != nil {
		if isInvalidPackLogoType(u.Upload.MediaType.String()) {
			db.AddError(fmt.Errorf("Invalid logo media type"))
		}
	}
}

// AbsolutePath generates the absolute path to the logo.
func (u *PackLogo) AbsolutePath() (string, error) {
	if config.Server.Storage == "" {
		return "", fmt.Errorf("Missing storage path for logo")
	}

	return path.Join(
		config.Server.Storage,
		u.RelativePath(),
	), nil
}

// RelativePath generates the relative path to the logo.
func (u *PackLogo) RelativePath() string {
	return path.Join(
		"logo",
		strconv.Itoa(u.PackID),
	)
}

// SetURL generates the absolute URL to access this logo.
func (u *PackLogo) SetURL(location string) {
	if config.S3.Enabled {
		if config.S3.Endpoint == "" {
			u.URL = fmt.Sprintf(
				"https://s3-%s.amazonaws.com/%s/%s",
				config.S3.Region,
				config.S3.Bucket,
				u.RelativePath(),
			)
		} else {
			u.URL = fmt.Sprintf(
				"%s/%s/%s",
				config.S3.Endpoint,
				config.S3.Bucket,
				u.RelativePath(),
			)
		}
	} else {
		u.URL = strings.Join(
			[]string{
				location,
				"storage",
				u.RelativePath(),
			},
			"/",
		)
	}
}

func isInvalidPackLogoType(mediaType string) bool {
	logrus.Debugf("Got %s pack logo media type", mediaType)
	return false
}
