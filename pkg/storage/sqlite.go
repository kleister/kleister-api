// +build cgo

package storage

import (
	// Register SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	EnableSQLite3 = true
}
