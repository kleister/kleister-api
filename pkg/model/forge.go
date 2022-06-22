package model

import (
	"time"
)

// Forge within Kleister.
type Forge struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Minecraft *Minecraft
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
