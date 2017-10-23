package dblog

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-xorm/core"
)

// New initializes a new XORM logger instance.
func New(logger log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Logger provides the implementation for core.ILogger
type Logger struct {
	logger log.Logger
	level  core.LogLevel
	show   bool
}

// Error implements core.ILogger.
func (s *Logger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		level.Error(s.logger).Log(
			"msg", fmt.Sprint(v...),
		)
	}
}

// Errorf implements core.ILogger.
func (s *Logger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		level.Error(s.logger).Log(
			"msg", fmt.Sprintf(format, v...),
		)
	}
}

// Debug implements core.ILogger.
func (s *Logger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		level.Debug(s.logger).Log(
			"msg", fmt.Sprint(v...),
		)
	}
}

// Debugf implements core.ILogger.
func (s *Logger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		level.Debug(s.logger).Log(
			"msg", fmt.Sprintf(format, v...),
		)
	}
}

// Info implements core.ILogger.
func (s *Logger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		level.Info(s.logger).Log(
			"msg", fmt.Sprint(v...),
		)
	}
}

// Infof implements core.ILogger.
func (s *Logger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		level.Info(s.logger).Log(
			"msg", fmt.Sprintf(format, v...),
		)
	}
}

// Warn implements core.ILogger.
func (s *Logger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		level.Warn(s.logger).Log(
			"msg", fmt.Sprint(v...),
		)
	}
}

// Warnf implements core.ILogger.
func (s *Logger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		level.Warn(s.logger).Log(
			"msg", fmt.Sprintf(format, v...),
		)
	}
}

// Level implements core.ILogger.
func (s *Logger) Level() core.LogLevel {
	return s.level
}

// SetLevel implements core.ILogger.
func (s *Logger) SetLevel(l core.LogLevel) {
	s.level = l
}

// ShowSQL implements core.ILogger.
func (s *Logger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.show = true
	} else {
		s.show = show[0]
	}
}

// IsShowSQL implements core.ILogger.
func (s *Logger) IsShowSQL() bool {
	return s.show
}
