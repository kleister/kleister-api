package solder

import (
	"github.com/kleister/kleister-api/model"
)

// Mod represents a solder mod model definition.
type Mod struct {
	Slug        string   `json:"name,omitempty"`
	Name        string   `json:"pretty_name,omitempty"`
	Description string   `json:"description,omitempty"`
	Author      string   `json:"author,omitempty"`
	Website     string   `json:"link,omitempty"`
	Donate      string   `json:"donate,omitempty"`
	Versions    []string `json:"versions,omitempty"`
}

// NewModFromModel generates a solder model from our used models.
func NewModFromModel(source *model.Mod, client *model.Client, key *model.Key, include string) *Mod {
	result := &Mod{}

	result.Slug = source.Slug
	result.Name = source.Name
	result.Description = source.Description
	result.Author = source.Author
	result.Website = source.Website
	result.Donate = source.Donate

	for _, version := range source.Versions {
		result.Versions = append(
			result.Versions,
			version.Slug,
		)
	}

	return result
}
