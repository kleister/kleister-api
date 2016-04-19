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
func GetMinecraftBuilds(c context.Context, id int) (*model.Builds, error) {
	return FromContext(c).GetMinecraftBuilds(id)
}

// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
func GetMinecraftHasBuild(c context.Context, parent, id int) bool {
	return FromContext(c).GetMinecraftHasBuild(parent, id)
}

func CreateMinecraftBuild(c context.Context, parent, id int) error {
	return FromContext(c).CreateMinecraftBuild(parent, id)
}

func DeleteMinecraftBuild(c context.Context, parent, id int) error {
	return FromContext(c).DeleteMinecraftBuild(parent, id)
}
