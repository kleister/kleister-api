package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetMods retrieves all available mods from the database.
func (db *data) GetMods() (*model.Mods, error) {
	records := &model.Mods{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Versions",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// CreateMod creates a new mod.
func (db *data) CreateMod(record *model.Mod) error {
	return db.Create(
		record,
	).Error
}

// UpdateMod updates a mod.
func (db *data) UpdateMod(record *model.Mod) error {
	return db.Save(
		record,
	).Error
}

// DeleteMod deletes a mod.
func (db *data) DeleteMod(record *model.Mod) error {
	return db.Delete(
		record,
	).Error
}

// GetMod retrieves a specific mod from the database.
func (db *data) GetMod(id string) (*model.Mod, *gorm.DB) {
	var (
		record = &model.Mod{}
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

	res := query.Model(
		record,
	).Preload(
		"Versions",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).First(
		record,
	)

	return record, res
}

// GetModUsers retrieves users for a mod.
func (db *data) GetModUsers(params *model.ModUserParams) (*model.Users, error) {
	mod, _ := db.GetMod(params.Mod)

	records := &model.Users{}

	err := db.Model(
		mod,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetModHasUser checks if a specific user is assigned to a mod.
func (db *data) GetModHasUser(params *model.ModUserParams) bool {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		mod,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateModUser(params *model.ModUserParams) error {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	return db.Model(
		mod,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteModUser(params *model.ModUserParams) error {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	return db.Model(
		mod,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetModTeams retrieves teams for a mod.
func (db *data) GetModTeams(params *model.ModTeamParams) (*model.Teams, error) {
	mod, _ := db.GetMod(params.Mod)

	records := &model.Teams{}

	err := db.Model(
		mod,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetModHasTeam checks if a specific team is assigned to a mod.
func (db *data) GetModHasTeam(params *model.ModTeamParams) bool {
	mod, _ := db.GetMod(params.Mod)
	team, _ := db.GetTeam(params.Team)

	res := db.Model(
		mod,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

func (db *data) CreateModTeam(params *model.ModTeamParams) error {
	mod, _ := db.GetMod(params.Mod)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		mod,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeleteModTeam(params *model.ModTeamParams) error {
	mod, _ := db.GetMod(params.Mod)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		mod,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
