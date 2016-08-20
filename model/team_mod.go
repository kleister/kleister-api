package model

// TeamMods is simply a collection of team mod structs.
type TeamMods []*TeamMod

// TeamMod represents a team mod model definition.
type TeamMod struct {
	Team *Team  `json:"team,omitempty"`
	Mod  *Mod   `json:"mod,omitempty"`
	Perm string `json:"perm,omitempty"`
}
