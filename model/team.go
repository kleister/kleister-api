package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Teams is simply a collection of registry structs.
type Teams []*Team

// Team represents a registry model definition.
type Team struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     Users     `json:"users,omitempty" gorm:"many2many:team_users;"`
	TeamUsers TeamUsers `json:"team_users,omitempty"`
	Mods      Mods      `json:"mods,omitempty" gorm:"many2many:team_mods;"`
	TeamMods  TeamMods  `json:"team_mods,omitempty"`
	Packs     Packs     `json:"packs,omitempty" gorm:"many2many:team_packs;"`
	TeamPacks TeamPacks `json:"team_packs,omitempty"`
}

// BeforeSave invokes required actions before persisting.
func (u *Team) BeforeSave(db *gorm.DB) (err error) {
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
				&Team{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// BeforeDelete invokes required actions before deletion.
func (u *Team) BeforeDelete(tx *gorm.DB) error {
	if err := tx.Model(u).Association("Users").Clear().Error; err != nil {
		return err
	}

	// TODO(tboerger): Prevent delete if team is last owner
	if err := tx.Model(u).Association("Mods").Clear().Error; err != nil {
		return err
	}

	// TODO(tboerger): Prevent delete if team is last owner
	if err := tx.Model(u).Association("Packs").Clear().Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Team) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, "1", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 1 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Team{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
