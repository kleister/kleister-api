package forge

import (
	"context"
	"time"

	"github.com/kleister/go-forge/version"
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
		logger:  log.With().Str("service", "forge").Logger(),
	}
}

// External implements the Service interface for logging.
func (s *loggingService) WithPrincipal(principal *model.User) Service {
	s.service.WithPrincipal(principal)
	return s
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, params model.ListParams) ([]*model.Forge, int64, error) {
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
			Msg("Failed to find all forge")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, counter, err
}

// Show implements the Service interface for logging.
func (s *loggingService) Show(ctx context.Context, name string) (*model.Forge, error) {
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
			Msg("Failed to find forge by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Sync implements the Service interface for logging.
func (s *loggingService) Sync(ctx context.Context, versions version.Versions) error {
	start := time.Now()
	err := s.service.Sync(ctx, versions)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "sync").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to sync forge versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// ListBuilds implements the Service interface for logging.
func (s *loggingService) ListBuilds(ctx context.Context, params model.ForgeBuildParams) ([]*model.Build, int64, error) {
	start := time.Now()
	records, counter, err := s.service.ListBuilds(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "listBuilds").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to find all forge builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, counter, err
}

// AttachBuild implements the Service interface for logging.
func (s *loggingService) AttachBuild(ctx context.Context, params model.ForgeBuildParams) error {
	start := time.Now()
	err := s.service.AttachBuild(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attachBuild").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach forge to build")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// DropBuild implements the Service interface for logging.
func (s *loggingService) DropBuild(ctx context.Context, params model.ForgeBuildParams) error {
	start := time.Now()
	err := s.service.DropBuild(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "dropBuild").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop forge from build")
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
