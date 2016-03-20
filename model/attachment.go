package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vincent-petithory/dataurl"
)

// Attachment represents any uploadable asset.
type Attachment struct {
	ID        int       `json:"id" gorm:"primary_key"`
	OwnerID   int       `json:"-"`
	OwnerType string    `json:"-"`
	URL       string    `json:"url" gorm:"-"`
	MD5       string    `json:"md5"`
	Content   string    `json:"-" gorm:"type:longtext"`
	Upload    string    `json:"upload,omitempty" gorm:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *Attachment) BeforeSave(db *gorm.DB) (err error) {
	if u.Upload != "" {
		decoded, err := dataurl.DecodeString(
			u.Upload,
		)

		if err != nil {
			return fmt.Errorf("Failed to decode upload")
		}

		check := md5.Sum(
			decoded.Data,
		)

		hash := hex.EncodeToString(
			check[:],
		)

		u.MD5 = hash
		u.Content = u.Upload
	}

	return nil
}
