package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*UserGroup)(nil)
)

const (
	// UserGroupOwnerPerm defines the permission for an owner on user groups.
	UserGroupOwnerPerm = OwnerPerm

	// UserGroupAdminPerm defines the permission for an admin on user groups.
	UserGroupAdminPerm = AdminPerm

	// UserGroupUserPerm defines the permission for an user on user groups.
	UserGroupUserPerm = UserPerm
)

// UserGroup defines the model for user_groups table.
type UserGroup struct {
	bun.BaseModel `bun:"table:user_groups"`

	UserID    string    `bun:",pk,type:varchar(20)"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	GroupID   string    `bun:",pk,type:varchar(20)"`
	Group     *Group    `bun:"rel:belongs-to,join:group_id=id"`
	Perm      string    `bun:"type:varchar(32)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *UserGroup) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
