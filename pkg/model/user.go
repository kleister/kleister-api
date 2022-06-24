package model

import (
	"time"
)

// User represents a user model definition.
type User struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Slug      string `storm:"unique" gorm:"unique;length:255"`
	Email     string `storm:"unique" gorm:"unique;length:255"`
	Username  string `storm:"unique" gorm:"unique;length:255"`
	Password  string `gorm:"length:255"`
	Avatar    string `sql:"-"`
	Active    bool   `gorm:"default:false"`
	Admin     bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Teams     []*TeamUser
	Mods      []*UserMod
	Packs     []*UserPack
}
