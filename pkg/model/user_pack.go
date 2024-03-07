package model

import (
	"time"
)

// UserPack within Kleister.
type UserPack struct {
	UserID    string `gorm:"index:idx_id,unique;length:36"`
	User      *User
	PackID    string `gorm:"index:idx_id,unique;length:36"`
	Pack      *Pack
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
