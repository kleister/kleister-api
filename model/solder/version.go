package solder

import (
	"github.com/solderapp/solder-api/model"
)

// Version represents a solder version model definition.
type Version struct {
	URL string `json:"url"`
	MD5 string `json:"md5"`
}

// NewVersionFromModel generates a solder model from our used models.
func NewVersionFromModel(source *model.Version) *Version {
	result := &Version{}

	result.URL = source.File.URL
	result.MD5 = source.File.MD5

	return result
}
