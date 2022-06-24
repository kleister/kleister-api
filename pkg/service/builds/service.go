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
	List(context.Context, *model.Pack) ([]*model.Build, error)
	Show(context.Context, *model.Pack, string) (*model.Build, error)
	Create(context.Context, *model.Pack, *model.Build) (*model.Build, error)
	Update(context.Context, *model.Pack, *model.Build) (*model.Build, error)
	Delete(context.Context, *model.Pack, string) error
}

// Store defines the interface to persist builds.
type Store interface {
	List(context.Context, string) ([]*model.Build, error)
	Show(context.Context, string, string) (*model.Build, error)
	Create(context.Context, string, *model.Build) (*model.Build, error)
	Update(context.Context, string, *model.Build) (*model.Build, error)
	Delete(context.Context, string, string) error
}

type service struct {
	builds Store
}

// NewService returns a Service that handles all interactions with builds.
func NewService(builds Store) Service {
	return &service{
		builds: builds,
	}
}

func (s *service) List(ctx context.Context, pack *model.Pack) ([]*model.Build, error) {
	return s.builds.List(ctx, pack.ID)
}

func (s *service) Show(ctx context.Context, pack *model.Pack, id string) (*model.Build, error) {
	return s.builds.Show(ctx, pack.ID, id)
}

func (s *service) Create(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	return s.builds.Create(ctx, pack.ID, build)
}

func (s *service) Update(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	return s.builds.Update(ctx, pack.ID, build)
}

func (s *service) Delete(ctx context.Context, pack *model.Pack, name string) error {
	return s.builds.Delete(ctx, pack.ID, name)
}
