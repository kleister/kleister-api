package fabric

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

type metricsService struct {
	service        Service
	requestLatency *prometheus.HistogramVec
	errorsCount    *prometheus.CounterVec
	requestCount   *prometheus.CounterVec
}

// NewMetricsService wraps the Service and provides metrics for its methods.
func NewMetricsService(s Service, m *metrics.Metrics) Service {
	return &metricsService{
		service: s,
		requestLatency: m.RegisterHistogram(
			prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: m.Namespace,
					Subsystem: "fabric_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the fabric service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "fabric_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the fabric service.",
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "fabric_service",
					Name:      "request_count",
					Help:      "Total number of requests to the fabric service.",
				},
				[]string{"method"},
			),
		),
	}
}

// Search implements the Service interface for metrics.
func (s *metricsService) Search(ctx context.Context, search string) ([]*model.Fabric, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("search").Add(1)
		s.requestLatency.WithLabelValues("search").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.Search(ctx, search)

	if err != nil {
		s.errorsCount.WithLabelValues("search").Add(1)
	}

	return records, err
}

// Update implements the Service interface for metrics.
func (s *metricsService) Update(ctx context.Context) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Update(ctx)

	if err != nil {
		s.errorsCount.WithLabelValues("update").Add(1)
	}

	return err
}
