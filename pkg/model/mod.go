package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Mod within Kleister.
type Mod struct {
	ID          string `gorm:"primaryKey;length:20"`
	Slug        string `gorm:"unique;length:255"`
	Name        string `gorm:"unique;length:255"`
	Side        string `gorm:"index;length:64"`
	Description string
	Author      string
	Website     string
	Donate      string
	Public      bool `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []*Version
	Users       []*UserMod
	Teams       []*TeamMod
}

// BeforeSave defines the hook executed before every save.
func (m *Mod) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
