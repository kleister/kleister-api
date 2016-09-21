package data

import (
	"github.com/kleister/kleister-api/model"
)

func (db *data) GetSolderPacks() (*model.Packs, error) {
	records := &model.Packs{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Builds",
	).Preload(
		"Icon",
	).Preload(
		"Background",
	).Preload(
		"Logo",
	).Preload(
		"Recommended",
	).Preload(
		"Latest",
	).Find(
		records,
	).Error

	return records, err
}

func (db *data) GetSolderPack(pack string) (*model.Pack, error) {
	record := &model.Pack{}

	err := db.Where(
		&model.Pack{
			Slug: pack,
		},
	).Model(
		record,
	).Preload(
		"Builds",
	).Preload(
		"Icon",
	).Preload(
		"Background",
	).Preload(
		"Logo",
	).Preload(
		"Recommended",
	).Preload(
		"Latest",
	).First(
		record,
	).Error

	return record, err
}

func (db *data) GetSolderBuild(pack, build string) (*model.Build, error) {
	record := &model.Build{}

	err := db.Where(
		&model.Build{
			Slug: build,
		},
	).Model(
		record,
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).Preload(
		"Versions",
	).Preload(
		"Versions.Mod",
	).Preload(
		"Versions.File",
	).Joins(
		"INNER JOIN packs ON packs.id = builds.pack_id AND packs.slug = ?",
		pack,
	).First(
		record,
	).Error

	return record, err
}
