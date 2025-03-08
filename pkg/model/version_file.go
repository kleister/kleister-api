package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*VersionFile)(nil)
)

// VersionFile defines the model for version_files table.
type VersionFile struct {
	bun.BaseModel `bun:"table:version_files"`

	ID          string    `bun:",pk,type:varchar(20)"`
	Version     *Version  `bun:"rel:belongs-to,join:version_id=id"`
	VersionID   string    `bun:"type:varchar(20)"`
	Slug        string    `bun:"type:varchar(255)"`
	ContentType string    `bun:"type:varchar(255)"`
	MD5         string    `bun:"column:md5,type:varchar(255)"`
	URL         string    `bun:"-"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *VersionFile) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
