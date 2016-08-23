package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetTeams retrieves all available teams from the database.
func GetTeams(c context.Context) (*model.Teams, error) {
	return FromContext(c).GetTeams()
}

// CreateTeam creates a new team.
func CreateTeam(c context.Context, record *model.Team) error {
	return FromContext(c).CreateTeam(record, Current(c))
}

// UpdateTeam updates a team.
func UpdateTeam(c context.Context, record *model.Team) error {
	return FromContext(c).UpdateTeam(record, Current(c))
}

// DeleteTeam deletes a team.
func DeleteTeam(c context.Context, record *model.Team) error {
	return FromContext(c).DeleteTeam(record, Current(c))
}

// GetTeam retrieves a specific team from the database.
func GetTeam(c context.Context, id string) (*model.Team, *gorm.DB) {
	return FromContext(c).GetTeam(id)
}

// GetTeamUsers retrieves users for a team.
func GetTeamUsers(c context.Context, params *model.TeamUserParams) (*model.TeamUsers, error) {
	return FromContext(c).GetTeamUsers(params)
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func GetTeamHasUser(c context.Context, params *model.TeamUserParams) bool {
	return FromContext(c).GetTeamHasUser(params)
}

// CreateTeamUser assigns a user to a specific team.
func CreateTeamUser(c context.Context, params *model.TeamUserParams) error {
	return FromContext(c).CreateTeamUser(params, Current(c))
}

// UpdateTeamUser updates the team user permission.
func UpdateTeamUser(c context.Context, params *model.TeamUserParams) error {
	return FromContext(c).UpdateTeamUser(params, Current(c))
}

// DeleteTeamUser removes a user from a specific team.
func DeleteTeamUser(c context.Context, params *model.TeamUserParams) error {
	return FromContext(c).DeleteTeamUser(params, Current(c))
}

// GetTeamPacks retrieves packs for a team.
func GetTeamPacks(c context.Context, params *model.TeamPackParams) (*model.TeamPacks, error) {
	return FromContext(c).GetTeamPacks(params)
}

// GetTeamHasPack checks if a specific pack is assigned to a team.
func GetTeamHasPack(c context.Context, params *model.TeamPackParams) bool {
	return FromContext(c).GetTeamHasPack(params)
}

// CreateTeamPack assigns a pack to a specific team.
func CreateTeamPack(c context.Context, params *model.TeamPackParams) error {
	return FromContext(c).CreateTeamPack(params, Current(c))
}

// UpdateTeamPack updates the team pack permission.
func UpdateTeamPack(c context.Context, params *model.TeamPackParams) error {
	return FromContext(c).UpdateTeamPack(params, Current(c))
}

// DeleteTeamPack removes a pack from a specific team.
func DeleteTeamPack(c context.Context, params *model.TeamPackParams) error {
	return FromContext(c).DeleteTeamPack(params, Current(c))
}

// GetTeamMods retrieves mods for a team.
func GetTeamMods(c context.Context, params *model.TeamModParams) (*model.TeamMods, error) {
	return FromContext(c).GetTeamMods(params)
}

// GetTeamHasMod checks if a specific mod is assigned to a team.
func GetTeamHasMod(c context.Context, params *model.TeamModParams) bool {
	return FromContext(c).GetTeamHasMod(params)
}

// CreateTeamMod assigns a mod to a specific team.
func CreateTeamMod(c context.Context, params *model.TeamModParams) error {
	return FromContext(c).CreateTeamMod(params, Current(c))
}

// UpdateTeamMod updates the team mod permission.
func UpdateTeamMod(c context.Context, params *model.TeamModParams) error {
	return FromContext(c).UpdateTeamMod(params, Current(c))
}

// DeleteTeamMod removes a mod from a specific team.
func DeleteTeamMod(c context.Context, params *model.TeamModParams) error {
	return FromContext(c).DeleteTeamMod(params, Current(c))
}
