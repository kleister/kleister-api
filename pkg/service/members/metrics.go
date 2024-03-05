package members

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
					Subsystem: "members_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the members service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "members_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the members service.",
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "members_service",
					Name:      "request_count",
					Help:      "Total number of requests to the members service.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx, teamID, userID)

	if err != nil {
		s.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Attach implements the Service interface for metrics.
func (s *metricsService) Attach(ctx context.Context, teamID, userID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("attach").Add(1)
		s.requestLatency.WithLabelValues("attach").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Attach(ctx, teamID, userID, perm)

	if err != nil {
		s.errorsCount.WithLabelValues("attach").Add(1)
	}

	return err
}

// Permit implements the Service interface for metrics.
func (s *metricsService) Permit(ctx context.Context, teamID, userID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("permit").Add(1)
		s.requestLatency.WithLabelValues("permit").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Permit(ctx, teamID, userID, perm)

	if err != nil {
		s.errorsCount.WithLabelValues("permit").Add(1)
	}

	return err
}

// Drop implements the Service interface for metrics.
func (s *metricsService) Drop(ctx context.Context, teamID, userID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("drop").Add(1)
		s.requestLatency.WithLabelValues("drop").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Drop(ctx, teamID, userID)

	if err != nil {
		s.errorsCount.WithLabelValues("drop").Add(1)
	}

	return err
}
