package model

import (
	"time"
)

// TeamMod within Kleister.
type TeamMod struct {
	TeamID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Team      *Team
	ModID     string `gorm:"primaryKey;autoIncrement:false;length:20"`
	Mod       *Mod
	Perm      string `gorm:"length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
