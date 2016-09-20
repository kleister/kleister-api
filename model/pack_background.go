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

// PackBackgrounds is simply a collection of pack background structs.
type PackBackgrounds []*PackBackground

// PackBackground represents a pack background model definition.
type PackBackground struct {
	ID          int64            `json:"-" gorm:"primary_key"`
	PackID      int64            `json:"-" sql:"index"`
	Pack        *Pack            `json:"-"`
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
func (u *PackBackground) AfterFind(db *gorm.DB) {
	u.SetURL()
}

// BeforeSave invokes required actions before persisting.
func (u *PackBackground) BeforeSave(db *gorm.DB) error {
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
				&PackBackground{},
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
func (u *PackBackground) AfterSave(db *gorm.DB) error {
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
			return fmt.Errorf("Failed to store background. %s", err)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *PackBackground) AfterDelete(tx *gorm.DB) error {
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
		return fmt.Errorf("Failed to remove background. %s", err)
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
		u.RelativePath(),
	), nil
}

// RelativePath generates the relative path to the background.
func (u *PackBackground) RelativePath() string {
	return path.Join(
		"background",
		u.Slug,
	)
}

// SetURL generates the absolute URL to access this background.
func (u *PackBackground) SetURL() {
	u.URL = AttachmentURL(
		u.RelativePath(),
	)
}

func isInvalidPackBackgroundType(mediaType string) bool {
	logrus.Debugf("Got %s pack background media type", mediaType)
	return false
}
