package model

import (
	"time"

	_ "github.com/Machiel/slugify"
)

// Packs is simply a collection of pack structs.
type Packs []*Pack

// Pack represents a pack model definition.
type Pack struct {
	ID            int64        `json:"id" gorm:"primary_key"`
	Icon          *Attachment  `json:"icon" gorm:"polymorphic:Owner"`
	Logo          *Attachment  `json:"logo" gorm:"polymorphic:Owner"`
	Background    *Attachment  `json:"background" gorm:"polymorphic:Owner"`
	Recommended   *Build       `json:"recommended"`
	RecommendedID int          `json:"recommended_id" sql:"index"`
	Latest        *Build       `json:"latest"`
	LatestID      int          `json:"latest_id" sql:"index"`
	Slug          string       `json:"slug" sql:"unique_index"`
	Name          string       `json:"name" sql:"unique_index"`
	Website       string       `json:"website"`
	Hidden        bool         `json:"hidden" sql:"default:true"`
	Private       bool         `json:"private" sql:"default:false"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Builds        *Builds      `json:"builds"`
	Permissions   *Permissions `json:"permissions" gorm:"many2many:permission_packs;"`
	Clients       *Clients     `json:"clients" gorm:"many2many:client_packs;"`
}
