package model

import (
	"time"
)

// User within Kleister.
type User struct {
	ID        string `gorm:"primaryKey;length:36"`
	Slug      string `gorm:"unique;length:255"`
	Username  string `gorm:"unique;length:255"`
	Password  string `gorm:"length:255"`
	Email     string `gorm:"unique;length:255"`
	Firstname string `gorm:"length:255"`
	Lastname  string `gorm:"length:255"`
	Active    bool   `gorm:"default:false"`
	Admin     bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Teams     []*Member
	Mods      []*UserMod
	Packs     []*UserPack
}
