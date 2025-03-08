package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Forge)(nil)
)

// Forge defines the model for forges table.
type Forge struct {
	bun.BaseModel `bun:"table:forges"`

	ID        string    `bun:",pk,type:varchar(20)"`
	Name      string    `bun:",unique,type:varchar(64)"`
	Minecraft string    `bun:"type:varchar(64)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Builds    []*Build  `bun:"rel:has-many,join:id=forge_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Forge) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
