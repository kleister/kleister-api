package model

import (
	"time"
)

// Team represents a team model definition.
type Team struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*TeamUser
	Mods      []*TeamMod
	Packs     []*TeamPack
}
