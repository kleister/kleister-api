package model

import (
	"time"
)

// TeamPack rwithin Kleister.
type TeamPack struct {
	TeamID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Team      *Team
	PackID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Pack      *Pack
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
