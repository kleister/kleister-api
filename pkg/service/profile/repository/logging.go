package repository

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggingRequestID returns the request ID as string for logging
type LoggingRequestID func(context.Context) string

// LoggingRepository implements ProfileRepository interface.
type LoggingRepository struct {
	upstream  ProfileRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the ProfileRepository and provides logging for its methods.
func NewLoggingRepository(repository ProfileRepository, requestID LoggingRequestID) ProfileRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "teams").Logger(),
	}
}
