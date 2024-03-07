package mods

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a mod was not found.
	ErrNotFound = errors.New("mod not found")

	// ErrAlreadyAssigned is returned when a mod is already assigned.
	ErrAlreadyAssigned = errors.New("mod is already assigned")

	// ErrNotAssigned is returned when a mod is not assigned.
	ErrNotAssigned = errors.New("mod is not assigned")
)

// Service handles all interactions with mods.
type Service interface {
	List(context.Context) ([]*model.Mod, error)
	Show(context.Context, string) (*model.Mod, error)
	Create(context.Context, *model.Mod) (*model.Mod, error)
	Update(context.Context, *model.Mod) (*model.Mod, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
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

// List implements the Service interface.
func (s *service) List(ctx context.Context) ([]*model.Mod, error) {
	return s.mods.List(ctx)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Mod, error) {
	return s.mods.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	return s.mods.Create(ctx, mod)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
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
