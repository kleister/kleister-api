package model

import (
	"time"
)

// TeamMod represents a team mod model definition.
type TeamMod struct {
	TeamID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	ModID     string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Mod       *Mod
	Perm      string `gorm:"length:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
