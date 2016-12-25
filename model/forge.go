package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Forges is simply a collection of forge structs.
type Forges []*Forge

// Filter searches for a name substring and returns a new collection.
func (u *Forges) Filter(term string) *Forges {
	res := Forges{}

	for _, record := range *u {
		if strings.Contains(record.Name, term) {
			res = append(res, record)
		}
	}

	return &res
}

// Forge represents a forge model definition.
type Forge struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    Builds    `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *Forge) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = u.Name
			} else {
				u.Slug = fmt.Sprintf("%s-%d", u.Name, i)
			}

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Forge{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// BeforeDelete invokes required actions before deletion.
func (u *Forge) BeforeDelete(tx *gorm.DB) error {
	builds := Builds{}

	tx.Model(
		u,
	).Related(
		&builds,
	)

	if len(builds) > 0 {
		return fmt.Errorf("Can't delete, still assigned to builds")
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Forge) Validate(db *gorm.DB) {

}
