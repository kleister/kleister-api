package data

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetMods retrieves all available mods from the database.
func (db *data) GetMods() (*model.Mods, error) {
	records := &model.Mods{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Versions",
	).Find(
		records,
	).Error

	return records, err
}

// CreateMod creates a new mod.
func (db *data) CreateMod(record *model.Mod) error {
	return db.Create(
		record,
	).Error
}

// UpdateMod updates a mod.
func (db *data) UpdateMod(record *model.Mod) error {
	return db.Save(
		record,
	).Error
}

// DeleteMod deletes a mod.
func (db *data) DeleteMod(record *model.Mod) error {
	return db.Delete(
		record,
	).Error
}

// GetMod retrieves a specific mod from the database.
func (db *data) GetMod(id string) (*model.Mod, *gorm.DB) {
	record := &model.Mod{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		record,
	).Preload(
		"Versions",
	).First(
		record,
	)

	return record, res
}

// GetModUsers retrieves users for a mod.
func (db *data) GetModUsers(params *model.ModUserParams) (*model.Users, error) {
	mod, _ := db.GetMod(params.Mod)

	records := &model.Users{}

	err := db.Model(
		mod,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetModHasUser checks if a specific user is assigned to a mod.
func (db *data) GetModHasUser(params *model.ModUserParams) bool {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	count := db.Model(
		mod,
	).Association(
		"Users",
	).Find(
		user,
	).Count()

	return count > 0
}

func (db *data) CreateModUser(params *model.ModUserParams) error {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	return db.Model(
		mod,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteModUser(params *model.ModUserParams) error {
	mod, _ := db.GetMod(params.Mod)
	user, _ := db.GetUser(params.User)

	return db.Model(
		mod,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}
