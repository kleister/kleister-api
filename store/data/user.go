package data

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
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
func (db *data) GetUserMods(params *model.UserModParams) (*model.Mods, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Mods{}

	err := db.Model(
		user,
	).Association(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func (db *data) GetUserHasMod(params *model.UserModParams) bool {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	count := db.Model(
		user,
	).Association(
		"Mods",
	).Find(
		mod,
	).Count()

	return count > 0
}

func (db *data) CreateUserMod(params *model.UserModParams) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		user,
	).Association(
		"Mods",
	).Append(
		mod,
	).Error
}

func (db *data) DeleteUserMod(params *model.UserModParams) error {
	user, _ := db.GetUser(params.User)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		user,
	).Association(
		"Mods",
	).Delete(
		mod,
	).Error
}

// GetUserPacks retrieves packs for a user.
func (db *data) GetUserPacks(params *model.UserPackParams) (*model.Packs, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Packs{}

	err := db.Model(
		user,
	).Association(
		"Packs",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasPack checks if a specific pack is assigned to a user.
func (db *data) GetUserHasPack(params *model.UserPackParams) bool {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	count := db.Model(
		user,
	).Association(
		"Packs",
	).Find(
		pack,
	).Count()

	return count > 0
}

func (db *data) CreateUserPack(params *model.UserPackParams) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		user,
	).Association(
		"Packs",
	).Append(
		pack,
	).Error
}

func (db *data) DeleteUserPack(params *model.UserPackParams) error {
	user, _ := db.GetUser(params.User)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		user,
	).Association(
		"Packs",
	).Delete(
		pack,
	).Error
}
