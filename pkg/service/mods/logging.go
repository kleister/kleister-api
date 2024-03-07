package mods

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggingRequestID returns the request ID as string for logging
type LoggingRequestID func(context.Context) string

type loggingService struct {
	service   Service
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(s Service, requestID LoggingRequestID) Service {
	return &loggingService{
		service:   s,
		requestID: requestID,
		logger:    log.With().Str("service", "mods").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context) ([]*model.Mod, error) {
	start := time.Now()
	records, err := s.service.List(ctx)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to find all mods")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Show implements the Service interface for logging.
func (s *loggingService) Show(ctx context.Context, name string) (*model.Mod, error) {
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
			Msg("Failed to find mod by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Create implements the Service interface for logging.
func (s *loggingService) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, mod)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to create mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the Service interface for logging.
func (s *loggingService) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, mod)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to update mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
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
			Msg("Failed to delete mod")
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

func (s *loggingService) extractIdentifier(record *model.Mod) string {
	if record == nil {
		return ""
	}

	if record.Slug != "" {
		return record.Slug
	}

	if record.ID != "" {
		return record.ID
	}

	return ""
}
