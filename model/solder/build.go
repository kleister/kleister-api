package solder

import (
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
	Name    string `json:"name"`
	Version string `json:"version"`
	MD5     string `json:"md5"`
	URL     string `json:"url"`
}

// NewBuildFromModel generates a solder model from our used models.
func NewBuildFromModel(source *model.Build) *Build {
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
		result.Mods = append(
			result.Mods,
			&Child{
				Name:    version.Mod.Slug,
				Version: version.Name,
				MD5:     version.File.MD5,
				URL:     version.File.URL,
			},
		)
	}

	return result
}
