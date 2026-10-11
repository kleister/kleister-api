package model

import (
	"time"
)

const (
	// UserGroupOwnerPerm defines the permission for an owner on user groups.
	UserGroupOwnerPerm = OwnerPerm

	// UserGroupAdminPerm defines the permission for an admin on user groups.
	UserGroupAdminPerm = AdminPerm

	// UserGroupUserPerm defines the permission for an user on user groups.
	UserGroupUserPerm = UserPerm
)

// UserGroup defines the membership of a user in a group, backed by an
// OpenFGA relationship tuple rather than a database table.
type UserGroup struct {
	UserID    string
	User      *User
	GroupID   string
	Group     *Group
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
