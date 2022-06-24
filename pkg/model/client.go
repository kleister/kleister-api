package model

import (
	"time"
)

// Client represents a client model definition.
type Client struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	Value     string `storm:"unique" gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Packs     []*ClientPack
}
