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

// LoggingRepository implements PacksRepository interface.
type LoggingRepository struct {
	upstream  PacksRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the PacksRepository and provides logging for its methods.
func NewLoggingRepository(repository PacksRepository, requestID LoggingRequestID) PacksRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "packs").Logger(),
	}
}

// List implements the PacksRepository interface.
func (r *LoggingRepository) List(ctx context.Context, search string) ([]*model.Pack, error) {
	start := time.Now()
	records, err := r.upstream.List(ctx, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch packs")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Create implements the PacksRepository interface.
func (r *LoggingRepository) Create(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	start := time.Now()
	record, err := r.upstream.Create(ctx, pack)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the PacksRepository interface.
func (r *LoggingRepository) Update(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	start := time.Now()
	record, err := r.upstream.Update(ctx, pack)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Show implements the PacksRepository interface.
func (r *LoggingRepository) Show(ctx context.Context, id string) (*model.Pack, error) {
	start := time.Now()
	record, err := r.upstream.Show(ctx, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the PacksRepository interface.
func (r *LoggingRepository) Delete(ctx context.Context, id string) error {
	start := time.Now()
	err := r.upstream.Delete(ctx, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the PacksRepository interface.
func (r *LoggingRepository) Exists(ctx context.Context, id string) (bool, string, error) {
	start := time.Now()
	exists, realID, err := r.upstream.Exists(ctx, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "exists").
		Dur("duration", time.Since(start)).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to check pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return exists, realID, err
}

// ListUsers implements the PacksRepository interface.
func (r *LoggingRepository) ListUsers(ctx context.Context, packID, search string) ([]*model.UserPack, error) {
	start := time.Now()
	records, err := r.upstream.ListUsers(ctx, packID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listUsers").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachUser implements the PacksRepository interface.
func (r *LoggingRepository) AttachUser(ctx context.Context, packID, userID string) error {
	start := time.Now()
	err := r.upstream.AttachUser(ctx, packID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachUser").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to attach user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// DropUser implements the PacksRepository interface.
func (r *LoggingRepository) DropUser(ctx context.Context, packID, userID string) error {
	start := time.Now()
	err := r.upstream.DropUser(ctx, packID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropUser").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("user", userID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// ListTeams implements the PacksRepository interface.
func (r *LoggingRepository) ListTeams(ctx context.Context, packID, search string) ([]*model.TeamPack, error) {
	start := time.Now()
	records, err := r.upstream.ListTeams(ctx, packID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listTeams").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete pack")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachTeam implements the PacksRepository interface.
func (r *LoggingRepository) AttachTeam(ctx context.Context, packID, teamID string) error {
	start := time.Now()
	err := r.upstream.AttachTeam(ctx, packID, teamID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachTeam").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("team", teamID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to attach team")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// DropTeam implements the PacksRepository interface.
func (r *LoggingRepository) DropTeam(ctx context.Context, packID, teamID string) error {
	start := time.Now()
	err := r.upstream.DropTeam(ctx, packID, teamID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropTeam").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("team", teamID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop team")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (r *LoggingRepository) extractIdentifier(record *model.Pack) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	if record.Slug != "" {
		return record.Slug
	}

	if record.Name != "" {
		return record.Name
	}

	return ""
}
