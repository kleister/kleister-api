package model

//go:generate mockery -all -case=underscore

// StoreAPI describes a store API.
type StoreAPI interface {
	// GetBuilds retrieves all available builds from the database.
	GetBuilds(int) (*Builds, error)

	// GetBuild retrieves a specific build from the database.
	GetBuild(int, string) (*Build, *Store)

	// GetClients retrieves all available clients from the database.
	GetClients() (*Clients, error)

	// GetClient retrieves a specific client from the database.
	GetClient(string) (*Client, *Store)

	// GetForges retrieves all available Forge versions from the database.
	GetForges() (*Forges, error)

	// GetForge retrieves a specific Forge version from the database.
	GetForge(string) (*Forge, *Store)

	// GetKeys retrieves all available keys from the database.
	GetKeys() (*Keys, error)

	// GetKey retrieves a specific key from the database.
	GetKey(string) (*Key, *Store)

	// GetMinecrafts retrieves all available Minecraft versions from the database.
	GetMinecrafts() (*Minecrafts, error)

	// GetMinecraft retrieves a specific Minecraft version from the database.
	GetMinecraft(string) (*Minecraft, *Store)

	// GetMods retrieves all available mods from the database.
	GetMods() (*Mods, error)

	// GetMod retrieves a specific mod from the database.
	GetMod(string) (*Mod, *Store)

	// GetPacks retrieves all available packs from the database.
	GetPacks() (*Packs, error)

	// GetPack retrieves a specific pack from the database.
	GetPack(string) (*Pack, *Store)

	// GetUsers retrieves all available users from the database.
	GetUsers() (*Users, error)

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*User, *Store)

	// GetVersions retrieves all available versions from the database.
	GetVersions(int) (*Versions, error)

	// GetVersion retrieves a specific version from the database.
	GetVersion(int, string) (*Version, *Store)
}
