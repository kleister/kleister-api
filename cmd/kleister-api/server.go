package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/router"
	"github.com/oklog/oklog/pkg/group"
	"github.com/rs/zerolog/log"
	"gopkg.in/urfave/cli.v2"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "start integrated server",
		Flags:  serverFlags(cfg),
		Before: serverBefore(cfg),
		Action: serverAction(cfg),
	}
}

func serverFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "metrics-addr",
			Value:       "0.0.0.0:8090",
			Usage:       "address to bind the metrics",
			EnvVars:     []string{"KLEISTER_API_METRICS_ADDR"},
			Destination: &cfg.Metrics.Addr,
		},
		&cli.StringFlag{
			Name:        "metrics-token",
			Value:       "",
			Usage:       "token to make metrics secure",
			EnvVars:     []string{"KLEISTER_API_METRICS_TOKEN"},
			Destination: &cfg.Metrics.Token,
		},
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       "0.0.0.0:8080",
			Usage:       "address to bind the server",
			EnvVars:     []string{"KLEISTER_API_SERVER_ADDR"},
			Destination: &cfg.Server.Addr,
		},
		&cli.BoolFlag{
			Name:        "server-pprof",
			Value:       false,
			Usage:       "enable pprof debugging",
			EnvVars:     []string{"KLEISTER_API_SERVER_PPROF"},
			Destination: &cfg.Server.Pprof,
		},
		&cli.BoolFlag{
			Name:        "server-docs",
			Value:       true,
			Usage:       "enable swagger documentation",
			EnvVars:     []string{"KLEISTER_API_SERVER_DOCS"},
			Destination: &cfg.Server.Docs,
		},
		&cli.StringFlag{
			Name:        "server-host",
			Value:       "http://localhost:8080",
			Usage:       "external access to server",
			EnvVars:     []string{"KLEISTER_API_SERVER_HOST"},
			Destination: &cfg.Server.Host,
		},
		&cli.StringFlag{
			Name:        "server-root",
			Value:       "/",
			Usage:       "path to access the server",
			EnvVars:     []string{"KLEISTER_API_SERVER_ROOT"},
			Destination: &cfg.Server.Root,
		},
		&cli.StringFlag{
			Name:        "db-dsn",
			Value:       "boltdb://kleister.db",
			Usage:       "database dsn",
			EnvVars:     []string{"KLEISTER_API_DB_DSN"},
			Destination: &cfg.Database.DSN,
		},
		&cli.StringFlag{
			Name:        "upload-dsn",
			Value:       "file://storage/",
			Usage:       "uploads dsn",
			EnvVars:     []string{"KLEISTER_API_UPLOAD_DSN"},
			Destination: &cfg.Upload.DSN,
		},
		&cli.BoolFlag{
			Name:        "admin-create",
			Value:       true,
			Usage:       "create an initial admin user",
			EnvVars:     []string{"KLEISTER_API_ADMIN_CREATE"},
			Destination: &cfg.Admin.Create,
		},
		&cli.StringFlag{
			Name:        "admin-username",
			Value:       "admin",
			Usage:       "initial admin username",
			EnvVars:     []string{"KLEISTER_API_ADMIN_USERNAME"},
			Destination: &cfg.Admin.Username,
		},
		&cli.StringFlag{
			Name:        "admin-password",
			Value:       "admin",
			Usage:       "initial admin password",
			EnvVars:     []string{"KLEISTER_API_ADMIN_PASSWORD"},
			Destination: &cfg.Admin.Username,
		},
		&cli.StringFlag{
			Name:        "admin-email",
			Value:       "",
			Usage:       "initial admin email",
			EnvVars:     []string{"KLEISTER_API_ADMIN_EMAIL"},
			Destination: &cfg.Admin.Email,
		},
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Value:       false,
			Usage:       "enable open tracing",
			EnvVars:     []string{"KLEISTER_API_TRACING_ENABLED"},
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "open tracing endpoint",
			EnvVars:     []string{"KLEISTER_API_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
	}
}

func serverBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		setupLogger(cfg)
		return nil
	}
}

func serverAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		tracing, err := setupTracing(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup tracing")
		}

		if tracing != nil {
			defer tracing.Close()
		}

		storage, err := setupStorage(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup database")
		}

		if storage != nil {
			defer storage.Close()
		}

		uploads, err := setupUploads(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup uploads")
		}

		log.Info().
			Msg(uploads.Info())

		if uploads != nil {
			defer uploads.Close()
		}

		var gr group.Group

		{
			server := &http.Server{
				Addr:         cfg.Server.Addr,
				Handler:      router.Server(cfg, storage, uploads),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Server.Addr).
					Msg("starting http server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("failed to shutdown http gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("http shutdown gracefully")
			})
		}

		{
			server := &http.Server{
				Addr:         cfg.Metrics.Addr,
				Handler:      router.Metrics(cfg, storage, uploads),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Metrics.Addr).
					Msg("starting metrics server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("failed to shutdown metrics gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("metrics shutdown gracefully")
			})
		}

		{
			stop := make(chan os.Signal, 1)

			gr.Add(func() error {
				signal.Notify(stop, os.Interrupt)

				<-stop

				return nil
			}, func(err error) {
				close(stop)
			})
		}

		return gr.Run()
	}
}
