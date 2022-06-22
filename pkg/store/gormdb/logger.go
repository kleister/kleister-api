package gormdb

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Logger defines a Gorm compatible logger.
type Logger struct {
	SlowThreshold time.Duration
}

// LogMode implements the logger.Interface.
func (l *Logger) LogMode(_ logger.LogLevel) logger.Interface {
	return l
}

// Info implements the logger.Interface.
func (l *Logger) Info(_ context.Context, msg string, data ...interface{}) {
	log.Info().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Warn implements the logger.Interface.
func (l *Logger) Warn(_ context.Context, msg string, data ...interface{}) {
	log.Warn().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Error implements the logger.Interface.
func (l *Logger) Error(_ context.Context, msg string, data ...interface{}) {
	log.Error().
		Str("source", utils.FileWithLineNum()).
		Msgf(msg, data...)
}

// Trace implements the logger.Interface.
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
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

// NewLogger prepares a Gorm compatible logger.
func NewLogger() *Logger {
	return &Logger{
		SlowThreshold: 200 * time.Millisecond,
	}
}
