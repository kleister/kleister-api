package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

const (
	// ClientNameMinLength is the minimum length of the client name.
	ClientNameMinLength = "3"

	// ClientNameMaxLength is the maximum length of the client name.
	ClientNameMaxLength = "255"
)

// ClientDefaultOrder is the default ordering for client listings.
func ClientDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"clients.name ASC",
	)
}

// Clients is simply a collection of client structs.
type Clients []*Client

// Client represents a client model definition.
type Client struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index"`
	Value     string    `json:"uuid" sql:"unique_index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Packs     Packs     `json:"packs" gorm:"many2many:client_packs;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Client) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		u.Slug = uuid.NewV4().String()
	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Client) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Client) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}

	if !govalidator.StringLength(u.Name, ClientNameMinLength, ClientNameMaxLength) {
		db.AddError(fmt.Errorf("Name should be longer than 3 characters"))
	}

	if u.Value == "" {
		db.AddError(fmt.Errorf("UUID is a required attribute"))
	}

	if u.Name != "" {
		count := 1

		db.Where("name = ?", u.Name).Not("id", u.ID).Find(
			&Client{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if u.Value != "" {
		count := 1

		db.Where("value = ?", u.Value).Not("id", u.ID).Find(
			&Client{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("UUID is already present"))
		}
	}
}
