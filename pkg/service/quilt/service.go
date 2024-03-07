package quilt

import (
	"context"
	"errors"

	quiltClient "github.com/kleister/kleister-api/pkg/internal/quilt"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a quilt was not found.
	ErrNotFound = errors.New("quilt not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("quilt version service is unavailable")
)

// Service handles all interactions with quilt.
type Service interface {
	Search(context.Context, string) ([]*model.Quilt, error)
	Show(context.Context, string) (*model.Quilt, error)
	Update(context.Context) error
}

// Store defines the functions to persist records.
type Store interface {
	Search(context.Context, string) ([]*model.Quilt, error)
	Show(context.Context, string) (*model.Quilt, error)
	Sync(context.Context, quiltClient.Versions) error
}

type service struct {
	quilt Store
}

// NewService returns a Service that handles all interactions with quilt.
func NewService(quilt Store) Service {
	return &service{
		quilt: quilt,
	}
}

// Search implements the Service interface.
func (s *service) Search(ctx context.Context, search string) ([]*model.Quilt, error) {
	return s.quilt.Search(ctx, search)
}

// Search implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Quilt, error) {
	return s.quilt.Show(ctx, id)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context) error {
	result, err := quiltClient.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "quilt").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	quiltClient.ByVersion(
		result.Versions,
	).Sort()

	return s.quilt.Sync(ctx, result.Versions)
}
