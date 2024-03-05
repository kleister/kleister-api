package members

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
		logger:    log.With().Str("service", "members").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	start := time.Now()
	records, err := s.service.List(ctx, teamID, userID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch members")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Attach implements the Service interface for logging.
func (s *loggingService) Attach(ctx context.Context, teamID, userID, perm string) error {
	start := time.Now()
	err := s.service.Attach(ctx, teamID, userID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attach").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Permit implements the Service interface for logging.
func (s *loggingService) Permit(ctx context.Context, teamID, userID, perm string) error {
	start := time.Now()
	err := s.service.Permit(ctx, teamID, userID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permit").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to permit member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the Service interface for logging.
func (s *loggingService) Drop(ctx context.Context, teamID, userID string) error {
	start := time.Now()
	err := s.service.Drop(ctx, teamID, userID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
