package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/migrations"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bunzerolog"
	"github.com/uptrace/bun/migrate"

	// generally import the mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

var (
	// ErrUnknownDriver defines a named error for unknown store drivers.
	ErrUnknownDriver = fmt.Errorf("unknown database driver")
)

// Store provides the general database abstraction layer.
type Store struct {
	scim   config.Scim
	upload upload.Upload

	driver          string
	username        string
	password        string
	host            string
	port            string
	database        string
	meta            url.Values
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
	handle          *bun.DB
	principal       *model.User

	Minecraft *Minecraft
	Forge     *Forge
	Neoforge  *Neoforge
	Quilt     *Quilt
	Fabric    *Fabric
	Auth      *Auth
	Mods      *Mods
	Versions  *Versions
	Packs     *Packs
	Builds    *Builds
	Groups    *Groups
	Users     *Users
}

// Handle returns a database handle.
func (s *Store) Handle() *bun.DB {
	return s.handle
}

// WithPrincipal integrates the current user.
func (s *Store) WithPrincipal(principal *model.User) *Store {
	s.principal = principal
	return s
}

// SearchQuery builds a query for search terms.
func (s *Store) SearchQuery(q *bun.SelectQuery, _ string) *bun.SelectQuery {
	// opts := queryparser.Options{
	// 	CutFn: searchCut,
	// 	Allowed: []string{
	// 		"slug",
	// 		"name",
	// 	},
	// }

	// parser := queryparser.New(
	// 	term,
	// 	opts,
	// ).Parse()

	// for _, name := range opts.Allowed {
	// 	if parser.Has(name) {

	// 		q = q.Where(
	// 			fmt.Sprintf(
	// 				"%s LIKE ?",
	// 				name,
	// 			),
	// 			strings.ReplaceAll(
	// 				parser.GetOne(name),
	// 				"*",
	// 				"%",
	// 			),
	// 		)
	// 	}
	// }

	return q
}

// Admin creates an initial admin user within the database.
func (s *Store) Admin(username, password, email string) error {
	ctx := context.Background()
	admin := &model.User{}

	if err := s.handle.NewSelect().
		Model(admin).
		Where("username = ?", username).
		Scan(ctx); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to check admin: %w", err)
	}

	admin.Username = username
	admin.Password = password
	admin.Email = email
	admin.Active = true
	admin.Admin = true

	if admin.Fullname == "" {
		admin.Fullname = "Admin"
	}

	if admin.ID == "" {
		_, err := s.handle.NewInsert().
			Model(admin).
			Exec(ctx)

		if err != nil {
			return fmt.Errorf("failed to create admin: %w", err)
		}
	} else {
		_, err := s.handle.NewUpdate().
			Model(admin).
			Where("username = ?", username).
			Exec(ctx)

		if err != nil {
			return fmt.Errorf("failed to update admin: %w", err)
		}
	}

	return nil
}

// Info returns some basic db informations.
func (s *Store) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = s.driver
	result["database"] = s.database

	if s.host != "" {
		result["host"] = s.host
	}

	if s.port != "" {
		result["port"] = s.port
	}

	if s.username != "" {
		result["username"] = s.username
	}

	return result
}

// Prepare is preparing some database behavior.
func (s *Store) Prepare() error {
	switch s.driver {
	case "mysql", "mariadb":
		s.handle.SetMaxOpenConns(s.maxOpenConns)
		s.handle.SetMaxIdleConns(s.maxIdleConns)
		s.handle.SetConnMaxLifetime(s.connMaxLifetime)
	case "postgres", "postgresql":
		s.handle.SetMaxOpenConns(s.maxOpenConns)
		s.handle.SetMaxIdleConns(s.maxIdleConns)
		s.handle.SetConnMaxLifetime(s.connMaxLifetime)
	case "sqlite", "sqlite3":
		if strings.Contains(s.database, ":memory:") {
			s.handle.SetMaxIdleConns(1000)
			s.handle.SetConnMaxLifetime(0)
		}
	}

	return nil
}

// Open simply opens the database connection.
func (s *Store) Open() (bool, error) {
	err := s.open()

	if err != nil {
		return false, err
	}

	s.handle.AddQueryHook(
		bunzerolog.NewQueryHook(
			bunzerolog.WithQueryLogLevel(zerolog.TraceLevel),
			bunzerolog.WithSlowQueryLogLevel(zerolog.WarnLevel),
			bunzerolog.WithErrorQueryLogLevel(zerolog.ErrorLevel),
			bunzerolog.WithSlowQueryThreshold(3*time.Second),
		),
	)

	if err = s.Prepare(); err != nil {
		return false, err
	}

	return true, nil
}

// Close simply closes the database connection.
func (s *Store) Close() (bool, error) {
	if s.handle != nil {
		if err := s.handle.Close(); err != nil {
			return false, err
		}
	}

	return true, nil
}

// Ping just tests the database connection.
func (s *Store) Ping() (bool, error) {
	if s.handle != nil {
		if err := s.handle.Ping(); err != nil {
			return false, err
		}
	}

	return true, nil
}

// Migrator provides the migration client.
func (s *Store) Migrator(ctx context.Context) (*migrate.Migrator, error) {
	migrator := migrate.NewMigrator(
		s.handle,
		migrations.Migrations,
	)

	if err := migrator.Init(ctx); err != nil {
		return nil, err
	}

	return migrator, nil
}

// Migrate handles a database migration.
func (s *Store) Migrate(ctx context.Context) (*migrate.MigrationGroup, error) {
	migrator, err := s.Migrator(ctx)

	if err != nil {
		return nil, err
	}

	if err := migrator.Lock(ctx); err != nil {
		return nil, err
	}

	defer func() {
		_ = migrator.Unlock(ctx)
	}()

	return migrator.Migrate(ctx)
}

// Rollback handles a database rollback.
func (s *Store) Rollback(ctx context.Context) (*migrate.MigrationGroup, error) {
	migrator, err := s.Migrator(ctx)

	if err != nil {
		return nil, err
	}

	if err := migrator.Lock(ctx); err != nil {
		return nil, err
	}

	defer func() {
		_ = migrator.Unlock(ctx)
	}()

	return migrator.Rollback(ctx)
}

func (s *Store) open() error {
	switch s.driver {
	case "sqlite", "sqlite3":
		sqldb, err := sql.Open(
			sqliteshim.ShimName,
			fmt.Sprintf(
				"%s?%s",
				s.database,
				s.meta.Encode(),
			),
		)

		if err != nil {
			return err
		}

		s.handle = bun.NewDB(
			sqldb,
			sqlitedialect.New(),
		)

		return nil
	case "mysql", "mariadb":
		var (
			sqldb *sql.DB
			err   error
		)

		if s.password != "" {
			sqldb, err = sql.Open(
				"mysql",
				fmt.Sprintf(
					"%s:%s@(%s:%s)/%s?%s",
					s.username,
					s.password,
					s.host,
					s.port,
					s.database,
					s.meta.Encode(),
				),
			)
		} else {
			sqldb, err = sql.Open(
				"mysql",
				fmt.Sprintf(
					"%s@(%s:%s)/%s?%s",
					s.username,
					s.host,
					s.port,
					s.database,
					s.meta.Encode(),
				),
			)
		}

		if err != nil {
			return err
		}

		s.handle = bun.NewDB(
			sqldb,
			mysqldialect.New(),
		)

		return nil
	case "postgres", "postgresql":
		var (
			sqldb *sql.DB
		)

		if s.password != "" {
			sqldb = sql.OpenDB(
				pgdriver.NewConnector(
					pgdriver.WithDSN(
						fmt.Sprintf(
							"%s:%s@(%s:%s)/%s?%s",
							s.username,
							s.password,
							s.host,
							s.port,
							s.database,
							s.meta.Encode(),
						),
					),
				),
			)
		} else {
			sqldb = sql.OpenDB(
				pgdriver.NewConnector(
					pgdriver.WithDSN(
						fmt.Sprintf(
							"%s@(%s:%s)/%s?%s",
							s.username,
							s.host,
							s.port,
							s.database,
							s.meta.Encode(),
						),
					),
				),
			)
		}

		s.handle = bun.NewDB(
			sqldb,
			pgdialect.New(),
		)

		return nil
	}

	return ErrUnknownDriver
}

// NewStore initializes a new Bun
func NewStore(cfg config.Database, scim config.Scim, uploads upload.Upload) (*Store, error) {
	username, err := config.Value(cfg.Username)

	if err != nil {
		return nil, fmt.Errorf("failed to parse username secret: %w", err)
	}

	password, err := config.Value(cfg.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to parse password secret: %w", err)
	}

	client := &Store{
		scim:     scim,
		upload:   uploads,
		driver:   cfg.Driver,
		database: cfg.Name,
		username: username,
		password: password,
		meta:     url.Values{},
	}

	if val, ok := cfg.Options["maxOpenConns"]; ok {
		cur, err := strconv.Atoi(
			val,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse maxOpenConns: %w", err)
		}

		client.maxOpenConns = cur
	} else {
		client.maxOpenConns = 25
	}

	if val, ok := cfg.Options["maxIdleConns"]; ok {
		cur, err := strconv.Atoi(
			val,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse maxIdleConns: %w", err)
		}

		client.maxIdleConns = cur
	} else {
		client.maxIdleConns = 25
	}

	if val, ok := cfg.Options["connMaxLifetime"]; ok {
		cur, err := time.ParseDuration(
			val,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to parse connMaxLifetime: %w", err)
		}

		client.connMaxLifetime = cur
	} else {
		client.connMaxLifetime = 5 * time.Minute
	}

	switch client.driver {
	case "sqlite", "sqlite3":
		client.driver = "sqlite3"

		client.meta.Add("_pragma", "journal_mode(WAL)")
		client.meta.Add("_pragma", "busy_timeout(5000)")
		client.meta.Add("_pragma", "foreign_keys(1)")
	case "mysql", "mariadb":
		client.driver = "mysql"

		client.host = cfg.Address
		client.port = "3306"

		if cfg.Port != "" {
			client.port = cfg.Port
		}

		if val, ok := cfg.Options["charset"]; ok {
			client.meta.Set("charset", val)
		} else {
			client.meta.Set("charset", "utf8")
		}

		if val, ok := cfg.Options["parseTime"]; ok {
			client.meta.Set("parseTime", val)
		} else {
			client.meta.Set("parseTime", "True")
		}

		if val, ok := cfg.Options["loc"]; ok {
			client.meta.Set("loc", val)
		} else {
			client.meta.Set("loc", "Local")
		}

	case "postgres", "postgresql":
		client.driver = "postgres"

		client.host = cfg.Address
		client.port = "5432"

		if cfg.Port != "" {
			client.port = cfg.Port
		}

		if val, ok := cfg.Options["sslmode"]; ok {
			client.meta.Set("sslmode", val)
		} else {
			client.meta.Set("sslmode", "disable")
		}
	}

	client.Minecraft = &Minecraft{
		client: client,
	}

	client.Forge = &Forge{
		client: client,
	}

	client.Neoforge = &Neoforge{
		client: client,
	}

	client.Quilt = &Quilt{
		client: client,
	}

	client.Fabric = &Fabric{
		client: client,
	}

	client.Auth = &Auth{
		client: client,
	}

	client.Mods = &Mods{
		client: client,
	}

	client.Versions = &Versions{
		client: client,
	}

	client.Packs = &Packs{
		client: client,
	}

	client.Builds = &Builds{
		client: client,
	}

	client.Groups = &Groups{
		client: client,
	}

	client.Users = &Users{
		client: client,
	}

	return client, nil
}

// MustStore simply calls NewStore and panics on an error.
func MustStore(cfg config.Database, scim config.Scim, uploads upload.Upload) *Store {
	s, err := NewStore(cfg, scim, uploads)

	if err != nil {
		panic(err)
	}

	return s
}
