package model

import (
	"time"
)

// Build represents a build model definition.
type Build struct {
	ID          string `storm:"id" gorm:"primaryKey;length:36"`
	Pack        *Pack
	PackID      string `storm:"index" gorm:"index;length:36"`
	Minecraft   *Minecraft
	MinecraftID string `storm:"index" gorm:"index;length:36"`
	Forge       *Forge
	ForgeID     string `storm:"index" gorm:"index;length:36"`
	Slug        string `storm:"unique" gorm:"unique;length:255"`
	Name        string `storm:"unique" gorm:"unique;length:255"`
	MinJava     string
	MinMemory   string
	Published   bool `gorm:"default:false"`
	Hidden      bool `json:"-" gorm:"-"`
	Private     bool `gorm:"default:false"`
	Public      bool `json:"-" gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*BuildVersion
}
