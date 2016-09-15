package model

// TeamMods is simply a collection of team mod structs.
type TeamMods []*TeamMod

// TeamMod represents a team mod model definition.
type TeamMod struct {
	TeamID int64  `json:"team_id" sql:"index"`
	Team   *Team  `json:"team,omitempty"`
	ModID  int64  `json:"mod_id" sql:"index"`
	Mod    *Mod   `json:"mod,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
