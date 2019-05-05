package storage

import (
	"github.com/kleister/kleister-api/pkg/model"
)

// Store implements all required database-layer functions.
type Store interface {
	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, error)
}
