package store

import (
	"golang.org/x/net/context"
)

const (
	key = "store"
)

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext gets the store from the context.
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext injects the store into the context.
func ToContext(c Setter, store Store) {
	c.Set(key, store)
}
