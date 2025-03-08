package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*GroupMod)(nil)
)

const (
	// GroupModOwnerPerm defines the permission for an owner on group mods.
	GroupModOwnerPerm = OwnerPerm

	// GroupModAdminPerm defines the permission for an admin on group mods.
	GroupModAdminPerm = AdminPerm

	// GroupModUserPerm defines the permission for an user on group mods.
	GroupModUserPerm = UserPerm
)

// GroupMod defines the model for group_mods table.
type GroupMod struct {
	bun.BaseModel `bun:"table:group_mods"`

	GroupID   string    `bun:",pk,type:varchar(20)"`
	Group     *Group    `bun:"rel:belongs-to,join:group_id=id"`
	ModID     string    `bun:",pk,type:varchar(20)"`
	Mod       *Mod      `bun:"rel:belongs-to,join:mod_id=id"`
	Perm      string    `bun:"type:varchar(32)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *GroupMod) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
