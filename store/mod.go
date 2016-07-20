package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetMods retrieves all available mods from the database.
func GetMods(c context.Context) (*model.Mods, error) {
	return FromContext(c).GetMods()
}

// CreateMod creates a new mod.
func CreateMod(c context.Context, record *model.Mod) error {
	return FromContext(c).CreateMod(record)
}

// UpdateMod updates a mod.
func UpdateMod(c context.Context, record *model.Mod) error {
	return FromContext(c).UpdateMod(record)
}

// DeleteMod deletes a mod.
func DeleteMod(c context.Context, record *model.Mod) error {
	return FromContext(c).DeleteMod(record)
}

// GetMod retrieves a specific mod from the database.
func GetMod(c context.Context, id string) (*model.Mod, *gorm.DB) {
	return FromContext(c).GetMod(id)
}

// GetModUsers retrieves users for a mod.
func GetModUsers(c context.Context, params *model.ModUserParams) (*model.Users, error) {
	return FromContext(c).GetModUsers(params)
}

// GetModHasUser checks if a specific user is assigned to a mod.
func GetModHasUser(c context.Context, params *model.ModUserParams) bool {
	return FromContext(c).GetModHasUser(params)
}

// CreateModUser assigns a user to a specific mod.
func CreateModUser(c context.Context, params *model.ModUserParams) error {
	return FromContext(c).CreateModUser(params)
}

// DeleteModUser removes a user from a specific mod.
func DeleteModUser(c context.Context, params *model.ModUserParams) error {
	return FromContext(c).DeleteModUser(params)
}
