package metrics

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kleister/kleister-api/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var (
	// ErrInvalidToken is returned when the request token is invalid.
	ErrInvalidToken = errors.New("invalid or missing token")
)

// Metrics simply defines the basic metrics including the registry.
type Metrics struct {
	Namespace string
	Token     string
	Registry  *prometheus.Registry
}

// RegisterHistogram is used to transparently register a new histogram.
func (m *Metrics) RegisterHistogram(histogram *prometheus.HistogramVec) *prometheus.HistogramVec {
	m.Registry.MustRegister(histogram)
	return histogram
}

// RegisterCounter is used to transparently register a new counter.
func (m *Metrics) RegisterCounter(counter *prometheus.CounterVec) *prometheus.CounterVec {
	m.Registry.MustRegister(counter)
	return counter
}

// Handler initializes the prometheus middleware.
func (m *Metrics) Handler() http.HandlerFunc {
	h := promhttp.HandlerFor(
		m.Registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
			ErrorLog:          Logger{},
		},
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if m.Token == "" {
			h.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		if header != "Bearer "+m.Token {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}
}

// New simply initializes the metrics handling including the go metrics.
func New(opts ...Option) *Metrics {
	options := newOptions(opts...)
	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace: options.Namespace,
	}))

	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(version.Collector(options.Namespace))

	return &Metrics{
		Namespace: options.Namespace,
		Token:     options.Token,
		Registry:  registry,
	}
}

// Logger implements the required logging interface for the http handler.
type Logger struct{}

// Println is used by the prometheus http handler to serve the registry.
func (Logger) Println(v ...interface{}) {
	log.Error().
		Msg(fmt.Sprintln(v...))
}
