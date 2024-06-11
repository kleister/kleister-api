package mods

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a mod was not found.
	ErrNotFound = errors.New("mod not found")
)

// Service handles all interactions with mods.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Mod, int64, error)
	Show(context.Context, string) (*model.Mod, error)
	Create(context.Context, *model.Mod) error
	Update(context.Context, *model.Mod) error
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
	Column(context.Context, string, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	mods Service
}

// NewService returns a Service that handles all interactions with mods.
func NewService(mods Service) Service {
	return &service{
		mods: mods,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.mods.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Mod, int64, error) {
	return s.mods.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Mod, error) {
	return s.mods.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, mod *model.Mod) error {
	return s.mods.Create(ctx, mod)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, mod *model.Mod) error {
	return s.mods.Update(ctx, mod)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, name string) error {
	return s.mods.Delete(ctx, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, name string) (bool, error) {
	return s.mods.Exists(ctx, name)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, name, col string, val any) error {
	return s.mods.Column(ctx, name, col, val)
}
