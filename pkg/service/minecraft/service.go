package minecraft

import (
	"context"
	"errors"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog/log"
)

var (
	// ErrNotFound is returned when a minecraft was not found.
	ErrNotFound = errors.New("minecraft not found")

	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("minecraft version service is unavailable")
)

// Service handles all interactions with minecrafts.
type Service interface {
	List(context.Context) ([]*model.Minecraft, error)
	Update(context.Context) error
}

// Store defines the interface to persist minecrafts.
type Store interface {
	List(context.Context) ([]*model.Minecraft, error)
	Sync(context.Context, version.Versions) error
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

func (s *service) Update(ctx context.Context) error {
	result, err := version.FromDefault()

	if err != nil {
		log.Debug().
			Str("service", "minecraft").
			Str("request", requestid.Get(ctx)).
			Str("method", "update").
			Err(err).
			Msg("failed to sync versions")

		return ErrSyncUnavailable
	}

	return s.minecrafts.Sync(ctx, result.Releases)
}
