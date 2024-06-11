package builds

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a build was not found.
	ErrNotFound = errors.New("build not found")
)

// Service handles all interactions with builds.
type Service interface {
	List(context.Context, model.BuildParams) ([]*model.Build, int64, error)
	Show(context.Context, model.BuildParams) (*model.Build, error)
	Create(context.Context, model.BuildParams, *model.Build) error
	Update(context.Context, model.BuildParams, *model.Build) error
	Delete(context.Context, model.BuildParams) error
	Exists(context.Context, model.BuildParams) (bool, error)
	Column(context.Context, model.BuildParams, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	builds Service
}

// NewService returns a Service that handles all interactions with builds.
func NewService(builds Service) Service {
	return &service{
		builds: builds,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.builds.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.BuildParams) ([]*model.Build, int64, error) {
	return s.builds.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, params model.BuildParams) (*model.Build, error) {
	return s.builds.Show(ctx, params)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, params model.BuildParams, build *model.Build) error {
	return s.builds.Create(ctx, params, build)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, params model.BuildParams, build *model.Build) error {
	return s.builds.Update(ctx, params, build)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, params model.BuildParams) error {
	return s.builds.Delete(ctx, params)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, params model.BuildParams) (bool, error) {
	return s.builds.Exists(ctx, params)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, params model.BuildParams, col string, val any) error {
	return s.builds.Column(ctx, params, col, val)
}
