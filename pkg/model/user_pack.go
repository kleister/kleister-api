package model

import (
	"time"
)

// UserPack represents a user pack model definition.
type UserPack struct {
	UserID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	User      *User
	PackID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Pack      *Pack
	Perm      string `gorm:"length:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
