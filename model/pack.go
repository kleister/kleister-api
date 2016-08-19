package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Packs is simply a collection of pack structs.
type Packs []*Pack

// Pack represents a pack model definition.
type Pack struct {
	ID            int             `json:"id" gorm:"primary_key"`
	Icon          *PackIcon       `json:"icon,omitempty"`
	Logo          *PackLogo       `json:"logo,omitempty"`
	Background    *PackBackground `json:"background,omitempty"`
	Recommended   *Build          `json:"recommended,omitempty"`
	RecommendedID int             `json:"recommended_id" sql:"index"`
	Latest        *Build          `json:"latest,omitempty"`
	LatestID      int             `json:"latest_id" sql:"index"`
	Slug          string          `json:"slug" sql:"unique_index"`
	Name          string          `json:"name" sql:"unique_index"`
	Website       string          `json:"website"`
	Published     bool            `json:"published" sql:"default:false"`
	Private       bool            `json:"private" sql:"default:false"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Builds        Builds          `json:"builds,omitempty"`
	Clients       Clients         `json:"clients,omitempty" gorm:"many2many:client_packs"`
	Users         Users           `json:"users,omitempty" gorm:"many2many:user_packs;"`
	Teams         Teams           `json:"teams,omitempty" gorm:"many2many:team_packs;"`
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

// AfterDelete invokes required actions after deletion.
func (u *Pack) AfterDelete(tx *gorm.DB) error {
	for _, build := range u.Builds {
		if err := tx.Delete(build).Error; err != nil {
			return err
		}
	}

	if err := tx.Model(u).Association("Clients").Clear().Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Users").Clear().Error; err != nil {
		return err
	}

	if u.Icon != nil {
		if err := tx.Delete(u.Icon).Error; err != nil {
			return err
		}
	}

	if u.Background != nil {
		if err := tx.Delete(u.Background).Error; err != nil {
			return err
		}
	}

	if u.Logo != nil {
		if err := tx.Delete(u.Logo).Error; err != nil {
			return err
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
			&Pack{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
