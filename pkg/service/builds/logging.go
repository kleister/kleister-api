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

func (s *loggingService) List(ctx context.Context, pack *model.Pack) ([]*model.Build, error) {
	start := time.Now()
	records, err := s.service.List(ctx, pack)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Str("pack", pack.ID).
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) Show(ctx context.Context, pack *model.Pack, name string) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, pack, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Str("pack", pack.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to find build by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Create(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, pack, build)

	name := ""

	if record != nil {
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Str("pack", pack.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Update(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, pack, build)

	id := ""
	name := ""

	if record != nil {
		id = record.ID
		name = record.Name
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Str("pack", pack.ID).
		Dur("duration", time.Since(start)).
		Str("id", id).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to update build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Delete(ctx context.Context, pack *model.Pack, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, pack, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Str("pack", pack.ID).
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to delete build")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
