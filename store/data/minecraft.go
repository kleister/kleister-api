package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/minecraft"
)

// GetMinecrafts retrieves all available minecrafts from the database.
func (db *data) GetMinecrafts() (*model.Minecrafts, error) {
	records := &model.Minecrafts{}

	err := db.Order(
		"type DESC",
	).Order(
		"name DESC",
	).Find(
		records,
	).Error

	return records, err
}

// SyncMinecraft creates or updates a minecraft record.
func (db *data) SyncMinecraft(version *minecraft.Version, current *model.User) (*model.Minecraft, error) {
	record := &model.Minecraft{}

	err := db.Where(
		&model.Minecraft{
			Name: version.ID,
		},
	).Attrs(
		&model.Minecraft{
			Type: version.Type,
		},
	).FirstOrCreate(
		record,
	).Error

	return record, err
}

// GetMinecraft retrieves a specific minecraft from the database.
func (db *data) GetMinecraft(id string) (*model.Minecraft, *gorm.DB) {
	var (
		record = &model.Minecraft{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Minecraft{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.Minecraft{
				Slug: id,
			},
		)
	}

	res := query.First(
		record,
	)

	return record, res
}

// GetMinecraftBuilds retrieves builds for a minecraft.
func (db *data) GetMinecraftBuilds(params *model.MinecraftBuildParams) (*model.Builds, error) {
	minecraft, _ := db.GetMinecraft(params.Minecraft)

	records := &model.Builds{}

	err := db.Model(
		minecraft,
	).Preload(
		"Pack",
	).Association(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}

// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
func (db *data) GetMinecraftHasBuild(params *model.MinecraftBuildParams) bool {
	minecraft, _ := db.GetMinecraft(params.Minecraft)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return build.MinecraftID.Int64 == minecraft.ID
}

func (db *data) CreateMinecraftBuild(params *model.MinecraftBuildParams, current *model.User) error {
	minecraft, _ := db.GetMinecraft(params.Minecraft)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return db.Model(
		build,
	).Update(
		"minecraft_id",
		minecraft.ID,
	).Error
}

func (db *data) DeleteMinecraftBuild(params *model.MinecraftBuildParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return db.Model(
		build,
	).Update(
		"minecraft_id",
		0,
	).Error
}
