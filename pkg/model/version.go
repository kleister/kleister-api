package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Versions is simply a collection of version structs.
type Versions []*Version

// Version represents a version model definition.
type Version struct {
	ID        int64        `json:"id" gorm:"primary_key"`
	File      *VersionFile `json:"file,omitempty"`
	Mod       *Mod         `json:"mod,omitempty"`
	ModID     int64        `json:"mod_id" sql:"index"`
	Slug      string       `json:"slug"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Builds    Builds       `json:"builds,omitempty" gorm:"many2many:build_versions"`
}

// BeforeSave invokes required actions before persisting.
func (u *Version) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = u.Name
			} else {
				u.Slug = fmt.Sprintf("%s-%d", u.Name, i)
			}

			notFound := db.Where(
				"mod_id = ?",
				u.ModID,
			).Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Version{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// BeforeDelete invokes required actions before deletion.
func (u *Version) BeforeDelete(tx *gorm.DB) error {
	if err := tx.Model(u).Association("Builds").Clear().Error; err != nil {
		return err
	}

	if err := tx.Delete(&VersionFile{}, "version_id = ?", u.ID).Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Version) Validate(db *gorm.DB) {
	if u.ModID == 0 {
		db.AddError(fmt.Errorf("A mod reference is required"))
	} else {
		res := db.Where(
			"id = ?",
			u.ModID,
		).First(
			&Mod{},
		)

		if res.RecordNotFound() {
			db.AddError(fmt.Errorf("Referenced mod does not exist"))
		}
	}

	if !govalidator.StringLength(u.Name, "2", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 2 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"mod_id = ?",
			u.ModID,
		).Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Version{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
