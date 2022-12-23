package repository

import (
	"context"
	"errors"
	"time"

	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

// MetricsRepository implements UsersRepository interface.
type MetricsRepository struct {
	upstream       UsersRepository
	requestLatency *prometheus.HistogramVec
	errorsCount    *prometheus.CounterVec
	requestCount   *prometheus.CounterVec
}

// NewMetricsRepository wraps the UsersRepository and provides metrics for its methods.
func NewMetricsRepository(repository UsersRepository, metricz *metrics.Metrics) UsersRepository {
	return &MetricsRepository{
		upstream: repository,
		requestLatency: metricz.RegisterHistogram(
			prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: metricz.Namespace,
					Subsystem: "users_repository",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the users repository.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "users_repository",
					Name:      "errors_count",
					Help:      "Total number of errors within the users repository.",
				},
				[]string{"method"},
			),
		),
		requestCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "users_repository",
					Name:      "request_count",
					Help:      "Total number of requests to the users repository.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the UsersRepository interface.
func (r *MetricsRepository) List(ctx context.Context) ([]*model.User, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("list").Add(1)
		r.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := r.upstream.List(ctx)

	if err != nil {
		r.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Create implements the UsersRepository interface.
func (r *MetricsRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("create").Add(1)
		r.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Create(ctx, user)

	if err != nil {
		r.errorsCount.WithLabelValues("create").Add(1)
	}

	return record, err
}

// Update implements the UsersRepository interface.
func (r *MetricsRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("update").Add(1)
		r.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Update(ctx, user)

	if err != nil && !errors.Is(err, ErrUserNotFound) {
		r.errorsCount.WithLabelValues("update").Add(1)
	}

	return record, err
}

// Show implements the UsersRepository interface.
func (r *MetricsRepository) Show(ctx context.Context, id string) (*model.User, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("show").Add(1)
		r.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Show(ctx, id)

	if err != nil && !errors.Is(err, ErrUserNotFound) {
		r.errorsCount.WithLabelValues("show").Add(1)
	}

	return record, err
}

// Delete implements the UsersRepository interface.
func (r *MetricsRepository) Delete(ctx context.Context, name string) error {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("delete").Add(1)
		r.requestLatency.WithLabelValues("delete").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := r.upstream.Delete(ctx, name)

	if err != nil {
		r.errorsCount.WithLabelValues("delete").Add(1)
	}

	return err
}

// Exists implements the UsersRepository interface.
func (r *MetricsRepository) Exists(ctx context.Context, name string) (bool, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("exists").Add(1)
		r.requestLatency.WithLabelValues("exists").Observe(time.Since(start).Seconds())
	}(time.Now())

	exists, err := r.upstream.Exists(ctx, name)

	if err != nil {
		r.errorsCount.WithLabelValues("exists").Add(1)
	}

	return exists, err
}
