package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/minecraft"
	"golang.org/x/net/context"
)

// GetMinecrafts retrieves all available minecrafts from the database.
func GetMinecrafts(c context.Context) (*model.Minecrafts, error) {
	return FromContext(c).GetMinecrafts()
}

// SyncMinecraft creates or updates a minecraft record.
func SyncMinecraft(c context.Context, number *minecraft.Version) (*model.Minecraft, error) {
	return FromContext(c).SyncMinecraft(number)
}

// GetMinecraft retrieves a specific minecraft from the database.
func GetMinecraft(c context.Context, id string) (*model.Minecraft, *gorm.DB) {
	return FromContext(c).GetMinecraft(id)
}

// GetMinecraftBuilds retrieves builds for a minecraft.
func GetMinecraftBuilds(c context.Context, params *model.MinecraftBuildParams) (*model.Builds, error) {
	return FromContext(c).GetMinecraftBuilds(params)
}

// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
func GetMinecraftHasBuild(c context.Context, params *model.MinecraftBuildParams) bool {
	return FromContext(c).GetMinecraftHasBuild(params)
}

// CreateMinecraftBuild assigns a build to a specific minecraft.
func CreateMinecraftBuild(c context.Context, params *model.MinecraftBuildParams) error {
	return FromContext(c).CreateMinecraftBuild(params)
}

// DeleteMinecraftBuild removes a build from a specific minecraft.
func DeleteMinecraftBuild(c context.Context, params *model.MinecraftBuildParams) error {
	return FromContext(c).DeleteMinecraftBuild(params)
}
