package model

// TeamUsers is simply a collection of team user structs.
type TeamUsers []*TeamUser

// TeamUser represents a team user model definition.
type TeamUser struct {
	Team *Team  `json:"team,omitempty"`
	User *User  `json:"user,omitempty"`
	Perm string `json:"perm,omitempty"`
}
