package session

import (
	"time"

	"github.com/alexedwards/scs/v2"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Store    scs.Store
	Lifetime time.Duration
	Path     string
	Secure   bool
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithStore provides a function to set the store option.
func WithStore(v scs.Store) Option {
	return func(o *Options) {
		o.Store = v
	}
}

// WithLifetime provides a function to set the lifetime option.
func WithLifetime(v time.Duration) Option {
	return func(o *Options) {
		o.Lifetime = v
	}
}

// WithPath provides a function to set the path option.
func WithPath(v string) Option {
	return func(o *Options) {
		o.Path = v
	}
}

// WithSecure provides a function to set the secure option.
func WithSecure(v bool) Option {
	return func(o *Options) {
		o.Secure = v
	}
}
