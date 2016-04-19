package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/minecraft"
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
func (db *data) SyncMinecraft(version *minecraft.Version) (*model.Minecraft, error) {
	record := &model.Minecraft{}

	err := db.Where(
		model.Minecraft{
			Name: version.ID,
		},
	).Attrs(
		model.Minecraft{
			Type: version.Type,
		},
	).FirstOrCreate(
		&record,
	).Error

	return record, err
}

// GetMinecraft retrieves a specific minecraft from the database.
func (db *data) GetMinecraft(id string) (*model.Minecraft, *gorm.DB) {
	record := &model.Minecraft{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		record,
	)

	return record, res
}

// GetMinecraftBuilds retrieves builds for a minecraft.
func (db *data) GetMinecraftBuilds(id int) (*model.Builds, error) {
	records := &model.Builds{}

	err := db.Model(
		&model.Minecraft{
			ID: id,
		},
	).Association(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}

// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
func (db *data) GetMinecraftHasBuild(parent, id int) bool {
	record := &model.Build{
		ID: id,
	}

	count := db.Model(
		&model.Minecraft{
			ID: parent,
		},
	).Association(
		"Builds",
	).Find(
		record,
	).Count()

	return count > 0
}

func (db *data) CreateMinecraftBuild(parent, id int) error {
	return db.Model(
		&model.Build{},
	).Where(
		"id = ?",
		parent,
	).Update(
		"minecraft_id",
		id,
	).Error
}

func (db *data) DeleteMinecraftBuild(parent, id int) error {
	return db.Model(
		&model.Build{},
	).Where(
		"id = ?",
		parent,
	).Update(
		"minecraft_id",
		0,
	).Error
}
