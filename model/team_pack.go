package model

// TeamPacks is simply a collection of team pack structs.
type TeamPacks []*TeamPack

// TeamPack represents a team pack model definition.
type TeamPack struct {
	Team *Team  `json:"team,omitempty"`
	Pack *Pack  `json:"pack,omitempty"`
	Perm string `json:"perm,omitempty"`
}
