// +build cgo

package storage

import (
	// Register SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	EnableSQLite3 = true
}
