package model

import (
	"time"

	"github.com/jinzhu/gorm"
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
	Name      string    `json:"name" sql:"unique_index"`
	Value     string    `json:"uuid" sql:"unique_index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
