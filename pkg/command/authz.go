package command

import (
	"context"

	"github.com/kleister/kleister-api/pkg/authz"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// maxTuplesPerWrite chunks tuple writes to stay under the authz server's
// default per-request limit.
const maxTuplesPerWrite = 100

var (
	authzCmd = &cobra.Command{
		Use:   "authz",
		Short: "Authorization maintenance",
		Args:  cobra.NoArgs,
	}

	authzMigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Copy authorization tuples from one backend to another",
		Long: `Copies every relationship tuple from a source authorization backend to a
destination backend, e.g. moving from the embedded sqlite3 datastore to an
external OpenFGA service, or back again. The --from-* and --to-* flags each
describe a full backend, following the same shape as the server's own
--authz-* flags (driver/address/port/username/password/name/options for an
embedded engine, or api/store/token for an external service).

CAUTION: this is a one-shot copy, not a live sync. Any tuple written to the
source after this command finishes reading — e.g. a permission change made
through the running API — will not be copied. Stop the server (or otherwise
pause writes) before running this command, then point the server's own
--authz-* flags at the destination backend and restart it once the copy has
finished. The copy is idempotent, so it's safe to re-run if interrupted.`,
		Run:  authzMigrateAction,
		Args: cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(authzCmd)

	authzCmd.AddCommand(authzMigrateCmd)

	bindAuthzFlags(authzMigrateCmd, "from", "Source")
	bindAuthzFlags(authzMigrateCmd, "to", "Destination")
}

func bindAuthzFlags(cmd *cobra.Command, prefix, label string) {
	cmd.Flags().String(prefix+"-driver", "sqlite3", label+" driver for the authorization engine (sqlite3, mysql, mariadb, postgres, postgresql or external)")
	cmd.Flags().String(prefix+"-address", "", label+" address for the embedded authorization datastore")
	cmd.Flags().String(prefix+"-port", "", label+" port for the embedded authorization datastore")
	cmd.Flags().String(prefix+"-username", "", label+" username for the embedded authorization datastore")
	cmd.Flags().String(prefix+"-password", "", label+" password for the embedded authorization datastore")
	cmd.Flags().String(prefix+"-name", "", label+" name of the authorization database or path for local databases")
	cmd.Flags().StringToString(prefix+"-options", map[string]string{}, label+" options for the embedded authorization datastore connection")
	cmd.Flags().String(prefix+"-api", "", label+" API URL for an external authorization service")
	cmd.Flags().String(prefix+"-store", "", label+" store ID for an external authorization service")
	cmd.Flags().String(prefix+"-token", "", label+" token for an external authorization service")
}

func authzConfigFromFlags(cmd *cobra.Command, prefix string) config.Authz {
	driver, _ := cmd.Flags().GetString(prefix + "-driver")
	address, _ := cmd.Flags().GetString(prefix + "-address")
	port, _ := cmd.Flags().GetString(prefix + "-port")
	username, _ := cmd.Flags().GetString(prefix + "-username")
	password, _ := cmd.Flags().GetString(prefix + "-password")
	name, _ := cmd.Flags().GetString(prefix + "-name")
	options, _ := cmd.Flags().GetStringToString(prefix + "-options")
	api, _ := cmd.Flags().GetString(prefix + "-api")
	store, _ := cmd.Flags().GetString(prefix + "-store")
	token, _ := cmd.Flags().GetString(prefix + "-token")

	return config.Authz{
		Driver:   driver,
		Address:  address,
		Port:     port,
		Username: username,
		Password: password,
		Name:     name,
		Options:  options,
		API:      api,
		Store:    store,
		Token:    token,
	}
}

func authzMigrateAction(ccmd *cobra.Command, _ []string) {
	from, err := authz.New(authzConfigFromFlags(ccmd, "from"))

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to open source authz backend")
	}

	defer func() { _ = from.Close() }()

	to, err := authz.New(authzConfigFromFlags(ccmd, "to"))

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to open destination authz backend")
	}

	defer func() { _ = to.Close() }()

	ctx := ccmd.Context()

	tuples, err := from.ReadAll(ctx)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to read tuples from source backend")
	}

	if err := writeTuplesInChunks(ctx, to, tuples); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to write tuples to destination backend")
	}

	log.Info().
		Int("tuples", len(tuples)).
		Msg("Successfully migrated authz tuples")
}

func writeTuplesInChunks(ctx context.Context, checker authz.Checker, tuples []authz.TupleKey) error {
	for i := 0; i < len(tuples); i += maxTuplesPerWrite {
		end := min(i+maxTuplesPerWrite, len(tuples))

		if err := checker.Write(ctx, tuples[i:end]...); err != nil {
			return err
		}
	}

	return nil
}
