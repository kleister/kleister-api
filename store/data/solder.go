package data

import (
	"strconv"
	"strings"

	"github.com/kleister/kleister-api/model"
)

func (db *data) GetSolderPacks() (*model.Packs, error) {
	records := &model.Packs{}

	err := db.Order(
		"name ASC",
	).Find(
		records,
	).Error

	return records, err
}

func (db *data) GetSolderPack(pack, location string) (*model.Pack, error) {
	record := &model.Pack{}

	err := db.Where(
		"packs.slug = ?",
		pack,
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

	if record.Logo != nil {
		record.Logo.URL = strings.Join(
			[]string{
				location,
				"storage",
				"logo",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	if record.Background != nil {
		record.Background.URL = strings.Join(
			[]string{
				location,
				"storage",
				"background",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	if record.Icon != nil {
		record.Icon.URL = strings.Join(
			[]string{
				location,
				"storage",
				"icon",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	return record, err
}

func (db *data) GetSolderBuild(pack, build, location string) (*model.Build, error) {
	record := &model.Build{}

	err := db.Where(
		"builds.slug = ?",
		build,
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

	for _, version := range record.Versions {
		if version.File != nil {
			version.File.URL = strings.Join(
				[]string{
					location,
					"storage",
					"version",
					strconv.Itoa(version.ID),
				},
				"/",
			)
		}
	}

	return record, err
}
