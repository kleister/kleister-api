package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Minecraft within Kleister.
type Minecraft struct {
	ID        string `gorm:"primaryKey;length:20"`
	Name      string `gorm:"unique;length:64"`
	Type      string `gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}

// BeforeSave defines the hook executed before every save.
func (m *Minecraft) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
