package model

import (
	"time"
)

// Build within Kleister.
type Build struct {
	ID          string `gorm:"primaryKey;length:36"`
	PackID      string `gorm:"index:idx_id;length:36"`
	Pack        *Pack
	Slug        string `gorm:"length:255"`
	Name        string `gorm:"length:255"`
	MinecraftID string `gorm:"index:idx_id;length:36"`
	Minecraft   *Minecraft
	ForgeID     string `gorm:"index:idx_id;length:36"`
	Forge       *Forge
	Java        string `gorm:"length:255"`
	Memory      string `gorm:"length:255"`
	Recommended bool   `gorm:"default:false"`
	Published   bool   `gorm:"default:true"`
	Private     bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*BuildVersion
}
