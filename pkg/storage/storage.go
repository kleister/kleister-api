package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/dblog"
	"github.com/kleister/kleister-api/pkg/model"

	// Register MySQL driver
	_ "github.com/go-sql-driver/mysql"

	// Register PostgreSQL driver
	_ "github.com/lib/pq"

	// Register MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	// EnableSQLite3 controls the SQLite3 integration.
	EnableSQLite3 bool
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
	engine, err := xorm.NewEngine(driver, config)

	if err != nil {
		level.Error(logger).Log(
			"msg", "failed to initialize connection",
			"err", err,
		)

		return nil, err
	}

	dblogger := dblog.New(logger)
	dblogger.ShowSQL(true)

	engine.SetLogger(dblogger)
	engine.SetMapper(core.GonicMapper{})

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
	engine *xorm.Engine
	logger log.Logger
}

func (db *data) prepare() error {
	if db.engine.DriverName() == "mysql" {
		db.engine.SetMaxIdleConns(0)
	}

	return nil
}

func (db *data) ping() error {
	for i := 0; i < config.Database.Timeout; i++ {
		err := db.engine.Ping()

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
	m := migrate.New(
		db.engine,
		&migrate.Options{
			TableName:    "migration",
			IDColumnName: "id",
		},
		migrations,
	)

	if err := m.Migrate(); err != nil {
		return err
	}

	count, _ := db.engine.Count(
		&model.User{},
	)

	if config.Admin.Create && count == 0 {
		record := &model.User{
			Username: "admin",
			Password: "admin",
			Email:    "admin@example.com",
			Active:   true,
			Admin:    true,
		}

		_, err := db.engine.Insert(
			record,
		)

		if err != nil {
			level.Warn(db.logger).Log(
				"msg", "failed to create admin",
				"err", err,
			)
		} else {
			level.Debug(db.logger).Log(
				"msg", "created admin user",
			)
		}
	}

	if len(config.Admin.Users) > 0 {
		_, err := db.engine.Table(
			&model.User{},
		).In(
			"username",
			config.Admin.Users,
		).Update(
			map[string]interface{}{
				"admin": true,
			},
		)

		if err != nil {
			level.Warn(db.logger).Log(
				"msg", "failed to enforce admins",
				"err", err,
			)
		} else {
			level.Debug(db.logger).Log(
				"msg", "enforced admin users",
				"admins", strings.Join(config.Admin.Users, ","),
			)
		}
	}

	return nil
}
