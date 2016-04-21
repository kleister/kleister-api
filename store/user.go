package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetUsers retrieves all available users from the database.
func GetUsers(c context.Context) (*model.Users, error) {
	return FromContext(c).GetUsers()
}

// CreateUser creates a new user.
func CreateUser(c context.Context, record *model.User) error {
	return FromContext(c).CreateUser(record)
}

// UpdateUser updates a user.
func UpdateUser(c context.Context, record *model.User) error {
	return FromContext(c).UpdateUser(record)
}

// DeleteUser deletes a user.
func DeleteUser(c context.Context, record *model.User) error {
	return FromContext(c).DeleteUser(record)
}

// GetUser retrieves a specific user from the database.
func GetUser(c context.Context, id string) (*model.User, *gorm.DB) {
	return FromContext(c).GetUser(id)
}

// GetUserMods retrieves mods for a user.
func GetUserMods(c context.Context, id int) (*model.Mods, error) {
	return FromContext(c).GetUserMods(id)
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func GetUserHasMod(c context.Context, parent, id int) bool {
	return FromContext(c).GetUserHasMod(parent, id)
}

func CreateUserMod(c context.Context, parent, id int) error {
	return FromContext(c).CreateUserMod(parent, id)
}

func DeleteUserMod(c context.Context, parent, id int) error {
	return FromContext(c).DeleteUserMod(parent, id)
}

// GetUserPacks retrieves packs for a user.
func GetUserPacks(c context.Context, id int) (*model.Packs, error) {
	return FromContext(c).GetUserPacks(id)
}

// GetUserHasPack checks if a specific pack is assigned to a user.
func GetUserHasPack(c context.Context, parent, id int) bool {
	return FromContext(c).GetUserHasPack(parent, id)
}

func CreateUserPack(c context.Context, parent, id int) error {
	return FromContext(c).CreateUserPack(parent, id)
}

func DeleteUserPack(c context.Context, parent, id int) error {
	return FromContext(c).DeleteUserPack(parent, id)
}
