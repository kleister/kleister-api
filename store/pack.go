package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetPacks retrieves all available packs from the database.
func GetPacks(c context.Context) (*model.Packs, error) {
	return FromContext(c).GetPacks()
}

// CreatePack creates a new pack.
func CreatePack(c context.Context, record *model.Pack) error {
	return FromContext(c).CreatePack(record, Current(c))
}

// UpdatePack updates a pack.
func UpdatePack(c context.Context, record *model.Pack) error {
	return FromContext(c).UpdatePack(record, Current(c))
}

// DeletePack deletes a pack.
func DeletePack(c context.Context, record *model.Pack) error {
	return FromContext(c).DeletePack(record, Current(c))
}

// GetPack retrieves a specific pack from the database.
func GetPack(c context.Context, id string) (*model.Pack, *gorm.DB) {
	return FromContext(c).GetPack(id)
}

// GetPackClients retrieves clients for a pack.
func GetPackClients(c context.Context, params *model.PackClientParams) (*model.ClientPacks, error) {
	return FromContext(c).GetPackClients(params)
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func GetPackHasClient(c context.Context, params *model.PackClientParams) bool {
	return FromContext(c).GetPackHasClient(params)
}

// CreatePackClient assigns a client to a specific pack.
func CreatePackClient(c context.Context, params *model.PackClientParams) error {
	return FromContext(c).CreatePackClient(params, Current(c))
}

// DeletePackClient removes a client from a specific pack.
func DeletePackClient(c context.Context, params *model.PackClientParams) error {
	return FromContext(c).DeletePackClient(params, Current(c))
}

// GetPackUsers retrieves users for a pack.
func GetPackUsers(c context.Context, params *model.PackUserParams) (*model.UserPacks, error) {
	return FromContext(c).GetPackUsers(params)
}

// GetPackHasUser checks if a specific user is assigned to a pack.
func GetPackHasUser(c context.Context, params *model.PackUserParams) bool {
	return FromContext(c).GetPackHasUser(params)
}

// CreatePackUser assigns a user to a specific pack.
func CreatePackUser(c context.Context, params *model.PackUserParams) error {
	return FromContext(c).CreatePackUser(params, Current(c))
}

// UpdatePackUser updates the pack user permission.
func UpdatePackUser(c context.Context, params *model.PackUserParams) error {
	return FromContext(c).UpdatePackUser(params, Current(c))
}

// DeletePackUser removes a user from a specific pack.
func DeletePackUser(c context.Context, params *model.PackUserParams) error {
	return FromContext(c).DeletePackUser(params, Current(c))
}

// GetPackTeams retrieves teams for a pack.
func GetPackTeams(c context.Context, params *model.PackTeamParams) (*model.TeamPacks, error) {
	return FromContext(c).GetPackTeams(params)
}

// GetPackHasTeam checks if a specific team is assigned to a pack.
func GetPackHasTeam(c context.Context, params *model.PackTeamParams) bool {
	return FromContext(c).GetPackHasTeam(params)
}

// CreatePackTeam assigns a team to a specific pack.
func CreatePackTeam(c context.Context, params *model.PackTeamParams) error {
	return FromContext(c).CreatePackTeam(params, Current(c))
}

// UpdatePackTeam updates the pack team permission.
func UpdatePackTeam(c context.Context, params *model.PackTeamParams) error {
	return FromContext(c).UpdatePackTeam(params, Current(c))
}

// DeletePackTeam removes a team from a specific pack.
func DeletePackTeam(c context.Context, params *model.PackTeamParams) error {
	return FromContext(c).DeletePackTeam(params, Current(c))
}
