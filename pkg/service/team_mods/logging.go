package teammods

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
		logger:  log.With().Str("service", "teammods").Logger(),
	}
}

// External implements the Service interface for logging.
func (s *loggingService) WithPrincipal(principal *model.User) Service {
	s.service.WithPrincipal(principal)
	return s
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, params model.TeamModParams) ([]*model.TeamMod, int64, error) {
	start := time.Now()
	records, counter, err := s.service.List(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("mod", params.ModID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch team mods")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, counter, err
}

// Attach implements the Service interface for logging.
func (s *loggingService) Attach(ctx context.Context, params model.TeamModParams) error {
	start := time.Now()
	err := s.service.Attach(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attach").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("mod", params.ModID).
		Str("perm", params.Perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach team mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Permit implements the Service interface for logging.
func (s *loggingService) Permit(ctx context.Context, params model.TeamModParams) error {
	start := time.Now()
	err := s.service.Permit(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permit").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("mod", params.ModID).
		Str("perm", params.Perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to permit team mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the Service interface for logging.
func (s *loggingService) Drop(ctx context.Context, params model.TeamModParams) error {
	start := time.Now()
	err := s.service.Drop(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("mod", params.ModID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop team mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) requestID(ctx context.Context) string {
	id, ok := hlog.IDFromCtx(ctx)

	if ok {
		return id.String()
	}

	return ""
}
