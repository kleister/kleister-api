package packs

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a pack was not found.
	ErrNotFound = errors.New("pack not found")
)

// Service handles all interactions with packs.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Pack, int64, error)
	Show(context.Context, string) (*model.Pack, error)
	Create(context.Context, *model.Pack) error
	Update(context.Context, *model.Pack) error
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
	Column(context.Context, string, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	packs Service
}

// NewService returns a Service that handles all interactions with packs.
func NewService(packs Service) Service {
	return &service{
		packs: packs,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.packs.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Pack, int64, error) {
	return s.packs.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Pack, error) {
	return s.packs.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, pack *model.Pack) error {
	return s.packs.Create(ctx, pack)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, pack *model.Pack) error {
	return s.packs.Update(ctx, pack)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, name string) error {
	return s.packs.Delete(ctx, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, name string) (bool, error) {
	return s.packs.Exists(ctx, name)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, name, col string, val any) error {
	return s.packs.Column(ctx, name, col, val)
}
