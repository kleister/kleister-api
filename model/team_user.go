package model

// TeamUsers is simply a collection of team user structs.
type TeamUsers []*TeamUser

// TeamUser represents a team user model definition.
type TeamUser struct {
	TeamID int64  `json:"team_id" sql:"index"`
	Team   *Team  `json:"team,omitempty"`
	UserID int64  `json:"user_id" sql:"index"`
	User   *User  `json:"user,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
