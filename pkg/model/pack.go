package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pack within Kleister.
type Pack struct {
	ID            string `gorm:"primaryKey;length:36"`
	Recommended   *Build
	RecommendedID *string `gorm:"index;length:36"`
	Latest        *Build
	LatestID      *string `gorm:"index;length:36"`
	Slug          string  `gorm:"unique;length:255"`
	Name          string  `gorm:"unique;length:255"`
	Back          *PackBack
	Icon          *PackIcon
	Logo          *PackLogo
	Website       string
	Public        bool `default:"true"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Builds        []*Build
	Users         []*UserPack
	Teams         []*TeamPack
}

// BeforeSave defines the hook executed before every save.
func (m *Pack) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return nil
}
