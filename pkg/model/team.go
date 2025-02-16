package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Team within Kleister.
type Team struct {
	ID        string `gorm:"primaryKey;length:20"`
	Scim      string `gorm:"length:255"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*UserTeam
	Mods      []*TeamMod
	Packs     []*TeamPack
}

// BeforeSave defines the hook executed before every save.
func (m *Team) BeforeSave(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	if m.Slug == "" {
		m.Slug = Slugify(
			tx.Model(&Team{}),
			m.Name,
			m.ID,
			false,
		)
	}

	return nil
}
