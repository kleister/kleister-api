package solder

import (
	"github.com/solderapp/solder-api/model"
)

type Mod struct {
	Slug        string   `json:"name"`
	Name        string   `json:"pretty_name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Website     string   `json:"link"`
	Donate      string   `json:"donate"`
	Versions    []string `json:"versions"`
}

func NewModFromModel(source *model.Mod) *Mod {
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
