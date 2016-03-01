package store

import (
	"github.com/Sirupsen/logrus"
	"github.com/solderapp/solder/config"
)

func Load(cfg *config.Config) *Store {
	logrus.Infof("using database driver %s", cfg.Database.Driver)
	logrus.Infof("using database config %s", cfg.Database.Config)

	return New(
		cfg.Database.Driver,
		cfg.Database.Config,
	)
}
