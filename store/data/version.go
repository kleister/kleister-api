package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
)

// GetVersions retrieves all available versions from the database.
func (db *data) GetVersions(mod int) (*model.Versions, error) {
	records := &model.Versions{}

	err := db.Order(
		"name ASC",
	).Where(
		"mod_id = ?",
		mod,
	).Preload(
		"Mod",
	).Preload(
		"File",
	).Find(
		records,
	).Error

	return records, err
}

// CreateVersion creates a new version.
func (db *data) CreateVersion(mod int, record *model.Version) error {
	record.ModID = mod

	return db.Create(
		record,
	).Error
}

// UpdateVersion updates a version.
func (db *data) UpdateVersion(mod int, record *model.Version) error {
	record.ModID = mod

	return db.Save(
		record,
	).Error
}

// DeleteVersion deletes a version.
func (db *data) DeleteVersion(mod int, record *model.Version) error {
	// tx := context.Store(c).Begin()
	// failed := false

	// if record.File != nil {
	// 	record.File.Path = config.Server.Storage

	// 	err := tx.Delete(
	// 		&record.File,
	// 	).Error

	// 	if err != nil {
	// 		failed = true
	// 	}
	// }

	// err := tx.Delete(
	// 	&record,
	// ).Error

	// if failed || err != nil {
	// 	tx.Rollback()

	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  http.StatusBadRequest,
	// 			"message": "Failed to delete version",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// tx.Commit()

	record.ModID = mod

	return db.Delete(
		record,
	).Error
}

// GetVersion retrieves a specific version from the database.
func (db *data) GetVersion(mod int, id string) (*model.Version, *gorm.DB) {
	record := &model.Version{}

	res := db.Where(
		"mod_id = ?",
		mod,
	).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		record,
	).Preload(
		"Mod",
	).Preload(
		"File",
	).First(
		record,
	)

	return record, res
}

// GetVersionBuilds retrieves builds for a version.
func (db *data) GetVersionBuilds(id int) (*model.Builds, error) {
	records := &model.Builds{}

	err := db.Model(
		&model.Version{
			ID: id,
		},
	).Association(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}
