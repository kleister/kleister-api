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
	List(context.Context) ([]*model.Pack, error)
	Show(context.Context, string) (*model.Pack, error)
	Create(context.Context, *model.Pack) (*model.Pack, error)
	Update(context.Context, *model.Pack) (*model.Pack, error)
	Delete(context.Context, string) error
}

// Store defines the interface to persist packs.
type Store interface {
	List(context.Context) ([]*model.Pack, error)
	Show(context.Context, string) (*model.Pack, error)
	Create(context.Context, *model.Pack) (*model.Pack, error)
	Update(context.Context, *model.Pack) (*model.Pack, error)
	Delete(context.Context, string) error
}

type service struct {
	packs Store
}

// NewService returns a Service that handles all interactions with packs.
func NewService(packs Store) Service {
	return &service{
		packs: packs,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Pack, error) {
	return s.packs.List(ctx)
}

func (s *service) Show(ctx context.Context, id string) (*model.Pack, error) {
	return s.packs.Show(ctx, id)
}

func (s *service) Create(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	return s.packs.Create(ctx, pack)
}

func (s *service) Update(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	return s.packs.Update(ctx, pack)
}

func (s *service) Delete(ctx context.Context, name string) error {
	return s.packs.Delete(ctx, name)
}
