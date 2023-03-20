package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrBuildNotFound defines the error if a build could not be found.
	ErrBuildNotFound = errors.New("build not found")
)

// BuildsRepository defines the required functions for the repository.
type BuildsRepository interface {
	List(context.Context, string) ([]*model.Build, error)
	Create(context.Context, *model.Build) (*model.Build, error)
	Update(context.Context, *model.Build) (*model.Build, error)
	Show(context.Context, string, string) (*model.Build, error)
	Delete(context.Context, string, string) error
	Exists(context.Context, string, string) (bool, string, error)

	ListVersions(context.Context, string, string, string) ([]*model.Version, error)
	AttachVersion(context.Context, string, string, string, string) error
	DropVersion(context.Context, string, string, string, string) error
}
