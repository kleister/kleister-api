package model

import (
	"time"
)

const (
	// GroupPackOwnerPerm defines the permission for an owner on group packs.
	GroupPackOwnerPerm = OwnerPerm

	// GroupPackAdminPerm defines the permission for an admin on group packs.
	GroupPackAdminPerm = AdminPerm

	// GroupPackUserPerm defines the permission for an user on group packs.
	GroupPackUserPerm = UserPerm
)

// GroupPack defines the assignment of a group to a pack, backed by an
// OpenFGA relationship tuple rather than a database table.
type GroupPack struct {
	GroupID   string
	Group     *Group
	PackID    string
	Pack      *Pack
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
