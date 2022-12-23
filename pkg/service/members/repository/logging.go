package repository

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggingRequestID returns the request ID as string for logging
type LoggingRequestID func(context.Context) string

// LoggingRepository implements MembersRepository interface.
type LoggingRepository struct {
	upstream  MembersRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the MembersRepository and provides logging for its methods.
func NewLoggingRepository(repository MembersRepository, requestID LoggingRequestID) MembersRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "members").Logger(),
	}
}

// List implements the MembersRepository interface.
func (r *LoggingRepository) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	start := time.Now()
	records, err := r.upstream.List(ctx, teamID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch members")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Append implements the MembersRepository interface.
func (r *LoggingRepository) Append(ctx context.Context, teamID, userID string) error {
	start := time.Now()
	err := r.upstream.Append(ctx, teamID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "append").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to append member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the MembersRepository interface.
func (r *LoggingRepository) Drop(ctx context.Context, teamID, userID string) error {
	start := time.Now()
	err := r.upstream.Drop(ctx, teamID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("team", teamID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
