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

// PackLogos is simply a collection of pack logo structs.
type PackLogos []*PackLogo

// PackLogo represents a pack logo model definition.
type PackLogo struct {
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
func (u *PackLogo) AfterFind(db *gorm.DB) {
	u.SetURL()
}

// BeforeSave invokes required actions before persisting.
func (u *PackLogo) BeforeSave(db *gorm.DB) error {
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
				&PackLogo{},
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
func (u *PackLogo) AfterSave(db *gorm.DB) error {
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
			return fmt.Errorf("Failed to store logo. %s", err)
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *PackLogo) AfterDelete(tx *gorm.DB) error {
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
		return fmt.Errorf("Failed to remove logo. %s", err)
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
		u.Slug,
	)
}

// SetURL generates the absolute URL to access this logo.
func (u *PackLogo) SetURL() {
	u.URL = AttachmentURL(
		u.RelativePath(),
	)
}

func isInvalidPackLogoType(mediaType string) bool {
	logrus.Debugf("Got %s pack logo media type", mediaType)
	return false
}
