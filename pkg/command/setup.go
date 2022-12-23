package command

import (
	"net/url"
	"os"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func setup(cfg *config.Config) error {
	if err := setupLogger(cfg); err != nil {
		return err
	}

	return setupConfig(cfg)
}

func setupLogger(cfg *config.Config) error {
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
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
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

	return nil
}

func setupConfig(cfg *config.Config) error {
	if cfg.File != "" {
		viper.SetConfigFile(cfg.File)
	} else {
		viper.SetConfigName("api")

		viper.AddConfigPath("/etc/kleister")
		viper.AddConfigPath("$HOME/.kleister")
		viper.AddConfigPath("./config")
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Info().
				Msg("Continue without config")
		case viper.UnsupportedConfigError:
			log.Fatal().
				Err(err).
				Msg("Unsupported config type")

			return err
		default:
			log.Fatal().
				Err(err).
				Msg("Failed to read config")

			return err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to parse config")

		return err
	}

	return nil
}

func setupUploads(cfg *config.Config) (upload.Upload, error) {
	parsed, err := url.Parse(cfg.Upload.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "file":
		return upload.NewFileUpload(cfg.Upload)
	case "s3":
		return upload.NewS3Upload(cfg.Upload)
	case "minio":
		return upload.NewS3Upload(cfg.Upload)
	}

	return nil, upload.ErrUnknownDriver
}

func setupStorage(cfg *config.Config) (store.Store, error) {
	parsed, err := url.Parse(cfg.Database.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "sqlite", "sqlite3":
		return store.NewGormStore(cfg.Database)
	case "mysql", "mariadb":
		return store.NewGormStore(cfg.Database)
	case "postgres", "postgresql":
		return store.NewGormStore(cfg.Database)
	}

	return nil, store.ErrUnknownDriver
}
