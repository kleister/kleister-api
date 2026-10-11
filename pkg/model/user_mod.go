package model

import (
	"time"
)

const (
	// UserModOwnerPerm defines the permission for an owner on user mods.
	UserModOwnerPerm = OwnerPerm

	// UserModAdminPerm defines the permission for an admin on user mods.
	UserModAdminPerm = AdminPerm

	// UserModUserPerm defines the permission for an user on user mods.
	UserModUserPerm = UserPerm
)

// UserMod defines the assignment of a user to a mod, backed by an OpenFGA
// relationship tuple rather than a database table.
type UserMod struct {
	UserID    string
	User      *User
	ModID     string
	Mod       *Mod
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
