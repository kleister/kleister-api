package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"gopkg.in/guregu/null.v3"
)

// Builds is simply a collection of build structs.
type Builds []*Build

// Build represents a build model definition.
type Build struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	Pack        *Pack      `json:"pack,omitempty"`
	PackID      int64      `json:"pack_id" sql:"index"`
	Minecraft   *Minecraft `json:"minecraft,omitempty"`
	MinecraftID null.Int   `json:"minecraft_id" sql:"index"`
	Forge       *Forge     `json:"forge,omitempty"`
	ForgeID     null.Int   `json:"forge_id" sql:"index"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	MinJava     string     `json:"min_java"`
	MinMemory   string     `json:"min_memory"`
	Published   bool       `json:"published" sql:"default:false"`
	Hidden      bool       `json:"hidden" sql:"-"`
	Private     bool       `json:"private" sql:"default:false"`
	Public      bool       `json:"public" sql:"-"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Versions    Versions   `json:"versions,omitempty" gorm:"many2many:build_versions"`
}

// AfterFind invokes required after loading a record from the database.
func (u *Build) AfterFind(db *gorm.DB) {
	u.Hidden = !u.Published
	u.Public = !u.Private
}

// BeforeSave invokes required actions before persisting.
func (u *Build) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = u.Name
			} else {
				u.Slug = fmt.Sprintf("%s-%d", u.Name, i)
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

// BeforeDelete invokes required actions before deletion.
func (u *Build) BeforeDelete(tx *gorm.DB) error {
	return tx.Model(u).Association("Versions").Clear().Error
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

	if u.MinecraftID.Int64 > 0 {
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

	if u.ForgeID.Int64 > 0 {
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

	if !govalidator.StringLength(u.Name, "2", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 2 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"pack_id = ?",
			u.PackID,
		).Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Build{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
