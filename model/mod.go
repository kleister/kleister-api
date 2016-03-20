package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// ModNameMinLength is the minimum length of the name.
	ModNameMinLength = "3"

	// ModNameMaxLength is the maximum length of the name.
	ModNameMaxLength = "255"
)

// ModDefaultOrder is the default ordering for mod listings.
func ModDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"mods.name ASC",
	)
}

// Mods is simply a collection of mod structs.
type Mods []*Mod

// Mod represents a mod model definition.
type Mod struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Slug        string    `json:"slug" sql:"unique_index"`
	Name        string    `json:"name" sql:"unique_index"`
	Description string    `json:"description" sql:"type:text"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Versions    Versions  `json:"versions"`
	Users       Users     `json:"users" gorm:"many2many:user_mods;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Mod) BeforeSave(db *gorm.DB) (err error) {
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
				&Mod{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Mod) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, ModNameMinLength, ModNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			ModNameMinLength,
			ModNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("name = ?", u.Name).Not("id", u.ID).First(&Mod{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
