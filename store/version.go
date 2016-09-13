package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetVersions retrieves all available versions from the database.
func GetVersions(c context.Context, mod int64) (*model.Versions, error) {
	return FromContext(c).GetVersions(mod)
}

// CreateVersion creates a new version.
func CreateVersion(c context.Context, mod int64, record *model.Version) error {
	return FromContext(c).CreateVersion(mod, record, Current(c))
}

// UpdateVersion updates a version.
func UpdateVersion(c context.Context, mod int64, record *model.Version) error {
	return FromContext(c).UpdateVersion(mod, record, Current(c))
}

// DeleteVersion deletes a version.
func DeleteVersion(c context.Context, mod int64, record *model.Version) error {
	return FromContext(c).DeleteVersion(mod, record, Current(c))
}

// GetVersion retrieves a specific version from the database.
func GetVersion(c context.Context, mod int64, id string) (*model.Version, *gorm.DB) {
	return FromContext(c).GetVersion(mod, id)
}

// GetVersionBuilds retrieves builds for a version.
func GetVersionBuilds(c context.Context, params *model.VersionBuildParams) (*model.BuildVersions, error) {
	return FromContext(c).GetVersionBuilds(params)
}

// GetVersionHasBuild checks if a specific build is assigned to a version.
func GetVersionHasBuild(c context.Context, params *model.VersionBuildParams) bool {
	return FromContext(c).GetVersionHasBuild(params)
}

// CreateVersionBuild assigns a build to a specific version.
func CreateVersionBuild(c context.Context, params *model.VersionBuildParams) error {
	return FromContext(c).CreateVersionBuild(params, Current(c))
}

// DeleteVersionBuild removes a build from a specific version.
func DeleteVersionBuild(c context.Context, params *model.VersionBuildParams) error {
	return FromContext(c).DeleteVersionBuild(params, Current(c))
}
