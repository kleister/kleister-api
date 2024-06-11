package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Version within Kleister.
type Version struct {
	ID        string `gorm:"primaryKey;length:20"`
	ModID     string `gorm:"index;length:20"`
	Mod       *Mod
	Name      string `gorm:"unique;length:255"`
	File      *VersionFile
	Public    bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*BuildVersion
}

// BeforeSave defines the hook executed before every save.
func (m *Version) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
