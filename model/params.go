package model

type ForgeBuildParams struct {
	Forge string
	Pack  string
	Build string
}

type MinecraftBuildParams struct {
	Minecraft string
	Pack      string
	Build     string
}

type BuildVersionParams struct {
	Pack    string
	Build   string
	Mod     string
	Version string
}

type VersionBuildParams struct {
	Mod     string
	Version string
	Pack    string
	Build   string
}

type PackClientParams struct {
	Pack   string
	Client string
}

type ClientPackParams struct {
	Client string
	Pack   string
}

type PackUserParams struct {
	Pack string
	User string
}

type UserPackParams struct {
	User string
	Pack string
}

type ModUserParams struct {
	Mod  string
	User string
}

type UserModParams struct {
	User string
	Mod  string
}
