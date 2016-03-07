package model

import (
	"time"

	"github.com/Machiel/slugify"
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder/store"
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
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index" validate:"gte=3,lte=255"`
	Value     string    `json:"uuid" sql:"unique_index" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeSave invokes required actions before persisting.
func (u *Client) BeforeSave() (err error) {
	if u.Slug == "" {
		u.Slug = slugify.Slugify(u.Name)
	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Client) Defaults() {

}

// Validate does some validation to be able to store the record.
func (u *Client) Validate(store store.Store) error {
	err := validate.Struct(u)

	if u.Name != "" {
		count := 1

		store.Where("clients.name = ?", u.Name).Find(
			&Client{},
		).Count(
			&count,
		)

		if count > 0 {
			// Invalid because it's bigger than 1
		}
	}

	return err
}
