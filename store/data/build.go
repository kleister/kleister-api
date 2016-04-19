package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
)

// GetBuilds retrieves all available builds from the database.
func (db *data) GetBuilds(pack int) (*model.Builds, error) {
	records := &model.Builds{}

	err := db.Order(
		"name ASC",
	).Where(
		"pack_id = ?",
		pack,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).Find(
		records,
	).Error

	return records, err
}

// CreateBuild creates a new build.
func (db *data) CreateBuild(pack int, record *model.Build) error {
	record.PackID = pack

	return db.Create(
		record,
	).Error
}

// UpdateBuild updates a build.
func (db *data) UpdateBuild(pack int, record *model.Build) error {
	record.PackID = pack

	return db.Save(
		record,
	).Error
}

// DeleteBuild deletes a build.
func (db *data) DeleteBuild(pack int, record *model.Build) error {
	record.PackID = pack

	return db.Delete(
		record,
	).Error
}

// GetBuild retrieves a specific build from the database.
func (db *data) GetBuild(pack int, id string) (*model.Build, *gorm.DB) {
	record := &model.Build{}

	res := db.Where(
		"pack_id = ?",
		pack,
	).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		record,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).First(
		record,
	)

	return record, res
}

// GetBuildVersions retrieves versions for a build.
func (db *data) GetBuildVersions(id int) (*model.Versions, error) {
	records := &model.Versions{}

	err := db.Model(
		&model.Build{
			ID: id,
		},
	).Association(
		"Versions",
	).Find(
		records,
	).Error

	return records, err
}

// GetBuildHasVersion checks if a specific version is assigned to a build.
func (db *data) GetBuildHasVersion(parent, id int) bool {
	record := &model.Version{
		ID: id,
	}

	count := db.Model(
		&model.Build{
			ID: parent,
		},
	).Association(
		"Versions",
	).Find(
		record,
	).Count()

	return count > 0
}

func (db *data) CreateBuildVersion(parent, id int) error {
	return db.Model(
		&model.Build{
			ID: parent,
		},
	).Association(
		"Versions",
	).Append(
		&model.Version{
			ID: id,
		},
	).Error
}

func (db *data) DeleteBuildVersion(parent, id int) error {
	return db.Model(
		&model.Build{
			ID: parent,
		},
	).Association(
		"Versions",
	).Delete(
		&model.Version{
			ID: id,
		},
	).Error
}
