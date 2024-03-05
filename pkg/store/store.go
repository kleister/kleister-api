package store

import (
	"fmt"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

var (
	// ErrUnknownDriver defines a named error for unknown store drivers.
	ErrUnknownDriver = fmt.Errorf("unknown database driver")
)

// Store provides the interface for the store implementations.
type Store interface {
	Info() map[string]interface{}
	Prepare() error
	Open() error
	Close() error
	Ping() error
	Migrate() error
	Handle() *gorm.DB
	Session() scs.Store
}
