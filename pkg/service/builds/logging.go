package builds

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
		logger:    log.With().Str("service", "builds").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, packID string) ([]*model.Build, error) {
	start := time.Now()
	records, err := s.service.List(ctx, packID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to find pack builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Show implements the Service interface for logging.
func (s *loggingService) Show(ctx context.Context, packID, name string) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, packID, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to find build by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Create implements the Service interface for logging.
func (s *loggingService) Create(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, packID, build)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to create build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the Service interface for logging.
func (s *loggingService) Update(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, packID, build)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("name", s.extractIdentifier(record)).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to update build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the Service interface for logging.
func (s *loggingService) Delete(ctx context.Context, packID, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, packID, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Error().
			Err(err).
			Msg("Failed to delete build")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *loggingService) Exists(ctx context.Context, packID, name string) (bool, error) {
	return s.service.Exists(ctx, packID, name)
}

// Column implements the Service interface.
func (s *loggingService) Column(ctx context.Context, packID, id, col string, val any) error {
	return s.service.Column(ctx, packID, id, col, val)
}

func (s *loggingService) extractIdentifier(record *model.Build) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	return ""
}
