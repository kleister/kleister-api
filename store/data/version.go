package data

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
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
func (db *data) GetVersionBuilds(params *model.VersionBuildParams) (*model.Builds, error) {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)

	records := &model.Builds{}

	err := db.Model(
		version,
	).Preload(
		"Pack",
	).Association(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}

// GetVersionHasBuild checks if a specific build is assigned to a version.
func (db *data) GetVersionHasBuild(params *model.VersionBuildParams) bool {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	count := db.Model(
		version,
	).Association(
		"Builds",
	).Find(
		build,
	).Count()

	return count > 0
}

func (db *data) CreateVersionBuild(params *model.VersionBuildParams) error {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return db.Model(
		version,
	).Association(
		"Builds",
	).Append(
		build,
	).Error
}

func (db *data) DeleteVersionBuild(params *model.VersionBuildParams) error {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return db.Model(
		version,
	).Association(
		"Builds",
	).Delete(
		build,
	).Error
}
