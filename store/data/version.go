package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetVersions retrieves all available versions from the database.
func (db *data) GetVersions(mod int64) (*model.Versions, error) {
	records := &model.Versions{}

	err := db.Order(
		"name ASC",
	).Where(
		&model.Version{
			ModID: mod,
		},
	).Preload(
		"Mod",
	).Preload(
		"File",
	).Preload(
		"Builds",
	).Find(
		records,
	).Error

	return records, err
}

// CreateVersion creates a new version.
func (db *data) CreateVersion(mod int64, record *model.Version, current *model.User) error {
	record.ModID = mod

	return db.Create(
		record,
	).Error
}

// UpdateVersion updates a version.
func (db *data) UpdateVersion(mod int64, record *model.Version, current *model.User) error {
	record.ModID = mod

	return db.Save(
		record,
	).Error
}

// DeleteVersion deletes a version.
func (db *data) DeleteVersion(mod int64, record *model.Version, current *model.User) error {
	record.ModID = mod

	return db.Delete(
		record,
	).Error
}

// GetVersion retrieves a specific version from the database.
func (db *data) GetVersion(mod int64, id string) (*model.Version, *gorm.DB) {
	var (
		record = &model.Version{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Version{
				ID:    val,
				ModID: mod,
			},
		)
	} else {
		query = db.Where(
			&model.Version{
				Slug:  id,
				ModID: mod,
			},
		)
	}

	res := query.Model(
		record,
	).Preload(
		"Mod",
	).Preload(
		"File",
	).Preload(
		"Builds",
	).Preload(
		"Builds.Pack",
	).First(
		record,
	)

	return record, res
}

// GetVersionBuilds retrieves builds for a version.
func (db *data) GetVersionBuilds(params *model.VersionBuildParams) (*model.BuildVersions, error) {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)
	records := &model.BuildVersions{}

	err := db.Where(
		&model.BuildVersion{
			VersionID: version.ID,
		},
	).Model(
		&model.BuildVersion{},
	).Preload(
		"Build",
	).Preload(
		"Build.Pack",
	).Preload(
		"Version",
	).Preload(
		"Version.Mod",
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

	res := db.Model(
		version,
	).Association(
		"Builds",
	).Find(
		build,
	).Error

	return res == nil
}

func (db *data) CreateVersionBuild(params *model.VersionBuildParams, current *model.User) error {
	mod, _ := db.GetMod(params.Mod)
	version, _ := db.GetVersion(mod.ID, params.Version)
	pack, _ := db.GetPack(params.Pack)
	build, _ := db.GetBuild(pack.ID, params.Build)

	return db.Create(
		&model.BuildVersion{
			VersionID: version.ID,
			BuildID:   build.ID,
		},
	).Error
}

func (db *data) DeleteVersionBuild(params *model.VersionBuildParams, current *model.User) error {
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
