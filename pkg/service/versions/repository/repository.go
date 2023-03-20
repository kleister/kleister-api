package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrVersionNotFound defines the error if a version could not be found.
	ErrVersionNotFound = errors.New("version not found")
)

// VersionsRepository defines the required functions for the repository.
type VersionsRepository interface {
	List(context.Context, string) ([]*model.Version, error)
	Create(context.Context, *model.Version) (*model.Version, error)
	Update(context.Context, *model.Version) (*model.Version, error)
	Show(context.Context, string, string) (*model.Version, error)
	Delete(context.Context, string, string) error
	Exists(context.Context, string, string) (bool, string, error)

	ListBuilds(context.Context, string, string, string) ([]*model.Build, error)
	AttachBuild(context.Context, string, string, string, string) error
	DropBuild(context.Context, string, string, string, string) error
}
