package model

import (
	"time"
)

// UserPack within Kleister.
type UserPack struct {
	UserID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	User      *User
	PackID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Pack      *Pack
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
