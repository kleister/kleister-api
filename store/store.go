package store

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/o1egl/gormrus"
	"github.com/solderapp/solder/model"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Store struct {
	*gorm.DB
}

func New(driver string, config string) *Store {
	db, err := gorm.Open(driver, config)

	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	return &Store{
		setupDatabase(driver, &db),
	}
}

func From(driver string, handle *sql.DB) *Store {
	db, err := gorm.Open(driver, handle)

	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	return &Store{
		setupDatabase(driver, &db),
	}
}

func setupDatabase(driver string, db *gorm.DB) *gorm.DB {
	db.LogMode(true)
	db.SetLogger(gormrus.New())

	if err := prepareDatabase(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database preparation failed")
	}

	if err := pingDatabase(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database ping attempts failed")
	}

	if err := migrateDatabase(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database auto migrate failed")
	}

	return db
}

func prepareDatabase(driver string, db *gorm.DB) error {
	if driver == "mysql" {
		db.DB().SetMaxIdleConns(0)
	}

	return nil
}

func pingDatabase(driver string, db *gorm.DB) error {
	for i := 0; i < 30; i++ {
		err := db.DB().Ping()

		if err == nil {
			return nil
		}

		logrus.Infof("database ping failed, retry in 1s")
		time.Sleep(time.Second)
	}

	return nil
}

func migrateDatabase(driver string, db *gorm.DB) error {
	db.AutoMigrate(
		&model.Attachment{},
		&model.Build{},
		&model.Client{},
		&model.Forge{},
		&model.Key{},
		&model.Minecraft{},
		&model.Mod{},
		&model.Pack{},
		&model.Permission{},
		&model.User{},
		&model.Version{},
	)

	return nil
}
