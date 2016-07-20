package solder

import (
	"github.com/kleister/kleister-api/model"
)

// Pack represents a solder pack model definition.
type Pack struct {
	Slug          string   `json:"name"`
	Name          string   `json:"display_name"`
	Website       string   `json:"url"`
	Recommended   string   `json:"recommended"`
	Latest        string   `json:"latest"`
	Icon          string   `json:"icon"`
	IconMD5       string   `json:"icon_md5"`
	Background    string   `json:"background"`
	BackgroundMD5 string   `json:"background_md5"`
	Logo          string   `json:"logo"`
	LogoMD5       string   `json:"logo_md5"`
	Builds        []string `json:"builds"`
}

// NewPackFromModel generates a solder model from our used models.
func NewPackFromModel(source *model.Pack) *Pack {
	result := &Pack{}

	result.Slug = source.Slug
	result.Name = source.Name
	result.Website = source.Website

	if source.Latest != nil {
		result.Latest = source.Latest.Slug
	}

	if source.Recommended != nil {
		result.Recommended = source.Recommended.Slug
	}

	if source.Icon != nil {
		result.Icon = source.Icon.URL
		result.IconMD5 = source.Icon.MD5
	}

	if source.Background != nil {
		result.Background = source.Background.URL
		result.BackgroundMD5 = source.Background.MD5
	}

	if source.Logo != nil {
		result.Logo = source.Logo.URL
		result.LogoMD5 = source.Logo.MD5
	}

	for _, build := range source.Builds {
		result.Builds = append(
			result.Builds,
			build.Slug,
		)
	}

	return result
}
