package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

const (
	// KeyNameMinLength is the minimum length of the key name.
	KeyNameMinLength = "3"

	// KeyNameMaxLength is the maximum length of the key name.
	KeyNameMaxLength = "255"
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
		u.Slug = uuid.NewV4().String()
	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Key) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Key) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}

	if !govalidator.StringLength(u.Name, KeyNameMinLength, KeyNameMaxLength) {
		db.AddError(fmt.Errorf("Name should be longer than 3 characters"))
	}

	if u.Value == "" {
		db.AddError(fmt.Errorf("Key is a required attribute"))
	}

	if u.Name != "" {
		count := 1

		db.Where("name = ?", u.Name).Not("id", u.ID).Find(
			&Key{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if u.Value != "" {
		count := 1

		db.Where("value = ?", u.Name).Not("id", u.ID).Find(
			&Key{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("Key is already present"))
		}
	}
}
