package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
)

// GetUsers retrieves all available users from the database.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.Order(
		"username ASC",
	).Preload(
		"Permission",
	).Find(
		records,
	).Error

	return records, err
}

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetUser retrieves a specific user from the database.
func (db *data) GetUser(id string) (*model.User, *gorm.DB) {
	record := &model.User{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		record,
	).Preload(
		"Permission",
	).First(
		record,
	)

	return record, res
}

// GetUserMods retrieves mods for a user.
func (db *data) GetUserMods(id int) (*model.Mods, error) {
	records := &model.Mods{}

	err := db.Model(
		&model.User{
			ID: id,
		},
	).Association(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func (db *data) GetUserHasMod(parent, id int) bool {
	record := &model.Mod{
		ID: id,
	}

	count := db.Model(
		&model.User{
			ID: parent,
		},
	).Association(
		"Mods",
	).Find(
		record,
	).Count()

	return count > 0
}

func (db *data) CreateUserMod(parent, id int) error {
	return db.Model(
		&model.User{
			ID: parent,
		},
	).Association(
		"Mods",
	).Append(
		&model.Mod{
			ID: id,
		},
	).Error
}

func (db *data) DeleteUserMod(parent, id int) error {
	return db.Model(
		&model.User{
			ID: parent,
		},
	).Association(
		"Mods",
	).Delete(
		&model.Mod{
			ID: id,
		},
	).Error
}
