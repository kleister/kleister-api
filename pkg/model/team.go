package model

import (
	"time"
)

// Team within Kleister.
type Team struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*TeamUser
}
