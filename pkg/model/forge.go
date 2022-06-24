package model

import (
	"time"
)

// Forge represents a forge model definition.
type Forge struct {
	ID        string `storm:"id" gorm:"primaryKey;length:36"`
	Name      string `storm:"unique" gorm:"unique;length:255"`
	Minecraft string `storm:"index" gorm:"index;length:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}
