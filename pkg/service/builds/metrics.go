package builds

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
					Subsystem: "builds_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the builds service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method", "pack"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "builds_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the builds service.",
				},
				[]string{"method", "pack"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "builds_service",
					Name:      "request_count",
					Help:      "Total number of requests to the builds service.",
				},
				[]string{"method", "pack"},
			),
		),
	}
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, packID string) ([]*model.Build, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list", packID).Add(1)
		s.requestLatency.WithLabelValues("list", packID).Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx, packID)

	if err != nil {
		s.errorsCount.WithLabelValues("list", packID).Add(1)
	}

	return records, err
}

// Show implements the Service interface for metrics.
func (s *metricsService) Show(ctx context.Context, packID, id string) (*model.Build, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show", packID).Add(1)
		s.requestLatency.WithLabelValues("show", packID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Show(ctx, packID, id)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("show", packID).Add(1)
	}

	return record, err
}

// Create implements the Service interface for metrics.
func (s *metricsService) Create(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create", packID).Add(1)
		s.requestLatency.WithLabelValues("create", packID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Create(ctx, packID, build)

	if err != nil {
		s.errorsCount.WithLabelValues("create", packID).Add(1)
	}

	return record, err
}

// Update implements the Service interface for metrics.
func (s *metricsService) Update(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update", packID).Add(1)
		s.requestLatency.WithLabelValues("update", packID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Update(ctx, packID, build)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("update", packID).Add(1)
	}

	return record, err
}

// Delete implements the Service interface for metrics.
func (s *metricsService) Delete(ctx context.Context, packID, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete", packID).Add(1)
		s.requestLatency.WithLabelValues("delete", packID).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Delete(ctx, packID, name)

	if err != nil {
		s.errorsCount.WithLabelValues("delete", packID).Add(1)
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *metricsService) Exists(ctx context.Context, packID, name string) (bool, error) {
	return s.service.Exists(ctx, packID, name)
}

// Column implements the Service interface.
func (s *metricsService) Column(ctx context.Context, packID, id, col string, val any) error {
	return s.service.Column(ctx, packID, id, col, val)
}
