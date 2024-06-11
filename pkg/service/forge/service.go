package forge

import (
	"context"
	"errors"

	"github.com/kleister/go-forge/version"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a forge was not found.
	ErrNotFound = errors.New("forge not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("forge version service is unavailable")

	// ErrAlreadyAssigned defines the error if a build is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a build is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with forge.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Forge, int64, error)
	Show(context.Context, string) (*model.Forge, error)
	Sync(context.Context, version.Versions) error
	ListBuilds(context.Context, model.ForgeBuildParams) ([]*model.Build, int64, error)
	AttachBuild(context.Context, model.ForgeBuildParams) error
	DropBuild(context.Context, model.ForgeBuildParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	forge Service
}

// NewService returns a Service that handles all interactions with forge.
func NewService(forge Service) Service {
	return &service{
		forge: forge,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.forge.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Forge, int64, error) {
	return s.forge.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Forge, error) {
	return s.forge.Show(ctx, id)
}

// Sync implements the Service interface.
func (s *service) Sync(ctx context.Context, versions version.Versions) error {
	return s.forge.Sync(ctx, versions)
}

// ListBuilds implements the Service interface.
func (s *service) ListBuilds(ctx context.Context, params model.ForgeBuildParams) ([]*model.Build, int64, error) {
	return s.forge.ListBuilds(ctx, params)
}

// AttachBuild implements the Service interface.
func (s *service) AttachBuild(ctx context.Context, params model.ForgeBuildParams) error {
	return s.forge.AttachBuild(ctx, params)
}

// DropBuild implements the Service interface.
func (s *service) DropBuild(ctx context.Context, params model.ForgeBuildParams) error {
	return s.forge.DropBuild(ctx, params)
}

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (version.Versions, error) {
	result, err := version.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "forge").
			Str("method", "fetch").
			Msg("Failed to sync versions")

		return nil, ErrSyncUnavailable
	}

	version.ByVersion(
		result.Releases,
	).Sort()

	return result.Releases.Filter(
		&version.Filter{
			Minecraft: ">=1.7.10",
		},
	), nil
}
