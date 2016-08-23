package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetUsers retrieves all available users from the database.
func GetUsers(c context.Context) (*model.Users, error) {
	return FromContext(c).GetUsers()
}

// CreateUser creates a new user.
func CreateUser(c context.Context, record *model.User) error {
	return FromContext(c).CreateUser(record, Current(c))
}

// UpdateUser updates a user.
func UpdateUser(c context.Context, record *model.User) error {
	return FromContext(c).UpdateUser(record, Current(c))
}

// DeleteUser deletes a user.
func DeleteUser(c context.Context, record *model.User) error {
	return FromContext(c).DeleteUser(record, Current(c))
}

// GetUser retrieves a specific user from the database.
func GetUser(c context.Context, id string) (*model.User, *gorm.DB) {
	return FromContext(c).GetUser(id)
}

// GetUserMods retrieves mods for a user.
func GetUserMods(c context.Context, params *model.UserModParams) (*model.UserMods, error) {
	return FromContext(c).GetUserMods(params)
}

// GetUserHasMod checks if a specific mod is assigned to a user.
func GetUserHasMod(c context.Context, params *model.UserModParams) bool {
	return FromContext(c).GetUserHasMod(params)
}

// CreateUserMod assigns a mod to a specific user.
func CreateUserMod(c context.Context, params *model.UserModParams) error {
	return FromContext(c).CreateUserMod(params, Current(c))
}

// UpdateUserMod updates the user mod permission.
func UpdateUserMod(c context.Context, params *model.UserModParams) error {
	return FromContext(c).UpdateUserMod(params, Current(c))
}

// DeleteUserMod removes a mod from a specific user.
func DeleteUserMod(c context.Context, params *model.UserModParams) error {
	return FromContext(c).DeleteUserMod(params, Current(c))
}

// GetUserPacks retrieves packs for a user.
func GetUserPacks(c context.Context, params *model.UserPackParams) (*model.UserPacks, error) {
	return FromContext(c).GetUserPacks(params)
}

// GetUserHasPack checks if a specific pack is assigned to a user.
func GetUserHasPack(c context.Context, params *model.UserPackParams) bool {
	return FromContext(c).GetUserHasPack(params)
}

// CreateUserPack assigns a pack to a specific user.
func CreateUserPack(c context.Context, params *model.UserPackParams) error {
	return FromContext(c).CreateUserPack(params, Current(c))
}

// UpdateUserPack updates the user mod permission.
func UpdateUserPack(c context.Context, params *model.UserPackParams) error {
	return FromContext(c).UpdateUserPack(params, Current(c))
}

// DeleteUserPack removes a pack from a specific user.
func DeleteUserPack(c context.Context, params *model.UserPackParams) error {
	return FromContext(c).DeleteUserPack(params, Current(c))
}

// GetUserTeams retrieves teams for a user.
func GetUserTeams(c context.Context, params *model.UserTeamParams) (*model.TeamUsers, error) {
	return FromContext(c).GetUserTeams(params)
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func GetUserHasTeam(c context.Context, params *model.UserTeamParams) bool {
	return FromContext(c).GetUserHasTeam(params)
}

// CreateUserTeam assigns a team to a specific user.
func CreateUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).CreateUserTeam(params, Current(c))
}

// UpdateUserTeam updates the user team permission.
func UpdateUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).UpdateUserTeam(params, Current(c))
}

// DeleteUserTeam removes a team from a specific user.
func DeleteUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).DeleteUserTeam(params, Current(c))
}
