package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
	"github.com/vincent-petithory/dataurl"
)

var (
	_ bun.BeforeAppendModelHook = (*Version)(nil)
)

// Version defines the model for versions table.
type Version struct {
	bun.BaseModel `bun:"table:versions"`

	ID         string           `bun:",pk,type:varchar(20)"`
	ModID      string           `bun:"type:varchar(20)"`
	Mod        *Mod             `bun:"rel:belongs-to,join:mod_id=id"`
	Name       string           `bun:",unique,type:varchar(255)"`
	File       *VersionFile     `bun:"rel:has-one,join:id=version_id"`
	FileUpload *dataurl.DataURL `bun:"-"`
	Public     bool             `bun:"default:true"`
	CreatedAt  time.Time        `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time        `bun:",nullzero,notnull,default:current_timestamp"`
	Builds     []*BuildVersion  `bun:"rel:has-many,join:id=version_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Version) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
