package metrics

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Namespace string
	Token     string
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithNamespace provides a function to set the namespace option.
func WithNamespace(v string) Option {
	return func(o *Options) {
		o.Namespace = v
	}
}

// WithToken provides a function to set the token option.
func WithToken(v string) Option {
	return func(o *Options) {
		o.Token = v
	}
}
