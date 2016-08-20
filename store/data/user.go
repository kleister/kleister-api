package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetUsers retrieves all available users from the database.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.Order(
		"username ASC",
	).Find(
		records,
	).Error

	return records, err
}

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetUser retrieves a specific user from the database.
func (db *data) GetUser(id string) (*model.User, *gorm.DB) {
	var (
		record = &model.User{}
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
	).First(
		record,
	)

	return record, res
}

// GetUserMods retrieves mods for a user.
func (db *data) GetUserMods(params *model.UserModParams) (*model.Mods, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Mods{}

	err := db.Model(
		user,
	).Association(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func (db *data) GetUserHasMod(params *model.UserModParams) bool {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	res := db.Model(
		user,
	).Association(
		"Mods",
	).Find(
		mod,
	).Error

	return res == nil
}

func (db *data) CreateUserMod(params *model.UserModParams) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		user,
	).Association(
		"Mods",
	).Append(
		mod,
	).Error
}

func (db *data) DeleteUserMod(params *model.UserModParams) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		user,
	).Association(
		"Mods",
	).Delete(
		mod,
	).Error
}

// GetUserPacks retrieves packs for a user.
func (db *data) GetUserPacks(params *model.UserPackParams) (*model.Packs, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Packs{}

	err := db.Model(
		user,
	).Association(
		"Packs",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasPack checks if a specific pack is assigned to a user.
func (db *data) GetUserHasPack(params *model.UserPackParams) bool {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	res := db.Model(
		user,
	).Association(
		"Packs",
	).Find(
		pack,
	).Error

	return res == nil
}

func (db *data) CreateUserPack(params *model.UserPackParams) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		user,
	).Association(
		"Packs",
	).Append(
		pack,
	).Error
}

func (db *data) DeleteUserPack(params *model.UserPackParams) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		user,
	).Association(
		"Packs",
	).Delete(
		pack,
	).Error
}

// GetUserTeams retrieves teams for a user.
func (db *data) GetUserTeams(params *model.UserTeamParams) (*model.Teams, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Teams{}

	err := db.Model(
		user,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func (db *data) GetUserHasTeam(params *model.UserTeamParams) bool {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	res := db.Model(
		user,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

func (db *data) CreateUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		user,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeleteUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		user,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
