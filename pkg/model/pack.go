package model

import (
	"time"
)

// Pack within Kleister.
type Pack struct {
	ID        string `gorm:"primaryKey;length:36"`
	Slug      string `gorm:"unique;length:255"`
	Name      string `gorm:"unique;length:255"`
	Back      *PackBack
	Icon      *PackIcon
	Logo      *PackLogo
	Website   string
	Published bool `gorm:"default:true"`
	Private   bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
	Users     []*UserPack
	Teams     []*TeamPack
}
