package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/forge"
	"github.com/kleister/kleister-api/model/minecraft"
)

// Store implements all required data-layer functions for Solder.
type Store interface {
	// GetKeys retrieves all available keys from the database.
	GetKeys() (*model.Keys, error)

	// CreateKey creates a new key.
	CreateKey(*model.Key, *model.User) error

	// UpdateKey updates a key.
	UpdateKey(*model.Key, *model.User) error

	// DeleteKey deletes a key.
	DeleteKey(*model.Key, *model.User) error

	// GetKey retrieves a specific key from the database.
	GetKey(string) (*model.Key, *gorm.DB)

	// GetKeyByValue retrieves a specific key by value from the database.
	GetKeyByValue(string) (*model.Key, *gorm.DB)

	// GetBuilds retrieves all available builds from the database.
	GetBuilds(int64) (*model.Builds, error)

	// CreateBuild creates a new build.
	CreateBuild(int64, *model.Build, *model.User) error

	// UpdateBuild updates a build.
	UpdateBuild(int64, *model.Build, *model.User) error

	// DeleteBuild deletes a build.
	DeleteBuild(int64, *model.Build, *model.User) error

	// GetBuild retrieves a specific build from the database.
	GetBuild(int64, string) (*model.Build, *gorm.DB)

	// GetBuildVersions retrieves versions for a build.
	GetBuildVersions(*model.BuildVersionParams) (*model.BuildVersions, error)

	// GetBuildHasVersion checks if a specific version is assigned to a build.
	GetBuildHasVersion(*model.BuildVersionParams) bool

	// CreateBuildVersion assigns a version to a specific build.
	CreateBuildVersion(*model.BuildVersionParams, *model.User) error

	// DeleteBuildVersion removes a version from a specific build.
	DeleteBuildVersion(*model.BuildVersionParams, *model.User) error

	// GetClients retrieves all available clients from the database.
	GetClients() (*model.Clients, error)

	// CreateClient creates a new client.
	CreateClient(*model.Client, *model.User) error

	// UpdateClient updates a client.
	UpdateClient(*model.Client, *model.User) error

	// DeleteClient deletes a client.
	DeleteClient(*model.Client, *model.User) error

	// GetClient retrieves a specific client from the database.
	GetClient(string) (*model.Client, *gorm.DB)

	// GetClientByValue retrieves a specific client by value from the database.
	GetClientByValue(string) (*model.Client, *gorm.DB)

	// GetClientPacks retrieves packs for a client.
	GetClientPacks(*model.ClientPackParams) (*model.ClientPacks, error)

	// GetClientHasPack checks if a specific pack is assigned to a client.
	GetClientHasPack(*model.ClientPackParams) bool

	// CreateClientPack assigns a pack to a specific client.
	CreateClientPack(*model.ClientPackParams, *model.User) error

	// DeleteClientPack removes a pack from a specific client.
	DeleteClientPack(*model.ClientPackParams, *model.User) error

	// GetMods retrieves all available mods from the database.
	GetMods() (*model.Mods, error)

	// CreateMod creates a new mod.
	CreateMod(*model.Mod, *model.User) error

	// UpdateMod updates a mod.
	UpdateMod(*model.Mod, *model.User) error

	// DeleteMod deletes a mod.
	DeleteMod(*model.Mod, *model.User) error

	// GetMod retrieves a specific mod from the database.
	GetMod(string) (*model.Mod, *gorm.DB)

	// GetModUsers retrieves users for a mod.
	GetModUsers(*model.ModUserParams) (*model.UserMods, error)

	// GetModHasUser checks if a specific user is assigned to a mod.
	GetModHasUser(*model.ModUserParams) bool

	// CreateModUser assigns a user to a specific mod.
	CreateModUser(*model.ModUserParams, *model.User) error

	// UpdateModUser updates the mod user permission.
	UpdateModUser(*model.ModUserParams, *model.User) error

	// DeleteModUser removes a user from a specific mod.
	DeleteModUser(*model.ModUserParams, *model.User) error

	// GetModTeams retrieves teams for a mod.
	GetModTeams(*model.ModTeamParams) (*model.TeamMods, error)

	// GetModHasTeam checks if a specific team is assigned to a mod.
	GetModHasTeam(*model.ModTeamParams) bool

	// CreateModTeam assigns a team to a specific mod.
	CreateModTeam(*model.ModTeamParams, *model.User) error

	// UpdateModTeam updates the mod team permission.
	UpdateModTeam(*model.ModTeamParams, *model.User) error

	// DeleteModTeam removes a team from a specific mod.
	DeleteModTeam(*model.ModTeamParams, *model.User) error

	// GetPacks retrieves all available packs from the database.
	GetPacks() (*model.Packs, error)

	// CreatePack creates a new pack.
	CreatePack(*model.Pack, *model.User) error

	// UpdatePack updates a pack.
	UpdatePack(*model.Pack, *model.User) error

	// DeletePack deletes a pack.
	DeletePack(*model.Pack, *model.User) error

	// GetPack retrieves a specific pack from the database.
	GetPack(string) (*model.Pack, *gorm.DB)

	// GetPackClients retrieves clients for a pack.
	GetPackClients(*model.PackClientParams) (*model.ClientPacks, error)

	// GetPackHasClient checks if a specific client is assigned to a pack.
	GetPackHasClient(*model.PackClientParams) bool

	// CreatePackClient assigns a client to a specific pack.
	CreatePackClient(*model.PackClientParams, *model.User) error

	// DeletePackClient removes a client from a specific pack.
	DeletePackClient(*model.PackClientParams, *model.User) error

	// GetPackUsers retrieves users for a pack.
	GetPackUsers(*model.PackUserParams) (*model.UserPacks, error)

	// GetPackHasUser checks if a specific user is assigned to a pack.
	GetPackHasUser(*model.PackUserParams) bool

	// CreatePackUser assigns a user to a specific pack.
	CreatePackUser(*model.PackUserParams, *model.User) error

	// UpdatePackUser updates the pack user permission.
	UpdatePackUser(*model.PackUserParams, *model.User) error

	// DeletePackUser removes a user from a specific pack.
	DeletePackUser(*model.PackUserParams, *model.User) error

	// GetPackTeams retrieves teams for a pack.
	GetPackTeams(*model.PackTeamParams) (*model.TeamPacks, error)

	// GetPackHasTeam checks if a specific team is assigned to a pack.
	GetPackHasTeam(*model.PackTeamParams) bool

	// CreatePackTeam assigns a team to a specific pack.
	CreatePackTeam(*model.PackTeamParams, *model.User) error

	// UpdatePackTeam updates the pack team permission.
	UpdatePackTeam(*model.PackTeamParams, *model.User) error

	// DeletePackTeam removes a team from a specific pack.
	DeletePackTeam(*model.PackTeamParams, *model.User) error

	// GetUsers retrieves all available users from the database.
	GetUsers() (*model.Users, error)

	// CreateUser creates a new user.
	CreateUser(*model.User, *model.User) error

	// UpdateUser updates a user.
	UpdateUser(*model.User, *model.User) error

	// DeleteUser deletes a user.
	DeleteUser(*model.User, *model.User) error

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, *gorm.DB)

	// GetUserMods retrieves mods for a user.
	GetUserMods(*model.UserModParams) (*model.UserMods, error)

	// GetUserHasMod checks if a specific mod is assigned to a user.
	GetUserHasMod(*model.UserModParams) bool

	// CreateUserMod assigns a mod to a specific user.
	CreateUserMod(*model.UserModParams, *model.User) error

	// UpdateUserMod updates the user mod permission.
	UpdateUserMod(*model.UserModParams, *model.User) error

	// DeleteUserMod removes a mod from a specific user.
	DeleteUserMod(*model.UserModParams, *model.User) error

	// GetUserPacks retrieves packs for a user.
	GetUserPacks(*model.UserPackParams) (*model.UserPacks, error)

	// GetUserHasPack checks if a specific pack is assigned to a user.
	GetUserHasPack(*model.UserPackParams) bool

	// CreateUserPack assigns a pack to a specific user.
	CreateUserPack(*model.UserPackParams, *model.User) error

	// UpdateUserPack updates the user pack permission.
	UpdateUserPack(*model.UserPackParams, *model.User) error

	// DeleteUserPack removes a pack from a specific user.
	DeleteUserPack(*model.UserPackParams, *model.User) error

	// GetUserTeams retrieves teams for a user.
	GetUserTeams(*model.UserTeamParams) (*model.TeamUsers, error)

	// GetUserHasTeam checks if a specific team is assigned to a user.
	GetUserHasTeam(*model.UserTeamParams) bool

	// CreateUserTeam assigns a team to a specific user.
	CreateUserTeam(*model.UserTeamParams, *model.User) error

	// UpdateUserTeam updates the user team permission.
	UpdateUserTeam(*model.UserTeamParams, *model.User) error

	// DeleteUserTeam removes a team from a specific user.
	DeleteUserTeam(*model.UserTeamParams, *model.User) error

	// GetVersions retrieves all available versions from the database.
	GetVersions(int64) (*model.Versions, error)

	// CreateVersion creates a new version.
	CreateVersion(int64, *model.Version, *model.User) error

	// UpdateVersion updates a version.
	UpdateVersion(int64, *model.Version, *model.User) error

	// DeleteVersion deletes a version.
	DeleteVersion(int64, *model.Version, *model.User) error

	// GetVersion retrieves a specific version from the database.
	GetVersion(int64, string) (*model.Version, *gorm.DB)

	// GetVersionBuilds retrieves builds for a version.
	GetVersionBuilds(*model.VersionBuildParams) (*model.BuildVersions, error)

	// GetVersionHasBuild checks if a specific build is assigned to a version.
	GetVersionHasBuild(*model.VersionBuildParams) bool

	// CreateVersionBuild assigns a build to a specific version.
	CreateVersionBuild(*model.VersionBuildParams, *model.User) error

	// DeleteVersionBuild removes a build from a specific version.
	DeleteVersionBuild(*model.VersionBuildParams, *model.User) error

	// GetTeams retrieves all available teams from the database.
	GetTeams() (*model.Teams, error)

	// CreateTeam creates a new team.
	CreateTeam(*model.Team, *model.User) error

	// UpdateTeam updates a team.
	UpdateTeam(*model.Team, *model.User) error

	// DeleteTeam deletes a team.
	DeleteTeam(*model.Team, *model.User) error

	// GetTeam retrieves a specific team from the database.
	GetTeam(string) (*model.Team, *gorm.DB)

	// GetTeamUsers retrieves users for a team.
	GetTeamUsers(*model.TeamUserParams) (*model.TeamUsers, error)

	// GetTeamHasUser checks if a specific user is assigned to a team.
	GetTeamHasUser(*model.TeamUserParams) bool

	// CreateTeamUser assigns a user to a specific team.
	CreateTeamUser(*model.TeamUserParams, *model.User) error

	// UpdateTeamUser updates the team user permission.
	UpdateTeamUser(*model.TeamUserParams, *model.User) error

	// DeleteTeamUser removes a user from a specific team.
	DeleteTeamUser(*model.TeamUserParams, *model.User) error

	// GetTeamPacks retrieves packs for a team.
	GetTeamPacks(*model.TeamPackParams) (*model.TeamPacks, error)

	// GetTeamHasPack checks if a specific pack is assigned to a team.
	GetTeamHasPack(*model.TeamPackParams) bool

	// CreateTeamPack assigns a pack to a specific team.
	CreateTeamPack(*model.TeamPackParams, *model.User) error

	// UpdateTeamPack updates the team pack permission.
	UpdateTeamPack(*model.TeamPackParams, *model.User) error

	// DeleteTeamPack removes a pack from a specific team.
	DeleteTeamPack(*model.TeamPackParams, *model.User) error

	// GetTeamMods retrieves mods for a team.
	GetTeamMods(*model.TeamModParams) (*model.TeamMods, error)

	// GetTeamHasMod checks if a specific mod is assigned to a team.
	GetTeamHasMod(*model.TeamModParams) bool

	// CreateTeamMod assigns a mod to a specific team.
	CreateTeamMod(*model.TeamModParams, *model.User) error

	// UpdateTeamMod updates the team mod permission.
	UpdateTeamMod(*model.TeamModParams, *model.User) error

	// DeleteTeamMod removes a mod from a specific team.
	DeleteTeamMod(*model.TeamModParams, *model.User) error

	// GetMinecrafts retrieves all available minecrafts from the database.
	GetMinecrafts() (*model.Minecrafts, error)

	// SyncMinecraft creates or updates a minecraft record.
	SyncMinecraft(*minecraft.Version, *model.User) (*model.Minecraft, error)

	// GetMinecraft retrieves a specific minecraft from the database.
	GetMinecraft(string) (*model.Minecraft, *gorm.DB)

	// GetMinecraftBuilds retrieves builds for a minecraft.
	GetMinecraftBuilds(*model.MinecraftBuildParams) (*model.Builds, error)

	// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
	GetMinecraftHasBuild(*model.MinecraftBuildParams) bool

	// CreateMinecraftBuild assigns a build to a specific minecraft.
	CreateMinecraftBuild(*model.MinecraftBuildParams, *model.User) error

	// DeleteMinecraftBuild removes a build from a specific minecraft.
	DeleteMinecraftBuild(*model.MinecraftBuildParams, *model.User) error

	// GetForges retrieves all available forges from the database.
	GetForges() (*model.Forges, error)

	// SyncForge creates or updates a forge record.
	SyncForge(*forge.Number, *model.User) (*model.Forge, error)

	// GetForge retrieves a specific forge from the database.
	GetForge(string) (*model.Forge, *gorm.DB)

	// GetForgeBuilds retrieves builds for a forge.
	GetForgeBuilds(*model.ForgeBuildParams) (*model.Builds, error)

	// GetForgeHasBuild checks if a specific build is assigned to a forge.
	GetForgeHasBuild(*model.ForgeBuildParams) bool

	// CreateForgeBuild assigns a build to a specific forge.
	CreateForgeBuild(*model.ForgeBuildParams, *model.User) error

	// DeleteForgeBuild removes a build from a specific forge.
	DeleteForgeBuild(*model.ForgeBuildParams, *model.User) error

	// GetSolderPacks retrieves all available packs optimized for the solder compatible API.
	GetSolderPacks() (*model.Packs, error)

	// GetSolderPack retrieves a specific pack optimized for the solder compatible API.
	GetSolderPack(string) (*model.Pack, error)

	// GetSolderBuild retrieves a specific build optimized for the solder compatible API.
	GetSolderBuild(string, string) (*model.Build, error)
}
