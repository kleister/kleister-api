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

// LoggingRepository implements MinecraftRepository interface.
type LoggingRepository struct {
	upstream  MinecraftRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the MinecraftRepository and provides logging for its methods.
func NewLoggingRepository(repository MinecraftRepository, requestID LoggingRequestID) MinecraftRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "teams").Logger(),
	}
}

// Search implements the MinecraftRepository interface.
func (r *LoggingRepository) Search(ctx context.Context, search string) ([]*model.Minecraft, error) {
	start := time.Now()
	records, err := r.upstream.Search(ctx, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "search").
		Dur("duration", time.Since(start)).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to search minecraft versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Update implements the MinecraftRepository interface.
func (r *LoggingRepository) Update(ctx context.Context) error {
	start := time.Now()
	err := r.upstream.Update(ctx)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
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

// ListBuilds implements the MinecraftRepository interface.
func (r *LoggingRepository) ListBuilds(ctx context.Context, id, search string) ([]*model.Build, error) {
	start := time.Now()
	records, err := r.upstream.ListBuilds(ctx, id, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listBuilds").
		Dur("duration", time.Since(start)).
		Str("minecraft", id).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to list builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}
