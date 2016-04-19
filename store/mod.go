package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
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
func GetModUsers(c context.Context, id int) (*model.Users, error) {
	return FromContext(c).GetModUsers(id)
}

// GetModHasUser checks if a specific user is assigned to a mod.
func GetModHasUser(c context.Context, parent, id int) bool {
	return FromContext(c).GetModHasUser(parent, id)
}

func CreateModUser(c context.Context, parent, id int) error {
	return FromContext(c).CreateModUser(parent, id)
}

func DeleteModUser(c context.Context, parent, id int) error {
	return FromContext(c).DeleteModUser(parent, id)
}
