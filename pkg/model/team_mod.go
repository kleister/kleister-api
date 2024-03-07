package model

import (
	"time"
)

// TeamMod within Kleister.
type TeamMod struct {
	TeamID    string `gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	ModID     string `gorm:"index:idx_id,unique;length:36"`
	Mod       *Mod
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
