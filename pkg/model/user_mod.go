package model

import (
	"time"
)

// UserMod within Kleister.
type UserMod struct {
	UserID    string `gorm:"index:idx_id,unique;length:36"`
	User      *User
	ModID     string `gorm:"index:idx_id,unique;length:36"`
	Mod       *Mod
	CreatedAt time.Time
	UpdatedAt time.Time
}
