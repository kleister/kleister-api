package buildVersions

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
		logger:    log.With().Str("service", "buildVersions").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, packID, buildID, modID, versionID string) ([]*model.BuildVersion, error) {
	start := time.Now()
	records, err := s.service.List(ctx, packID, buildID, modID, versionID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Str("mod", modID).
		Str("version", versionID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch build versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Attach implements the Service interface for logging.
func (s *loggingService) Attach(ctx context.Context, packID, buildID, modID, versionID string) error {
	start := time.Now()
	err := s.service.Attach(ctx, packID, buildID, modID, versionID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attach").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Str("mod", modID).
		Str("version", versionID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach build version")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the Service interface for logging.
func (s *loggingService) Drop(ctx context.Context, packID, buildID, modID, versionID string) error {
	start := time.Now()
	err := s.service.Drop(ctx, packID, buildID, modID, versionID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Str("mod", modID).
		Str("version", versionID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop build version")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
