package data

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetTeams retrieves all available teams from the database.
func (db *data) GetTeams() (*model.Teams, error) {
	records := &model.Teams{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Users",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// CreateTeam creates a new team.
func (db *data) CreateTeam(record *model.Team, current *model.User) error {
	record.TeamUsers = model.TeamUsers{
		&model.TeamUser{
			UserID: current.ID,
			Perm:   "owner",
		},
	}

	return db.Create(
		record,
	).Error
}

// UpdateTeam updates a team.
func (db *data) UpdateTeam(record *model.Team, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteTeam deletes a team.
func (db *data) DeleteTeam(record *model.Team, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetTeam retrieves a specific team from the database.
func (db *data) GetTeam(id string) (*model.Team, *gorm.DB) {
	var (
		record = &model.Team{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Team{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.Team{
				Slug: id,
			},
		)
	}

	res := query.Model(
		record,
	).Preload(
		"Users",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).First(
		record,
	)

	return record, res
}

// GetTeamUsers retrieves users for a team.
func (db *data) GetTeamUsers(params *model.TeamUserParams) (*model.TeamUsers, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamUsers{}

	err := db.Where(
		&model.TeamUser{
			TeamID: team.ID,
		},
	).Model(
		&model.TeamUser{},
	).Preload(
		"Team",
	).Preload(
		"User",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func (db *data) GetTeamHasUser(params *model.TeamUserParams) bool {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamUser{
					TeamID: team.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		&model.TeamUser{},
	).Where(
		&model.TeamUser{
			TeamID: team.ID,
			UserID: user.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetTeamPacks retrieves packs for a team.
func (db *data) GetTeamPacks(params *model.TeamPackParams) (*model.TeamPacks, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamPacks{}

	err := db.Where(
		&model.TeamPack{
			TeamID: team.ID,
		},
	).Model(
		&model.TeamPack{},
	).Preload(
		"Team",
	).Preload(
		"Pack",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasPack checks if a specific pack is assigned to a team.
func (db *data) GetTeamHasPack(params *model.TeamPackParams) bool {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	res := db.Model(
		team,
	).Association(
		"Packs",
	).Find(
		pack,
	).Error

	return res == nil
}

func (db *data) CreateTeamPack(params *model.TeamPackParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamPack{
					TeamID: team.ID,
					PackID: pack.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateTeamPack(params *model.TeamPackParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		&model.TeamPack{},
	).Where(
		&model.TeamPack{
			TeamID: team.ID,
			PackID: pack.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteTeamPack(params *model.TeamPackParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		team,
	).Association(
		"Packs",
	).Delete(
		pack,
	).Error
}

// GetTeamMods retrieves mods for a team.
func (db *data) GetTeamMods(params *model.TeamModParams) (*model.TeamMods, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamMods{}

	err := db.Where(
		&model.TeamMod{
			TeamID: team.ID,
		},
	).Model(
		&model.TeamMod{},
	).Preload(
		"Team",
	).Preload(
		"Mod",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasMod checks if a specific mod is assigned to a team.
func (db *data) GetTeamHasMod(params *model.TeamModParams) bool {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	res := db.Model(
		team,
	).Association(
		"Mods",
	).Find(
		mod,
	).Error

	return res == nil
}

func (db *data) CreateTeamMod(params *model.TeamModParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamMod{
					TeamID: team.ID,
					ModID:  mod.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateTeamMod(params *model.TeamModParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		&model.TeamMod{},
	).Where(
		&model.TeamMod{
			TeamID: team.ID,
			ModID:  mod.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteTeamMod(params *model.TeamModParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		team,
	).Association(
		"Mods",
	).Delete(
		mod,
	).Error
}
