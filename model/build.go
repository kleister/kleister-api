package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// BuildNameMinLength is the minimum length of the name.
	BuildNameMinLength = "2"

	// BuildNameMaxLength is the maximum length of the name.
	BuildNameMaxLength = "255"
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
	ID          uint       `json:"id" gorm:"primary_key"`
	Pack        *Pack      `json:"pack"`
	PackID      uint       `json:"pack_id" sql:"index"`
	Minecraft   *Minecraft `json:"minecraft"`
	MinecraftID uint       `json:"minecraft_id" sql:"index"`
	Forge       *Forge     `json:"forge"`
	ForgeID     uint       `json:"forge_id" sql:"index"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	MinJava     string     `json:"min_java"`
	MinMemory   string     `json:"min_memory"`
	Published   bool       `json:"published" sql:"default:false"`
	Private     bool       `json:"private" sql:"default:false"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Versions    Versions   `json:"versions" gorm:"many2many:build_versions;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Build) BeforeSave(db *gorm.DB) (err error) {
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
				"pack_id = ?",
				u.PackID,
			).Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Build{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Build) Validate(db *gorm.DB) {
	if u.PackID == 0 {
		db.AddError(fmt.Errorf("A pack reference is required"))
	} else {
		res := db.Where(
			"id = ?",
			u.PackID,
		).First(
			&Pack{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced pack does not exist"))
		}
	}

	if u.MinecraftID > 0 {
		res := db.Where(
			"id = ?",
			u.MinecraftID,
		).First(
			&Minecraft{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced minecraft does not exist"))
		}
	}

	if u.ForgeID > 0 {
		res := db.Where(
			"id = ?",
			u.ForgeID,
		).First(
			&Forge{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced forge does not exist"))
		}
	}

	if !govalidator.StringLength(u.Name, BuildNameMinLength, BuildNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			BuildNameMinLength,
			BuildNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("pack_id = ?", u.PackID).Where("name = ?", u.Name).Not("id", u.ID).First(&Build{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
