package model

// UserPacks is simply a collection of user pack structs.
type UserPacks []*UserPack

// UserPack represents a user pack model definition.
type UserPack struct {
	User *User  `json:"user,omitempty"`
	Pack *Pack  `json:"pack,omitempty"`
	Perm string `json:"perm,omitempty"`
}
