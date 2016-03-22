package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vincent-petithory/dataurl"
)

// VersionFiles is simply a collection of version file structs.
type VersionFiles []*VersionFile

// VersionFile represents a version file model definition.
type VersionFile struct {
	VersionID int       `json:"-" gorm:"primary_key"`
	Version   *Pack     `json:"-"`
	URL       string    `json:"url" sql:"-"`
	MD5       string    `json:"md5"`
	Content   string    `json:"-" gorm:"type:longtext"`
	Upload    string    `json:"upload,omitempty" sql:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *VersionFile) BeforeSave(db *gorm.DB) (err error) {
	if u.Upload != "" {
		decoded, err := dataurl.DecodeString(
			u.Upload,
		)

		if err != nil {
			return fmt.Errorf("Failed to decode file")
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
