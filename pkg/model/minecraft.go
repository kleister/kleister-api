package model

import (
	"time"
)

// Minecraft within Kleister.
type Minecraft struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	Type      string `storm:"index" gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
