package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/forge"
)

// GetForges retrieves all available forges from the database.
func (db *data) GetForges() (*model.Forges, error) {
	records := &model.Forges{}

	err := db.Order(
		"minecraft DESC",
	).Order(
		"name DESC",
	).Find(
		records,
	).Error

	return records, err
}

// SyncForge creates or updates a forge record.
func (db *data) SyncForge(number *forge.Number) (*model.Forge, error) {
	record := &model.Forge{}

	err := db.Where(
		model.Forge{
			Name: number.ID,
		},
	).Attrs(
		model.Forge{
			Minecraft: number.Minecraft,
		},
	).FirstOrCreate(
		&record,
	).Error

	return record, err
}

// GetForge retrieves a specific forge from the database.
func (db *data) GetForge(id string) (*model.Forge, *gorm.DB) {
	record := &model.Forge{}

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

// GetForgeBuilds retrieves builds for a forge.
func (db *data) GetForgeBuilds(id int) (*model.Builds, error) {
	records := &model.Builds{}

	err := db.Model(
		&model.Forge{
			ID: id,
		},
	).Association(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}

// GetForgeHasBuild checks if a specific build is assigned to a minecraft.
func (db *data) GetForgeHasBuild(parent, id int) bool {
	record := &model.Build{
		ID: id,
	}

	count := db.Model(
		&model.Forge{
			ID: parent,
		},
	).Association(
		"Builds",
	).Find(
		record,
	).Count()

	return count > 0
}

func (db *data) CreateForgeBuild(parent, id int) error {
	return db.Model(
		&model.Build{},
	).Where(
		"id = ?",
		parent,
	).Update(
		"forge_id",
		id,
	).Error
}

func (db *data) DeleteForgeBuild(parent, id int) error {
	return db.Model(
		&model.Build{},
	).Where(
		"id = ?",
		parent,
	).Update(
		"forge_id",
		0,
	).Error
}
