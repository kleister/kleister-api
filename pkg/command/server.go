package command

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/providers"
	"github.com/kleister/kleister-api/pkg/router"
	"github.com/kleister/kleister-api/pkg/secret"
	buildversions "github.com/kleister/kleister-api/pkg/service/build_versions"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/fabric"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/neoforge"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/quilt"
	teammods "github.com/kleister/kleister-api/pkg/service/team_mods"
	teampacks "github.com/kleister/kleister-api/pkg/service/team_packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	usermods "github.com/kleister/kleister-api/pkg/service/user_mods"
	userpacks "github.com/kleister/kleister-api/pkg/service/user_packs"
	userteams "github.com/kleister/kleister-api/pkg/service/user_teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/session"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start integrated server",
		Run:   serverAction,
		Args:  cobra.NoArgs,
	}

	defaultMetricsAddr      = "0.0.0.0:8000"
	defaultMetricsToken     = ""
	defaultServerAddr       = "0.0.0.0:8080"
	defaultMetricsPprof     = false
	defaultServerHost       = "http://localhost:8080"
	defaultServerRoot       = "/api"
	defaultServerCert       = ""
	defaultServerKey        = ""
	defaultDatabaseDriver   = "sqlite3"
	defaultDatabaseAddress  = ""
	defaultDatabasePort     = ""
	defaultDatabaseUsername = ""
	defaultDatabasePassword = ""
	defaultDatabaseName     = "storage/kleister.sqlite3"
	defaultDatabaseOptions  = make(map[string]string, 0)
	defaultUploadDriver     = "file"
	defaultUploadEndpoint   = ""
	defaultUploadPath       = "storage/uploads/"
	defaultUploadAccess     = ""
	defaultUploadSecret     = ""
	defaultUploadBucket     = ""
	defaultUploadRegion     = "us-east-1"
	defaultUploadPerms      = "0755"
	defaultSessionName      = "kleister"
	defaultSessionSecret    = secret.Generate(32)
	defaultSessionExpire    = time.Hour * 24
	defaultSessionSecure    = false
	defaultAdminCreate      = true
	defaultAdminUsername    = "admin"
	defaultAdminPassword    = "admin"
	defaultAdminEmail       = "admin@localhost"
	defaultAuthConfig       = ""
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().String("metrics-addr", defaultMetricsAddr, "Address to bind the metrics")
	viper.SetDefault("metrics.addr", defaultMetricsAddr)
	_ = viper.BindPFlag("metrics.addr", serverCmd.PersistentFlags().Lookup("metrics-addr"))

	serverCmd.PersistentFlags().String("metrics-token", defaultMetricsToken, "Token to make metrics secure")
	viper.SetDefault("metrics.token", defaultMetricsToken)
	_ = viper.BindPFlag("metrics.token", serverCmd.PersistentFlags().Lookup("metrics-token"))

	serverCmd.PersistentFlags().Bool("metrics-pprof", defaultMetricsPprof, "Enable pprof debugging")
	viper.SetDefault("metrics.pprof", defaultMetricsPprof)
	_ = viper.BindPFlag("metrics.pprof", serverCmd.PersistentFlags().Lookup("metrics-pprof"))

	serverCmd.PersistentFlags().String("server-addr", defaultServerAddr, "Address to bind the server")
	viper.SetDefault("server.addr", defaultServerAddr)
	_ = viper.BindPFlag("server.addr", serverCmd.PersistentFlags().Lookup("server-addr"))

	serverCmd.PersistentFlags().String("server-host", defaultServerHost, "External access to server")
	viper.SetDefault("server.host", defaultServerHost)
	_ = viper.BindPFlag("server.host", serverCmd.PersistentFlags().Lookup("server-host"))

	serverCmd.PersistentFlags().String("server-root", defaultServerRoot, "Path to access the server")
	viper.SetDefault("server.root", defaultServerRoot)
	_ = viper.BindPFlag("server.root", serverCmd.PersistentFlags().Lookup("server-root"))

	serverCmd.PersistentFlags().String("server-cert", defaultServerCert, "Path to SSL cert")
	viper.SetDefault("server.cert", defaultServerCert)
	_ = viper.BindPFlag("server.cert", serverCmd.PersistentFlags().Lookup("server-cert"))

	serverCmd.PersistentFlags().String("server-key", defaultServerKey, "Path to SSL key")
	viper.SetDefault("server.key", defaultServerKey)
	_ = viper.BindPFlag("server.key", serverCmd.PersistentFlags().Lookup("server-key"))

	serverCmd.PersistentFlags().String("database-driver", defaultDatabaseDriver, "Driver for the database")
	viper.SetDefault("database.driver", defaultDatabaseDriver)
	_ = viper.BindPFlag("database.driver", serverCmd.PersistentFlags().Lookup("database-driver"))

	serverCmd.PersistentFlags().String("database-address", defaultDatabaseAddress, "Address for the database")
	viper.SetDefault("database.address", defaultDatabaseAddress)
	_ = viper.BindPFlag("database.address", serverCmd.PersistentFlags().Lookup("database-address"))

	serverCmd.PersistentFlags().String("database-port", defaultDatabasePort, "Port for the database")
	viper.SetDefault("database.port", defaultDatabasePort)
	_ = viper.BindPFlag("database.port", serverCmd.PersistentFlags().Lookup("database-port"))

	serverCmd.PersistentFlags().String("database-username", defaultDatabaseUsername, "Username for the database")
	viper.SetDefault("database.username", defaultDatabaseUsername)
	_ = viper.BindPFlag("database.username", serverCmd.PersistentFlags().Lookup("database-username"))

	serverCmd.PersistentFlags().String("database-password", defaultDatabasePassword, "Password for the database")
	viper.SetDefault("database.password", defaultDatabasePassword)
	_ = viper.BindPFlag("database.password", serverCmd.PersistentFlags().Lookup("database-password"))

	serverCmd.PersistentFlags().String("database-name", defaultDatabaseName, "Name of the database or path for local databases")
	viper.SetDefault("database.name", defaultDatabaseName)
	_ = viper.BindPFlag("database.name", serverCmd.PersistentFlags().Lookup("database-name"))

	serverCmd.PersistentFlags().StringToString("database-options", defaultDatabaseOptions, "Options for the database connection")
	viper.SetDefault("database.options", defaultDatabaseOptions)
	_ = viper.BindPFlag("database.options", serverCmd.PersistentFlags().Lookup("database-options"))

	serverCmd.PersistentFlags().String("upload-driver", defaultUploadDriver, "Driver for the uploads")
	viper.SetDefault("upload.driver", defaultUploadDriver)
	_ = viper.BindPFlag("upload.driver", serverCmd.PersistentFlags().Lookup("upload-driver"))

	serverCmd.PersistentFlags().String("upload-endpoint", defaultUploadEndpoint, "Endpoint for uploads")
	viper.SetDefault("upload.endpoint", defaultUploadEndpoint)
	_ = viper.BindPFlag("upload.endpoint", serverCmd.PersistentFlags().Lookup("upload-endpoint"))

	serverCmd.PersistentFlags().String("upload-path", defaultUploadPath, "Path to store uploads")
	viper.SetDefault("upload.path", defaultUploadPath)
	_ = viper.BindPFlag("upload.path", serverCmd.PersistentFlags().Lookup("upload-path"))

	serverCmd.PersistentFlags().String("upload-access", defaultUploadAccess, "Access key for uploads")
	viper.SetDefault("upload.access", defaultUploadAccess)
	_ = viper.BindPFlag("upload.access", serverCmd.PersistentFlags().Lookup("upload-access"))

	serverCmd.PersistentFlags().String("upload-secret", defaultUploadSecret, "Secret key for uploads")
	viper.SetDefault("upload.secret", defaultUploadSecret)
	_ = viper.BindPFlag("upload.secret", serverCmd.PersistentFlags().Lookup("upload-secret"))

	serverCmd.PersistentFlags().String("upload-bucket", defaultUploadBucket, "Bucket to store uploads")
	viper.SetDefault("upload.bucket", defaultUploadBucket)
	_ = viper.BindPFlag("upload.bucket", serverCmd.PersistentFlags().Lookup("upload-bucket"))

	serverCmd.PersistentFlags().String("upload-region", defaultUploadRegion, "Region to store uploads")
	viper.SetDefault("upload.region", defaultUploadRegion)
	_ = viper.BindPFlag("upload.region", serverCmd.PersistentFlags().Lookup("upload-region"))

	serverCmd.PersistentFlags().String("upload-perms", defaultUploadPerms, "Chmod value for upload path")
	viper.SetDefault("upload.perms", defaultUploadPerms)
	_ = viper.BindPFlag("upload.perms", serverCmd.PersistentFlags().Lookup("upload-perms"))

	serverCmd.PersistentFlags().String("session-name", defaultSessionName, "Session cookie name")
	viper.SetDefault("session.name", defaultSessionName)
	_ = viper.BindPFlag("session.name", serverCmd.PersistentFlags().Lookup("session-name"))

	serverCmd.PersistentFlags().String("session-secret", defaultSessionSecret, "Session encryption secret")
	viper.SetDefault("session.secret", defaultSessionSecret)
	_ = viper.BindPFlag("session.secret", serverCmd.PersistentFlags().Lookup("session-secret"))

	serverCmd.PersistentFlags().Duration("session-expire", defaultSessionExpire, "Session expire duration")
	viper.SetDefault("session.expire", defaultSessionExpire)
	_ = viper.BindPFlag("session.expire", serverCmd.PersistentFlags().Lookup("session-expire"))

	serverCmd.PersistentFlags().Bool("session-secure", defaultSessionSecure, "Enable secure cookie on HTTPS")
	viper.SetDefault("session.secure", defaultSessionSecure)
	_ = viper.BindPFlag("session.secure", serverCmd.PersistentFlags().Lookup("session-secure"))

	serverCmd.PersistentFlags().Bool("admin-create", defaultAdminCreate, "Create an initial admin user")
	viper.SetDefault("admin.create", defaultAdminCreate)
	_ = viper.BindPFlag("admin.create", serverCmd.PersistentFlags().Lookup("admin-create"))

	serverCmd.PersistentFlags().String("admin-username", defaultAdminUsername, "Initial admin username")
	viper.SetDefault("admin.username", defaultAdminUsername)
	_ = viper.BindPFlag("admin.username", serverCmd.PersistentFlags().Lookup("admin-username"))

	serverCmd.PersistentFlags().String("admin-password", defaultAdminPassword, "Initial admin password")
	viper.SetDefault("admin.password", defaultAdminPassword)
	_ = viper.BindPFlag("admin.password", serverCmd.PersistentFlags().Lookup("admin-password"))

	serverCmd.PersistentFlags().String("admin-email", defaultAdminEmail, "Initial admin email")
	viper.SetDefault("admin.email", defaultAdminEmail)
	_ = viper.BindPFlag("admin.email", serverCmd.PersistentFlags().Lookup("admin-email"))

	serverCmd.PersistentFlags().String("auth-config", defaultAuthConfig, "Path to authentication config for OAuth2/OIDC")
	viper.SetDefault("auth.config", defaultAuthConfig)
	_ = viper.BindPFlag("auth.config", serverCmd.PersistentFlags().Lookup("auth-config"))
}

func serverAction(_ *cobra.Command, _ []string) {
	if err := providers.Register(
		providers.WithConfig(cfg.Auth.Config),
	); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to load providers")

		os.Exit(1)
	}

	uploads, err := setupUploads(cfg)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to setup uploads")

		os.Exit(1)
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

		os.Exit(1)
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
		func(_ error, dur time.Duration) {
			log.Warn().
				Dur("retry", dur).
				Msg("Database open failed")
		},
	); err != nil {
		log.Fatal().
			Err(err).
			Msg("Giving up to connect to db")

		os.Exit(1)
	}

	if err := backoff.RetryNotify(
		storage.Ping,
		backoff.NewExponentialBackOff(),
		func(_ error, dur time.Duration) {
			log.Warn().
				Dur("retry", dur).
				Msg("Database ping failed")
		},
	); err != nil {
		log.Fatal().
			Err(err).
			Msg("Giving up to ping the db")

		os.Exit(1)
	}

	if err := storage.Migrate(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to migrate database")
	}

	if cfg.Admin.Create {
		username, err := config.Value(cfg.Admin.Username)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to parse admin username secret")

			os.Exit(1)
		}

		password, err := config.Value(cfg.Admin.Password)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to parse admin password secret")

			os.Exit(1)
		}

		email, err := config.Value(cfg.Admin.Email)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to parse admin email secret")

			os.Exit(1)
		}

		if err := storage.Admin(
			username,
			password,
			email,
		); err != nil {
			log.Warn().
				Err(err).
				Str("username", username).
				Str("email", email).
				Msg("Failed to create admin")
		} else {
			log.Info().
				Str("username", username).
				Str("email", email).
				Msg("Admin successfully stored")
		}
	}

	token, err := config.Value(cfg.Metrics.Token)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to parse metrics token secret")

		os.Exit(1)
	}

	registry := metrics.New(
		metrics.WithNamespace("kleister_api"),
		metrics.WithToken(token),
	)

	sess := session.New(
		session.WithStore(storage.Session()),
		session.WithLifetime(cfg.Session.Expire),
		session.WithName(cfg.Session.Name),
		session.WithPath(cfg.Server.Root),
		session.WithSecure(cfg.Session.Secure),
	)

	gr := run.Group{}

	{
		minecraftService := minecraft.NewMetricsService(
			minecraft.NewLoggingService(
				minecraft.NewService(
					minecraft.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		forgeService := forge.NewMetricsService(
			forge.NewLoggingService(
				forge.NewService(
					forge.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		neoforgeService := neoforge.NewMetricsService(
			neoforge.NewLoggingService(
				neoforge.NewService(
					neoforge.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		quiltService := quilt.NewMetricsService(
			quilt.NewLoggingService(
				quilt.NewService(
					quilt.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		fabricService := fabric.NewMetricsService(
			fabric.NewLoggingService(
				fabric.NewService(
					fabric.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		teamsService := teams.NewMetricsService(
			teams.NewLoggingService(
				teams.NewService(
					teams.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		usersService := users.NewMetricsService(
			users.NewLoggingService(
				users.NewService(
					users.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		userteamsService := userteams.NewMetricsService(
			userteams.NewLoggingService(
				userteams.NewService(
					userteams.NewGormService(
						storage.Handle(),
						cfg,
						teamsService,
						usersService,
					),
				),
			),
			registry,
		)

		modsService := mods.NewMetricsService(
			mods.NewLoggingService(
				mods.NewService(
					mods.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		usermodsService := usermods.NewMetricsService(
			usermods.NewLoggingService(
				usermods.NewService(
					usermods.NewGormService(
						storage.Handle(),
						cfg,
						usersService,
						modsService,
					),
				),
			),
			registry,
		)

		teammodsService := teammods.NewMetricsService(
			teammods.NewLoggingService(
				teammods.NewService(
					teammods.NewGormService(
						storage.Handle(),
						cfg,
						teamsService,
						modsService,
					),
				),
			),
			registry,
		)

		versionsService := versions.NewMetricsService(
			versions.NewLoggingService(
				versions.NewService(
					versions.NewGormService(
						storage.Handle(),
						cfg,
						modsService,
					),
				),
			),
			registry,
		)

		packsService := packs.NewMetricsService(
			packs.NewLoggingService(
				packs.NewService(
					packs.NewGormService(
						storage.Handle(),
						cfg,
					),
				),
			),
			registry,
		)

		userpacksService := userpacks.NewMetricsService(
			userpacks.NewLoggingService(
				userpacks.NewService(
					userpacks.NewGormService(
						storage.Handle(),
						cfg,
						usersService,
						packsService,
					),
				),
			),
			registry,
		)

		teampacksService := teampacks.NewMetricsService(
			teampacks.NewLoggingService(
				teampacks.NewService(
					teampacks.NewGormService(
						storage.Handle(),
						cfg,
						teamsService,
						packsService,
					),
				),
			),
			registry,
		)

		buildsService := builds.NewMetricsService(
			builds.NewLoggingService(
				builds.NewService(
					builds.NewGormService(
						storage.Handle(),
						cfg,
						packsService,
					),
				),
			),
			registry,
		)

		buildversionsService := buildversions.NewMetricsService(
			buildversions.NewLoggingService(
				buildversions.NewService(
					buildversions.NewGormService(
						storage.Handle(),
						cfg,
						packsService,
						buildsService,
						modsService,
						versionsService,
					),
				),
			),
			registry,
		)

		server := &http.Server{
			Addr: cfg.Server.Addr,
			Handler: router.Server(
				cfg,
				registry,
				sess,
				uploads,
				storage,
				minecraftService,
				forgeService,
				neoforgeService,
				quiltService,
				fabricService,
				teamsService,
				usersService,
				userteamsService,
				modsService,
				usermodsService,
				teammodsService,
				versionsService,
				packsService,
				userpacksService,
				teampacksService,
				buildsService,
				buildversionsService,
			),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		gr.Add(func() error {
			log.Info().
				Str("addr", cfg.Server.Addr).
				Msg("Starting application server")

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
					Msg("Failed to shutdown application gracefully")

				return
			}

			log.Info().
				Err(reason).
				Msg("Shutdown application gracefully")
		})
	}

	{
		server := &http.Server{
			Addr:         cfg.Metrics.Addr,
			Handler:      router.Metrics(cfg, registry),
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
		}, func(_ error) {
			close(stop)
		})
	}

	if err := gr.Run(); err != nil {
		os.Exit(1)
	}
}
