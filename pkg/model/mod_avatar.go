package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*ModAvatar)(nil)
)

// ModAvatar defines the model for mod_avatars table.
type ModAvatar struct {
	bun.BaseModel `bun:"table:mod_avatars"`

	ID        string    `bun:",pk,type:varchar(20)"`
	Mod       *Mod      `bun:"rel:belongs-to,join:mod_id=id"`
	ModID     string    `bun:"type:varchar(20)"`
	Slug      string    `bun:"type:varchar(255)"`
	URL       string    `bun:"-"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *ModAvatar) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
