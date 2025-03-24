package command

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/kleister/kleister-api/pkg/authn"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/router"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/store"
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
	defaultMetricsPprof     = false
	defaultServerAddr       = "0.0.0.0:8080"
	defaultServerHost       = "http://localhost:8080"
	defaultServerRoot       = "/"
	defaultServerCert       = ""
	defaultServerKey        = ""
	defaultServerTemplates  = ""
	defaultServerFrontend   = ""
	defaultServerDocs       = true
	defaultDatabaseDriver   = "sqlite3"
	defaultDatabaseAddress  = ""
	defaultDatabasePort     = ""
	defaultDatabaseUsername = ""
	defaultDatabasePassword = ""
	defaultDatabaseName     = "kleister.sqlite3"
	defaultDatabaseOptions  = make(map[string]string, 0)
	defaultUploadDriver     = "file"
	defaultUploadEndpoint   = ""
	defaultUploadPath       = ""
	defaultUploadAccess     = ""
	defaultUploadSecret     = ""
	defaultUploadBucket     = ""
	defaultUploadRegion     = "us-east-1"
	defaultUploadPerms      = "0755"
	defaultUploadPathstyle  = false
	defaultUploadProxy      = true
	defaultTokenSecret      = secret.Generate(32)
	defaultTokenExpire      = time.Hour * 1
	defaultScimEnabled      = false
	defaultScimToken        = ""
	defaultCleanupEnabled   = true
	defaultCleanupInterval  = 30 * time.Minute
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

	serverCmd.PersistentFlags().String("server-templates", defaultServerTemplates, "Path to custom template filrs")
	viper.SetDefault("server.templates", defaultServerTemplates)
	_ = viper.BindPFlag("server.templates", serverCmd.PersistentFlags().Lookup("server-templates"))

	serverCmd.PersistentFlags().String("server-frontend", defaultServerFrontend, "Path to custom frontend files")
	viper.SetDefault("server.frontend", defaultServerFrontend)
	_ = viper.BindPFlag("server.frontend", serverCmd.PersistentFlags().Lookup("server-frontend"))

	serverCmd.PersistentFlags().Bool("server-docs", defaultServerDocs, "Enable OpenAPI docs")
	viper.SetDefault("server.docs", defaultServerDocs)
	_ = viper.BindPFlag("server.docs", serverCmd.PersistentFlags().Lookup("server-docs"))

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

	serverCmd.PersistentFlags().Bool("upload-pathstyle", defaultUploadPathstyle, "Enable S3 pathstyle access")
	viper.SetDefault("upload.pathstyle", defaultUploadPathstyle)
	_ = viper.BindPFlag("upload.pathstyle", serverCmd.PersistentFlags().Lookup("upload-pathstyle"))

	serverCmd.PersistentFlags().Bool("upload-proxy", defaultUploadProxy, "Proxy S3 access through server")
	viper.SetDefault("upload.proxy", defaultUploadProxy)
	_ = viper.BindPFlag("upload.proxy", serverCmd.PersistentFlags().Lookup("upload-proxy"))

	serverCmd.PersistentFlags().String("token-secret", defaultTokenSecret, "Token encryption secret")
	viper.SetDefault("token.secret", defaultTokenSecret)
	_ = viper.BindPFlag("token.secret", serverCmd.PersistentFlags().Lookup("token-secret"))

	serverCmd.PersistentFlags().Duration("token-expire", defaultTokenExpire, "Token expire duration")
	viper.SetDefault("token.expire", defaultTokenExpire)
	_ = viper.BindPFlag("token.expire", serverCmd.PersistentFlags().Lookup("token-expire"))

	serverCmd.PersistentFlags().Bool("scim-enabled", defaultScimEnabled, "Enable SCIM provisioning integration")
	viper.SetDefault("scim.enabled", defaultScimEnabled)
	_ = viper.BindPFlag("scim.enabled", serverCmd.PersistentFlags().Lookup("scim-enabled"))

	serverCmd.PersistentFlags().String("scim-token", defaultScimToken, "Bearer token for SCIM authentication")
	viper.SetDefault("scim.token", defaultScimToken)
	_ = viper.BindPFlag("scim.token", serverCmd.PersistentFlags().Lookup("scim-token"))

	serverCmd.PersistentFlags().Bool("cleanup-enabled", defaultCleanupEnabled, "Enable periodic cleanup tasks")
	viper.SetDefault("cleanup.enabled", defaultCleanupEnabled)
	_ = viper.BindPFlag("cleanup.enabled", serverCmd.PersistentFlags().Lookup("cleanup-enabled"))

	serverCmd.PersistentFlags().Duration("cleanup-interval", defaultCleanupInterval, "Interval for cleanup task")
	viper.SetDefault("cleanup.interval", defaultCleanupInterval)
	_ = viper.BindPFlag("cleanup.interval", serverCmd.PersistentFlags().Lookup("cleanup-interval"))

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

func serverAction(ccmd *cobra.Command, _ []string) {
	identity, err := authn.New(
		authn.WithConfig(cfg.Auth.Config),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to setup identity")

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

	defer uploads.Close()

	storage, err := store.NewStore(
		cfg.Database,
		cfg.Scim,
		uploads,
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to setup database")

		os.Exit(1)
	}

	log.Info().
		Fields(storage.Info()).
		Msg("Preparing database")

	defer storage.Close()

	if val, err := backoff.Retry(
		ccmd.Context(),
		storage.Open,
		backoff.WithBackOff(backoff.NewExponentialBackOff()),
		backoff.WithNotify(func(err error, dur time.Duration) {
			log.Warn().
				Err(err).
				Dur("retry", dur).
				Msg("Database open failed")
		}),
	); err != nil || !val {
		log.Fatal().
			Err(err).
			Msg("Giving up to connect to db")

		os.Exit(1)
	}

	if val, err := backoff.Retry(
		ccmd.Context(),
		storage.Ping,
		backoff.WithBackOff(backoff.NewExponentialBackOff()),
		backoff.WithNotify(func(err error, dur time.Duration) {
			log.Warn().
				Err(err).
				Dur("retry", dur).
				Msg("Database ping failed")
		}),
	); err != nil || !val {
		log.Fatal().
			Err(err).
			Msg("Giving up to ping the db")

		os.Exit(1)
	}

	if _, err := storage.Migrate(ccmd.Context()); err != nil {
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

	gr := run.Group{}

	{
		server := &http.Server{
			Addr: cfg.Server.Addr,
			Handler: router.Server(
				cfg,
				registry,
				identity,
				uploads,
				storage,
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
			Addr: cfg.Metrics.Addr,
			Handler: router.Metrics(
				cfg,
				registry,
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

	if cfg.Cleanup.Enabled {
		ticker := time.NewTicker(cfg.Cleanup.Interval)
		stop := make(chan struct{})

		gr.Add(func() error {
			defer ticker.Stop()

			log.Info().
				Str("interval", cfg.Cleanup.Interval.String()).
				Msg("Starting periodic cleanup")

			for {
				select {
				case <-ticker.C:
					log.Debug().
						Msg("Running periodic cleanup")

					if err := storage.Users.CleanupRedirectTokens(
						context.Background(),
					); err != nil {
						log.Error().
							Err(err).
							Msg("Failed to cleanup redirect tokens")
					}
				case <-stop:
					log.Info().
						Msg("Shutdown periodic cleanup")

					return nil
				}
			}
		}, func(_ error) {
			close(stop)
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
