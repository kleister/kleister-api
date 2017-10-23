package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Keys is simply a collection of key structs.
type Keys []*Key

// Key represents a key model definition.
type Key struct {
	ID        int64     `json:"id" gorm:"primary_key"`
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
	if !govalidator.StringLength(u.Name, "2", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 2 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Key{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if !govalidator.StringLength(u.Value, "2", "255") {
		db.AddError(fmt.Errorf("Key should be longer than 2 and shorter than 255"))
	}

	if u.Value != "" {
		notFound := db.Where(
			"value = ?",
			u.Value,
		).Not(
			"id",
			u.ID,
		).First(
			&Key{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Key is already present"))
		}
	}
}
