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

// Service handles all interactions with forges.
type Service interface {
	List(context.Context) ([]*model.Forge, error)
	Update(context.Context) error
}

// Store defines the interface to persist forges.
type Store interface {
	List(context.Context) ([]*model.Forge, error)
	Sync(context.Context, version.Versions) error
}

type service struct {
	forges Store
}

// NewService returns a Service that handles all interactions with forges.
func NewService(s Store) Service {
	return &service{
		forges: s,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Forge, error) {
	return s.forges.List(ctx)
}

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

	return s.forges.Sync(ctx, result.Releases)
}
