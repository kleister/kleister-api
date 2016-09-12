package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetClients retrieves all available clients from the database.
func GetClients(c context.Context) (*model.Clients, error) {
	return FromContext(c).GetClients()
}

// CreateClient creates a new client.
func CreateClient(c context.Context, record *model.Client) error {
	return FromContext(c).CreateClient(record, Current(c))
}

// UpdateClient updates a client.
func UpdateClient(c context.Context, record *model.Client) error {
	return FromContext(c).UpdateClient(record, Current(c))
}

// DeleteClient deletes a client.
func DeleteClient(c context.Context, record *model.Client) error {
	return FromContext(c).DeleteClient(record, Current(c))
}

// GetClient retrieves a specific client from the database.
func GetClient(c context.Context, id string) (*model.Client, *gorm.DB) {
	return FromContext(c).GetClient(id)
}

// GetClientByValue retrieves a specific client by value from the database.
func GetClientByValue(c context.Context, id string) (*model.Client, *gorm.DB) {
	return FromContext(c).GetClientByValue(id)
}

// GetClientPacks retrieves packs for a client.
func GetClientPacks(c context.Context, params *model.ClientPackParams) (*model.ClientPacks, error) {
	return FromContext(c).GetClientPacks(params)
}

// GetClientHasPack checks if a specific pack is assigned to a client.
func GetClientHasPack(c context.Context, params *model.ClientPackParams) bool {
	return FromContext(c).GetClientHasPack(params)
}

// CreateClientPack assigns a pack to a specific client.
func CreateClientPack(c context.Context, params *model.ClientPackParams) error {
	return FromContext(c).CreateClientPack(params, Current(c))
}

// DeleteClientPack removes a pack from a specific client.
func DeleteClientPack(c context.Context, params *model.ClientPackParams) error {
	return FromContext(c).DeleteClientPack(params, Current(c))
}
