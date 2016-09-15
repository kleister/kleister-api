package data

import (
	"fmt"
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
	).Preload(
		"Teams",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User, current *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User, current *model.User) error {
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
			&model.User{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.User{
				Slug: id,
			},
		)
	}

	res := query.Model(
		record,
	).Preload(
		"Teams",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).First(
		record,
	)

	return record, res
}

// GetUserMods retrieves mods for a user.
func (db *data) GetUserMods(params *model.UserModParams) (*model.UserMods, error) {
	user, _ := db.GetUser(params.User)
	records := &model.UserMods{}

	err := db.Where(
		&model.UserMod{
			UserID: user.ID,
		},
	).Model(
		&model.UserMod{},
	).Preload(
		"User",
	).Preload(
		"Mod",
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

func (db *data) CreateUserMod(params *model.UserModParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.UserMod{
					UserID: user.ID,
					ModID:  mod.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateUserMod(params *model.UserModParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		&model.UserMod{},
	).Where(
		&model.UserMod{
			UserID: user.ID,
			ModID:  mod.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteUserMod(params *model.UserModParams, current *model.User) error {
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
func (db *data) GetUserPacks(params *model.UserPackParams) (*model.UserPacks, error) {
	user, _ := db.GetUser(params.User)
	records := &model.UserPacks{}

	err := db.Where(
		&model.UserPack{
			UserID: user.ID,
		},
	).Model(
		&model.UserPack{},
	).Preload(
		"User",
	).Preload(
		"Pack",
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

func (db *data) CreateUserPack(params *model.UserPackParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.UserPack{
					UserID: user.ID,
					PackID: pack.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateUserPack(params *model.UserPackParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		&model.UserPack{},
	).Where(
		&model.UserPack{
			UserID: user.ID,
			PackID: pack.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteUserPack(params *model.UserPackParams, current *model.User) error {
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
func (db *data) GetUserTeams(params *model.UserTeamParams) (*model.TeamUsers, error) {
	user, _ := db.GetUser(params.User)
	records := &model.TeamUsers{}

	err := db.Where(
		&model.TeamUser{
			UserID: user.ID,
		},
	).Model(
		&model.TeamUser{},
	).Preload(
		"User",
	).Preload(
		"Team",
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

func (db *data) CreateUserTeam(params *model.UserTeamParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamUser{
					UserID: user.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateUserTeam(params *model.UserTeamParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		&model.TeamUser{},
	).Where(
		&model.TeamUser{
			UserID: user.ID,
			TeamID: team.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteUserTeam(params *model.UserTeamParams, current *model.User) error {
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
