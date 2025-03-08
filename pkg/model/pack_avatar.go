package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*PackAvatar)(nil)
)

// PackAvatar defines the model for pack_avatars table.
type PackAvatar struct {
	bun.BaseModel `bun:"table:pack_avatars"`

	ID        string    `bun:",pk,type:varchar(20)"`
	Pack      *Pack     `bun:"rel:belongs-to,join:pack_id=id"`
	PackID    string    `bun:"type:varchar(20)"`
	Slug      string    `bun:"type:varchar(255)"`
	URL       string    `bun:"-"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *PackAvatar) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
