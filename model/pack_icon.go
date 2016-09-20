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

// PackIcons is simply a collection of pack icon structs.
type PackIcons []*PackIcon

// PackIcon represents a pack icon model definition.
type PackIcon struct {
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
func (u *PackIcon) AfterFind(db *gorm.DB) {
	u.SetURL()
}

// BeforeSave invokes required actions before persisting.
func (u *PackIcon) BeforeSave(db *gorm.DB) error {
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
				&PackIcon{},
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
func (u *PackIcon) AfterSave(db *gorm.DB) error {
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
			return fmt.Errorf("Failed to store icon. %s", err)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *PackIcon) AfterDelete(tx *gorm.DB) error {
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
		return fmt.Errorf("Failed to remove icon. %s", err)
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *PackIcon) Validate(db *gorm.DB) {
	if u.Upload != nil {
		if isInvalidPackIconType(u.Upload.MediaType.String()) {
			db.AddError(fmt.Errorf("Invalid icon media type"))
		}
	}
}

// AbsolutePath generates the absolute path to the icon.
func (u *PackIcon) AbsolutePath() (string, error) {
	if config.Server.Storage == "" {
		return "", fmt.Errorf("Missing storage path for icon")
	}

	return path.Join(
		config.Server.Storage,
		u.RelativePath(),
	), nil
}

// RelativePath generates the relative path to the icon.
func (u *PackIcon) RelativePath() string {
	return path.Join(
		"icon",
		u.Slug,
	)
}

// SetURL generates the absolute URL to access this icon.
func (u *PackIcon) SetURL() {
	u.URL = AttachmentURL(
		u.RelativePath(),
	)
}

func isInvalidPackIconType(mediaType string) bool {
	logrus.Debugf("Got %s pack icon media type", mediaType)
	return false
}
