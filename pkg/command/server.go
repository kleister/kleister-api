package command

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/dchest/uniuri"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/router"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "start integrated server",
		Flags:  serverFlags(cfg),
		Action: serverAction(cfg),
	}
}

func serverFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "metrics-addr",
			Value:       defaultMetricsAddr,
			Usage:       "Address to bind the metrics",
			EnvVars:     []string{"KLEISTER_API_METRICS_ADDR"},
			Destination: &cfg.Metrics.Addr,
		},
		&cli.StringFlag{
			Name:        "metrics-token",
			Value:       "",
			Usage:       "Token to make metrics secure",
			EnvVars:     []string{"KLEISTER_API_METRICS_TOKEN"},
			Destination: &cfg.Metrics.Token,
		},
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       defaultServerAddress,
			Usage:       "Address to bind the server",
			EnvVars:     []string{"KLEISTER_API_SERVER_ADDR"},
			Destination: &cfg.Server.Addr,
		},
		&cli.BoolFlag{
			Name:        "server-pprof",
			Value:       false,
			Usage:       "Enable pprof debugging",
			EnvVars:     []string{"KLEISTER_API_SERVER_PPROF"},
			Destination: &cfg.Server.Pprof,
		},
		&cli.BoolFlag{
			Name:        "server-docs",
			Value:       true,
			Usage:       "Enable swagger documentation",
			EnvVars:     []string{"KLEISTER_API_SERVER_DOCS"},
			Destination: &cfg.Server.Docs,
		},
		&cli.StringFlag{
			Name:        "server-host",
			Value:       "http://localhost:8080",
			Usage:       "External access to server",
			EnvVars:     []string{"KLEISTER_API_SERVER_HOST"},
			Destination: &cfg.Server.Host,
		},
		&cli.StringFlag{
			Name:        "server-root",
			Value:       "/",
			Usage:       "Path to access the server",
			EnvVars:     []string{"KLEISTER_API_SERVER_ROOT"},
			Destination: &cfg.Server.Root,
		},
		&cli.StringFlag{
			Name:        "db-dsn",
			Value:       "boltdb://storage/kleister.db",
			Usage:       "Database dsn",
			EnvVars:     []string{"KLEISTER_API_DB_DSN"},
			Destination: &cfg.Database.DSN,
		},
		&cli.StringFlag{
			Name:        "upload-dsn",
			Value:       "file://storage/uploads/",
			Usage:       "Uploads dsn",
			EnvVars:     []string{"KLEISTER_API_UPLOAD_DSN"},
			Destination: &cfg.Upload.DSN,
		},
		&cli.DurationFlag{
			Name:        "session-expire",
			Value:       time.Hour * 24,
			Usage:       "Session expire duration",
			EnvVars:     []string{"KLEISTER_API_SESSION_EXPIRE"},
			Destination: &cfg.Session.Expire,
		},
		&cli.StringFlag{
			Name:        "session-secret",
			Value:       uniuri.NewLen(32),
			Usage:       "Session encryption secret",
			EnvVars:     []string{"KLEISTER_API_SESSION_SECRET"},
			Destination: &cfg.Session.Secret,
		},
		&cli.BoolFlag{
			Name:        "admin-create",
			Value:       true,
			Usage:       "Create an initial admin user",
			EnvVars:     []string{"KLEISTER_API_ADMIN_CREATE"},
			Destination: &cfg.Admin.Create,
		},
		&cli.StringFlag{
			Name:        "admin-username",
			Value:       "admin",
			Usage:       "Initial admin username",
			EnvVars:     []string{"KLEISTER_API_ADMIN_USERNAME"},
			Destination: &cfg.Admin.Username,
		},
		&cli.StringFlag{
			Name:        "admin-password",
			Value:       "admin",
			Usage:       "Initial admin password",
			EnvVars:     []string{"KLEISTER_API_ADMIN_PASSWORD"},
			Destination: &cfg.Admin.Password,
		},
		&cli.StringFlag{
			Name:        "admin-email",
			Value:       "",
			Usage:       "Initial admin email",
			EnvVars:     []string{"KLEISTER_API_ADMIN_EMAIL"},
			Destination: &cfg.Admin.Email,
		},
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Value:       false,
			Usage:       "Enable open tracing",
			EnvVars:     []string{"KLEISTER_API_TRACING_ENABLED"},
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Open tracing endpoint",
			EnvVars:     []string{"KLEISTER_API_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
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

		uploads, err := setupUploads(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup uploads")

			return err
		}

		log.Info().
			Fields(uploads.Info()).
			Msg("preparing uploads")

		if uploads != nil {
			defer uploads.Close()
		}

		storage, err := setupStorage(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup database")

			return err
		}

		log.Info().
			Fields(storage.Info()).
			Msg("preparing database")

		if storage != nil {
			defer storage.Close()
		}

		if err := backoff.RetryNotify(
			storage.Open,
			backoff.NewExponentialBackOff(),
			func(err error, dur time.Duration) {
				log.Warn().
					Dur("retry", dur).
					Msg("database open failed")
			},
		); err != nil {
			log.Fatal().
				Err(err).
				Msg("giving up to connect to db")

			return err
		}

		if err := backoff.RetryNotify(
			storage.Ping,
			backoff.NewExponentialBackOff(),
			func(err error, dur time.Duration) {
				log.Warn().
					Dur("retry", dur).
					Msg("database ping failed")
			},
		); err != nil {
			log.Fatal().
				Err(err).
				Msg("giving up to ping the db")

			return err
		}

		if err := storage.Migrate(); err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to migrate database")
		}

		if cfg.Admin.Create {
			err := storage.Admin(
				cfg.Admin.Username,
				cfg.Admin.Password,
				cfg.Admin.Email,
			)

			if err != nil {
				log.Warn().
					Err(err).
					Str("username", cfg.Admin.Username).
					Str("password", cfg.Admin.Password).
					Str("email", cfg.Admin.Email).
					Msg("failed to create admin")
			} else {
				log.Info().
					Str("username", cfg.Admin.Username).
					Str("password", cfg.Admin.Password).
					Str("email", cfg.Admin.Email).
					Msg("admin successfully stored")
			}
		}

		metrics := metrics.New()

		teamsService := teams.NewTracingService(
			teams.NewMetricsService(
				teams.NewLoggingService(
					teams.NewService(
						storage.Teams(),
					),
					requestid.Get,
				),
				metrics,
			),
			requestid.Get,
		)

		usersService := users.NewTracingService(
			users.NewMetricsService(
				users.NewLoggingService(
					users.NewService(
						storage.Users(),
					),
					requestid.Get,
				),
				metrics,
			),
			requestid.Get,
		)

		minecraftService := minecraft.NewTracingService(
			minecraft.NewMetricsService(
				minecraft.NewLoggingService(
					minecraft.NewService(
						storage.Minecraft(),
					),
					requestid.Get,
				),
				metrics,
			),
			requestid.Get,
		)

		forgeService := forge.NewTracingService(
			forge.NewMetricsService(
				forge.NewLoggingService(
					forge.NewService(
						storage.Forge(),
					),
					requestid.Get,
				),
				metrics,
			),
			requestid.Get,
		)

		var gr run.Group

		{
			server := &http.Server{
				Addr: cfg.Server.Addr,
				Handler: router.Server(
					cfg,
					uploads,
					usersService,
					teamsService,
					minecraftService,
					forgeService,
				),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Server.Addr).
					Msg("starting http server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
				Addr: cfg.Metrics.Addr,
				Handler: router.Metrics(
					cfg,
					metrics,
				),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Metrics.Addr).
					Msg("starting metrics server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
