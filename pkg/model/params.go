package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int64
	Offset int64
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

// GroupModParams defines parameters for group mods.
type GroupModParams struct {
	ListParams

	GroupID string
	ModID   string
	Perm    string
}

// GroupPackParams defines parameters for group packs.
type GroupPackParams struct {
	ListParams

	GroupID string
	PackID  string
	Perm    string
}

// UserGroupParams defines parameters for user groups.
type UserGroupParams struct {
	ListParams

	UserID  string
	GroupID string
	Perm    string
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
