package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/shared/s3client"
	"github.com/vincent-petithory/dataurl"
)

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	ID          int              `json:"-" gorm:"primary_key"`
	VersionID   int              `json:"-"`
	Version     *Pack            `json:"-"`
	Slug        string           `json:"slug" sql:"unique_index"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5"`
	Upload      *dataurl.DataURL `json:"upload,omitempty" sql:"-"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
}

// AfterFind invokes required after loading a record from the database.
func (u *VersionFile) AfterFind(db *gorm.DB) {
	u.SetURL()
}

// BeforeSave invokes required actions before persisting.
func (u *VersionFile) BeforeSave(db *gorm.DB) error {
	if u.Slug == "" {
		for i := 0; true; i++ {
			hash := md5.Sum([]byte(fmt.Sprintf("%s", time.Now().Unix())))
			u.Slug = hex.EncodeToString(hash[:])

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&VersionFile{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

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
		if config.S3.Enabled {
			_, err := s3client.New().Upload(
				u.RelativePath(),
				u.ContentType,
				u.Upload.Data,
			)

			if err != nil {
				return fmt.Errorf("Failed to upload version to S3. %s", err)
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
				return fmt.Errorf("Failed to create version directory")
			}

			errWrite := ioutil.WriteFile(
				absolutePath,
				u.Upload.Data,
				0644,
			)

			if errWrite != nil {
				return fmt.Errorf("Failed to write version at %s", absolutePath)
			}
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *VersionFile) AfterDelete(tx *gorm.DB) error {
	if config.S3.Enabled {
		_, err := s3client.New().Delete(
			u.RelativePath(),
		)

		return fmt.Errorf("Failed to remove version from S3. %s", err)
	} else {
		absolutePath, err := u.AbsolutePath()

		if err != nil {
			return err
		}

		errRemove := os.Remove(
			absolutePath,
		)

		if errRemove != nil {
			return fmt.Errorf("Failed to remove version")
		}
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
		return "", fmt.Errorf("Missing storage path for version")
	}

	return path.Join(
		config.Server.Storage,
		u.RelativePath(),
	), nil
}

// RelativePath generates the relative path to the file.
func (u *VersionFile) RelativePath() string {
	return path.Join(
		"version",
		u.Slug,
	)
}

// SetURL generates the absolute URL to access this file.
func (u *VersionFile) SetURL() {
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
				config.Server.Host,
				"storage",
				u.RelativePath(),
			},
			"/",
		)
	}
}

func isInvalidVersionFileType(mediaType string) bool {
	logrus.Debugf("Got %s version file media type", mediaType)
	return false
}
