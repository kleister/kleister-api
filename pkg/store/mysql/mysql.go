package mysql

import (
	"net/url"

	"github.com/kleister/kleister-api/pkg/store"
)

type mysql struct {
	dsn *url.URL
}

// Close simply closes the MySQL connection.
func (s *mysql) Close() error {
	return nil
}

// New initializes a new MySQL connection.
func New(dsn *url.URL) (store.Store, error) {
	return &mysql{
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
