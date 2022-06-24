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

func (s *loggingService) List(ctx context.Context, mod *model.Mod) ([]*model.Version, error) {
	start := time.Now()
	records, err := s.service.List(ctx, mod)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Str("mod", mod.ID).
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) Show(ctx context.Context, mod *model.Mod, name string) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, mod, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Str("mod", mod.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to find version by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Create(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, mod, version)

	name := ""

	if record != nil {
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Str("mod", mod.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Update(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, mod, version)

	id := ""
	name := ""

	if record != nil {
		id = record.ID
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Str("mod", mod.ID).
		Dur("duration", time.Since(start)).
		Str("id", id).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to update version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Delete(ctx context.Context, mod *model.Mod, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, mod, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Str("mod", mod.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to delete version")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
