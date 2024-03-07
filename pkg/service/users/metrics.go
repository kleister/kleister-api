package users

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
					Subsystem: "users_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the users service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "users_service",
					Name:      "errors_count",
					Help:      "Total number of errors within the users service.",
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "users_service",
					Name:      "request_count",
					Help:      "Total number of requests to the users service.",
				},
				[]string{"method"},
			),
		),
	}
}

// ByBasicAuth implements the Service interface for metrics.
func (s *metricsService) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("byBasicAuth").Add(1)
		s.requestLatency.WithLabelValues("byBasicAuth").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.ByBasicAuth(ctx, username, password)

	if err != nil {
		s.errorsCount.WithLabelValues("byBasicAuth").Add(1)
	}

	return record, err
}

// List implements the Service interface for metrics.
func (s *metricsService) List(ctx context.Context) ([]*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := s.service.List(ctx)

	if err != nil {
		s.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Show implements the Service interface for metrics.
func (s *metricsService) Show(ctx context.Context, id string) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show").Add(1)
		s.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Show(ctx, id)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("show").Add(1)
	}

	return record, err
}

// Create implements the Service interface for metrics.
func (s *metricsService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create").Add(1)
		s.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Create(ctx, user)

	if err != nil {
		s.errorsCount.WithLabelValues("create").Add(1)
	}

	return record, err
}

// Update implements the Service interface for metrics.
func (s *metricsService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := s.service.Update(ctx, user)

	if err != nil && !errors.Is(err, ErrNotFound) {
		s.errorsCount.WithLabelValues("update").Add(1)
	}

	return record, err
}

// Delete implements the Service interface for metrics.
func (s *metricsService) Delete(ctx context.Context, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete").Add(1)
		s.requestLatency.WithLabelValues("delete").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := s.service.Delete(ctx, name)

	if err != nil {
		s.errorsCount.WithLabelValues("delete").Add(1)
	}

	return err
}

// Exists implements the Service interface for logging.
func (s *metricsService) Exists(ctx context.Context, name string) (bool, error) {
	return s.service.Exists(ctx, name)
}

// External implements the Service interface for database persistence.
func (s *metricsService) External(ctx context.Context, username, email, fullname string, admin bool) (*model.User, error) {
	return s.service.External(ctx, username, email, fullname, admin)
}
