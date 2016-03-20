// +build sqlite3
package model

import (
	// Register SQLite driver for GORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
