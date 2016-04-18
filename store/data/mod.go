package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
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
func (db *data) GetModUsers(id int) (*model.Users, error) {
	records := &model.Users{}

	err := db.Model(
		&model.Mod{
			ID: id,
		},
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetModHasUser checks if a specific user is assigned to a mod.
func (db *data) GetModHasUser(parent, id int) bool {
	record := &model.User{
		ID: id,
	}

	count := db.Model(
		&model.Mod{
			ID: parent,
		},
	).Association(
		"Users",
	).Find(
		record,
	).Count()

	return count > 0
}
