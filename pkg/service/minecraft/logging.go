package minecraft

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
		logger:    log.With().Str("service", "minecraft").Logger(),
	}
}

func (s *loggingService) List(ctx context.Context) ([]*model.Minecraft, error) {
	start := time.Now()
	records, err := s.service.List(ctx)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all minecrafts")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) Update(ctx context.Context) error {
	start := time.Now()
	err := s.service.Update(ctx)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update minecraft versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
