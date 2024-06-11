package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Pack within Kleister.
type Pack struct {
	ID        string `gorm:"primaryKey;length:20"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	Back      *PackBack
	Icon      *PackIcon
	Logo      *PackLogo
	Website   string
	Public    bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
	Users     []*UserPack
	Teams     []*TeamPack
}

// BeforeSave defines the hook executed before every save.
func (m *Pack) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
