package storage

import (
	"errors"
	"fmt"
	// "strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/dblog"
	// "github.com/kleister/kleister-api/pkg/model"
	"github.com/qor/validations"
	"gopkg.in/gormigrate.v1"

	// Register MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Register PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// Register MSSQL driver
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	// EnableSQLite3 just indicates if SQLite3 is enabled.
	EnableSQLite3 = false

	// ErrRecordNotFound defines an error for non-existant records.
	ErrRecordNotFound = errors.New("record not found")
)

// Load prepares a database connection.
func Load(logger log.Logger) (Store, error) {
	logger = log.WithPrefix(logger, "storage", "db")

	driver := config.Database.Driver
	connect := ""

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
		if !EnableSQLite3 {
			level.Error(logger).Log(
				"msg", "sqlite3 driver is not available",
			)

			return nil, fmt.Errorf("sqlite3 driver is not available")
		}

		connect = config.Database.Name
	}

	return New(
		logger,
		driver,
		connect,
	)
}

// New initializes a database connection.
func New(logger log.Logger, driver string, config string) (Store, error) {
	engine, err := gorm.Open(driver, config)

	if err != nil {
		level.Error(logger).Log(
			"msg", "failed to initialize connection",
			"err", err,
		)

		return nil, err
	}

	engine.SetLogger(dblog.New(logger))
	validations.RegisterCallbacks(engine)

	handle := &data{
		engine: engine,
		logger: logger,
	}

	if err := handle.prepare(); err != nil {
		return nil, err
	}

	if err := handle.ping(); err != nil {
		return nil, err
	}

	if err := handle.migrate(); err != nil {
		return nil, err
	}

	return handle, nil
}

type data struct {
	engine *gorm.DB
	logger log.Logger
}

func (db *data) prepare() error {
	if db.engine.Dialect().GetName() == "mysql" {
		db.engine.DB().SetMaxIdleConns(0)
	}

	return nil
}

func (db *data) ping() error {
	for i := 0; i < config.Database.Timeout; i++ {
		err := db.engine.DB().Ping()

		if err == nil {
			return nil
		}

		level.Info(db.logger).Log(
			"msg", "database ping failed, retry in 1s",
		)

		time.Sleep(time.Second)
	}

	return fmt.Errorf("database ping attempts failed")
}

func (db *data) migrate() error {
	m := gormigrate.New(
		db.engine,
		&gormigrate.Options{
			TableName:    "migrations",
			IDColumnName: "id",
		},
		migrations,
	)

	if err := m.Migrate(); err != nil {
		return err
	}

	// if config.Admin.Create && db.engine.First(&model.User{}).RecordNotFound() {
	// 	record := &model.User{
	// 		Username: "admin",
	// 		Password: "admin",
	// 		Email:    "admin@example.com",
	// 		Active:   true,
	// 		Admin:    true,
	// 	}

	// 	err := db.engine.Create(
	// 		record,
	// 	).Error

	// 	if err != nil {
	// 		level.Warn(db.logger).Log(
	// 			"msg", "failed to create admin",
	// 			"err", err,
	// 		)
	// 	} else {
	// 		level.Debug(db.logger).Log(
	// 			"msg", "created admin user",
	// 		)
	// 	}
	// }

	// if len(config.Admin.Users) > 0 {
	// 	err := db.engine.Model(
	// 		&model.User{},
	// 	).Where(
	// 		"username IN (?)",
	// 		config.Admin.Users,
	// 	).UpdateColumn(
	// 		"admin",
	// 		true,
	// 	).Error

	// 	if err != nil {
	// 		level.Warn(db.logger).Log(
	// 			"msg", "failed to enforce admins",
	// 			"err", err,
	// 		)
	// 	} else {
	// 		level.Debug(db.logger).Log(
	// 			"msg", "enforced admin users",
	// 			"admins", strings.Join(config.Admin.Users, ","),
	// 		)
	// 	}
	// }

	return nil
}
