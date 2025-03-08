package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*BuildVersion)(nil)
)

// BuildVersion defines the model for build_versions table.
type BuildVersion struct {
	bun.BaseModel `bun:"table:build_versions"`

	BuildID   string    `bun:",pk,type:varchar(20)"`
	Build     *Build    `bun:"rel:belongs-to,join:build_id=id"`
	VersionID string    `bun:",pk,type:varchar(20)"`
	Version   *Version  `bun:"rel:belongs-to,join:version_id=id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *BuildVersion) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
