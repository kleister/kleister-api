package model

// ForgeBuildParams represents the parameters to connect builds with forges.
type ForgeBuildParams struct {
	Forge string
	Pack  string
	Build string
}

// MinecraftBuildParams represents the parameters to connect builds with minecrafts.
type MinecraftBuildParams struct {
	Minecraft string
	Pack      string
	Build     string
}

// BuildVersionParams represents the parameters to connect versions with builds.
type BuildVersionParams struct {
	Pack    string
	Build   string
	Mod     string
	Version string
}

// VersionBuildParams represents the parameters to connect builds with versions.
type VersionBuildParams struct {
	Mod     string
	Version string
	Pack    string
	Build   string
}

// PackClientParams represents the parameters to connect clients with packs.
type PackClientParams struct {
	Pack   string
	Client string
}

// ClientPackParams represents the parameters to connect packs with clients.
type ClientPackParams struct {
	Client string
	Pack   string
}

// PackUserParams represents the parameters to connect users with packs.
type PackUserParams struct {
	Pack string
	User string
}

// UserPackParams represents the parameters to connect packs with users.
type UserPackParams struct {
	User string
	Pack string
}

// ModUserParams represents the parameters to connect users with mods.
type ModUserParams struct {
	Mod  string
	User string
}

// UserModParams represents the parameters to connect mods with users.
type UserModParams struct {
	User string
	Mod  string
}
