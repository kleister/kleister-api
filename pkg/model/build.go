package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Build within Kleister.
type Build struct {
	ID          string `gorm:"primaryKey;length:20"`
	PackID      string `gorm:"index;length:20"`
	Pack        *Pack
	Name        string  `gorm:"length:255"`
	MinecraftID *string `gorm:"index:idx_id;length:20"`
	Minecraft   *Minecraft
	ForgeID     *string `gorm:"index:idx_id;length:20"`
	Forge       *Forge
	NeoforgeID  *string `gorm:"index:idx_id;length:20"`
	Neoforge    *Neoforge
	QuiltID     *string `gorm:"index:idx_id;length:20"`
	Quilt       *Quilt
	FabricID    *string `gorm:"index:idx_id;length:20"`
	Fabric      *Fabric
	Java        string `gorm:"length:255"`
	Memory      string `gorm:"length:255"`
	Latest      bool   `gorm:"default:false"`
	Recommended bool   `gorm:"default:false"`
	Public      bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*BuildVersion
}

// BeforeSave defines the hook executed before every save.
func (m *Build) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
