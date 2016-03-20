package model

import (
	"github.com/jinzhu/gorm"
)

// Store is a basic struct to represent the database handle.
type Store struct {
	*gorm.DB
}
