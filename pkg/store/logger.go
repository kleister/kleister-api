package store

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

// GormLogger defines a Gorm compatible logger.
type GormLogger struct {
	SlowThreshold time.Duration
	Logger        zerolog.Logger
	Level         logger.LogLevel
}

// LogMode implements the logger.Interface.
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.Level = level
	return l
}

// Info implements the logger.Interface.
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Info {
		l.Logger.Log().Ctx(ctx).
			// Str("source", utils.FileWithLineNum()).
			Msgf(msg, data...)
	}
}

// Warn implements the logger.Interface.
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Warn {
		l.Logger.Log().Ctx(ctx).
			// Str("source", utils.FileWithLineNum()).
			Msgf(msg, data...)
	}
}

// Error implements the logger.Interface.
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Error {
		l.Logger.Log().Ctx(ctx).
			// Str("source", utils.FileWithLineNum()).
			Msgf(msg, data...)
	}
}

// Trace implements the logger.Interface.
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)

	switch {
	case err != nil && l.Level >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()

		l.Logger.Error().Ctx(ctx).
			Err(err).
			Dur("duration", elapsed).
			Int64("rows", rows).
			// Str("source", utils.FileWithLineNum()).
			Str("sql", sql).
			Msg("")
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.Level >= logger.Warn:
		sql, rows := fc()

		l.Logger.Log().Ctx(ctx).
			Dur("duration", elapsed).
			Int64("rows", rows).
			// Str("source", utils.FileWithLineNum()).
			Str("sql", sql).
			Dur("treshold", l.SlowThreshold).
			Msg("slow query")
	case l.Level == logger.Info:
		sql, rows := fc()

		l.Logger.Log().Ctx(ctx).
			Dur("duration", elapsed).
			Int64("rows", rows).
			// Str("source", utils.FileWithLineNum()).
			Msg(sql)
	}
}

// NewGormLogger prepares a Gorm compatible logger.
func NewGormLogger() *GormLogger {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond,
		Logger:        log.With().Str("service", "store").Logger(),
		Level:         logger.Warn,
	}
}
