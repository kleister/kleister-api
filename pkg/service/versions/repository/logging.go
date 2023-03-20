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

// LoggingRepository implements VersionsRepository interface.
type LoggingRepository struct {
	upstream  VersionsRepository
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingRepository wraps the VersionsRepository and provides logging for its methods.
func NewLoggingRepository(repository VersionsRepository, requestID LoggingRequestID) VersionsRepository {
	return &LoggingRepository{
		upstream:  repository,
		requestID: requestID,
		logger:    log.With().Str("service", "teams").Logger(),
	}
}

// List implements the VersionsRepository interface.
func (r *LoggingRepository) List(ctx context.Context, modID string) ([]*model.Version, error) {
	start := time.Now()
	records, err := r.upstream.List(ctx, modID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch versions")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// Create implements the VersionsRepository interface.
func (r *LoggingRepository) Create(ctx context.Context, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := r.upstream.Create(ctx, version)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("pack", r.extractMod(record)).
		Str("id", r.extractIdentifier(record)).
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

// Update implements the VersionsRepository interface.
func (r *LoggingRepository) Update(ctx context.Context, version *model.Version) (*model.Version, error) {
	start := time.Now()
	record, err := r.upstream.Update(ctx, version)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("mod", r.extractMod(record)).
		Str("id", r.extractIdentifier(record)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Show implements the VersionsRepository interface.
func (r *LoggingRepository) Show(ctx context.Context, modID string, id string) (*model.Version, error) {
	start := time.Now()
	record, err := r.upstream.Show(ctx, modID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch version")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

// Delete implements the VersionsRepository interface.
func (r *LoggingRepository) Delete(ctx context.Context, modID string, id string) error {
	start := time.Now()
	err := r.upstream.Delete(ctx, modID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to delete version")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Exists implements the VersionsRepository interface.
func (r *LoggingRepository) Exists(ctx context.Context, modID, id string) (bool, string, error) {
	start := time.Now()
	exists, realID, err := r.upstream.Exists(ctx, modID, id)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "exists").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("id", id).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to check version")
	} else {
		logger.Debug().
			Msg("")
	}

	return exists, realID, err
}

// ListBuilds implements the VersionsRepository interface.
func (r *LoggingRepository) ListBuilds(ctx context.Context, modID, versionID, search string) ([]*model.Build, error) {
	start := time.Now()
	records, err := r.upstream.ListBuilds(ctx, modID, versionID, search)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "listBuilds").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("version", versionID).
		Str("search", search).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to fetch version builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

// AttachBuild implements the VersionsRepository interface.
func (r *LoggingRepository) AttachBuild(ctx context.Context, modID, versionID, packID, buildID string) error {
	start := time.Now()
	err := r.upstream.AttachBuild(ctx, modID, versionID, packID, buildID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "attachBuild").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("version", versionID).
		Str("pack", packID).
		Str("build", buildID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to attach version builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// DropBuild implements the VersionsRepository interface.
func (r *LoggingRepository) DropBuild(ctx context.Context, modID, versionID, packID, buildID string) error {
	start := time.Now()
	err := r.upstream.DropBuild(ctx, modID, versionID, packID, buildID)

	logger := r.logger.With().
		Str("request", r.requestID(ctx)).
		Str("method", "dropBuild").
		Dur("duration", time.Since(start)).
		Str("mod", modID).
		Str("version", versionID).
		Str("pack", packID).
		Str("build", buildID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop version builds")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (r *LoggingRepository) extractIdentifier(record *model.Version) string {
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

func (r *LoggingRepository) extractMod(record *model.Version) string {
	if record == nil {
		return ""
	}

	if record.ModID != "" {
		return record.ModID
	}

	return ""
}
