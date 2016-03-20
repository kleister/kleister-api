// +build mssql
package model

import (
	// Register MSSQL driver for GORM
	_ "github.com/jinzhu/gorm/dialects/mssql"
)
