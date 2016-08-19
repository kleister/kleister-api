package model

// ForgeBuildParams represents the parameters to connect builds with forges.
type ForgeBuildParams struct {
	Forge string `json:"forge"`
	Pack  string `json:"pack"`
	Build string `json:"build"`
}

// MinecraftBuildParams represents the parameters to connect builds with minecrafts.
type MinecraftBuildParams struct {
	Minecraft string `json:"minecraft"`
	Pack      string `json:"pack"`
	Build     string `json:"build"`
}

// BuildVersionParams represents the parameters to connect versions with builds.
type BuildVersionParams struct {
	Pack    string `json:"pack"`
	Build   string `json:"build"`
	Mod     string `json:"mod"`
	Version string `json:"version"`
}

// VersionBuildParams represents the parameters to connect builds with versions.
type VersionBuildParams struct {
	Mod     string `json:"mod"`
	Version string `json:"version"`
	Pack    string `json:"pack"`
	Build   string `json:"build"`
}

// PackClientParams represents the parameters to connect clients with packs.
type PackClientParams struct {
	Pack   string `json:"pack"`
	Client string `json:"client"`
}

// ClientPackParams represents the parameters to connect packs with clients.
type ClientPackParams struct {
	Client string `json:"client"`
	Pack   string `json:"pack"`
}

// PackUserParams represents the parameters to connect users with packs.
type PackUserParams struct {
	Pack string `json:"pack"`
	User string `json:"user"`
}

// UserPackParams represents the parameters to connect packs with users.
type UserPackParams struct {
	User string `json:"user"`
	Pack string `json:"pack"`
}

// ModUserParams represents the parameters to connect users with mods.
type ModUserParams struct {
	Mod  string `json:"mod"`
	User string `json:"user"`
}

// UserModParams represents the parameters to connect mods with users.
type UserModParams struct {
	User string `json:"user"`
	Mod  string `json:"mod"`
}

// ModTeamParams represents the parameters to connect teams with mods.
type ModTeamParams struct {
	Mod  string `json:"mod"`
	Team string `json:"team"`
}

// TeamModParams represents the parameters to connect mods with teams.
type TeamModParams struct {
	Team string `json:"team"`
	Mod  string `json:"mod"`
}

// PackTeamParams represents the parameters to connect teams with packs.
type PackTeamParams struct {
	Pack string `json:"pack"`
	Team string `json:"team"`
}

// TeamPackParams represents the parameters to connect packs with teams.
type TeamPackParams struct {
	Team string `json:"team"`
	Pack string `json:"pack"`
}

// UserTeamParams represents the parameters to connect teams with users.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
}

// TeamUserParams represents the parameters to connect users with teams.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
}
