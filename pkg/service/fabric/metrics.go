package fabric

import (
	"context"
	"errors"
	"time"

	fabricClient "github.com/kleister/kleister-api/pkg/internal/fabric"
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

// External implements the Service interface for metrics.
func (s *metricsService) WithPrincipal(principal *model.User) Service {
	s.service.WithPrincipal(principal)
	return s
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, params model.ListParams) ([]*model.Fabric, int64, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("search").Add(1)
		s.requestLatency.WithLabelValues("search").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, counter, err := s.service.List(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("search").Add(1)
	}

	return records, counter, err
}

// Show implements the Service interface for metrics.
func (s *metricsService) Show(ctx context.Context, name string) (*model.Fabric, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show").Add(1)
		s.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Show(ctx, name)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("show").Add(1)
	}

	return record, err
}

// Sync implements the Service interface for metrics.
func (s *metricsService) Sync(ctx context.Context, versions fabricClient.Versions) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("sync").Add(1)
		s.requestLatency.WithLabelValues("sync").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Sync(ctx, versions)

	if err != nil {
		s.errorsCount.WithLabelValues("sync").Add(1)
	}

	return err
}

// ListBuilds implements the Service interface for metrics.
func (s *metricsService) ListBuilds(ctx context.Context, params model.FabricBuildParams) ([]*model.Build, int64, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("listBuilds").Add(1)
		s.requestLatency.WithLabelValues("listBuilds").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, counter, err := s.service.ListBuilds(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("listBuilds").Add(1)
	}

	return records, counter, err
}

// AttachBuild implements the Service interface for metrics.
func (s *metricsService) AttachBuild(ctx context.Context, params model.FabricBuildParams) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("attachBuild").Add(1)
		s.requestLatency.WithLabelValues("attachBuild").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.AttachBuild(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("attachBuild").Add(1)
	}

	return err
}

// DropBuild implements the Service interface for metrics.
func (s *metricsService) DropBuild(ctx context.Context, params model.FabricBuildParams) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("dropBuild").Add(1)
		s.requestLatency.WithLabelValues("dropBuild").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.DropBuild(ctx, params)

	if err != nil {
		s.errorsCount.WithLabelValues("dropBuild").Add(1)
	}

	return err
}
