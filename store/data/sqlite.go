// +build cgo

package data

import (
	// Register SQLite driver for GORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
