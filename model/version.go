package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/jinzhu/gorm"
)

const (
	// VersionUsernameMinLength is the minimum length of the username.
	VersionUsernameMinLength = "3"

	// VersionUsernameMaxLength is the maximum length of the username.
	VersionUsernameMaxLength = "255"
)

// VersionDefaultOrder is the default ordering for pack listings.
func VersionDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"versions.name ASC",
	)
}

// Versions is simply a collection of version structs.
type Versions []*Version

// Version represents a version model definition.
type Version struct {
	ID        uint        `json:"id" gorm:"primary_key"`
	File      *Attachment `json:"file" gorm:"polymorphic:Owner"`
	Mod       *Mod        `json:"mod"`
	ModID     uint        `json:"mod_id" sql:"index"`
	Slug      string      `json:"slug" sql:"unique_index"`
	Name      string      `json:"name" sql:"unique_index"`
	MD5       string      `json:"md5"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Builds    Builds      `json:"builds"`
}

// BeforeSave invokes required actions before persisting.
func (u *Version) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {

		u.Slug = slugify.Slugify(u.Name)
		// Fill the slug with slugified name

	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Version) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Version) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}
}
