package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int
	Offset int
}

// BuildParams defines parameters for builds.
type BuildParams struct {
	ListParams

	PackID  string
	BuildID string
}

// VersionParams defines parameters for versions.
type VersionParams struct {
	ListParams

	ModID     string
	VersionID string
}

// BuildVersionParams defines parameters for build versions.
type BuildVersionParams struct {
	ListParams

	PackID    string
	BuildID   string
	ModID     string
	VersionID string
	Perm      string
}

// TeamModParams defines parameters for team mods.
type TeamModParams struct {
	ListParams

	TeamID string
	ModID  string
	Perm   string
}

// TeamPackParams defines parameters for team packs.
type TeamPackParams struct {
	ListParams

	TeamID string
	PackID string
	Perm   string
}

// UserTeamParams defines parameters for user teams.
type UserTeamParams struct {
	ListParams

	UserID string
	TeamID string
	Perm   string
}

// UserModParams defines parameters for user mods.
type UserModParams struct {
	ListParams

	UserID string
	ModID  string
	Perm   string
}

// UserPackParams defines parameters for user packs.
type UserPackParams struct {
	ListParams

	UserID string
	PackID string
	Perm   string
}

// MinecraftBuildParams defines parameters for minecraft builds.
type MinecraftBuildParams struct {
	ListParams

	MinecraftID string
	PackID      string
	BuildID     string
}

// ForgeBuildParams defines parameters for forge builds.
type ForgeBuildParams struct {
	ListParams

	ForgeID string
	PackID  string
	BuildID string
}

// NeoforgeBuildParams defines parameters for neoforge builds.
type NeoforgeBuildParams struct {
	ListParams

	NeoforgeID string
	PackID     string
	BuildID    string
}

// QuiltBuildParams defines parameters for quilt builds.
type QuiltBuildParams struct {
	ListParams

	QuiltID string
	PackID  string
	BuildID string
}

// FabricBuildParams defines parameters for fabric builds.
type FabricBuildParams struct {
	ListParams

	FabricID string
	PackID   string
	BuildID  string
}
