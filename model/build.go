package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/jinzhu/gorm"
)

const (
	// BuildUsernameMinLength is the minimum length of the username.
	BuildUsernameMinLength = "3"

	// BuildUsernameMaxLength is the maximum length of the username.
	BuildUsernameMaxLength = "255"
)

// BuildDefaultOrder is the default ordering for pack listings.
func BuildDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"builds.name ASC",
	)
}

// Builds is simply a collection of build structs.
type Builds []*Build

// Build represents a build model definition.
type Build struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	Pack        *Pack      `json:"pack"`
	PackID      int64      `json:"pack_id" sql:"index"`
	Minecraft   *Minecraft `json:"minecraft"`
	MinecraftID int64      `json:"minecraft_id" sql:"index"`
	Forge       *Forge     `json:"forge"`
	ForgeID     int64      `json:"forge_id" sql:"index"`
	Slug        string     `json:"slug" sql:"unique_index"`
	Name        string     `json:"name" sql:"unique_index"`
	MinJava     string     `json:"min_java"`
	MinMemory   string     `json:"min_memory"`
	Published   bool       `json:"published" sql:"default:false"`
	Private     bool       `json:"private" sql:"default:false"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Mods        *Mods      `json:"mods" gorm:"many2many:build_mods;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Build) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {

		u.Slug = slugify.Slugify(u.Name)
		// Fill the slug with slugified name

	}

	return nil
}

// Defaults prefills the struct with some default values.
func (u *Build) Defaults() {
	// Currently no default values required.
}

// Validate does some validation to be able to store the record.
func (u *Build) Validate(db *gorm.DB) {
	if u.Name == "" {
		db.AddError(fmt.Errorf("Name is a required attribute"))
	}
}
