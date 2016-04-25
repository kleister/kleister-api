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
func GetUserMods(c context.Context, params *model.UserModParams) (*model.Mods, error) {
	return FromContext(c).GetUserMods(params)
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func GetUserHasMod(c context.Context, params *model.UserModParams) bool {
	return FromContext(c).GetUserHasMod(params)
}

// CreateUserMod assigns a mod to a specific user.
func CreateUserMod(c context.Context, params *model.UserModParams) error {
	return FromContext(c).CreateUserMod(params)
}

// DeleteUserMod removes a mod from a specific user.
func DeleteUserMod(c context.Context, params *model.UserModParams) error {
	return FromContext(c).DeleteUserMod(params)
}

// GetUserPacks retrieves packs for a user.
func GetUserPacks(c context.Context, params *model.UserPackParams) (*model.Packs, error) {
	return FromContext(c).GetUserPacks(params)
}

// GetUserHasPack checks if a specific pack is assigned to a user.
func GetUserHasPack(c context.Context, params *model.UserPackParams) bool {
	return FromContext(c).GetUserHasPack(params)
}

// CreateUserPack assigns a pack to a specific user.
func CreateUserPack(c context.Context, params *model.UserPackParams) error {
	return FromContext(c).CreateUserPack(params)
}

// DeleteUserPack removes a pack from a specific user.
func DeleteUserPack(c context.Context, params *model.UserPackParams) error {
	return FromContext(c).DeleteUserPack(params)
}
