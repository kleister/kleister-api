package minecraft

import (
	"context"
	"errors"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a minecraft was not found.
	ErrNotFound = errors.New("minecraft not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("minecraft version service is unavailable")
)

// Service handles all interactions with minecraft.
type Service interface {
	Search(context.Context, string) ([]*model.Minecraft, error)
	Show(context.Context, string) (*model.Minecraft, error)
	Update(context.Context) error
}

// Store defines the functions to persist records.
type Store interface {
	Search(context.Context, string) ([]*model.Minecraft, error)
	Show(context.Context, string) (*model.Minecraft, error)
	Sync(context.Context, version.Versions) error
}

type service struct {
	minecraft Store
}

// NewService returns a Service that handles all interactions with minecraft.
func NewService(minecraft Store) Service {
	return &service{
		minecraft: minecraft,
	}
}

// Search implements the Service interface.
func (s *service) Search(ctx context.Context, search string) ([]*model.Minecraft, error) {
	return s.minecraft.Search(ctx, search)
}

// Search implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Minecraft, error) {
	return s.minecraft.Show(ctx, id)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context) error {
	result, err := version.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "minecraft").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	version.ByVersion(
		result.Releases,
	).Sort()

	return s.minecraft.Sync(ctx, result.Releases.Filter(
		&version.Filter{
			Version: ">=1.7.10",
		},
	))
}
