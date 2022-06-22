package forge

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a forge was not found.
	ErrNotFound = errors.New("forge not found")
)

// Service handles all interactions with forges.
type Service interface {
	List(context.Context) ([]*model.Forge, error)
}

// Store defines the interface to persist forges.
type Store interface {
	List(context.Context) ([]*model.Forge, error)
}

type service struct {
	forges Store
}

// NewService returns a Service that handles all interactions with forges.
func NewService(s Store) Service {
	return &service{
		forges: s,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Forge, error) {
	return s.forges.List(ctx)
}
