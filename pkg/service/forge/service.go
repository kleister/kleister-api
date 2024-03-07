package forge

import (
	"context"
	"errors"

	"github.com/kleister/go-forge/version"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a forge was not found.
	ErrNotFound = errors.New("forge not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("forge version service is unavailable")
)

// Service handles all interactions with forge.
type Service interface {
	Search(context.Context, string) ([]*model.Forge, error)
	Show(context.Context, string) (*model.Forge, error)
	Update(context.Context) error
}

// Store defines the functions to persist records.
type Store interface {
	Search(context.Context, string) ([]*model.Forge, error)
	Show(context.Context, string) (*model.Forge, error)
	Sync(context.Context, version.Versions) error
}

type service struct {
	forge Store
}

// NewService returns a Service that handles all interactions with forge.
func NewService(forge Store) Service {
	return &service{
		forge: forge,
	}
}

// Search implements the Service interface.
func (s *service) Search(ctx context.Context, search string) ([]*model.Forge, error) {
	return s.forge.Search(ctx, search)
}

// Search implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Forge, error) {
	return s.forge.Show(ctx, id)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context) error {
	result, err := version.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "forge").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	version.ByVersion(
		result.Releases,
	).Sort()

	return s.forge.Sync(ctx, result.Releases.Filter(
		&version.Filter{
			Minecraft: ">=1.7.10",
		},
	))
}
