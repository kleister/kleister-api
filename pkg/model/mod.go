package model

import (
	"time"
)

// Mod represents a mod model definition.
type Mod struct {
	ID          string `storm:"id" gorm:"primaryKey;length:36"`
	Slug        string `storm:"unique" gorm:"unique;length:255"`
	Name        string `storm:"unique" gorm:"unique;length:255"`
	Side        string `gorm:"index;length:36"`
	Description string
	Author      string
	Website     string
	Donate      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*Version
	Users       []*UserMod
	Teams       []*TeamMod
}
