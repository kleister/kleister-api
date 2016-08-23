package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/forge"
	"golang.org/x/net/context"
)

// GetForges retrieves all available forges from the database.
func GetForges(c context.Context) (*model.Forges, error) {
	return FromContext(c).GetForges()
}

// SyncForge creates or updates a forge record.
func SyncForge(c context.Context, number *forge.Number) (*model.Forge, error) {
	return FromContext(c).SyncForge(number, Current(c))
}

// GetForge retrieves a specific forge from the database.
func GetForge(c context.Context, id string) (*model.Forge, *gorm.DB) {
	return FromContext(c).GetForge(id)
}

// GetForgeBuilds retrieves builds for a forge.
func GetForgeBuilds(c context.Context, params *model.ForgeBuildParams) (*model.Builds, error) {
	return FromContext(c).GetForgeBuilds(params)
}

// GetForgeHasBuild checks if a specific build is assigned to a forge.
func GetForgeHasBuild(c context.Context, params *model.ForgeBuildParams) bool {
	return FromContext(c).GetForgeHasBuild(params)
}

// CreateForgeBuild assigns a build to a specific forge.
func CreateForgeBuild(c context.Context, params *model.ForgeBuildParams) error {
	return FromContext(c).CreateForgeBuild(params, Current(c))
}

// DeleteForgeBuild removes a build from a specific forge.
func DeleteForgeBuild(c context.Context, params *model.ForgeBuildParams) error {
	return FromContext(c).DeleteForgeBuild(params, Current(c))
}
