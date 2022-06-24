package model

import (
	"time"
)

// TeamPack represents a team pack model definition.
type TeamPack struct {
	TeamID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	PackID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Pack      *Pack
	Perm      string `gorm:"length:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
