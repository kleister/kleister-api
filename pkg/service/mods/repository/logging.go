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

// LoggingRepository implements ModsRepository interface.
type LoggingRepository struct {
	upstream  ModsRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the ModsRepository and provides logging for its methods.
func NewLoggingRepository(repository ModsRepository, requestID LoggingRequestID) ModsRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "mods").Logger(),
	}
}

// List implements the ModsRepository interface.
func (r *LoggingRepository) List(ctx context.Context, search string) ([]*model.Mod, error) {
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
			Msg("failed to fetch mods")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Create implements the ModsRepository interface.
func (r *LoggingRepository) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	start := time.Now()
	record, err := r.upstream.Create(ctx, mod)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Update implements the ModsRepository interface.
func (r *LoggingRepository) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	start := time.Now()
	record, err := r.upstream.Update(ctx, mod)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Show implements the ModsRepository interface.
func (r *LoggingRepository) Show(ctx context.Context, id string) (*model.Mod, error) {
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
			Msg("failed to fetch mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the ModsRepository interface.
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
			Msg("failed to delete mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the ModsRepository interface.
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
			Msg("failed to check mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return exists, realID, err
}

// ListUsers implements the ModsRepository interface.
func (r *LoggingRepository) ListUsers(ctx context.Context, modID, search string) ([]*model.UserMod, error) {
	start := time.Now()
	records, err := r.upstream.ListUsers(ctx, modID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listUsers").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachUser implements the ModsRepository interface.
func (r *LoggingRepository) AttachUser(ctx context.Context, modID, userID string) error {
	start := time.Now()
	err := r.upstream.AttachUser(ctx, modID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachUser").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
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

// DropUser implements the ModsRepository interface.
func (r *LoggingRepository) DropUser(ctx context.Context, modID, userID string) error {
	start := time.Now()
	err := r.upstream.DropUser(ctx, modID, userID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropUser").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
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

// ListTeams implements the ModsRepository interface.
func (r *LoggingRepository) ListTeams(ctx context.Context, modID, search string) ([]*model.TeamMod, error) {
	start := time.Now()
	records, err := r.upstream.ListTeams(ctx, modID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listTeams").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete mod")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachTeam implements the ModsRepository interface.
func (r *LoggingRepository) AttachTeam(ctx context.Context, modID, teamID string) error {
	start := time.Now()
	err := r.upstream.AttachTeam(ctx, modID, teamID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachTeam").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
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

// DropTeam implements the ModsRepository interface.
func (r *LoggingRepository) DropTeam(ctx context.Context, modID, teamID string) error {
	start := time.Now()
	err := r.upstream.DropTeam(ctx, modID, teamID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropTeam").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
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

func (r *LoggingRepository) extractIdentifier(record *model.Mod) string {
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
