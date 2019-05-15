package boltdb

import (
	"net/url"

	"github.com/kleister/kleister-api/pkg/store"
)

type boltdb struct {
	dsn *url.URL
}

// Close simply closes the BoltDB connection.
func (s *boltdb) Close() error {
	return nil
}

// New initializes a new BoltDB connection.
func New(dsn *url.URL) (store.Store, error) {
	return &boltdb{
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
