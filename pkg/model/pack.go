package model

import (
	"time"
)

// Pack represents a pack model definition.
type Pack struct {
	ID            string `storm:"id" gorm:"primaryKey;length:36"`
	Icon          *PackIcon
	Logo          *PackLogo
	Background    *PackBackground
	Recommended   *Build
	RecommendedID string `storm:"index" gorm:"index;length:36"`
	Latest        *Build
	LatestID      string `storm:"index" gorm:"index;length:36"`
	Slug          string `storm:"unique" gorm:"unique;length:255"`
	Name          string `storm:"unique" gorm:"unique;length:255"`
	Website       string
	Published     bool `gorm:"default:false"`
	Hidden        bool `json:"-" gorm:"-"`
	Private       bool `gorm:"default:false"`
	Public        bool `json:"-" gorm:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Builds        []*Build
	Users         []*UserPack
	Teams         []*TeamPack
}
