package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Quilt)(nil)
)

// Quilt defines the model for quilts table.
type Quilt struct {
	bun.BaseModel `bun:"table:quilts"`

	ID        string    `bun:",pk,type:varchar(20)"`
	Name      string    `bun:",unique,type:varchar(64)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Builds    []*Build  `bun:"rel:has-many,join:id=quilt_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Quilt) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
