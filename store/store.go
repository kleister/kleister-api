package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/forge"
	"github.com/solderapp/solder-api/model/minecraft"
)

//go:generate mockery -all -case=underscore

type Store interface {
	// GetBuilds retrieves all available builds from the database.
	GetBuilds(int) (*model.Builds, error)

	// CreateBuild creates a new build.
	CreateBuild(int, *model.Build) error

	// UpdateBuild updates a build.
	UpdateBuild(int, *model.Build) error

	// DeleteBuild deletes a build.
	DeleteBuild(int, *model.Build) error

	// GetBuild retrieves a specific build from the database.
	GetBuild(int, string) (*model.Build, *gorm.DB)

	// GetBuildVersions retrieves versions for a build.
	GetBuildVersions(*model.BuildVersionParams) (*model.Versions, error)

	// GetBuildHasVersion checks if a specific version is assigned to a build.
	GetBuildHasVersion(*model.BuildVersionParams) bool

	CreateBuildVersion(*model.BuildVersionParams) error

	DeleteBuildVersion(*model.BuildVersionParams) error

	// GetClients retrieves all available clients from the database.
	GetClients() (*model.Clients, error)

	// CreateClient creates a new client.
	CreateClient(*model.Client) error

	// UpdateClient updates a client.
	UpdateClient(*model.Client) error

	// DeleteClient deletes a client.
	DeleteClient(*model.Client) error

	// GetClient retrieves a specific client from the database.
	GetClient(string) (*model.Client, *gorm.DB)

	// GetClientPacks retrieves packs for a client.
	GetClientPacks(*model.ClientPackParams) (*model.Packs, error)

	// GetClientHasPack checks if a specific pack is assigned to a client.
	GetClientHasPack(*model.ClientPackParams) bool

	CreateClientPack(*model.ClientPackParams) error

	DeleteClientPack(*model.ClientPackParams) error

	// GetKeys retrieves all available keys from the database.
	GetKeys() (*model.Keys, error)

	// CreateKey creates a new key.
	CreateKey(*model.Key) error

	// UpdateKey updates a key.
	UpdateKey(*model.Key) error

	// DeleteKey deletes a key.
	DeleteKey(*model.Key) error

	// GetKey retrieves a specific key from the database.
	GetKey(string) (*model.Key, *gorm.DB)

	// GetMods retrieves all available mods from the database.
	GetMods() (*model.Mods, error)

	// CreateMod creates a new mod.
	CreateMod(*model.Mod) error

	// UpdateMod updates a mod.
	UpdateMod(*model.Mod) error

	// DeleteMod deletes a mod.
	DeleteMod(*model.Mod) error

	// GetMod retrieves a specific mod from the database.
	GetMod(string) (*model.Mod, *gorm.DB)

	// GetModUsers retrieves users for a mod.
	GetModUsers(*model.ModUserParams) (*model.Users, error)

	// GetModHasUser checks if a specific user is assigned to a mod.
	GetModHasUser(*model.ModUserParams) bool

	CreateModUser(*model.ModUserParams) error

	DeleteModUser(*model.ModUserParams) error

	// GetPacks retrieves all available packs from the database.
	GetPacks() (*model.Packs, error)

	// CreatePack creates a new pack.
	CreatePack(*model.Pack) error

	// UpdatePack updates a pack.
	UpdatePack(*model.Pack) error

	// DeletePack deletes a pack.
	DeletePack(*model.Pack) error

	// GetPack retrieves a specific pack from the database.
	GetPack(string) (*model.Pack, *gorm.DB)

	// GetPackClients retrieves clients for a pack.
	GetPackClients(*model.PackClientParams) (*model.Clients, error)

	// GetPackHasClient checks if a specific client is assigned to a pack.
	GetPackHasClient(*model.PackClientParams) bool

	CreatePackClient(*model.PackClientParams) error

	DeletePackClient(*model.PackClientParams) error

	// GetPackUsers retrieves users for a pack.
	GetPackUsers(*model.PackUserParams) (*model.Users, error)

	// GetPackHasUser checks if a specific user is assigned to a pack.
	GetPackHasUser(*model.PackUserParams) bool

	CreatePackUser(*model.PackUserParams) error

	DeletePackUser(*model.PackUserParams) error

	// GetUsers retrieves all available users from the database.
	GetUsers() (*model.Users, error)

	// CreateUser creates a new user.
	CreateUser(*model.User) error

	// UpdateUser updates a user.
	UpdateUser(*model.User) error

	// DeleteUser deletes a user.
	DeleteUser(*model.User) error

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, *gorm.DB)

	// GetUserMods retrieves mods for a user.
	GetUserMods(*model.UserModParams) (*model.Mods, error)

	// GetUserHasMod checks if a specific mod is assigned to a user.
	GetUserHasMod(*model.UserModParams) bool

	CreateUserMod(*model.UserModParams) error

	DeleteUserMod(*model.UserModParams) error

	// GetUserPacks retrieves packs for a user.
	GetUserPacks(*model.UserPackParams) (*model.Packs, error)

	// GetUserHasPack checks if a specific pack is assigned to a user.
	GetUserHasPack(*model.UserPackParams) bool

	CreateUserPack(*model.UserPackParams) error

	DeleteUserPack(*model.UserPackParams) error

	// GetVersions retrieves all available versions from the database.
	GetVersions(int) (*model.Versions, error)

	// CreateVersion creates a new version.
	CreateVersion(int, *model.Version) error

	// UpdateVersion updates a version.
	UpdateVersion(int, *model.Version) error

	// DeleteVersion deletes a version.
	DeleteVersion(int, *model.Version) error

	// GetVersion retrieves a specific version from the database.
	GetVersion(int, string) (*model.Version, *gorm.DB)

	// GetVersionBuilds retrieves builds for a version.
	GetVersionBuilds(*model.VersionBuildParams) (*model.Builds, error)

	// GetVersionHasBuild checks if a specific build is assigned to a version.
	GetVersionHasBuild(*model.VersionBuildParams) bool

	CreateVersionBuild(*model.VersionBuildParams) error

	DeleteVersionBuild(*model.VersionBuildParams) error

	// GetMinecrafts retrieves all available minecrafts from the database.
	GetMinecrafts() (*model.Minecrafts, error)

	// SyncMinecraft creates or updates a minecraft record.
	SyncMinecraft(*minecraft.Version) (*model.Minecraft, error)

	// GetMinecraft retrieves a specific minecraft from the database.
	GetMinecraft(string) (*model.Minecraft, *gorm.DB)

	// GetMinecraftBuilds retrieves builds for a minecraft.
	GetMinecraftBuilds(*model.MinecraftBuildParams) (*model.Builds, error)

	// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
	GetMinecraftHasBuild(*model.MinecraftBuildParams) bool

	CreateMinecraftBuild(*model.MinecraftBuildParams) error

	DeleteMinecraftBuild(*model.MinecraftBuildParams) error

	// GetForges retrieves all available forges from the database.
	GetForges() (*model.Forges, error)

	// SyncForge creates or updates a forge record.
	SyncForge(*forge.Number) (*model.Forge, error)

	// GetForge retrieves a specific forge from the database.
	GetForge(string) (*model.Forge, *gorm.DB)

	// GetForgeBuilds retrieves builds for a forge.
	GetForgeBuilds(*model.ForgeBuildParams) (*model.Builds, error)

	// GetForgeHasBuild checks if a specific build is assigned to a forge.
	GetForgeHasBuild(*model.ForgeBuildParams) bool

	CreateForgeBuild(*model.ForgeBuildParams) error

	DeleteForgeBuild(*model.ForgeBuildParams) error

	GetSolderPacks() (*model.Packs, error)

	GetSolderPack(string, string) (*model.Pack, error)

	GetSolderBuild(string, string, string) (*model.Build, error)
}
