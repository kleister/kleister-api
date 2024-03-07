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
	List(context.Context, string) ([]*model.Version, error)
	Show(context.Context, string, string) (*model.Version, error)
	Create(context.Context, string, *model.Version) (*model.Version, error)
	Update(context.Context, string, *model.Version) (*model.Version, error)
	Delete(context.Context, string, string) error
	Exists(context.Context, string, string) (bool, error)
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

// List implements the Service interface.
func (s *service) List(ctx context.Context, modID string) ([]*model.Version, error) {
	return s.versions.List(ctx, modID)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, modID, id string) (*model.Version, error) {
	return s.versions.Show(ctx, modID, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	return s.versions.Create(ctx, modID, version)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	return s.versions.Update(ctx, modID, version)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, modID, name string) error {
	return s.versions.Delete(ctx, modID, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, modID, name string) (bool, error) {
	return s.versions.Exists(ctx, modID, name)
}
