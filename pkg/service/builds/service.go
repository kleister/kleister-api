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
	List(context.Context, string) ([]*model.Build, error)
	Show(context.Context, string, string) (*model.Build, error)
	Create(context.Context, string, *model.Build) (*model.Build, error)
	Update(context.Context, string, *model.Build) (*model.Build, error)
	Delete(context.Context, string, string) error
	Exists(context.Context, string, string) (bool, error)
	Column(context.Context, string, string, string, any) error
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

// List implements the Service interface.
func (s *service) List(ctx context.Context, packID string) ([]*model.Build, error) {
	return s.builds.List(ctx, packID)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, packID, id string) (*model.Build, error) {
	return s.builds.Show(ctx, packID, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	return s.builds.Create(ctx, packID, build)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	return s.builds.Update(ctx, packID, build)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, packID, name string) error {
	return s.builds.Delete(ctx, packID, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, packID, name string) (bool, error) {
	return s.builds.Exists(ctx, packID, name)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, packID, id, col string, val any) error {
	return s.builds.Column(ctx, packID, id, col, val)
}
