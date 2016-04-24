package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetPacks retrieves all available packs from the database.
func GetPacks(c context.Context) (*model.Packs, error) {
	return FromContext(c).GetPacks()
}

// CreatePack creates a new pack.
func CreatePack(c context.Context, record *model.Pack) error {
	return FromContext(c).CreatePack(record)
}

// UpdatePack updates a pack.
func UpdatePack(c context.Context, record *model.Pack) error {
	return FromContext(c).UpdatePack(record)
}

// DeletePack deletes a pack.
func DeletePack(c context.Context, record *model.Pack) error {
	return FromContext(c).DeletePack(record)
}

// GetPack retrieves a specific pack from the database.
func GetPack(c context.Context, id string) (*model.Pack, *gorm.DB) {
	return FromContext(c).GetPack(id)
}

// GetPackClients retrieves clients for a pack.
func GetPackClients(c context.Context, params *model.PackClientParams) (*model.Clients, error) {
	return FromContext(c).GetPackClients(params)
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func GetPackHasClient(c context.Context, params *model.PackClientParams) bool {
	return FromContext(c).GetPackHasClient(params)
}

func CreatePackClient(c context.Context, params *model.PackClientParams) error {
	return FromContext(c).CreatePackClient(params)
}

func DeletePackClient(c context.Context, params *model.PackClientParams) error {
	return FromContext(c).DeletePackClient(params)
}

// GetPackUsers retrieves users for a pack.
func GetPackUsers(c context.Context, params *model.PackUserParams) (*model.Users, error) {
	return FromContext(c).GetPackUsers(params)
}

// GetPackHasUser checks if a specific user is assigned to a pack.
func GetPackHasUser(c context.Context, params *model.PackUserParams) bool {
	return FromContext(c).GetPackHasUser(params)
}

func CreatePackUser(c context.Context, params *model.PackUserParams) error {
	return FromContext(c).CreatePackUser(params)
}

func DeletePackUser(c context.Context, params *model.PackUserParams) error {
	return FromContext(c).DeletePackUser(params)
}
