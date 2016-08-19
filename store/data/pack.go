package data

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetPacks retrieves all available packs from the database.
func (db *data) GetPacks() (*model.Packs, error) {
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
		"Builds.Forge",
	).Preload(
		"Builds.Minecraft",
	).Find(
		records,
	).Error

	return records, err
}

// CreatePack creates a new pack.
func (db *data) CreatePack(record *model.Pack) error {
	return db.Create(
		record,
	).Error
}

// UpdatePack updates a pack.
func (db *data) UpdatePack(record *model.Pack) error {
	return db.Save(
		record,
	).Error
}

// DeletePack deletes a pack.
func (db *data) DeletePack(record *model.Pack) error {
	return db.Delete(
		record,
	).Error
}

// GetPack retrieves a specific pack from the database.
func (db *data) GetPack(id string) (*model.Pack, *gorm.DB) {
	record := &model.Pack{}

	res := db.Where(
		"packs.id = ?",
		id,
	).Or(
		"packs.slug = ?",
		id,
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
	).First(
		record,
	)

	return record, res
}

// GetPackClients retrieves clients for a pack.
func (db *data) GetPackClients(params *model.PackClientParams) (*model.Clients, error) {
	pack, _ := db.GetPack(params.Pack)

	records := &model.Clients{}

	err := db.Model(
		pack,
	).Association(
		"Clients",
	).Find(
		records,
	).Error

	return records, err
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func (db *data) GetPackHasClient(params *model.PackClientParams) bool {
	pack, _ := db.GetPack(params.Pack)
	client, _ := db.GetClient(params.Client)

	count := db.Model(
		pack,
	).Association(
		"Clients",
	).Find(
		client,
	).Count()

	return count > 0
}

func (db *data) CreatePackClient(params *model.PackClientParams) error {
	pack, _ := db.GetPack(params.Pack)
	client, _ := db.GetClient(params.Client)

	return db.Model(
		pack,
	).Association(
		"Clients",
	).Append(
		client,
	).Error
}

func (db *data) DeletePackClient(params *model.PackClientParams) error {
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
func (db *data) GetPackUsers(params *model.PackUserParams) (*model.Users, error) {
	pack, _ := db.GetPack(params.Pack)

	records := &model.Users{}

	err := db.Model(
		pack,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetPackHasUser checks if a specific user is assigned to a pack.
func (db *data) GetPackHasUser(params *model.PackUserParams) bool {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	count := db.Model(
		pack,
	).Association(
		"Users",
	).Find(
		user,
	).Count()

	return count > 0
}

func (db *data) CreatePackUser(params *model.PackUserParams) error {
	pack, _ := db.GetPack(params.Pack)
	user, _ := db.GetUser(params.User)

	return db.Model(
		pack,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeletePackUser(params *model.PackUserParams) error {
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
func (db *data) GetPackTeams(params *model.PackTeamParams) (*model.Teams, error) {
	pack, _ := db.GetPack(params.Pack)

	records := &model.Teams{}

	err := db.Model(
		pack,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetPackHasTeam checks if a specific team is assigned to a pack.
func (db *data) GetPackHasTeam(params *model.PackTeamParams) bool {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	count := db.Model(
		pack,
	).Association(
		"Teams",
	).Find(
		team,
	).Count()

	return count > 0
}

func (db *data) CreatePackTeam(params *model.PackTeamParams) error {
	pack, _ := db.GetPack(params.Pack)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		pack,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeletePackTeam(params *model.PackTeamParams) error {
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
