package teamPacks

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
		logger:    log.With().Str("service", "teamPacks").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, teamID, packID string) ([]*model.TeamPack, error) {
	start := time.Now()
	records, err := s.service.List(ctx, teamID, packID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch team packs")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Attach implements the Service interface for logging.
func (s *loggingService) Attach(ctx context.Context, teamID, packID, perm string) error {
	start := time.Now()
	err := s.service.Attach(ctx, teamID, packID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attach").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach team pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Permit implements the Service interface for logging.
func (s *loggingService) Permit(ctx context.Context, teamID, packID, perm string) error {
	start := time.Now()
	err := s.service.Permit(ctx, teamID, packID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permit").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to permit team pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the Service interface for logging.
func (s *loggingService) Drop(ctx context.Context, teamID, packID string) error {
	start := time.Now()
	err := s.service.Drop(ctx, teamID, packID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("pack", packID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop team pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
