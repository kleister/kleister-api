package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// MinecraftDefaultOrder is the default ordering for minecraft listings.
func MinecraftDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"minecrafts.name DESC",
	)
}

// Minecrafts is simply a collection of minecraft structs.
type Minecrafts []*Minecraft

// Minecraft represents a minecraft model definition.
type Minecraft struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    Builds    `json:"builds"`
}

// BeforeSave invokes required actions before persisting.
func (u *Minecraft) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for {
			u.Slug = uuid.NewV4().String()

			notFound := db.Where(
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
