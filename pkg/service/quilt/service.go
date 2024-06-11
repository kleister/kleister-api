package quilt

import (
	"context"
	"errors"

	quiltClient "github.com/kleister/kleister-api/pkg/internal/quilt"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a quilt was not found.
	ErrNotFound = errors.New("quilt not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("quilt version service is unavailable")

	// ErrAlreadyAssigned defines the error if a build is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a build is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with quilt.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Quilt, int64, error)
	Show(context.Context, string) (*model.Quilt, error)
	Sync(context.Context, quiltClient.Versions) error
	ListBuilds(context.Context, model.QuiltBuildParams) ([]*model.Build, int64, error)
	AttachBuild(context.Context, model.QuiltBuildParams) error
	DropBuild(context.Context, model.QuiltBuildParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	quilt Service
}

// NewService returns a Service that handles all interactions with quilt.
func NewService(quilt Service) Service {
	return &service{
		quilt: quilt,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.quilt.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Quilt, int64, error) {
	return s.quilt.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Quilt, error) {
	return s.quilt.Show(ctx, id)
}

// Sync implements the Service interface.
func (s *service) Sync(ctx context.Context, versions quiltClient.Versions) error {
	return s.quilt.Sync(ctx, versions)
}

// ListBuilds implements the Service interface.
func (s *service) ListBuilds(ctx context.Context, params model.QuiltBuildParams) ([]*model.Build, int64, error) {
	return s.quilt.ListBuilds(ctx, params)
}

// AttachBuild implements the Service interface.
func (s *service) AttachBuild(ctx context.Context, params model.QuiltBuildParams) error {
	return s.quilt.AttachBuild(ctx, params)
}

// DropBuild implements the Service interface.
func (s *service) DropBuild(ctx context.Context, params model.QuiltBuildParams) error {
	return s.quilt.DropBuild(ctx, params)
}

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (quiltClient.Versions, error) {
	result, err := quiltClient.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "quilt").
			Str("method", "fetch").
			Msg("Failed to sync versions")

		return nil, ErrSyncUnavailable
	}

	quiltClient.ByVersion(
		result.Versions,
	).Sort()

	return result.Versions, nil
}
