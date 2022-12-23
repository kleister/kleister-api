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

// LoggingRepository implements UsersRepository interface.
type LoggingRepository struct {
	upstream  UsersRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the UsersRepository and provides logging for its methods.
func NewLoggingRepository(repository UsersRepository, requestID LoggingRequestID) UsersRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "users").Logger(),
	}
}

// List implements the UsersRepository interface.
func (r *LoggingRepository) List(ctx context.Context) ([]*model.User, error) {
	start := time.Now()
	records, err := r.upstream.List(ctx)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch users")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Create implements the UsersRepository interface.
func (r *LoggingRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	start := time.Now()
	record, err := r.upstream.Create(ctx, user)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("name", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create user")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the UsersRepository interface.
func (r *LoggingRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	start := time.Now()
	record, err := r.upstream.Update(ctx, user)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("name", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update user")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Show implements the UsersRepository interface.
func (r *LoggingRepository) Show(ctx context.Context, name string) (*model.User, error) {
	start := time.Now()
	record, err := r.upstream.Show(ctx, name)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch user")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the UsersRepository interface.
func (r *LoggingRepository) Delete(ctx context.Context, name string) error {
	start := time.Now()
	err := r.upstream.Delete(ctx, name)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the UsersRepository interface.
func (r *LoggingRepository) Exists(ctx context.Context, name string) (bool, error) {
	start := time.Now()
	exists, err := r.upstream.Exists(ctx, name)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "exists").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to check user")
	} else {
		logger.Debug().
			Msg("")
	}

	return exists, err
}

func (r *LoggingRepository) extractIdentifier(record *model.User) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	if record.Slug != "" {
		return record.Slug
	}

	if record.Username != "" {
		return record.Username
	}

	return ""
}
