package model

import (
	"time"
)

const (
	// UserPackOwnerPerm defines the permission for an owner on user packs.
	UserPackOwnerPerm = OwnerPerm

	// UserPackAdminPerm defines the permission for an admin on user packs.
	UserPackAdminPerm = AdminPerm

	// UserPackUserPerm defines the permission for an user on user packs.
	UserPackUserPerm = UserPerm
)

// UserPack defines the assignment of a user to a pack, backed by an OpenFGA
// relationship tuple rather than a database table.
type UserPack struct {
	UserID    string
	User      *User
	PackID    string
	Pack      *Pack
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
