package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrForgeNotFound defines the error if a forge could not be found.
	ErrForgeNotFound = errors.New("forge not found")
)

// ForgeRepository defines the required functions for the repository.
type ForgeRepository interface {
	Search(context.Context, string) ([]*model.Forge, error)
	Update(context.Context) error

	ListBuilds(context.Context, string, string) ([]*model.Build, error)
}
