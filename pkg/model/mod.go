package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Mod)(nil)
)

// Mod defines the model for mods table.
type Mod struct {
	bun.BaseModel `bun:"table:mods"`

	ID          string      `bun:",pk,type:varchar(20)"`
	Slug        string      `bun:",unique,type:varchar(255)"`
	Name        string      `bun:",unique,type:varchar(255)"`
	Avatar      *ModAvatar  `bun:"rel:has-one,join:id=mod_id"`
	Side        string      `bun:"type:varchar(64)"`
	Description string      `bun:"type:text"`
	Author      string      `bun:"type:varchar(255)"`
	Website     string      `bun:"type:varchar(255)"`
	Donate      string      `bun:"type:varchar(255)"`
	Public      bool        `bun:"default:true"`
	CreatedAt   time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	Versions    []*Version  `bun:"rel:has-many,join:id=mod_id"`
	Users       []*UserMod  `bun:"rel:has-many,join:id=mod_id"`
	Groups      []*GroupMod `bun:"rel:has-many,join:id=mod_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Mod) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if m.ID == "" {
			m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
		}

		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		if m.ID == "" {
			m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
		}

		m.UpdatedAt = time.Now()
	}

	return nil
}
