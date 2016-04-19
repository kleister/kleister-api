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
func GetPackClients(c context.Context, id int) (*model.Clients, error) {
	return FromContext(c).GetPackClients(id)
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func GetPackHasClient(c context.Context, parent, id int) bool {
	return FromContext(c).GetPackHasClient(parent, id)
}

func CreatePackClient(c context.Context, parent, id int) error {
	return FromContext(c).CreatePackClient(parent, id)
}

func DeletePackClient(c context.Context, parent, id int) error {
	return FromContext(c).DeletePackClient(parent, id)
}
