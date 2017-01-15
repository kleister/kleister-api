package data

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
	"github.com/o1egl/gormrus"
	"github.com/qor/validations"
	"gopkg.in/gormigrate.v1"

	// Register MySQL driver for GORM
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Register Postgres driver for GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// Register MSSQL driver for GORM
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	// EnableSQLite3 controls the SQLite3 integration.
	EnableSQLite3 bool
)

// Store is a basic struct to represent the database handle.
type data struct {
	*gorm.DB
}

// Test creates an in-memory database connection.
func Test() store.Store {
	var (
		driver = "sqlite3"
		config = ":memory:"
	)

	if os.Getenv("DATABASE_DRIVER") != "" && os.Getenv("DATABASE_CONFIG") != "" {
		driver = os.Getenv("DATABASE_DRIVER")
		config = os.Getenv("DATABASE_CONFIG")
	}

	db, err := gorm.Open(driver, config)

	if err != nil {
		logrus.Fatalf("Database connection failed. %s", err)
	}

	return &data{
		setupDatabase(driver, db),
	}
}

// New initializes a new database connection.
func New(driver string, config string) store.Store {
	db, err := gorm.Open(driver, config)

	if err != nil {
		logrus.Fatalf("Database connection failed. %s", err)
	}

	return &data{
		setupDatabase(driver, db),
	}
}

// From takes an existing database connection.
func From(driver string, handle *sql.DB) store.Store {
	db, err := gorm.Open(driver, handle)

	if err != nil {
		logrus.Fatalf("Database connection failed. %s", err)
	}

	return &data{
		setupDatabase(driver, db),
	}
}

// Load initializes the database connection.
func Load() store.Store {
	driver := config.Database.Driver
	connect := ""

	if invalidDriver(driver) {
		logrus.Fatalf("Unknown database driver %s selected", driver)
	}

	switch driver {
	case "mysql":
		connect = fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Name,
		)
	case "postgres":
		connect = fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Name,
		)
	case "mssql":
		connect = fmt.Sprintf(
			"mssql://%s:%s@%s/%s",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Name,
		)
	case "sqlite3":
		connect = config.Database.Name
	}

	logrus.Infof("Using database driver %s", driver)
	logrus.Infof("Using database config %s", connect)

	return New(
		driver,
		connect,
	)
}

func invalidDriver(driver string) bool {
	logrus.Debugf("Checking %s driver for validity", driver)

	if EnableSQLite3 && driver == "sqlite3" {
		logrus.Debugf("Detected successfully SQLite driver")
		return false
	}

	if driver == "mysql" {
		logrus.Debugf("Detected successfully MySQL driver")
		return false
	}

	if driver == "postgres" {
		logrus.Debugf("Detected successfully Postgres driver")
		return false
	}

	if driver == "mssql" {
		logrus.Debugf("Detected successfully MSSQL driver")
		return false
	}

	return true
}

func setupDatabase(driver string, db *gorm.DB) *gorm.DB {
	if config.Debug {
		db.LogMode(true)
		db.SetLogger(gormrus.New())
	}

	if err := prepareDatabase(driver, db); err != nil {
		logrus.Fatalln(err)
	}

	if err := pingDatabase(driver, db); err != nil {
		logrus.Fatalln(err)
	}

	if err := migrateDatabase(driver, db); err != nil {
		logrus.Fatalln(err)
	}

	return db
}

func prepareDatabase(driver string, db *gorm.DB) error {
	if driver == "mysql" {
		db.DB().SetMaxIdleConns(0)
	}

	validations.RegisterCallbacks(
		db,
	)

	return nil
}

func pingDatabase(driver string, db *gorm.DB) error {
	for i := 0; i < 30; i++ {
		err := db.DB().Ping()

		if err == nil {
			return nil
		}

		logrus.Infof("Database ping failed, retry in 1s")
		time.Sleep(time.Second)
	}

	return fmt.Errorf("Database ping attempts failed")
}

func migrateDatabase(driver string, db *gorm.DB) error {
	migrate := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		migrations,
	)

	if err := migrate.Migrate(); err != nil {
		return fmt.Errorf(
			"Failed to migrate database. %s",
			err,
		)
	}

	if config.Admin.Create && db.First(&model.User{}).RecordNotFound() {
		record := &model.User{
			Username: "admin",
			Password: "admin",
			Email:    "admin@example.com",
			Active:   true,
			Admin:    true,
		}

		err := db.Create(
			record,
		).Error

		if err != nil {
			return fmt.Errorf(
				"Failed to create initial user. %s",
				err.Error(),
			)
		}
	}

	logrus.Debugf("%v", config.Admin.Users)

	if len(config.Admin.Users) > 0 {
		logrus.Infof(
			"Enforcing admin users: %s",
			strings.Join(config.Admin.Users, ", "),
		)

		err := db.Model(
			&model.User{},
		).Where(
			"username IN (?)",
			config.Admin.Users,
		).UpdateColumn(
			"admin",
			true,
		).Error

		if err != nil {
			return fmt.Errorf(
				"Failed to enforce admin users. %s",
				err.Error(),
			)
		}
	}

	return nil
}
