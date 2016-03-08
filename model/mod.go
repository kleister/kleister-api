package model

import (
	"time"

	_ "github.com/Machiel/slugify"
)

// Mods is simply a collection of mod structs.
type Mods []*Mod

// Mod represents a mod model definition.
type Mod struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	Slug        string    `json:"slug" sql:"unique_index"`
	Name        string    `json:"name" sql:"unique_index"`
	Description string    `json:"description" sql:"type:text"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Versions    *Versions `json:"versions"`
}
