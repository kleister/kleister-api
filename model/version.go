package model

import (
	"time"

	_ "github.com/Machiel/slugify"
)

// Versions is simply a collection of version structs.
type Versions []*Version

// Version represents a version model definition.
type Version struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Mod       *Mod      `json:"mod"`
	ModID     int       `json:"mod_id" sql:"index"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
