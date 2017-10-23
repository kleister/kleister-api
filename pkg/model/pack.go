package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"gopkg.in/guregu/null.v3"
)

// Packs is simply a collection of pack structs.
type Packs []*Pack

// Pack represents a pack model definition.
type Pack struct {
	ID            int64           `json:"id" gorm:"primary_key"`
	Icon          *PackIcon       `json:"icon,omitempty"`
	Logo          *PackLogo       `json:"logo,omitempty"`
	Background    *PackBackground `json:"background,omitempty"`
	Recommended   *Build          `json:"recommended,omitempty"`
	RecommendedID null.Int        `json:"recommended_id" sql:"index"`
	Latest        *Build          `json:"latest,omitempty"`
	LatestID      null.Int        `json:"latest_id" sql:"index"`
	Slug          string          `json:"slug" sql:"unique_index"`
	Name          string          `json:"name" sql:"unique_index"`
	Website       string          `json:"website"`
	Published     bool            `json:"published" sql:"default:false"`
	Hidden        bool            `json:"hidden" sql:"-"`
	Private       bool            `json:"private" sql:"default:false"`
	Public        bool            `json:"public" sql:"-"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Builds        Builds          `json:"builds,omitempty"`
	Clients       Clients         `json:"clients,omitempty" gorm:"many2many:client_packs"`
	Users         Users           `json:"users,omitempty" gorm:"many2many:user_packs;"`
	UserPacks     UserPacks       `json:"user_packs,omitempty"`
	Teams         Teams           `json:"teams,omitempty" gorm:"many2many:team_packs;"`
	TeamPacks     TeamPacks       `json:"team_packs,omitempty"`
}

// AfterFind invokes required after loading a record from the database.
func (u *Pack) AfterFind(db *gorm.DB) {
	u.Hidden = !u.Published
	u.Public = !u.Private
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

// BeforeDelete invokes required actions before deletion.
func (u *Pack) BeforeDelete(tx *gorm.DB) error {
	builds := Builds{}

	tx.Model(
		u,
	).Related(
		&builds,
	)

	if len(builds) > 0 {
		return fmt.Errorf("Can't delete, still assigned to builds")
	}

	if err := tx.Model(u).Association("Users").Clear().Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Teams").Clear().Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Clients").Clear().Error; err != nil {
		return err
	}

	if err := tx.Delete(&PackIcon{}, "pack_id = ?", u.ID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&PackBackground{}, "pack_id = ?", u.ID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&PackLogo{}, "pack_id = ?", u.ID).Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Pack) Validate(db *gorm.DB) {
	if u.RecommendedID.Int64 > 0 {
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

	if u.LatestID.Int64 > 0 {
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
