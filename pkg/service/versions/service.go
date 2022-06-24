package versions

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a version was not found.
	ErrNotFound = errors.New("version not found")

	// ErrAlreadyAssigned is returned when a version is already assigned.
	ErrAlreadyAssigned = errors.New("version is already assigned")

	// ErrNotAssigned is returned when a version is not assigned.
	ErrNotAssigned = errors.New("version is not assigned")
)

// Service handles all interactions with versions.
type Service interface {
	List(context.Context, *model.Mod) ([]*model.Version, error)
	Show(context.Context, *model.Mod, string) (*model.Version, error)
	Create(context.Context, *model.Mod, *model.Version) (*model.Version, error)
	Update(context.Context, *model.Mod, *model.Version) (*model.Version, error)
	Delete(context.Context, *model.Mod, string) error
}

// Store defines the interface to persist versions.
type Store interface {
	List(context.Context, string) ([]*model.Version, error)
	Show(context.Context, string, string) (*model.Version, error)
	Create(context.Context, string, *model.Version) (*model.Version, error)
	Update(context.Context, string, *model.Version) (*model.Version, error)
	Delete(context.Context, string, string) error
}

type service struct {
	versions Store
}

// NewService returns a Service that handles all interactions with versions.
func NewService(versions Store) Service {
	return &service{
		versions: versions,
	}
}

func (s *service) List(ctx context.Context, mod *model.Mod) ([]*model.Version, error) {
	return s.versions.List(ctx, mod.ID)
}

func (s *service) Show(ctx context.Context, mod *model.Mod, id string) (*model.Version, error) {
	return s.versions.Show(ctx, mod.ID, id)
}

func (s *service) Create(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	return s.versions.Create(ctx, mod.ID, version)
}

func (s *service) Update(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	return s.versions.Update(ctx, mod.ID, version)
}

func (s *service) Delete(ctx context.Context, mod *model.Mod, name string) error {
	return s.versions.Delete(ctx, mod.ID, name)
}
