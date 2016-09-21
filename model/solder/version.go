package solder

import (
	"github.com/kleister/kleister-api/model"
)

// Version represents a solder version model definition.
type Version struct {
	URL string `json:"url,omitempty"`
	MD5 string `json:"md5,omitempty"`
}

// NewVersionFromModel generates a solder model from our used models.
func NewVersionFromModel(source *model.Version, client *model.Client, key *model.Key, include string) *Version {
	result := &Version{}

	result.URL = source.File.URL
	result.MD5 = source.File.MD5

	return result
}
