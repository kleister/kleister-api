package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetBuilds retrieves all available builds from the database.
func GetBuilds(c context.Context, pack int) (*model.Builds, error) {
	return FromContext(c).GetBuilds(pack)
}

// CreateBuild creates a new build.
func CreateBuild(c context.Context, pack int, record *model.Build) error {
	return FromContext(c).CreateBuild(pack, record)
}

// UpdateBuild updates a build.
func UpdateBuild(c context.Context, pack int, record *model.Build) error {
	return FromContext(c).UpdateBuild(pack, record)
}

// DeleteBuild deletes a build.
func DeleteBuild(c context.Context, pack int, record *model.Build) error {
	return FromContext(c).DeleteBuild(pack, record)
}

// GetBuild retrieves a specific build from the database.
func GetBuild(c context.Context, pack int, id string) (*model.Build, *gorm.DB) {
	return FromContext(c).GetBuild(pack, id)
}

// GetBuildVersions retrieves versions for a build.
func GetBuildVersions(c context.Context, id int) (*model.Versions, error) {
	return FromContext(c).GetBuildVersions(id)
}
