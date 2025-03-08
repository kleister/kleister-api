package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*UserPack)(nil)
)

const (
	// UserPackOwnerPerm defines the permission for an owner on user packs.
	UserPackOwnerPerm = OwnerPerm

	// UserPackAdminPerm defines the permission for an admin on user packs.
	UserPackAdminPerm = AdminPerm

	// UserPackUserPerm defines the permission for an user on user packs.
	UserPackUserPerm = UserPerm
)

// UserPack defines the model for user_packs table.
type UserPack struct {
	bun.BaseModel `bun:"table:user_packs"`

	UserID    string    `bun:",pk,type:varchar(20)"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	PackID    string    `bun:",pk,type:varchar(20)"`
	Pack      *Pack     `bun:"rel:belongs-to,join:pack_id=id"`
	Perm      string    `bun:"type:varchar(32)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *UserPack) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
