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
	"github.com/kleister/kleister-api/pkg/service/members"
	membersRepository "github.com/kleister/kleister-api/pkg/service/members/repository"
	"github.com/kleister/kleister-api/pkg/service/teams"
	teamsRepository "github.com/kleister/kleister-api/pkg/service/teams/repository"
	"github.com/kleister/kleister-api/pkg/service/users"
	usersRepository "github.com/kleister/kleister-api/pkg/service/users/repository"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "Start integrated server",
		Flags:  ServerFlags(cfg),
		Action: ServerAction(cfg),
	}
}

// ServerFlags defines server flags.
func ServerFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Value:       false,
			Usage:       "Enable pprof debugging",
			EnvVars:     []string{"KLEISTER_API_DEBUG_PPROF"},
			Destination: &cfg.Metrics.Pprof,
		},
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
			FilePath:    "/etc/kleister/secrets/metrics-token",
		},
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       defaultServerAddress,
			Usage:       "Address to bind the server",
			EnvVars:     []string{"KLEISTER_API_SERVER_ADDR"},
			Destination: &cfg.Server.Addr,
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
			Name:        "server-cert",
			Value:       "",
			Usage:       "Path to SSL cert",
			EnvVars:     []string{"KLEISTER_API_SERVER_CERT"},
			Destination: &cfg.Server.Cert,
		},
		&cli.StringFlag{
			Name:        "server-key",
			Value:       "",
			Usage:       "Path to SSL key",
			EnvVars:     []string{"KLEISTER_API_SERVER_KEY"},
			Destination: &cfg.Server.Key,
		},
		&cli.StringFlag{
			Name:        "database-dsn",
			Value:       "sqlite3://storage/kleister.db",
			Usage:       "Database dsn",
			EnvVars:     []string{"KLEISTER_API_DATABASE_DSN"},
			Destination: &cfg.Database.DSN,
			FilePath:    "/etc/kleister/secrets/database-dsn",
		},
		&cli.StringFlag{
			Name:        "upload-dsn",
			Value:       "file://storage/uploads/",
			Usage:       "Uploads dsn",
			EnvVars:     []string{"KLEISTER_API_UPLOAD_DSN"},
			Destination: &cfg.Upload.DSN,
			FilePath:    "/etc/kleister/secrets/upload-dsn",
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
			FilePath:    "/etc/kleister/secrets/session-secret",
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
			FilePath:    "/etc/kleister/secrets/admin-username",
		},
		&cli.StringFlag{
			Name:        "admin-password",
			Value:       "admin",
			Usage:       "Initial admin password",
			EnvVars:     []string{"KLEISTER_API_ADMIN_PASSWORD"},
			Destination: &cfg.Admin.Password,
			FilePath:    "/etc/kleister/secrets/admin-password",
		},
		&cli.StringFlag{
			Name:        "admin-email",
			Value:       "",
			Usage:       "Initial admin email",
			EnvVars:     []string{"KLEISTER_API_ADMIN_EMAIL"},
			Destination: &cfg.Admin.Email,
			FilePath:    "/etc/kleister/secrets/admin-email",
		},
	}
}

// ServerAction defines server action.
func ServerAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		uploads, err := setupUploads(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to setup uploads")

			return err
		}

		log.Info().
			Fields(uploads.Info()).
			Msg("Preparing uploads")

		if uploads != nil {
			defer uploads.Close()
		}

		storage, err := setupStorage(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to setup database")

			return err
		}

		log.Info().
			Fields(storage.Info()).
			Msg("Preparing database")

		if storage != nil {
			defer storage.Close()
		}

		if err := backoff.RetryNotify(
			storage.Open,
			backoff.NewExponentialBackOff(),
			func(err error, dur time.Duration) {
				log.Warn().
					Dur("retry", dur).
					Msg("Database open failed")
			},
		); err != nil {
			log.Fatal().
				Err(err).
				Msg("Giving up to connect to db")

			return err
		}

		if err := backoff.RetryNotify(
			storage.Ping,
			backoff.NewExponentialBackOff(),
			func(err error, dur time.Duration) {
				log.Warn().
					Dur("retry", dur).
					Msg("Database ping failed")
			},
		); err != nil {
			log.Fatal().
				Err(err).
				Msg("Giving up to ping the db")

			return err
		}

		if err := storage.Migrate(); err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to migrate database")
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
					Msg("Failed to create admin")
			} else {
				log.Info().
					Str("username", cfg.Admin.Username).
					Str("password", cfg.Admin.Password).
					Str("email", cfg.Admin.Email).
					Msg("Admin successfully stored")
			}
		}

		metricz := metrics.New()
		gr := run.Group{}

		{
			routing := router.Server(
				cfg,
				uploads,
			)

			usersRepo := usersRepository.NewMetricsRepository(
				usersRepository.NewLoggingRepository(
					usersRepository.NewGormRepository(
						storage.Handle(),
					),
					requestid.Get,
				),
				metricz,
			)

			users.RegisterServer(
				cfg,
				uploads,
				metricz,
				usersRepo,
				routing,
			)

			teamsRepo := teamsRepository.NewMetricsRepository(
				teamsRepository.NewLoggingRepository(
					teamsRepository.NewGormRepository(
						storage.Handle(),
					),
					requestid.Get,
				),
				metricz,
			)

			teams.RegisterServer(
				cfg,
				uploads,
				metricz,
				teamsRepo,
				routing,
			)

			membersRepo := membersRepository.NewMetricsRepository(
				membersRepository.NewLoggingRepository(
					membersRepository.NewGormRepository(
						storage.Handle(),
						teamsRepo,
						usersRepo,
					),
					requestid.Get,
				),
				metricz,
			)

			members.RegisterServer(
				cfg,
				uploads,
				metricz,
				membersRepo,
				routing,
			)

			server := &http.Server{
				Addr: cfg.Server.Addr,
				Handler: h2c.NewHandler(
					routing,
					&http2.Server{},
				),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Server.Addr).
					Msg("Starting HTTP server")

				if cfg.Server.Cert != "" && cfg.Server.Key != "" {
					return server.ListenAndServeTLS(
						cfg.Server.Cert,
						cfg.Server.Key,
					)
				}

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("Failed to shutdown HTTP gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("HTTP shutdown gracefully")
			})
		}

		{
			routing := router.Metrics(
				cfg,
				metricz,
			)

			server := &http.Server{
				Addr: cfg.Metrics.Addr,
				Handler: h2c.NewHandler(
					routing,
					&http2.Server{},
				),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Metrics.Addr).
					Msg("Starting metrics server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("Failed to shutdown metrics gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("Metrics shutdown gracefully")
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
