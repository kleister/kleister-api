package model

import (
	"time"
)

// Mod within Kleister.
type Mod struct {
	ID          string `gorm:"primaryKey;length:36"`
	Slug        string `gorm:"unique;length:255"`
	Name        string `gorm:"unique;length:255"`
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
