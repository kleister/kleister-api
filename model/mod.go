package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/jinzhu/gorm"
)

const (
	// ModUsernameMinLength is the minimum length of the username.
	ModUsernameMinLength = "3"

	// ModUsernameMaxLength is the maximum length of the username.
	ModUsernameMaxLength = "255"
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
	ID          int64     `json:"id" gorm:"primary_key"`
	Slug        string    `json:"slug" sql:"unique_index"`
	Name        string    `json:"name" sql:"unique_index"`
	Description string    `json:"description" sql:"type:text"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Versions    *Versions `json:"versions"`
}

// BeforeSave invokes required actions before persisting.
func (u *Mod) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {

		u.Slug = slugify.Slugify(u.Name)
		// Fill the slug with slugified name

	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Mod) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Mod) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}
}
