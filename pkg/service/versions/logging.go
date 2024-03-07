package versions

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
		logger:    log.With().Str("service", "versions").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, modID string) ([]*model.Version, error) {
	start := time.Now()
	records, err := s.service.List(ctx, modID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to find mod versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Show implements the Service interface for logging.
func (s *loggingService) Show(ctx context.Context, modID, name string) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, modID, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to find version by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Create implements the Service interface for logging.
func (s *loggingService) Create(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, modID, version)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to create version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the Service interface for logging.
func (s *loggingService) Update(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, modID, version)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to update version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the Service interface for logging.
func (s *loggingService) Delete(ctx context.Context, modID, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, modID, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to delete version")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *loggingService) Exists(ctx context.Context, modID, name string) (bool, error) {
	return s.service.Exists(ctx, modID, name)
}

func (s *loggingService) extractIdentifier(record *model.Version) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	return ""
}
