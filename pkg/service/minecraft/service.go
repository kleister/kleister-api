package minecraft

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a minecraft was not found.
	ErrNotFound = errors.New("minecraft not found")
)

// Service handles all interactions with minecrafts.
type Service interface {
	List(context.Context) ([]*model.Minecraft, error)
}

// Store defines the interface to persist minecrafts.
type Store interface {
	List(context.Context) ([]*model.Minecraft, error)
}

type service struct {
	minecrafts Store
}

// NewService returns a Service that handles all interactions with minecrafts.
func NewService(s Store) Service {
	return &service{
		minecrafts: s,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Minecraft, error) {
	return s.minecrafts.List(ctx)
}
