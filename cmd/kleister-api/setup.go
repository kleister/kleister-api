package main

import (
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/store/boltdb"
	"github.com/kleister/kleister-api/pkg/store/mysql"
	"github.com/kleister/kleister-api/pkg/store/postgres"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/kleister/kleister-api/pkg/upload/file"
	"github.com/kleister/kleister-api/pkg/upload/s3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	jaeger "github.com/uber/jaeger-client-go"
	tracecfg "github.com/uber/jaeger-client-go/config"
)

func setupLogger(cfg *config.Config) {
	switch strings.ToLower(cfg.Logs.Level) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if cfg.Logs.Pretty {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !cfg.Logs.Color,
			},
		)
	}
}

func setupTracing(cfg *config.Config) (io.Closer, error) {
	switch {
	case cfg.Tracing.Enabled:
		closer, err := tracecfg.Configuration{
			Sampler: &tracecfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &tracecfg.ReporterConfig{
				LocalAgentHostPort: cfg.Tracing.Endpoint,
			},
		}.InitGlobalTracer("kleister-api")

		if err != nil {
			return nil, err
		}

		log.Info().
			Str("addr", cfg.Tracing.Endpoint).
			Msg("application tracing is enabled")

		return closer, nil
	default:
		log.Info().
			Msg("application tracing is disabled")

		return nil, nil
	}
}

func setupUploads(cfg *config.Config) (upload.Upload, error) {
	parsed, err := url.Parse(cfg.Upload.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "file":
		return file.New(parsed)
	case "s3":
		return s3.New(parsed)
	case "minio":
		return s3.New(parsed)
	}

	return nil, upload.ErrUnknownDriver
}

func setupStorage(cfg *config.Config) (store.Store, error) {
	parsed, err := url.Parse(cfg.Database.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "boltdb":
		return boltdb.New(parsed)
	case "postgres":
		return postgres.New(parsed)
	case "mysql":
		return mysql.New(parsed)
	case "mariadb":
		return mysql.New(parsed)
	}

	return nil, store.ErrUnknownDriver
}
