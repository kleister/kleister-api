package fabric

import (
	"context"
	"errors"

	fabricClient "github.com/kleister/kleister-api/pkg/internal/fabric"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a fabric was not found.
	ErrNotFound = errors.New("fabric not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("fabric version service is unavailable")

	// ErrAlreadyAssigned defines the error if a build is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a build is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with fabric.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Fabric, int64, error)
	Show(context.Context, string) (*model.Fabric, error)
	Sync(context.Context, fabricClient.Versions) error
	ListBuilds(context.Context, model.FabricBuildParams) ([]*model.Build, int64, error)
	AttachBuild(context.Context, model.FabricBuildParams) error
	DropBuild(context.Context, model.FabricBuildParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	fabric Service
}

// NewService returns a Service that handles all interactions with fabric.
func NewService(fabric Service) Service {
	return &service{
		fabric: fabric,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.fabric.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Fabric, int64, error) {
	return s.fabric.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Fabric, error) {
	return s.fabric.Show(ctx, id)
}

// Sync implements the Service interface.
func (s *service) Sync(ctx context.Context, versions fabricClient.Versions) error {
	return s.fabric.Sync(ctx, versions)
}

// ListBuilds implements the Service interface.
func (s *service) ListBuilds(ctx context.Context, params model.FabricBuildParams) ([]*model.Build, int64, error) {
	return s.fabric.ListBuilds(ctx, params)
}

// AttachBuild implements the Service interface.
func (s *service) AttachBuild(ctx context.Context, params model.FabricBuildParams) error {
	return s.fabric.AttachBuild(ctx, params)
}

// DropBuild implements the Service interface.
func (s *service) DropBuild(ctx context.Context, params model.FabricBuildParams) error {
	return s.fabric.DropBuild(ctx, params)
}

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (fabricClient.Versions, error) {
	result, err := fabricClient.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "fabric").
			Str("method", "fetch").
			Msg("Failed to sync versions")

		return nil, ErrSyncUnavailable
	}

	fabricClient.ByVersion(
		result.Versions,
	).Sort()

	return result.Versions, nil
}
