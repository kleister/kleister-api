package command

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbCmd = &cobra.Command{
		Use:   "database",
		Short: "Database migrations",
		Args:  cobra.NoArgs,
	}

	dbCleanupCmd = &cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup expired content",
		Run:   dbCleanupAction,
		Args:  cobra.NoArgs,
	}

	dbMigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Execute migrations",
		Run:   dbMigrateAction,
		Args:  cobra.NoArgs,
	}

	dbRollbackCmd = &cobra.Command{
		Use:   "rollback",
		Short: "Rollback migrations",
		Run:   dbRollbackAction,
		Args:  cobra.NoArgs,
	}

	dbLockCmd = &cobra.Command{
		Use:   "lock",
		Short: "Lock migrations",
		Run:   dbLockAction,
		Args:  cobra.NoArgs,
	}

	dbUnlockCmd = &cobra.Command{
		Use:   "unlock",
		Short: "Unlock migrations",
		Run:   dbUnlockAction,
		Args:  cobra.NoArgs,
	}

	dbStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Status of migrations",
		Run:   dbStatusAction,
		Args:  cobra.NoArgs,
	}

	dbCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new migration",
		Run:   dbCreateAction,
	}
)

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.AddCommand(dbCleanupCmd)
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbRollbackCmd)
	dbCmd.AddCommand(dbLockCmd)
	dbCmd.AddCommand(dbUnlockCmd)
	dbCmd.AddCommand(dbStatusCmd)
	dbCmd.AddCommand(dbCreateCmd)

	dbCmd.PersistentFlags().String("database-driver", defaultDatabaseDriver, "Driver for the database")
	viper.SetDefault("database.driver", defaultDatabaseDriver)
	_ = viper.BindPFlag("database.driver", serverCmd.PersistentFlags().Lookup("database-driver"))

	dbCmd.PersistentFlags().String("database-address", defaultDatabaseAddress, "Address for the database")
	viper.SetDefault("database.address", defaultDatabaseAddress)
	_ = viper.BindPFlag("database.address", serverCmd.PersistentFlags().Lookup("database-address"))

	dbCmd.PersistentFlags().String("database-port", defaultDatabasePort, "Port for the database")
	viper.SetDefault("database.port", defaultDatabasePort)
	_ = viper.BindPFlag("database.port", serverCmd.PersistentFlags().Lookup("database-port"))

	dbCmd.PersistentFlags().String("database-username", defaultDatabaseUsername, "Username for the database")
	viper.SetDefault("database.username", defaultDatabaseUsername)
	_ = viper.BindPFlag("database.username", serverCmd.PersistentFlags().Lookup("database-username"))

	dbCmd.PersistentFlags().String("database-password", defaultDatabasePassword, "Password for the database")
	viper.SetDefault("database.password", defaultDatabasePassword)
	_ = viper.BindPFlag("database.password", serverCmd.PersistentFlags().Lookup("database-password"))

	dbCmd.PersistentFlags().String("database-name", defaultDatabaseName, "Name of the database or path for local databases")
	viper.SetDefault("database.name", defaultDatabaseName)
	_ = viper.BindPFlag("database.name", serverCmd.PersistentFlags().Lookup("database-name"))

	dbCmd.PersistentFlags().StringToString("database-options", defaultDatabaseOptions, "Options for the database connection")
	viper.SetDefault("database.options", defaultDatabaseOptions)
	_ = viper.BindPFlag("database.options", serverCmd.PersistentFlags().Lookup("database-options"))
}

func dbCleanupAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	if err := storage.Users.CleanupRedirectTokens(
		context.Background(),
	); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to cleanup redirect tokens")

		os.Exit(1)
	}

	log.Info().
		Msg("Finished cleanup task")
}

func dbMigrateAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	group, err := storage.Migrate(ccmd.Context())

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to migrate database")

		os.Exit(1)
	}

	if group.IsZero() {
		log.Info().
			Msg("Noting to migrate")
	} else {
		log.Debug().
			Msg("Finished migrate")
	}
}

func dbRollbackAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	group, err := storage.Rollback(ccmd.Context())

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to rollback database")

		os.Exit(1)
	}

	if group.IsZero() {
		log.Info().
			Msg("Noting to rollback")
	} else {
		log.Debug().
			Msg("Finished rollback")
	}
}

func dbLockAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	migrator, err := storage.Migrator(
		ccmd.Context(),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init migrator")

		os.Exit(1)
	}

	if err := migrator.Lock(ccmd.Context()); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to lock database")

		os.Exit(1)
	}

	log.Debug().
		Msg("Finished locking")
}

func dbUnlockAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	migrator, err := storage.Migrator(
		ccmd.Context(),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init migrator")

		os.Exit(1)
	}

	if err := migrator.Unlock(ccmd.Context()); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to unlock database")

		os.Exit(1)
	}

	log.Debug().
		Msg("Finished unlocking")
}

func dbStatusAction(ccmd *cobra.Command, _ []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	migrator, err := storage.Migrator(
		ccmd.Context(),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init migrator")

		os.Exit(1)
	}

	m, err := migrator.MigrationsWithStatus(
		ccmd.Context(),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to check migrations")

		os.Exit(1)
	}

	pending := []string{}

	for _, row := range m.Unapplied() {
		pending = append(pending, row.String())
	}

	applied := []string{}

	for _, row := range m.Applied() {
		applied = append(applied, row.String())
	}

	log.Info().
		Strs("pending", pending).
		Strs("applied", applied).
		Msg("Migrations")
}

func dbCreateAction(ccmd *cobra.Command, args []string) {
	storage := prepareStorage(ccmd.Context())
	defer storage.Close()

	migrator, err := storage.Migrator(
		ccmd.Context(),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init migrator")

		os.Exit(1)
	}

	m, err := migrator.CreateGoMigration(
		ccmd.Context(),
		strings.Join(
			args,
			"_",
		),
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to generate migration")

		os.Exit(1)
	}

	log.Info().
		Str("name", m.Name).
		Str("path", m.Path).
		Msg("Finished generating")
}

func prepareStorage(ctx context.Context) *store.Store {
	storage, err := store.NewStore(
		cfg.Database,
		cfg.Scim,
		nil,
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

	if val, err := backoff.Retry(
		ctx,
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
		ctx,
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

	return storage
}
