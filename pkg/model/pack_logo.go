package model

import (
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/vincent-petithory/dataurl"
	"gorm.io/gorm"
)

// PackLogo within Kleister.
type PackLogo struct {
	ID          string `gorm:"primaryKey;length:20"`
	Pack        *Pack
	PackID      string `gorm:"index;length:20"`
	Slug        string `gorm:"unique;length:255"`
	ContentType string
	MD5         string           `gorm:"column:md5"`
	Path        string           `gorm:"-"`
	URL         string           `gorm:"-"`
	Upload      *dataurl.DataURL `json:"-" gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeSave defines the hook executed before every save.
func (m *PackLogo) BeforeSave(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
	}

	return nil
}
