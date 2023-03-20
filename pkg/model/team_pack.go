package model

import (
	"time"
)

// TeamPack rwithin Kleister.
type TeamPack struct {
	TeamID    string `gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	PackID    string `gorm:"index:idx_id,unique;length:36"`
	Pack      *Pack
	CreatedAt time.Time
	UpdatedAt time.Time
}
