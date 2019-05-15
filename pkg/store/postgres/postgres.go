package postgres

import (
	"net/url"

	"github.com/kleister/kleister-api/pkg/store"
)

type postgres struct {
	dsn *url.URL
}

// Close simply closes the PostgreSQL connection.
func (s *postgres) Close() error {
	return nil
}

// New initializes a new PostgreSQL connection.
func New(dsn *url.URL) (store.Store, error) {
	return &postgres{
		dsn: dsn,
	}, nil
}

// Must simply calls New and panics on an error.
func Must(dsn *url.URL) store.Store {
	db, err := New(dsn)

	if err != nil {
		panic(err)
	}

	return db
}
