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

// LoggingRepository implements BuildsRepository interface.
type LoggingRepository struct {
	upstream  BuildsRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the BuildsRepository and provides logging for its methods.
func NewLoggingRepository(repository BuildsRepository, requestID LoggingRequestID) BuildsRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "builds").Logger(),
	}
}

// List implements the BuildsRepository interface.
func (r *LoggingRepository) List(ctx context.Context, packID string) ([]*model.Build, error) {
	start := time.Now()
	records, err := r.upstream.List(ctx, packID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Create implements the BuildsRepository interface.
func (r *LoggingRepository) Create(ctx context.Context, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := r.upstream.Create(ctx, build)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("pack", r.extractPack(record)).
		Str("id", r.extractIdentifier(record)).
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

// Update implements the BuildsRepository interface.
func (r *LoggingRepository) Update(ctx context.Context, build *model.Build) (*model.Build, error) {
	start := time.Now()
	record, err := r.upstream.Update(ctx, build)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("pack", r.extractPack(record)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Show implements the BuildsRepository interface.
func (r *LoggingRepository) Show(ctx context.Context, packID string, id string) (*model.Build, error) {
	start := time.Now()
	record, err := r.upstream.Show(ctx, packID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch build")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the BuildsRepository interface.
func (r *LoggingRepository) Delete(ctx context.Context, packID string, id string) error {
	start := time.Now()
	err := r.upstream.Delete(ctx, packID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete build")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the BuildsRepository interface.
func (r *LoggingRepository) Exists(ctx context.Context, packID, id string) (bool, string, error) {
	start := time.Now()
	exists, realID, err := r.upstream.Exists(ctx, packID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "exists").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to check build")
	} else {
		logger.Debug().
			Msg("")
	}

	return exists, realID, err
}

// ListVersions implements the BuildsRepository interface.
func (r *LoggingRepository) ListVersions(ctx context.Context, packID, buildID, search string) ([]*model.Version, error) {
	start := time.Now()
	records, err := r.upstream.ListVersions(ctx, packID, buildID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch build versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachVersion implements the BuildsRepository interface.
func (r *LoggingRepository) AttachVersion(ctx context.Context, packID, buildID, modID, versionID string) error {
	start := time.Now()
	err := r.upstream.AttachVersion(ctx, packID, buildID, modID, versionID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachVersion").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Str("mod", modID).
		Str("version", versionID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to attach build versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// DropVersion implements the BuildsRepository interface.
func (r *LoggingRepository) DropVersion(ctx context.Context, packID, buildID, modID, versionID string) error {
	start := time.Now()
	err := r.upstream.DropVersion(ctx, packID, buildID, modID, versionID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropVersion").
		Dur("duration", time.Since(start)).
		Str("pack", packID).
		Str("build", buildID).
		Str("mod", modID).
		Str("version", versionID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop build versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (r *LoggingRepository) extractIdentifier(record *model.Build) string {
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

func (r *LoggingRepository) extractPack(record *model.Build) string {
	if record == nil {
		return ""
	}

	if record.PackID != "" {
		return record.PackID
	}

	return ""
}
