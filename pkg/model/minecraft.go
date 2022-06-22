package model

import (
	"time"
)

// Minecraft within Kleister.
type Minecraft struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	Type      string `gorm:"length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
