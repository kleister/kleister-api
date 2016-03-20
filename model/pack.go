package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// PackNameMinLength is the minimum length of the name.
	PackNameMinLength = "3"

	// PackNameMaxLength is the maximum length of the name.
	PackNameMaxLength = "255"
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
	ID            int         `json:"id" gorm:"primary_key"`
	Icon          *Attachment `json:"icon" gorm:"polymorphic:Owner"`
	Logo          *Attachment `json:"logo" gorm:"polymorphic:Owner"`
	Background    *Attachment `json:"background" gorm:"polymorphic:Owner"`
	Recommended   *Build      `json:"recommended"`
	RecommendedID int         `json:"recommended_id" sql:"index"`
	Latest        *Build      `json:"latest"`
	LatestID      int         `json:"latest_id" sql:"index"`
	Slug          string      `json:"slug" sql:"unique_index"`
	Name          string      `json:"name" sql:"unique_index"`
	Website       string      `json:"website"`
	Published     bool        `json:"published" sql:"default:false"`
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
func (u *Pack) Validate(db *gorm.DB) {
	if u.RecommendedID > 0 {
		res := db.Where(
			"pack_id = ?",
			u.ID,
		).Where(
			"id = ?",
			u.RecommendedID,
		).First(
			&Build{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced recommended build does not exist"))
		}
	}

	if u.LatestID > 0 {
		res := db.Where(
			"pack_id = ?",
			u.ID,
		).Where(
			"id = ?",
			u.LatestID,
		).First(
			&Build{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced latest build does not exist"))
		}
	}

	if !govalidator.StringLength(u.Name, PackNameMinLength, PackNameMaxLength) {
		db.AddError(fmt.Errorf(
			"Name should be longer than %s and shorter than %s",
			PackNameMinLength,
			PackNameMaxLength,
		))
	}

	if u.Name != "" {
		notFound := db.Where("name = ?", u.Name).Not("id", u.ID).First(&Pack{}).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
