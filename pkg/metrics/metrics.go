package metrics

import (
	"fmt"

	"github.com/kleister/kleister-api/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/rs/zerolog/log"
)

const (
	namespace = "kleister_api"
)

// Metrics simply defines the basic metrics including the registry.
type Metrics struct {
	Namespace string
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

// New simply initializes the metrics handling including the go metrics.
func New() *Metrics {
	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace: namespace,
	}))

	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(version.Collector(namespace))

	return &Metrics{
		Namespace: namespace,
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
