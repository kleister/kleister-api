package minecraft

import (
	"context"
	"errors"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a minecraft was not found.
	ErrNotFound = errors.New("minecraft not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("minecraft version service is unavailable")

	// ErrAlreadyAssigned defines the error if a build is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a build is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with minecraft.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Minecraft, int64, error)
	Show(context.Context, string) (*model.Minecraft, error)
	Sync(context.Context, version.Versions) error
	ListBuilds(context.Context, model.MinecraftBuildParams) ([]*model.Build, int64, error)
	AttachBuild(context.Context, model.MinecraftBuildParams) error
	DropBuild(context.Context, model.MinecraftBuildParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	minecraft Service
}

// NewService returns a Service that handles all interactions with minecraft.
func NewService(minecraft Service) Service {
	return &service{
		minecraft: minecraft,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.minecraft.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Minecraft, int64, error) {
	return s.minecraft.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Minecraft, error) {
	return s.minecraft.Show(ctx, id)
}

// Sync implements the Service interface.
func (s *service) Sync(ctx context.Context, versions version.Versions) error {
	return s.minecraft.Sync(ctx, versions)
}

// ListBuilds implements the Service interface.
func (s *service) ListBuilds(ctx context.Context, params model.MinecraftBuildParams) ([]*model.Build, int64, error) {
	return s.minecraft.ListBuilds(ctx, params)
}

// AttachBuild implements the Service interface.
func (s *service) AttachBuild(ctx context.Context, params model.MinecraftBuildParams) error {
	return s.minecraft.AttachBuild(ctx, params)
}

// DropBuild implements the Service interface.
func (s *service) DropBuild(ctx context.Context, params model.MinecraftBuildParams) error {
	return s.minecraft.DropBuild(ctx, params)
}

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (version.Versions, error) {
	result, err := version.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "minecraft").
			Str("method", "fetch").
			Msg("Failed to fetch versions")

		return nil, ErrSyncUnavailable
	}

	version.ByVersion(
		result.Releases,
	).Sort()

	return result.Releases.Filter(
		&version.Filter{
			Version: ">=1.7.10",
		},
	), nil
}
