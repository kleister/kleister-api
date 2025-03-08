package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*GroupPack)(nil)
)

const (
	// GroupPackOwnerPerm defines the permission for an owner on group packs.
	GroupPackOwnerPerm = OwnerPerm

	// GroupPackAdminPerm defines the permission for an admin on group packs.
	GroupPackAdminPerm = AdminPerm

	// GroupPackUserPerm defines the permission for an user on group packs.
	GroupPackUserPerm = UserPerm
)

// GroupPack defines the packel for group_packs table.
type GroupPack struct {
	bun.BaseModel `bun:"table:group_packs"`

	GroupID   string    `bun:",pk,type:varchar(20)"`
	Group     *Group    `bun:"rel:belongs-to,join:group_id=id"`
	PackID    string    `bun:",pk,type:varchar(20)"`
	Pack      *Pack     `bun:"rel:belongs-to,join:pack_id=id"`
	Perm      string    `bun:"type:varchar(32)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *GroupPack) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
