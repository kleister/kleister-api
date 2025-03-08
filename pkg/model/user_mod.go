package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*UserMod)(nil)
)

const (
	// UserModOwnerPerm defines the permission for an owner on user mods.
	UserModOwnerPerm = OwnerPerm

	// UserModAdminPerm defines the permission for an admin on user mods.
	UserModAdminPerm = AdminPerm

	// UserModUserPerm defines the permission for an user on user mods.
	UserModUserPerm = UserPerm
)

// UserMod defines the model for user_mods table.
type UserMod struct {
	bun.BaseModel `bun:"table:user_mods"`

	UserID    string    `bun:",pk,type:varchar(20)"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	ModID     string    `bun:",pk,type:varchar(20)"`
	Mod       *Mod      `bun:"rel:belongs-to,join:mod_id=id"`
	Perm      string    `bun:"type:varchar(32)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *UserMod) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
