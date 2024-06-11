package neoforge

import (
	"context"
	"errors"

	neoforgeClient "github.com/kleister/kleister-api/pkg/internal/neoforge"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a neoforge was not found.
	ErrNotFound = errors.New("neoforge not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("neoforge version service is unavailable")

	// ErrAlreadyAssigned defines the error if a build is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a build is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with neoforge.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Neoforge, int64, error)
	Show(context.Context, string) (*model.Neoforge, error)
	Sync(context.Context, neoforgeClient.Versions) error
	ListBuilds(context.Context, model.NeoforgeBuildParams) ([]*model.Build, int64, error)
	AttachBuild(context.Context, model.NeoforgeBuildParams) error
	DropBuild(context.Context, model.NeoforgeBuildParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	neoforge Service
}

// NewService returns a Service that handles all interactions with neoforge.
func NewService(neoforge Service) Service {
	return &service{
		neoforge: neoforge,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.neoforge.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Neoforge, int64, error) {
	return s.neoforge.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Neoforge, error) {
	return s.neoforge.Show(ctx, id)
}

// Sync implements the Service interface.
func (s *service) Sync(ctx context.Context, versions neoforgeClient.Versions) error {
	return s.neoforge.Sync(ctx, versions)
}

// ListBuilds implements the Service interface.
func (s *service) ListBuilds(ctx context.Context, params model.NeoforgeBuildParams) ([]*model.Build, int64, error) {
	return s.neoforge.ListBuilds(ctx, params)
}

// AttachBuild implements the Service interface.
func (s *service) AttachBuild(ctx context.Context, params model.NeoforgeBuildParams) error {
	return s.neoforge.AttachBuild(ctx, params)
}

// DropBuild implements the Service interface.
func (s *service) DropBuild(ctx context.Context, params model.NeoforgeBuildParams) error {
	return s.neoforge.DropBuild(ctx, params)
}

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (neoforgeClient.Versions, error) {
	result, err := neoforgeClient.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "neoforge").
			Str("method", "fetch").
			Msg("Failed to sync versions")

		return nil, ErrSyncUnavailable
	}

	neoforgeClient.ByVersion(
		result.Versions,
	).Sort()

	return result.Versions, nil
}
