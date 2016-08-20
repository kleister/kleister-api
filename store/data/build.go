package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
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
	).Preload(
		"Versions",
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
	var (
		record = &model.Build{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			"id = ?",
			val,
		)
	} else {
		query = db.Where(
			"slug = ?",
			id,
		)
	}

	res := query.Where(
		"pack_id = ?",
		pack,
	).Model(
		record,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).Preload(
		"Versions",
	).First(
		record,
	)

	return record, res
}

// GetBuildVersions retrieves versions for a build.
func (db *data) GetBuildVersions(params *model.BuildVersionParams) (*model.Versions, error) {
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	records := &model.Versions{}

	err := db.Model(
		build,
	).Preload(
		"Mod",
	).Association(
		"Versions",
	).Find(
		records,
	).Error

	return records, err
}

// GetBuildHasVersion checks if a specific version is assigned to a build.
func (db *data) GetBuildHasVersion(params *model.BuildVersionParams) bool {
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)

	res := db.Model(
		build,
	).Association(
		"Versions",
	).Find(
		version,
	).Error

	return res == nil
}

func (db *data) CreateBuildVersion(params *model.BuildVersionParams) error {
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)

	return db.Model(
		build,
	).Association(
		"Versions",
	).Append(
		version,
	).Error
}

func (db *data) DeleteBuildVersion(params *model.BuildVersionParams) error {
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)

	return db.Model(
		build,
	).Association(
		"Versions",
	).Delete(
		version,
	).Error
}
