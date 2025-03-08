package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Pack)(nil)
)

// Pack defines the model for packs table.
type Pack struct {
	bun.BaseModel `bun:"table:packs"`

	ID        string       `bun:",pk,type:varchar(20)"`
	Slug      string       `bun:",unique,type:varchar(255)"`
	Name      string       `bun:",unique,type:varchar(255)"`
	Avatar    *PackAvatar  `bun:"rel:has-one,join:id=pack_id"`
	Website   string       `bun:"type:varchar(255)"`
	Public    bool         `bun:"default:true"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	Builds    []*Build     `bun:"rel:has-many,join:id=pack_id"`
	Users     []*UserPack  `bun:"rel:has-many,join:id=pack_id"`
	Groups    []*GroupPack `bun:"rel:has-many,join:id=pack_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Pack) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
