package solder

import (
	"sort"

	"github.com/kleister/kleister-api/model"
)

// Build represents a solder build model definition.
type Build struct {
	Minecraft string   `json:"minecraft"`
	Forge     string   `json:"forge"`
	Mods      []*Child `json:"mods"`
}

// Child represents a build mod model definition.
type Child struct {
	Slug        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	MD5         string `json:"md5,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"pretty_name,omitempty"`
	Author      string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Donate      string `json:"donate,omitempty"`
}

// ChildBySlug sorts a list of children be the slug.
type ChildBySlug []*Child

// Len is part of the child sorting algorithm.
func (u ChildBySlug) Len() int {
	return len(u)
}

// Swap is part of the child sorting algorithm.
func (u ChildBySlug) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the child sorting algorithm.
func (u ChildBySlug) Less(i, j int) bool {
	return u[i].Slug < u[j].Slug
}

// NewBuildFromModel generates a solder model from our used models.
func NewBuildFromModel(source *model.Build, client *model.Client, key *model.Key, include string) *Build {
	result := &Build{
		Mods: make([]*Child, 0),
	}

	if source.Minecraft != nil {
		result.Minecraft = source.Minecraft.Name
	}

	if source.Forge != nil {
		result.Forge = source.Forge.Name
	}

	for _, version := range source.Versions {
		switch include {
		case "mods":
			result.Mods = append(
				result.Mods,
				&Child{
					Slug:        version.Mod.Slug,
					Version:     version.Name,
					MD5:         version.File.MD5,
					URL:         version.File.URL,
					Name:        version.Mod.Name,
					Author:      version.Mod.Author,
					Description: version.Mod.Description,
					Link:        version.Mod.Website,
					Donate:      version.Mod.Donate,
				},
			)
		default:
			result.Mods = append(
				result.Mods,
				&Child{
					Slug:    version.Mod.Slug,
					Version: version.Name,
					MD5:     version.File.MD5,
					URL:     version.File.URL,
				},
			)
		}
	}

	sort.Sort(
		ChildBySlug(
			result.Mods,
		),
	)

	return result
}
