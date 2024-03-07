package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	NeoforgeID  string `gorm:"index:idx_id;length:36"`
	Neoforge    *Neoforge
	QuiltID     string `gorm:"index:idx_id;length:36"`
	Quilt       *Quilt
	FabricID    string `gorm:"index:idx_id;length:36"`
	Fabric      *Fabric
	Java        string `gorm:"length:255"`
	Memory      string `gorm:"length:255"`
	Public      bool   `default:"true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*BuildVersion
}

// BeforeSave defines the hook executed before every save.
func (m *Build) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
