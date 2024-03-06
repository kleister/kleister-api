package neoforge

import (
	"context"
	"errors"

	neoforgeClient "github.com/kleister/kleister-api/pkg/internal/neoforge"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a neoforge was not found.
	ErrNotFound = errors.New("neoforge not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("neoforge version service is unavailable")
)

// Service handles all interactions with neoforge.
type Service interface {
	Search(context.Context, string) ([]*model.Neoforge, error)
	Update(context.Context) error
}

// Store defines the functions to persist records.
type Store interface {
	Search(context.Context, string) ([]*model.Neoforge, error)
	Sync(context.Context, neoforgeClient.Versions) error
}

type service struct {
	neoforge Store
}

// NewService returns a Service that handles all interactions with neoforge.
func NewService(neoforge Store) Service {
	return &service{
		neoforge: neoforge,
	}
}

// Search implements the Service interface.
func (s *service) Search(ctx context.Context, search string) ([]*model.Neoforge, error) {
	return s.neoforge.Search(ctx, search)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context) error {
	result, err := neoforgeClient.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "neoforge").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	neoforgeClient.ByVersion(
		result.Versions,
	).Sort()

	return s.neoforge.Sync(ctx, result.Versions)
}
