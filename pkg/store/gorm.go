package store

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/gormstore"
	"github.com/alexedwards/scs/v2"
	"github.com/glebarez/sqlite"
	"github.com/go-gormigrate/gormigrate/v2"
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
	session         *gormstore.GORMStore
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

	if admin.Fullname == "" {
		admin.Fullname = "Admin"
	}

	tx := s.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if admin.ID == "" {
		if err := tx.Create(admin).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}
	} else {
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
func (s *GormStore) Open() (bool, error) {
	dialect, err := s.open()

	if err != nil {
		return false, err
	}

	handle, err := gorm.Open(
		dialect,
		&gorm.Config{
			Logger:               NewGormLogger(),
			DisableAutomaticPing: true,
		},
	)

	if err != nil {
		return false, err
	}

	session, err := gormstore.New(handle)

	if err != nil {
		return false, err
	}

	s.handle = handle
	s.session = session

	if err := s.Prepare(); err != nil {
		return false, err
	}

	return true, nil
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
func (s *GormStore) Ping() (bool, error) {
	sqldb, err := s.handle.DB()

	if err != nil {
		return false, err
	}

	if err := sqldb.Ping(); err != nil {
		return false, err
	}

	return true, nil
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

// Session defines a db handler for sessions.
func (s *GormStore) Session() scs.Store {
	return s.session
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
	username, err := config.Value(cfg.Username)

	if err != nil {
		return nil, fmt.Errorf("failed to parse username secret: %w", err)
	}

	password, err := config.Value(cfg.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to parse password secret: %w", err)
	}

	client := &GormStore{
		driver:   cfg.Driver,
		database: cfg.Name,

		username: username,
		password: password,

		meta: url.Values{},
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
