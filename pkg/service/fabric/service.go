package fabric

import (
	"context"
	"errors"

	fabricClient "github.com/kleister/kleister-api/pkg/internal/fabric"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a fabric was not found.
	ErrNotFound = errors.New("fabric not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("fabric version service is unavailable")
)

// Service handles all interactions with fabric.
type Service interface {
	Search(context.Context, string) ([]*model.Fabric, error)
	Update(context.Context) error
}

// Store defines the functions to persist records.
type Store interface {
	Search(context.Context, string) ([]*model.Fabric, error)
	Sync(context.Context, fabricClient.Versions) error
}

type service struct {
	fabric Store
}

// NewService returns a Service that handles all interactions with fabric.
func NewService(fabric Store) Service {
	return &service{
		fabric: fabric,
	}
}

// Search implements the Service interface.
func (s *service) Search(ctx context.Context, search string) ([]*model.Fabric, error) {
	return s.fabric.Search(ctx, search)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context) error {
	result, err := fabricClient.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "fabric").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	fabricClient.ByVersion(
		result.Versions,
	).Sort()

	return s.fabric.Sync(ctx, result.Versions)
}
