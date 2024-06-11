package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Forge rwithin Kleister.
type Forge struct {
	ID        string `gorm:"primaryKey;length:20"`
	Name      string `gorm:"unique;length:64"`
	Minecraft string `gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}

// BeforeCreate defines the hook executed before every create.
func (m *Forge) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
