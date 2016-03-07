package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Forges is simply a collection of forge structs.
type Forges []*Forge

// Forge represents a forge model definition.
type Forge struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    []*Build  `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *Forge) BeforeSave() (err error) {
	if u.Slug == "" {
		u.Slug = uuid.NewV4().String()
	}

	return nil
}
