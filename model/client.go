package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// ClientNameMinLength is the minimum length of the client name.
	ClientNameMinLength = "3"

	// ClientNameMaxLength is the maximum length of the client name.
	ClientNameMaxLength = "255"

	// ClientValueMinLength is the minimum length of the client value.
	ClientValueMinLength = "3"

	// ClientValueMaxLength is the maximum length of the client value.
	ClientValueMaxLength = "255"
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
	ID        int       `json:"id" gorm:"primary_key"`
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
				&Client{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Client) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, ClientNameMinLength, ClientNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			ClientNameMinLength,
			ClientNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("name = ?", u.Name).Not("id", u.ID).First(&Client{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if !govalidator.StringLength(u.Value, ClientValueMinLength, ClientValueMaxLength) {
		db.AddError(fmt.Errorf(
			"UUID should be longer than %s and shorter than %s",
			ClientValueMinLength,
			ClientValueMaxLength,
		))
	}

	if u.Value != "" {
		notFound := db.Where("value = ?", u.Value).Not("id", u.ID).First(&Client{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("UUID is already present"))
		}
	}
}
