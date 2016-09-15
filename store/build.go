package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetBuilds retrieves all available builds from the database.
func GetBuilds(c context.Context, pack int64) (*model.Builds, error) {
	return FromContext(c).GetBuilds(pack)
}

// CreateBuild creates a new build.
func CreateBuild(c context.Context, pack int64, record *model.Build) error {
	return FromContext(c).CreateBuild(pack, record, Current(c))
}

// UpdateBuild updates a build.
func UpdateBuild(c context.Context, pack int64, record *model.Build) error {
	return FromContext(c).UpdateBuild(pack, record, Current(c))
}

// DeleteBuild deletes a build.
func DeleteBuild(c context.Context, pack int64, record *model.Build) error {
	return FromContext(c).DeleteBuild(pack, record, Current(c))
}

// GetBuild retrieves a specific build from the database.
func GetBuild(c context.Context, pack int64, id string) (*model.Build, *gorm.DB) {
	return FromContext(c).GetBuild(pack, id)
}

// GetBuildVersions retrieves versions for a build.
func GetBuildVersions(c context.Context, params *model.BuildVersionParams) (*model.BuildVersions, error) {
	return FromContext(c).GetBuildVersions(params)
}

// GetBuildHasVersion checks if a specific version is assigned to a build.
func GetBuildHasVersion(c context.Context, params *model.BuildVersionParams) bool {
	return FromContext(c).GetBuildHasVersion(params)
}

// CreateBuildVersion assigns a version to a specific build.
func CreateBuildVersion(c context.Context, params *model.BuildVersionParams) error {
	return FromContext(c).CreateBuildVersion(params, Current(c))
}

// DeleteBuildVersion removes a version from a specific build.
func DeleteBuildVersion(c context.Context, params *model.BuildVersionParams) error {
	return FromContext(c).DeleteBuildVersion(params, Current(c))
}
