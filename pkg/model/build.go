package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*Build)(nil)
)

// Build defines the model for builds table.
type Build struct {
	bun.BaseModel `bun:"table:builds"`

	ID          string          `bun:",pk,type:varchar(20)"`
	PackID      string          `bun:"type:varchar(20)"`
	Pack        *Pack           `bun:"rel:belongs-to,join:pack_id=id"`
	Name        string          `bun:"type:varchar(255)"`
	MinecraftID *string         `bun:"type:varchar(20)"`
	Minecraft   *Minecraft      `bun:"rel:belongs-to,join:minecraft_id=id"`
	ForgeID     *string         `bun:"type:varchar(20)"`
	Forge       *Forge          `bun:"rel:belongs-to,join:forge_id=id"`
	NeoforgeID  *string         `bun:"type:varchar(20)"`
	Neoforge    *Neoforge       `bun:"rel:belongs-to,join:neoforge_id=id"`
	QuiltID     *string         `bun:"type:varchar(20)"`
	Quilt       *Quilt          `bun:"rel:belongs-to,join:quilt_id=id"`
	FabricID    *string         `bun:"type:varchar(20)"`
	Fabric      *Fabric         `bun:"rel:belongs-to,join:fabric_id=id"`
	Java        string          `bun:"type:varchar(255)"`
	Memory      string          `bun:"type:varchar(255)"`
	Latest      bool            `bun:"default:false"`
	Recommended bool            `bun:"default:false"`
	Public      bool            `bun:"default:true"`
	CreatedAt   time.Time       `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time       `bun:",nullzero,notnull,default:current_timestamp"`
	Versions    []*BuildVersion `bun:"rel:has-many,join:id=build_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *Build) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
