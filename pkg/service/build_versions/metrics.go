package buildVersions

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
					Subsystem: "build_versions_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the buildVersions service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "build_versions_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the buildVersions service.",
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "build_versions_service",
					Name:      "request_count",
					Help:      "Total number of requests to the buildVersions service.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, packID, buildID, modID, versionID string) ([]*model.BuildVersion, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx, packID, buildID, modID, versionID)

	if err != nil {
		s.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Attach implements the Service interface for metrics.
func (s *metricsService) Attach(ctx context.Context, packID, buildID, modID, versionID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("attach").Add(1)
		s.requestLatency.WithLabelValues("attach").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Attach(ctx, packID, buildID, modID, versionID)

	if err != nil {
		s.errorsCount.WithLabelValues("attach").Add(1)
	}

	return err
}

// Drop implements the Service interface for metrics.
func (s *metricsService) Drop(ctx context.Context, packID, buildID, modID, versionID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("drop").Add(1)
		s.requestLatency.WithLabelValues("drop").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Drop(ctx, packID, buildID, modID, versionID)

	if err != nil {
		s.errorsCount.WithLabelValues("drop").Add(1)
	}

	return err
}
