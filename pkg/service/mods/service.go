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
	List(context.Context) ([]*model.Mod, error)
	Show(context.Context, string) (*model.Mod, error)
	Create(context.Context, *model.Mod) (*model.Mod, error)
	Update(context.Context, *model.Mod) (*model.Mod, error)
	Delete(context.Context, string) error
}

// Store defines the interface to persist mods.
type Store interface {
	List(context.Context) ([]*model.Mod, error)
	Show(context.Context, string) (*model.Mod, error)
	Create(context.Context, *model.Mod) (*model.Mod, error)
	Update(context.Context, *model.Mod) (*model.Mod, error)
	Delete(context.Context, string) error
}

type service struct {
	mods Store
}

// NewService returns a Service that handles all interactions with mods.
func NewService(mods Store) Service {
	return &service{
		mods: mods,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Mod, error) {
	return s.mods.List(ctx)
}

func (s *service) Show(ctx context.Context, id string) (*model.Mod, error) {
	return s.mods.Show(ctx, id)
}

func (s *service) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	return s.mods.Create(ctx, mod)
}

func (s *service) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	return s.mods.Update(ctx, mod)
}

func (s *service) Delete(ctx context.Context, name string) error {
	return s.mods.Delete(ctx, name)
}
