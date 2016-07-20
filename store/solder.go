package store

import (
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetSolderPacks retrieves all available packs optimized for the solder compatible API.
func GetSolderPacks(c context.Context) (*model.Packs, error) {
	return FromContext(c).GetSolderPacks()
}

// GetSolderPack retrieves a specific pack optimized for the solder compatible API.
func GetSolderPack(c context.Context, pack, location string) (*model.Pack, error) {
	return FromContext(c).GetSolderPack(pack, location)
}

// GetSolderBuild retrieves a specific build optimized for the solder compatible API.
func GetSolderBuild(c context.Context, pack, build, location string) (*model.Build, error) {
	return FromContext(c).GetSolderBuild(pack, build, location)
}
