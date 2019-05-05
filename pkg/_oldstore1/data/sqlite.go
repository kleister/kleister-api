// +build cgo

package data

import (
	// Register SQLite driver for GORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	EnableSQLite3 = true
}
