package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Group)(nil)
)

// Group defines the model for groups table.
type Group struct {
	bun.BaseModel `bun:"table:groups"`

	ID        string       `bun:",pk,type:varchar(20)"`
	Scim      string       `bun:"type:varchar(255)"`
	Slug      string       `bun:",unique,type:varchar(255)"`
	Name      string       `bun:"type:varchar(255)"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	Users     []*UserGroup `bun:"rel:has-many,join:id=group_id"`
	Mods      []*GroupMod  `bun:"rel:has-many,join:id=group_id"`
	Packs     []*GroupPack `bun:"rel:has-many,join:id=group_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Group) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
