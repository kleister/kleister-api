package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// KeyNameMinLength is the minimum length of the key name.
	KeyNameMinLength = "3"

	// KeyNameMaxLength is the maximum length of the key name.
	KeyNameMaxLength = "255"

	// KeyValueMinLength is the minimum length of the key value.
	KeyValueMinLength = "3"

	// KeyValueMaxLength is the maximum length of the key value.
	KeyValueMaxLength = "255"
)

// KeyDefaultOrder is the default ordering for key listings.
func KeyDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"keys.name ASC",
	)
}

// Keys is simply a collection of key structs.
type Keys []*Key

// Key represents a key model definition.
type Key struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index"`
	Value     string    `json:"key" sql:"unique_index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeSave invokes required actions before persisting.
func (u *Key) BeforeSave(db *gorm.DB) (err error) {
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
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Key{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Key) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, KeyNameMinLength, KeyNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			KeyNameMinLength,
			KeyNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("name = ?", u.Name).Not("id", u.ID).First(&Key{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if !govalidator.StringLength(u.Value, KeyValueMinLength, KeyValueMaxLength) {
		db.AddError(fmt.Errorf(
			"Key should be longer than %s and shorter than %s",
			KeyValueMinLength,
			KeyValueMaxLength,
		))
	}

	if u.Value != "" {
		notFound := db.Where("value = ?", u.Value).Not("id", u.ID).First(&Key{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Key is already present"))
		}
	}
}
