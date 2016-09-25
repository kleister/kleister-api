package model

import (
	"fmt"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/vincent-petithory/dataurl"
)

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	ID          int64            `json:"-" gorm:"primary_key"`
	VersionID   int64            `json:"-"`
	Version     *Pack            `json:"-"`
	Slug        string           `json:"slug" sql:"unique_index"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5" gorm:"column:md5"`
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
			u.Slug = RandomHash()

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
		u.MD5, u.ContentType = AttachmentHash(u.Upload)
	}

	return nil
}

// AfterSave invokes required actions after persisting.
func (u *VersionFile) AfterSave(db *gorm.DB) error {
	if u.Upload != nil {
		var (
			err  error
			dest string
		)

		if config.S3.Enabled {
			dest = u.RelativePath()
		} else {
			dest, err = u.AbsolutePath()

			if err != nil {
				return err
			}
		}

		if err := AttachmentUpload(dest, u.ContentType, u.Upload.Data); err != nil {
			return fmt.Errorf("Failed to store version. %s", err)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *VersionFile) AfterDelete(tx *gorm.DB) error {
	var (
		err  error
		dest string
	)

	if config.S3.Enabled {
		dest = u.RelativePath()
	} else {
		dest, err = u.AbsolutePath()

		if err != nil {
			return err
		}
	}

	if err := AttachmentDelete(dest); err != nil {
		return fmt.Errorf("Failed to remove version. %s", err)
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
	u.URL = AttachmentURL(
		u.RelativePath(),
	)
}

func isInvalidVersionFileType(mediaType string) bool {
	switch mediaType {
	case "application/zip":
		logrus.Debugf("Detected valid %s media type", mediaType)
		return false
	}

	logrus.Debugf("Failed to validate %s media type", mediaType)
	return true
}
