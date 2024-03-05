package model

import (
	"time"
)

// Member within Kleister.
type Member struct {
	TeamID    string `gorm:"index:idx_id,unique;length:20"`
	Team      *Team
	UserID    string `gorm:"index:idx_id,unique;length:20"`
	User      *User
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
