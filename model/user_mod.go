package model

// UserMods is simply a collection of user mod structs.
type UserMods []*UserMod

// UserMod represents a user mod model definition.
type UserMod struct {
	UserID int64  `json:"user_id" sql:"index"`
	User   *User  `json:"user,omitempty"`
	ModID  int64  `json:"mod_id" sql:"index"`
	Mod    *Mod   `json:"mod,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
