package store

import (
	"fmt"
	"net"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormStore implements the Store interface.
type GormStore struct {
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
	handle          *gorm.DB
}

// Handle returns a database handle.
func (s *GormStore) Handle() *gorm.DB {
	return s.handle
}

// Admin creates an initial admin user within the database.
func (s *GormStore) Admin(username, password, email string) error {
	admin := &model.User{}

	if err := s.handle.Where(
		&model.User{
			Username: username,
		},
	).First(
		admin,
	).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	admin.Username = username
	admin.Password = password
	admin.Email = email
	admin.Active = true
	admin.Admin = true

	tx := s.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if admin.ID == "" {
		if admin.Slug == "" {
			admin.Slug = Slugify(
				tx.Model(&model.User{}),
				admin.Username,
				"",
			)
		}

		admin.ID = uuid.New().String()

		if err := tx.Create(admin).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}
	} else {
		if admin.Slug == "" {
			admin.Slug = Slugify(
				tx.Model(&model.User{}),
				admin.Username,
				admin.ID,
			)
		}

		if err := tx.Save(admin).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}
	}

	return nil
}

// Info returns some basic db informations.
func (s *GormStore) Info() map[string]interface{} {
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

	if s.password != "" {
		result["password"] = s.password
	}

	return result
}

// Prepare is preparing some database behavior.
func (s *GormStore) Prepare() error {
	sqldb, err := s.handle.DB()

	if err != nil {
		return err
	}

	switch s.driver {
	case "mysql", "mariadb":
		sqldb.SetMaxOpenConns(s.maxOpenConns)
		sqldb.SetMaxIdleConns(s.maxIdleConns)
		sqldb.SetConnMaxLifetime(s.connMaxLifetime)
	case "postgres", "postgresql":
		sqldb.SetMaxOpenConns(s.maxOpenConns)
		sqldb.SetMaxIdleConns(s.maxIdleConns)
		sqldb.SetConnMaxLifetime(s.connMaxLifetime)
	}

	return nil
}

// Open simply opens the database connection.
func (s *GormStore) Open() error {
	dialect, err := s.open()

	if err != nil {
		return err
	}

	handle, err := gorm.Open(
		dialect,
		&gorm.Config{
			Logger:               NewGormLogger(),
			DisableAutomaticPing: true,
		},
	)

	if err != nil {
		return err
	}

	s.handle = handle
	return s.Prepare()
}

// Close simply closes the database connection.
func (s *GormStore) Close() error {
	sqldb, err := s.handle.DB()

	if err != nil {
		return err
	}

	return sqldb.Close()
}

// Ping just tests the database connection.
func (s *GormStore) Ping() error {
	sqldb, err := s.handle.DB()

	if err != nil {
		return err
	}

	return sqldb.Ping()
}

// Migrate executes required db migrations.
func (s *GormStore) Migrate() error {
	migrate := gormigrate.New(
		s.handle,
		gormigrate.DefaultOptions,
		Migrations,
	)

	return migrate.Migrate()
}

func (s *GormStore) open() (gorm.Dialector, error) {
	switch s.driver {
	case "sqlite", "sqlite3":
		return sqlite.Open(fmt.Sprintf(
			"%s?%s",
			s.database,
			s.meta.Encode(),
		)), nil
	case "mysql", "mariadb":
		if s.password != "" {
			return mysql.Open(fmt.Sprintf(
				"%s:%s@(%s:%s)/%s?%s",
				s.username,
				s.password,
				s.host,
				s.port,
				s.database,
				s.meta.Encode(),
			)), nil
		}

		return mysql.Open(fmt.Sprintf(
			"%s@(%s:%s)/%s?%s",
			s.username,
			s.host,
			s.port,
			s.database,
			s.meta.Encode(),
		)), nil
	case "postgres", "postgresql":
		dsn := fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s",
			s.host,
			s.port,
			s.database,
			s.username,
		)

		if s.password != "" {
			dsn = fmt.Sprintf(
				"%s password=%s",
				dsn,
				s.password,
			)
		}

		for key, val := range s.meta {
			dsn = fmt.Sprintf("%s %s=%s", dsn, key, strings.Join(val, ""))
		}

		return postgres.Open(
			dsn,
		), nil
	}

	return nil, nil
}

// NewGormStore initializes a new GORM
func NewGormStore(cfg config.Database) (Store, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	client := &GormStore{
		driver:   parsed.Scheme,
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
	case "sqlite", "sqlite3":
		client.driver = "sqlite3"
		client.database = path.Join(parsed.Host, parsed.Path)

		client.meta.Add("_pragma", "journal_mode(WAL)")
		client.meta.Add("_pragma", "busy_timeout(5000)")
		client.meta.Add("_pragma", "foreign_keys(1)")
	case "mysql", "mariadb":
		client.driver = "mysql"
		client.database = strings.TrimPrefix(parsed.Path, "/")

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
			client.meta.Set("charset", "utf8")
		}

		if val := client.meta.Get("parseTime"); val == "" {
			client.meta.Set("parseTime", "True")
		}

		if val := client.meta.Get("loc"); val == "" {
			client.meta.Set("loc", "Local")
		}
	case "postgres", "postgresql":
		client.driver = "postgres"
		client.database = strings.TrimPrefix(parsed.Path, "/")

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

	return client, nil
}

// MustGormStore simply calls NewGormStore and panics on an error.
func MustGormStore(cfg config.Database) Store {
	s, err := NewGormStore(cfg)

	if err != nil {
		panic(err)
	}

	return s
}
