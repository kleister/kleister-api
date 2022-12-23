package model

import (
	"time"
)

// Member within Kleister.
type Member struct {
	TeamID    string `gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	UserID    string `gorm:"index:idx_id,unique;length:36"`
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
