package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User within Kleister.
type User struct {
	ID        string `gorm:"primaryKey;length:20"`
	Scim      string `gorm:"length:255"`
	Username  string `gorm:"unique;length:255"`
	Password  string `gorm:"-"`
	Hashword  string `gorm:"lenght:255"`
	Email     string `gorm:"length:255"`
	Fullname  string `gorm:"length:255"`
	Profile   string `gorm:"-"`
	Active    bool   `gorm:"default:false"`
	Admin     bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Auths     []*UserAuth
	Teams     []*UserTeam
	Mods      []*UserMod
	Packs     []*UserPack
}

// BeforeSave defines the hook executed before every save.
func (m *User) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	if m.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword(
			[]byte(m.Password),
			10,
		)

		if err != nil {
			return err
		}

		m.Hashword = string(bytes)
	}

	return nil
}
