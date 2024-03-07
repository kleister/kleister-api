package packs

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a pack was not found.
	ErrNotFound = errors.New("pack not found")

	// ErrAlreadyAssigned is returned when a pack is already assigned.
	ErrAlreadyAssigned = errors.New("pack is already assigned")

	// ErrNotAssigned is returned when a pack is not assigned.
	ErrNotAssigned = errors.New("pack is not assigned")
)

// Service handles all interactions with packs.
type Service interface {
	List(context.Context) ([]*model.Pack, error)
	Show(context.Context, string) (*model.Pack, error)
	Create(context.Context, *model.Pack) (*model.Pack, error)
	Update(context.Context, *model.Pack) (*model.Pack, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
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

// List implements the Service interface.
func (s *service) List(ctx context.Context) ([]*model.Pack, error) {
	return s.packs.List(ctx)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Pack, error) {
	return s.packs.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	return s.packs.Create(ctx, pack)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
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
