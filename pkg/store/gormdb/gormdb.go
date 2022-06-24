package gormdb

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormdbStore struct {
	driver   string
	username string
	password string
	host     string
	port     string
	database string
	meta     url.Values

	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration

	handle    *gorm.DB
	teams     teams.Store
	users     users.Store
	minecraft minecraft.Store
	forge     forge.Store
	packs     packs.Store
	builds    builds.Store
	mods      mods.Store
	versions  versions.Store
}

func (db *gormdbStore) Teams() teams.Store {
	return db.teams
}

func (db *gormdbStore) Users() users.Store {
	return db.users
}

func (db *gormdbStore) Minecraft() minecraft.Store {
	return db.minecraft
}

func (db *gormdbStore) Forge() forge.Store {
	return db.forge
}

func (db *gormdbStore) Packs() packs.Store {
	return db.packs
}

func (db *gormdbStore) Builds() builds.Store {
	return db.builds
}

func (db *gormdbStore) Mods() mods.Store {
	return db.mods
}

func (db *gormdbStore) Versions() versions.Store {
	return db.versions
}

func (db *gormdbStore) Admin(username, password, email string) error {
	admin := &model.User{}

	if err := db.handle.Where(
		&model.User{
			Username: username,
		},
	).First(
		admin,
	).Error; err != nil && err != gorm.ErrRecordNotFound {
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
func (db *gormdbStore) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = db.driver
	result["database"] = db.database

	if db.host != "" {
		result["host"] = db.host
	}

	if db.port != "" {
		result["port"] = db.port
	}

	if db.username != "" {
		result["username"] = db.username
	}

	if db.password != "" {
		result["password"] = db.password
	}

	for key, value := range db.meta {
		result[key] = strings.Join(value, "&")
	}

	return result
}

// Prepare is preparing some database behavior.
func (db *gormdbStore) Prepare() error {
	sqldb, err := db.handle.DB()

	if err != nil {
		return err
	}

	switch db.driver {
	case "mysql", "mariadb":
		sqldb.SetMaxOpenConns(db.maxOpenConns)
		sqldb.SetMaxIdleConns(db.maxIdleConns)
		sqldb.SetConnMaxLifetime(db.connMaxLifetime)
	case "postgres", "postgresql":
		sqldb.SetMaxOpenConns(db.maxOpenConns)
		sqldb.SetMaxIdleConns(db.maxIdleConns)
		sqldb.SetConnMaxLifetime(db.connMaxLifetime)
	}

	return nil
}

// Open simply opens the database connection.
func (db *gormdbStore) Open() error {
	dialect, err := db.open()

	if err != nil {
		return err
	}

	handle, err := gorm.Open(
		dialect,
		&gorm.Config{
			Logger:               NewLogger(),
			DisableAutomaticPing: true,
		},
	)

	if err != nil {
		return err
	}

	db.handle = handle
	return db.Prepare()
}

// Close simply closes the database connection.
func (db *gormdbStore) Close() error {
	sqldb, err := db.handle.DB()

	if err != nil {
		return err
	}

	return sqldb.Close()
}

// Ping checks the connection to database.
func (db *gormdbStore) Ping() error {
	sqldb, err := db.handle.DB()

	if err != nil {
		return err
	}

	return sqldb.Ping()
}

// Migrate executes required db migrations.
func (db *gormdbStore) Migrate() error {
	migrate := gormigrate.New(
		db.handle,
		gormigrate.DefaultOptions,
		migrations,
	)

	return migrate.Migrate()
}

func (db *gormdbStore) open() (gorm.Dialector, error) {
	switch db.driver {
	case "mysql", "mariadb":
		if db.password != "" {
			return mysql.Open(fmt.Sprintf(
				"%s:%s@(%s:%s)/%s?%s",
				db.username,
				db.password,
				db.host,
				db.port,
				db.database,
				db.meta.Encode(),
			)), nil
		}

		return mysql.Open(fmt.Sprintf(
			"%s@(%s:%s)/%s?%s",
			db.username,
			db.host,
			db.port,
			db.database,
			db.meta.Encode(),
		)), nil
	case "postgres", "postgresql":
		dsn := fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s",
			db.host,
			db.port,
			db.database,
			db.username,
		)

		if db.password != "" {
			dsn = fmt.Sprintf(
				"%s password=%s",
				dsn,
				db.password,
			)
		}

		for key, val := range db.meta {
			dsn = fmt.Sprintf("%s %s=%s", dsn, key, strings.Join(val, ""))
		}

		return postgres.Open(
			dsn,
		), nil
	}

	return nil, nil
}

// New initializes a new MySQL connection.
func New(cfg config.Database) (store.Store, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	client := &gormdbStore{
		driver:   parsed.Scheme,
		database: strings.TrimPrefix(parsed.Path, "/"),
		username: parsed.User.Username(),
		meta:     parsed.Query(),
	}

	if password, ok := parsed.User.Password(); ok {
		client.password = password
	}

	if client.meta.Has("maxOpenConns") {
		val, err := strconv.Atoi(
			client.meta.Get("maxOpenConns"),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse maxOpenConns: %w", err)
		}

		client.maxOpenConns = val
		client.meta.Del("maxOpenConns")
	} else {
		client.maxOpenConns = 25
	}

	if client.meta.Has("maxIdleConns") {
		val, err := strconv.Atoi(
			client.meta.Get("maxIdleConns"),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse maxIdleConns: %w", err)
		}

		client.maxIdleConns = val
		client.meta.Del("maxIdleConns")
	} else {
		client.maxIdleConns = 25
	}

	if client.meta.Has("connMaxLifetime") {
		val, err := time.ParseDuration(
			client.meta.Get("connMaxLifetime"),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse connMaxLifetime: %w", err)
		}

		client.connMaxLifetime = val
		client.meta.Del("connMaxLifetime")
	} else {
		client.connMaxLifetime = 5 * time.Minute
	}

	switch client.driver {
	case "mysql", "mariadb":
		host, port, err := net.SplitHostPort(parsed.Host)

		if err != nil && strings.Contains(err.Error(), "missing port in address") {
			client.host = parsed.Host
			client.port = "3306"
		} else if err != nil {
			return nil, err
		} else {
			client.host = host
			client.port = port
		}

		if val := client.meta.Get("charset"); val == "" {
			client.meta.Add("charset", "utf8")
		}

		if val := client.meta.Get("parseTime"); val == "" {
			client.meta.Set("parseTime", "True")
		}

		if val := client.meta.Get("loc"); val == "" {
			client.meta.Set("loc", "Local")
		}
	case "postgres", "postgresql":
		host, port, err := net.SplitHostPort(parsed.Host)

		if err != nil && strings.Contains(err.Error(), "missing port in address") {
			client.host = parsed.Host
			client.port = "5432"
		} else if err != nil {
			return nil, err
		} else {
			client.host = host
			client.port = port
		}

		if val := client.meta.Get("sslmode"); val == "" {
			client.meta.Set("sslmode", "disable")
		}
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

	client.packs = &Packs{
		client: client,
	}

	client.builds = &Builds{
		client: client,
	}

	client.mods = &Mods{
		client: client,
	}

	client.versions = &Versions{
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
