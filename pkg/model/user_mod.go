package model

import (
	"time"
)

// UserMod within Kleister.
type UserMod struct {
	UserID    string `gorm:"index:idx_id,unique;length:20"`
	User      *User
	ModID     string `gorm:"index:idx_id,unique;length:20"`
	Mod       *Mod
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
