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
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    []*Build  `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *Minecraft) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		u.Slug = uuid.NewV4().String()
	}

	return nil
}
