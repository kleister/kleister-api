package store

import (
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

func GetSolderPacks(c context.Context) (*model.Packs, error) {
	return FromContext(c).GetSolderPacks()
}

func GetSolderPack(c context.Context, pack, location string) (*model.Pack, error) {
	return FromContext(c).GetSolderPack(pack, location)
}

func GetSolderBuild(c context.Context, pack, build, location string) (*model.Build, error) {
	return FromContext(c).GetSolderBuild(pack, build, location)
}
