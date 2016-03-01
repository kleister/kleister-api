package model

import (
	"github.com/jinzhu/gorm"
)

type Build struct {
	gorm.Model

	Pack        *Pack      `json:"pack"`
	PackID      int        `json:"pack_id" sql:"index"`
	Minecraft   *Minecraft `json:"minecraft"`
	MinecraftID int        `json:"minecraft_id" sql:"index"`
	Forge       *Forge     `json:"forge"`
	ForgeID     int        `json:"forge_id" sql:"index"`
	Slug        string     `json:"slug" sql:"unique_index"`
	Name        string     `json:"name" sql:"unique_index"`
	MinJava     string     `json:"min_java"`
	MinMemory   string     `json:"min_memory"`
	Published   bool       `json:"published" sql:"default:false"`
	Private     bool       `json:"private" sql:"default:false"`
	Mods        []*Mod     `json:"mods" gorm:"many2many:build_mods;"`
}
