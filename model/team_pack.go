package model

// TeamPacks is simply a collection of team pack structs.
type TeamPacks []*TeamPack

// TeamPack represents a team pack model definition.
type TeamPack struct {
	TeamID int64  `json:"team_id" sql:"index"`
	Team   *Team  `json:"team,omitempty"`
	PackID int64  `json:"pack_id" sql:"index"`
	Pack   *Pack  `json:"pack,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
