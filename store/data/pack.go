package data

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetPacks retrieves all available packs from the database.
func (db *data) GetPacks() (*model.Packs, error) {
	records := &model.Packs{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Clients",
	).Preload(
		"Users",
	).Preload(
		"Teams",
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
	).Preload(
		"Builds.Forge",
	).Preload(
		"Builds.Minecraft",
	).Find(
		records,
	).Error

	return records, err
}

// CreatePack creates a new pack.
func (db *data) CreatePack(record *model.Pack, current *model.User) error {
	record.UserPacks = model.UserPacks{
		&model.UserPack{
			UserID: current.ID,
			Perm:   "owner",
		},
	}

	return db.Create(
		record,
	).Error
}

// UpdatePack updates a pack.
func (db *data) UpdatePack(record *model.Pack, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeletePack deletes a pack.
func (db *data) DeletePack(record *model.Pack, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetPack retrieves a specific pack from the database.
func (db *data) GetPack(id string) (*model.Pack, *gorm.DB) {
	var (
		record = &model.Pack{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Pack{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.Pack{
				Slug: id,
			},
		)
	}

	res := query.Model(
		record,
	).Preload(
		"Clients",
	).Preload(
		"Users",
	).Preload(
		"Teams",
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
	).Preload(
		"Builds.Forge",
	).Preload(
		"Builds.Minecraft",
	).First(
		record,
	)

	return record, res
}

// GetPackClients retrieves clients for a pack.
func (db *data) GetPackClients(params *model.PackClientParams) (*model.ClientPacks, error) {
	pack, _ := db.GetPack(params.Pack)
	records := &model.ClientPacks{}

	err := db.Where(
		&model.ClientPack{
			PackID: pack.ID,
		},
	).Model(
		&model.ClientPack{},
	).Preload(
		"Client",
	).Preload(
		"Pack",
	).Find(
		records,
	).Error

	return records, err
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func (db *data) GetPackHasClient(params *model.PackClientParams) bool {
	pack, _ := db.GetPack(params.Pack)
	client, _ := db.GetClient(params.Client)

	res := db.Model(
		pack,
	).Association(
		"Clients",
	).Find(
		client,
	).Error

	return res == nil
}

func (db *data) CreatePackClient(params *model.PackClientParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	client, _ := db.GetClient(params.Client)

	return db.Create(
		&model.ClientPack{
			PackID:   pack.ID,
			ClientID: client.ID,
		},
	).Error
}

func (db *data) DeletePackClient(params *model.PackClientParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	client, _ := db.GetClient(params.Client)

	return db.Model(
		pack,
	).Association(
		"Clients",
	).Delete(
		client,
	).Error
}

// GetPackUsers retrieves users for a pack.
func (db *data) GetPackUsers(params *model.PackUserParams) (*model.UserPacks, error) {
	pack, _ := db.GetPack(params.Pack)
	records := &model.UserPacks{}

	err := db.Where(
		&model.UserPack{
			PackID: pack.ID,
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

// GetPackHasUser checks if a specific user is assigned to a pack.
func (db *data) GetPackHasUser(params *model.PackUserParams) bool {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		pack,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreatePackUser(params *model.PackUserParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.UserPack{
					PackID: pack.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdatePackUser(params *model.PackUserParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	return db.Model(
		&model.UserPack{},
	).Where(
		&model.UserPack{
			PackID: pack.ID,
			UserID: user.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeletePackUser(params *model.PackUserParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	return db.Model(
		pack,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetPackTeams retrieves teams for a pack.
func (db *data) GetPackTeams(params *model.PackTeamParams) (*model.TeamPacks, error) {
	pack, _ := db.GetPack(params.Pack)
	records := &model.TeamPacks{}

	err := db.Where(
		&model.TeamPack{
			PackID: pack.ID,
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

// GetPackHasTeam checks if a specific team is assigned to a pack.
func (db *data) GetPackHasTeam(params *model.PackTeamParams) bool {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	res := db.Model(
		pack,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

func (db *data) CreatePackTeam(params *model.PackTeamParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamPack{
					PackID: pack.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdatePackTeam(params *model.PackTeamParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		&model.TeamPack{},
	).Where(
		&model.TeamPack{
			PackID: pack.ID,
			TeamID: team.ID,
		},
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeletePackTeam(params *model.PackTeamParams, current *model.User) error {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		pack,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
