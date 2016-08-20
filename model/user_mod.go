package model

// UserMods is simply a collection of user mod structs.
type UserMods []*UserMod

// UserMod represents a user mod model definition.
type UserMod struct {
	User *User  `json:"user,omitempty"`
	Mod  *Mod   `json:"mod,omitempty"`
	Perm string `json:"perm,omitempty"`
}
