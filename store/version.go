package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetVersions retrieves all available versions from the database.
func GetVersions(c context.Context, mod int) (*model.Versions, error) {
	return FromContext(c).GetVersions(mod)
}

// CreateVersion creates a new version.
func CreateVersion(c context.Context, mod int, record *model.Version) error {
	return FromContext(c).CreateVersion(mod, record)
}

// UpdateVersion updates a version.
func UpdateVersion(c context.Context, mod int, record *model.Version) error {
	return FromContext(c).UpdateVersion(mod, record)
}

// DeleteVersion deletes a version.
func DeleteVersion(c context.Context, mod int, record *model.Version) error {
	return FromContext(c).DeleteVersion(mod, record)
}

// GetVersion retrieves a specific version from the database.
func GetVersion(c context.Context, mod int, id string) (*model.Version, *gorm.DB) {
	return FromContext(c).GetVersion(mod, id)
}

// GetVersionBuilds retrieves builds for a version.
func GetVersionBuilds(c context.Context, id int) (*model.Builds, error) {
	return FromContext(c).GetVersionBuilds(id)
}

// GetVersionHasBuild checks if a specific build is assigned to a version.
func GetVersionHasBuild(c context.Context, parent, id int) bool {
	return FromContext(c).GetVersionHasBuild(parent, id)
}
