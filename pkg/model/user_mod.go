package model

import (
	"time"
)

// UserMod within Kleister.
type UserMod struct {
	UserID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	User      *User
	ModID     string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Mod       *Mod
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
