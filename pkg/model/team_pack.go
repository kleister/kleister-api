package model

import (
	"time"
)

// TeamPack rwithin Kleister.
type TeamPack struct {
	TeamID    string `gorm:"index:idx_id,unique;length:20"`
	Team      *Team
	PackID    string `gorm:"index:idx_id,unique;length:20"`
	Pack      *Pack
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
