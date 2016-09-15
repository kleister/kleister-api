package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetClients retrieves all available clients from the database.
func (db *data) GetClients() (*model.Clients, error) {
	records := &model.Clients{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Packs",
	).Find(
		records,
	).Error

	return records, err
}

// CreateClient creates a new client.
func (db *data) CreateClient(record *model.Client, current *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateClient updates a client.
func (db *data) UpdateClient(record *model.Client, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteClient deletes a client.
func (db *data) DeleteClient(record *model.Client, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetClient retrieves a specific client from the database.
func (db *data) GetClient(id string) (*model.Client, *gorm.DB) {
	var (
		record = &model.Client{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Client{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.Client{
				Slug: id,
			},
		)
	}

	res := query.Preload(
		"Packs",
	).First(
		record,
	)

	return record, res
}

// GetClientByValue retrieves a specific client by value from the database.
func (db *data) GetClientByValue(id string) (*model.Client, *gorm.DB) {
	record := &model.Client{}

	res := db.Where(
		&model.Client{
			Value: id,
		},
	).Model(
		record,
	).Preload(
		"Packs",
	).First(
		record,
	)

	return record, res
}

// GetClientPacks retrieves packs for a client.
func (db *data) GetClientPacks(params *model.ClientPackParams) (*model.ClientPacks, error) {
	client, _ := db.GetClient(params.Client)
	records := &model.ClientPacks{}

	err := db.Where(
		&model.ClientPack{
			ClientID: client.ID,
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

// GetClientHasPack checks if a specific pack is assigned to a client.
func (db *data) GetClientHasPack(params *model.ClientPackParams) bool {
	client, _ := db.GetClient(params.Client)
	pack, _ := db.GetPack(params.Pack)

	res := db.Model(
		client,
	).Association(
		"Packs",
	).Find(
		pack,
	).Error

	return res == nil
}

func (db *data) CreateClientPack(params *model.ClientPackParams, current *model.User) error {
	client, _ := db.GetClient(params.Client)
	pack, _ := db.GetPack(params.Pack)

	return db.Create(
		&model.ClientPack{
			ClientID: client.ID,
			PackID:   pack.ID,
		},
	).Error
}

func (db *data) DeleteClientPack(params *model.ClientPackParams, current *model.User) error {
	client, _ := db.GetClient(params.Client)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		client,
	).Association(
		"Packs",
	).Delete(
		pack,
	).Error
}
