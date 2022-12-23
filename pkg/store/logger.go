package store

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GormLogger defines a Gorm compatible logger.
type GormLogger struct {
	SlowThreshold time.Duration
}

// LogMode implements the logger.Interface.
func (l *GormLogger) LogMode(_ logger.LogLevel) logger.Interface {
	return l
}

// Info implements the logger.Interface.
func (l *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	log.Info().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Warn implements the logger.Interface.
func (l *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	log.Warn().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Error implements the logger.Interface.
func (l *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	log.Error().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Trace implements the logger.Interface.
func (l *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	if err != nil && !errors.Is(err, logger.ErrRecordNotFound) {
		sql, rows := fc()

		log.Error().
			Err(err).
			Dur("duration", elapsed).
			Int64("rows", rows).
			Str("source", utils.FileWithLineNum()).
			Str("sql", sql).
			Msg("")

		return
	}

	if elapsed > l.SlowThreshold && l.SlowThreshold != 0 {
		sql, rows := fc()

		log.Trace().
			Dur("duration", elapsed).
			Int64("rows", rows).
			Str("source", utils.FileWithLineNum()).
			Str("sql", sql).
			Msg("slow query")

		return
	}

	sql, rows := fc()

	log.Trace().
		Dur("duration", elapsed).
		Int64("rows", rows).
		Str("source", utils.FileWithLineNum()).
		Msg(sql)
}

// NewGormLogger prepares a Gorm compatible logger.
func NewGormLogger() *GormLogger {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond,
	}
}
