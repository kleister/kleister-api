package model

import (
	"time"
)

// Team within Kleister.
type Team struct {
	ID        string `gorm:"primaryKey;length:36"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*Member
}
