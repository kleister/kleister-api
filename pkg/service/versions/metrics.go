package versions

import (
	"context"
	"errors"
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
					Subsystem: "versions_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the versions service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method", "mod"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "versions_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the versions service.",
				},
				[]string{"method", "mod"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "versions_service",
					Name:      "request_count",
					Help:      "Total number of requests to the versions service.",
				},
				[]string{"method", "mod"},
			),
		),
	}
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, modID string) ([]*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list", modID).Add(1)
		s.requestLatency.WithLabelValues("list", modID).Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx, modID)

	if err != nil {
		s.errorsCount.WithLabelValues("list", modID).Add(1)
	}

	return records, err
}

// Show implements the Service interface for metrics.
func (s *metricsService) Show(ctx context.Context, modID, id string) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show", modID).Add(1)
		s.requestLatency.WithLabelValues("show", modID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Show(ctx, modID, id)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("show", modID).Add(1)
	}

	return record, err
}

// Create implements the Service interface for metrics.
func (s *metricsService) Create(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create", modID).Add(1)
		s.requestLatency.WithLabelValues("create", modID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Create(ctx, modID, version)

	if err != nil {
		s.errorsCount.WithLabelValues("create", modID).Add(1)
	}

	return record, err
}

// Update implements the Service interface for metrics.
func (s *metricsService) Update(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update", modID).Add(1)
		s.requestLatency.WithLabelValues("update", modID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Update(ctx, modID, version)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("update", modID).Add(1)
	}

	return record, err
}

// Delete implements the Service interface for metrics.
func (s *metricsService) Delete(ctx context.Context, modID, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete", modID).Add(1)
		s.requestLatency.WithLabelValues("delete", modID).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Delete(ctx, modID, name)

	if err != nil {
		s.errorsCount.WithLabelValues("delete", modID).Add(1)
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *metricsService) Exists(ctx context.Context, modID, name string) (bool, error) {
	return s.service.Exists(ctx, modID, name)
}
