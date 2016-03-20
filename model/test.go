package model

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// Test creates an in-memory database connection.
func Test() *Store {
	var (
		driver = "sqlite3"
		config = ":memory:"
	)

	if os.Getenv("DATABASE_DRIVER") != "" && os.Getenv("DATABASE_DRIVER") != "" {
		driver = os.Getenv("DATABASE_DRIVER")
		config = os.Getenv("DATABASE_CONFIG")
	}

	db, err := gorm.Open(driver, config)

	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	return &Store{
		setupDatabase(driver, db),
	}
}
