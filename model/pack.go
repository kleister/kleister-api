package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/jinzhu/gorm"
)

const (
	// PackUsernameMinLength is the minimum length of the username.
	PackUsernameMinLength = "3"

	// PackUsernameMaxLength is the maximum length of the username.
	PackUsernameMaxLength = "255"
)

// PackDefaultOrder is the default ordering for pack listings.
func PackDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"packs.name ASC",
	)
}

// Packs is simply a collection of pack structs.
type Packs []*Pack

// Pack represents a pack model definition.
type Pack struct {
	ID            uint        `json:"id" gorm:"primary_key"`
	Icon          *Attachment `json:"icon" gorm:"polymorphic:Owner"`
	Logo          *Attachment `json:"logo" gorm:"polymorphic:Owner"`
	Background    *Attachment `json:"background" gorm:"polymorphic:Owner"`
	Recommended   *Build      `json:"recommended"`
	RecommendedID uint        `json:"recommended_id" sql:"index"`
	Latest        *Build      `json:"latest"`
	LatestID      uint        `json:"latest_id" sql:"index"`
	Slug          string      `json:"slug" sql:"unique_index"`
	Name          string      `json:"name" sql:"unique_index"`
	Website       string      `json:"website"`
	Hidden        bool        `json:"hidden" sql:"default:true"`
	Private       bool        `json:"private" sql:"default:false"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Builds        Builds      `json:"builds"`
	Permissions   Permissions `json:"permissions" gorm:"many2many:permission_packs;"`
	Clients       Clients     `json:"clients" gorm:"many2many:client_packs;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Pack) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {

		u.Slug = slugify.Slugify(u.Name)
		// Fill the slug with slugified name

	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Pack) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Pack) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}
}
