package model

import (
	"time"
)

// UserMod represents a user mod model definition.
type UserMod struct {
	UserID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	User      *User
	ModID     string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Mod       *Mod
	Perm      string `gorm:"length:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
