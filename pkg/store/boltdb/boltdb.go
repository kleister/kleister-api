package boltdb

import (
	"context"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	bolt "go.etcd.io/bbolt"
)

type botldbStore struct {
	path    string
	perms   os.FileMode
	timeout time.Duration

	handle    *storm.DB
	teams     teams.Store
	users     users.Store
	minecraft minecraft.Store
	forge     forge.Store
}

func (db *botldbStore) Teams() teams.Store {
	return db.teams
}

func (db *botldbStore) Users() users.Store {
	return db.users
}

func (db *botldbStore) Minecraft() minecraft.Store {
	return db.minecraft
}

func (db *botldbStore) Forge() forge.Store {
	return db.forge
}

func (db *botldbStore) Admin(username, password, email string) error {
	admin := &model.User{}

	if err := db.handle.Select(
		q.Eq("Username", username),
	).First(admin); err != nil && err != storm.ErrNotFound {
		return err
	}

	admin.Username = username
	admin.Password = password
	admin.Email = email
	admin.Active = true
	admin.Admin = true

	if admin.ID == "" {
		if _, err := db.users.Create(
			context.Background(),
			admin,
		); err != nil {
			return err
		}
	} else {
		if _, err := db.users.Update(
			context.Background(),
			admin,
		); err != nil {
			return err
		}
	}

	return nil
}

// Info returns some basic db informations.
func (db *botldbStore) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "boltdb"
	result["path"] = db.path
	result["perms"] = db.perms.String()
	result["timeout"] = db.timeout.String()

	return result
}

// Prepare is preparing some database behavior.
func (db *botldbStore) Prepare() error {
	return nil
}

// Close simply closes the BoltDB connection.
func (db *botldbStore) Open() error {
	handle, err := storm.Open(
		db.path,
		storm.BoltOptions(
			db.perms,
			&bolt.Options{
				Timeout: db.timeout,
			},
		),
	)

	if err != nil {
		log.Error().
			Err(err).
			Msg("")

		return err
	}

	db.handle = handle
	return db.Prepare()
}

// Close simply closes the BoltDB connection.
func (db *botldbStore) Close() error {
	return db.handle.Close()
}

// Close simply closes the BoltDB connection.
func (db *botldbStore) Ping() error {
	return nil
}

// Migrate executes required db migrations.
func (db *botldbStore) Migrate() error {
	return nil
}

// New initializes a new BoltDB connection.
func New(cfg config.Database) (store.Store, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	client := &botldbStore{
		path: path.Join(
			parsed.Host,
			parsed.EscapedPath(),
		),
	}

	if val := parsed.Query().Get("perms"); val != "" {
		res, err := strconv.ParseUint(val, 8, 32)

		if err != nil {
			client.perms = os.FileMode(0600)
		} else {
			client.perms = os.FileMode(res)
		}
	} else {
		client.perms = os.FileMode(0600)
	}

	if val := parsed.Query().Get("timeout"); val != "" {
		res, err := time.ParseDuration(val)

		if err != nil {
			client.timeout = 1 * time.Second
		} else {
			client.timeout = res
		}
	} else {
		client.timeout = 1 * time.Second
	}

	client.teams = &Teams{
		client: client,
	}

	client.users = &Users{
		client: client,
	}

	client.minecraft = &Minecraft{
		client: client,
	}

	client.forge = &Forge{
		client: client,
	}

	return client, nil
}

// Must simply calls New and panics on an error.
func Must(cfg config.Database) store.Store {
	db, err := New(cfg)

	if err != nil {
		panic(err)
	}

	return db
}
