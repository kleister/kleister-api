package versions

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a version was not found.
	ErrNotFound = errors.New("version not found")
)

// Service handles all interactions with versions.
type Service interface {
	List(context.Context, model.VersionParams) ([]*model.Version, int64, error)
	Show(context.Context, model.VersionParams) (*model.Version, error)
	Create(context.Context, model.VersionParams, *model.Version) error
	Update(context.Context, model.VersionParams, *model.Version) error
	Delete(context.Context, model.VersionParams) error
	Exists(context.Context, model.VersionParams) (bool, error)
	Column(context.Context, model.VersionParams, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	versions Service
}

// NewService returns a Service that handles all interactions with versions.
func NewService(versions Service) Service {
	return &service{
		versions: versions,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.versions.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.VersionParams) ([]*model.Version, int64, error) {
	return s.versions.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, params model.VersionParams) (*model.Version, error) {
	return s.versions.Show(ctx, params)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, params model.VersionParams, version *model.Version) error {
	return s.versions.Create(ctx, params, version)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, params model.VersionParams, version *model.Version) error {
	return s.versions.Update(ctx, params, version)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, params model.VersionParams) error {
	return s.versions.Delete(ctx, params)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, params model.VersionParams) (bool, error) {
	return s.versions.Exists(ctx, params)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, params model.VersionParams, col string, val any) error {
	return s.versions.Column(ctx, params, col, val)
}
