package scim

import (
	"github.com/kleister/kleister-api/pkg/config"
	"gorm.io/gorm"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Root   string
	Config config.Scim
	Store  *gorm.DB
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithRoot provides a function to set the root option.
func WithRoot(v string) Option {
	return func(o *Options) {
		o.Root = v
	}
}

// WithConfig provides a function to set the config option.
func WithConfig(v config.Scim) Option {
	return func(o *Options) {
		o.Config = v
	}
}

// WithStore provides a function to set the store option.
func WithStore(v *gorm.DB) Option {
	return func(o *Options) {
		o.Store = v
	}
}
