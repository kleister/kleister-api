package model

import (
	"time"
)

const (
	// GroupModOwnerPerm defines the permission for an owner on group mods.
	GroupModOwnerPerm = OwnerPerm

	// GroupModAdminPerm defines the permission for an admin on group mods.
	GroupModAdminPerm = AdminPerm

	// GroupModUserPerm defines the permission for an user on group mods.
	GroupModUserPerm = UserPerm
)

// GroupMod defines the assignment of a group to a mod, backed by an OpenFGA
// relationship tuple rather than a database table.
type GroupMod struct {
	GroupID   string
	Group     *Group
	ModID     string
	Mod       *Mod
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
