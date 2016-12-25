package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Mods is simply a collection of mod structs.
type Mods []*Mod

// Mod represents a mod model definition.
type Mod struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	Slug        string    `json:"slug" sql:"unique_index"`
	Name        string    `json:"name" sql:"unique_index"`
	Side        string    `json:"side"`
	Description string    `json:"description" sql:"type:text"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Versions    Versions  `json:"versions,omitempty"`
	Users       Users     `json:"users,omitempty" gorm:"many2many:user_mods"`
	UserMods    UserMods  `json:"user_mods,omitempty"`
	Teams       Teams     `json:"teams,omitempty" gorm:"many2many:team_mods;"`
	TeamMods    TeamMods  `json:"team_mods,omitempty"`
}

// BeforeSave invokes required actions before persisting.
func (u *Mod) BeforeSave(db *gorm.DB) (err error) {
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

// BeforeDelete invokes required actions before deletion.
func (u *Mod) BeforeDelete(tx *gorm.DB) error {
	versions := Versions{}

	tx.Model(
		u,
	).Related(
		&versions,
	)

	if len(versions) > 0 {
		return fmt.Errorf("Can't delete, still assigned to versions")
	}

	if err := tx.Model(u).Association("Users").Clear().Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Teams").Clear().Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Mod) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, "2", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 2 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Mod{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if u.Side != "" {
		u.Side = strings.ToLower(u.Side)
		invalidSide := true

		for _, side := range []string{"client", "server", "both"} {
			if u.Side == side {
				invalidSide = false
			}
		}

		if invalidSide {
			db.AddError(fmt.Errorf("Selected side is invalid"))
		}
	}
}
