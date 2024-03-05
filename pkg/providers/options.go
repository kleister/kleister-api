package providers

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Config string
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithConfig provides a function to set the config option.
func WithConfig(v string) Option {
	return func(o *Options) {
		o.Config = v
	}
}
