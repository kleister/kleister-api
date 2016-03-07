package model

import (
	"time"
)

// Mod represents a mod model definition.
type Mod struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	Slug        string     `json:"slug" sql:"unique_index"`
	Name        string     `json:"name" sql:"unique_index"`
	Description string     `json:"description" sql:"type:text"`
	Author      string     `json:"author"`
	Website     string     `json:"website"`
	Donate      string     `json:"donate"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Versions    []*Version `json:"versions"`
}
