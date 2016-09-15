package model

// UserPacks is simply a collection of user pack structs.
type UserPacks []*UserPack

// UserPack represents a user pack model definition.
type UserPack struct {
	UserID int64  `json:"user_id" sql:"index"`
	User   *User  `json:"user,omitempty"`
	PackID int64  `json:"pack_id" sql:"index"`
	Pack   *Pack  `json:"pack,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
