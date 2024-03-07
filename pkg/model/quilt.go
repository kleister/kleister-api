package model

import (
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Quilt within Kleister.
type Quilt struct {
	ID        string `gorm:"primaryKey;length:20"`
	Name      string `gorm:"unique;length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Builds    []*Build
}

// BeforeCreate defines the hook executed before every create.
func (m *Quilt) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uniuri.NewLen(uniuri.UUIDLen)
	}

	return nil
}
