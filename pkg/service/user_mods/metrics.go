package userMods

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
					Subsystem: "user_mods_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the userMods service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "user_mods_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the userMods service.",
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "user_mods_service",
					Name:      "request_count",
					Help:      "Total number of requests to the userMods service.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, userID, modID string) ([]*model.UserMod, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx, userID, modID)

	if err != nil {
		s.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Attach implements the Service interface for metrics.
func (s *metricsService) Attach(ctx context.Context, userID, modID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("attach").Add(1)
		s.requestLatency.WithLabelValues("attach").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Attach(ctx, userID, modID, perm)

	if err != nil {
		s.errorsCount.WithLabelValues("attach").Add(1)
	}

	return err
}

// Permit implements the Service interface for metrics.
func (s *metricsService) Permit(ctx context.Context, userID, modID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("permit").Add(1)
		s.requestLatency.WithLabelValues("permit").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Permit(ctx, userID, modID, perm)

	if err != nil {
		s.errorsCount.WithLabelValues("permit").Add(1)
	}

	return err
}

// Drop implements the Service interface for metrics.
func (s *metricsService) Drop(ctx context.Context, userID, modID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("drop").Add(1)
		s.requestLatency.WithLabelValues("drop").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Drop(ctx, userID, modID)

	if err != nil {
		s.errorsCount.WithLabelValues("drop").Add(1)
	}

	return err
}
