package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/forge"
	"golang.org/x/net/context"
)

// GetForges retrieves all available forges from the database.
func GetForges(c context.Context) (*model.Forges, error) {
	return FromContext(c).GetForges()
}

// SyncForge creates or updates a forge record.
func SyncForge(c context.Context, number *forge.Number) (*model.Forge, error) {
	return FromContext(c).SyncForge(number)
}

// GetForge retrieves a specific forge from the database.
func GetForge(c context.Context, id string) (*model.Forge, *gorm.DB) {
	return FromContext(c).GetForge(id)
}

// GetForgeBuilds retrieves builds for a forge.
func GetForgeBuilds(c context.Context, id int) (*model.Builds, error) {
	return FromContext(c).GetForgeBuilds(id)
}

// GetForgeHasBuild checks if a specific build is assigned to a forge.
func GetForgeHasBuild(c context.Context, parent, id int) bool {
	return FromContext(c).GetForgeHasBuild(parent, id)
}
