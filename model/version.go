package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// VersionNameMinLength is the minimum length of the name.
	VersionNameMinLength = "3"

	// VersionNameMaxLength is the maximum length of the name.
	VersionNameMaxLength = "255"
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
	ID        int         `json:"id" gorm:"primary_key"`
	File      *Attachment `json:"file" gorm:"polymorphic:Owner"`
	Mod       *Mod        `json:"mod"`
	ModID     int         `json:"mod_id" sql:"index"`
	Slug      string      `json:"slug"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Builds    Builds      `json:"builds"`
}

// BeforeSave invokes required actions before persisting.
func (u *Version) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = slugify.Slugify(u.Name)
			} else {
				u.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", u.Name, i),
				)
			}

			notFound := db.Where(
				"mod_id = ?",
				u.ModID,
			).Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Version{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Version) Validate(db *gorm.DB) {
	if u.ModID == 0 {
		db.AddError(fmt.Errorf("A mod reference is required"))
	} else {
		res := db.Where(
			"id = ?",
			u.ModID,
		).First(
			&Mod{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced mod does not exist"))
		}
	}

	if !govalidator.StringLength(u.Name, VersionNameMinLength, VersionNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			VersionNameMinLength,
			VersionNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("mod_id = ?", u.ModID).Where("name = ?", u.Name).Not("id", u.ID).First(&Version{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
