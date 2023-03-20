package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrMinecraftNotFound defines the error if a minecraft could not be found.
	ErrMinecraftNotFound = errors.New("minecraft not found")
)

// MinecraftRepository defines the required functions for the repository.
type MinecraftRepository interface {
	Search(context.Context, string) ([]*model.Minecraft, error)
	Update(context.Context) error

	ListBuilds(context.Context, string, string) ([]*model.Build, error)
}
