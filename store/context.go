package store

import (
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

const (
	currentKey = "current"
	storeKey   = "store"
)

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// Current gets the user from the context.
func Current(c context.Context) *model.User {
	return c.Value(currentKey).(*model.User)
}

// FromContext gets the store from the context.
func FromContext(c context.Context) Store {
	return c.Value(storeKey).(Store)
}

// ToContext injects the store into the context.
func ToContext(c Setter, store Store) {
	c.Set(storeKey, store)
}
