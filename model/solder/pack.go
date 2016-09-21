package solder

import (
	"github.com/kleister/kleister-api/model"
)

// Pack represents a solder pack model definition.
type Pack struct {
	Slug          string   `json:"name,omitempty"`
	Name          string   `json:"display_name,omitempty"`
	Website       string   `json:"url,omitempty"`
	Recommended   string   `json:"recommended,omitempty"`
	Latest        string   `json:"latest,omitempty"`
	Icon          string   `json:"icon,omitempty"`
	IconMD5       string   `json:"icon_md5,omitempty"`
	Background    string   `json:"background,omitempty"`
	BackgroundMD5 string   `json:"background_md5,omitempty"`
	Logo          string   `json:"logo,omitempty"`
	LogoMD5       string   `json:"logo_md5,omitempty"`
	Builds        []string `json:"builds,omitempty"`
}

// NewPackFromModel generates a solder model from our used models.
func NewPackFromModel(source *model.Pack, client *model.Client, key *model.Key, include string) *Pack {
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

	if source.Icon == nil {
		result.Icon, result.IconMD5 = model.AttachmentDefault("icon")
	} else {
		result.Icon = source.Icon.URL
		result.IconMD5 = source.Icon.MD5
	}

	if source.Background == nil {
		result.Background, result.BackgroundMD5 = model.AttachmentDefault("background")
	} else {
		result.Background = source.Background.URL
		result.BackgroundMD5 = source.Background.MD5
	}

	if source.Logo == nil {
		result.Logo, result.LogoMD5 = model.AttachmentDefault("logo")
	} else {
		result.Logo = source.Logo.URL
		result.LogoMD5 = source.Logo.MD5
	}

	for _, build := range source.Builds {
		if build.Hidden {
			continue
		}

		if build.Public || key != nil {
			result.Builds = append(
				result.Builds,
				build.Slug,
			)
		} else {
			if client != nil {
				for _, pack := range client.Packs {
					if build.PackID == pack.ID {
						result.Builds = append(
							result.Builds,
							build.Slug,
						)

						break
					}
				}
			}
		}
	}

	return result
}
