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
	return FromContext(c).CreateMod(record, Current(c))
}

// UpdateMod updates a mod.
func UpdateMod(c context.Context, record *model.Mod) error {
	return FromContext(c).UpdateMod(record, Current(c))
}

// DeleteMod deletes a mod.
func DeleteMod(c context.Context, record *model.Mod) error {
	return FromContext(c).DeleteMod(record, Current(c))
}

// GetMod retrieves a specific mod from the database.
func GetMod(c context.Context, id string) (*model.Mod, *gorm.DB) {
	return FromContext(c).GetMod(id)
}

// GetModUsers retrieves users for a mod.
func GetModUsers(c context.Context, params *model.ModUserParams) (*model.UserMods, error) {
	return FromContext(c).GetModUsers(params)
}

// GetModHasUser checks if a specific user is assigned to a mod.
func GetModHasUser(c context.Context, params *model.ModUserParams) bool {
	return FromContext(c).GetModHasUser(params)
}

// CreateModUser assigns a user to a specific mod.
func CreateModUser(c context.Context, params *model.ModUserParams) error {
	return FromContext(c).CreateModUser(params, Current(c))
}

// UpdateModUser updates the mod user permission.
func UpdateModUser(c context.Context, params *model.ModUserParams) error {
	return FromContext(c).UpdateModUser(params, Current(c))
}

// DeleteModUser removes a user from a specific mod.
func DeleteModUser(c context.Context, params *model.ModUserParams) error {
	return FromContext(c).DeleteModUser(params, Current(c))
}

// GetModTeams retrieves teams for a mod.
func GetModTeams(c context.Context, params *model.ModTeamParams) (*model.TeamMods, error) {
	return FromContext(c).GetModTeams(params)
}

// GetModHasTeam checks if a specific team is assigned to a mod.
func GetModHasTeam(c context.Context, params *model.ModTeamParams) bool {
	return FromContext(c).GetModHasTeam(params)
}

// CreateModTeam assigns a team to a specific mod.
func CreateModTeam(c context.Context, params *model.ModTeamParams) error {
	return FromContext(c).CreateModTeam(params, Current(c))
}

// UpdateModTeam updates the mod team permission.
func UpdateModTeam(c context.Context, params *model.ModTeamParams) error {
	return FromContext(c).UpdateModTeam(params, Current(c))
}

// DeleteModTeam removes a team from a specific mod.
func DeleteModTeam(c context.Context, params *model.ModTeamParams) error {
	return FromContext(c).DeleteModTeam(params, Current(c))
}
