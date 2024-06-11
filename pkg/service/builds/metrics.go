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

// External implements the Service interface for metrics.
func (s *metricsService) WithPrincipal(principal *model.User) Service {
	s.service.WithPrincipal(principal)
	return s
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, params model.BuildParams) ([]*model.Build, int64, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list", params.PackID).Add(1)
		s.requestLatency.WithLabelValues("list", params.PackID).Observe(time.Since(start).Seconds())
	}(time.Now())

	records, counter, err := s.service.List(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("list", params.PackID).Add(1)
	}

	return records, counter, err
}

// Show implements the Service interface for metrics.
func (s *metricsService) Show(ctx context.Context, params model.BuildParams) (*model.Build, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show", params.PackID).Add(1)
		s.requestLatency.WithLabelValues("show", params.PackID).Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Show(ctx, params)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("show", params.PackID).Add(1)
	}

	return record, err
}

// Create implements the Service interface for metrics.
func (s *metricsService) Create(ctx context.Context, params model.BuildParams, build *model.Build) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create", params.PackID).Add(1)
		s.requestLatency.WithLabelValues("create", params.PackID).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Create(ctx, params, build)

	if err != nil {
		s.errorsCount.WithLabelValues("create", params.PackID).Add(1)
	}

	return err
}

// Update implements the Service interface for metrics.
func (s *metricsService) Update(ctx context.Context, params model.BuildParams, build *model.Build) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update", params.PackID).Add(1)
		s.requestLatency.WithLabelValues("update", params.PackID).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Update(ctx, params, build)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("update", params.PackID).Add(1)
	}

	return err
}

// Delete implements the Service interface for metrics.
func (s *metricsService) Delete(ctx context.Context, params model.BuildParams) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete", params.PackID).Add(1)
		s.requestLatency.WithLabelValues("delete", params.PackID).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Delete(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("delete", params.PackID).Add(1)
	}

	return err
}

// Exists implements the Service interface for metrics.
func (s *metricsService) Exists(ctx context.Context, params model.BuildParams) (bool, error) {
	return s.service.Exists(ctx, params)
}

// Column implements the Service interface for metrics.
func (s *metricsService) Column(ctx context.Context, params model.BuildParams, col string, val any) error {
	return s.service.Column(ctx, params, col, val)
}
