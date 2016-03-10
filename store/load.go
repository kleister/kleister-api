package store

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"gopkg.in/solderapp/solder-api.v0/config"
)

// Load initializes the database connection.
func Load(cfg *config.Config) *Store {
	driver := cfg.Database.Driver
	connect := ""

	switch driver {
	case "mysql":
		connect = fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Name,
		)
	case "postgres":
		connect = fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Name,
		)
	case "sqlite":
		connect = cfg.Database.Name
	default:
		logrus.Fatal("Unknown database driver selected")
	}

	logrus.Infof("using database driver %s", driver)
	logrus.Infof("using database config %s", connect)

	return New(
		driver,
		connect,
	)
}
