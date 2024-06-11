package model

import (
	"time"
)

// UserTeam within Kleister.
type UserTeam struct {
	TeamID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Team      *Team
	UserID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	User      *User
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
