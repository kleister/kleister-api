package model

import (
	"time"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// UserAuth provides the model definition for a user auth.
type UserAuth struct {
	ID        string `gorm:"primaryKey;length:20"`
	UserID    string `gorm:"length:20"`
	User      *User
	Provider  string `gorm:"length:255"`
	Ref       string `gorm:"length:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeSave defines the hook executed before every save.
func (m *UserAuth) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uniuri.NewLen(uniuri.UUIDLen)
	}

	return nil
}
