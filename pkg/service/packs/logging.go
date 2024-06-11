package packs

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

type loggingService struct {
	service Service
	logger  zerolog.Logger
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(s Service) Service {
	return &loggingService{
		service: s,
		logger:  log.With().Str("service", "packs").Logger(),
	}
}

// External implements the Service interface for logging.
func (s *loggingService) WithPrincipal(principal *model.User) Service {
	s.service.WithPrincipal(principal)
	return s
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, params model.ListParams) ([]*model.Pack, int64, error) {
	start := time.Now()
	records, counter, err := s.service.List(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to find all packs")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, counter, err
}

// Show implements the Service interface for logging.
func (s *loggingService) Show(ctx context.Context, name string) (*model.Pack, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to find pack by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Create implements the Service interface for logging.
func (s *loggingService) Create(ctx context.Context, pack *model.Pack) error {
	start := time.Now()
	err := s.service.Create(ctx, pack)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("name", pack.ID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to create pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Update implements the Service interface for logging.
func (s *loggingService) Update(ctx context.Context, pack *model.Pack) error {
	start := time.Now()
	err := s.service.Update(ctx, pack)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("name", pack.ID).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to update pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Delete implements the Service interface for logging.
func (s *loggingService) Delete(ctx context.Context, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to delete pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *loggingService) Exists(ctx context.Context, name string) (bool, error) {
	return s.service.Exists(ctx, name)
}

// Column implements the Service interface for logging.
func (s *loggingService) Column(ctx context.Context, name, col string, val any) error {
	return s.service.Column(ctx, name, col, val)
}

func (s *loggingService) requestID(ctx context.Context) string {
	id, ok := hlog.IDFromCtx(ctx)

	if ok {
		return id.String()
	}

	return ""
}
