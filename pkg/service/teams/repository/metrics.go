package repository

import (
	"context"
	"errors"
	"time"

	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

// MetricsTeamsRepository implements TeamsRepository interface.
type MetricsTeamsRepository struct {
	upstream       TeamsRepository
	requestLatency *prometheus.HistogramVec
	errorsCount    *prometheus.CounterVec
	requestCount   *prometheus.CounterVec
}

// NewMetricsRepository wraps the TeamsRepository and provides metrics for its methods.
func NewMetricsRepository(repository TeamsRepository, metricz *metrics.Metrics) TeamsRepository {
	return &MetricsTeamsRepository{
		upstream: repository,
		requestLatency: metricz.RegisterHistogram(
			prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: metricz.Namespace,
					Subsystem: "teams_repository",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the teams repository.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "teams_repository",
					Name:      "errors_count",
					Help:      "Total number of errors within the teams repository.",
				},
				[]string{"method"},
			),
		),
		requestCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "teams_repository",
					Name:      "request_count",
					Help:      "Total number of requests to the teams repository.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) List(ctx context.Context) ([]*model.Team, error) {
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

// Create implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("create").Add(1)
		r.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Create(ctx, team)

	if err != nil {
		r.errorsCount.WithLabelValues("create").Add(1)
	}

	return record, err
}

// Update implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("update").Add(1)
		r.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Update(ctx, team)

	if err != nil && !errors.Is(err, ErrTeamNotFound) {
		r.errorsCount.WithLabelValues("update").Add(1)
	}

	return record, err
}

// Show implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) Show(ctx context.Context, id string) (*model.Team, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("show").Add(1)
		r.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	record, err := r.upstream.Show(ctx, id)

	if err != nil && !errors.Is(err, ErrTeamNotFound) {
		r.errorsCount.WithLabelValues("show").Add(1)
	}

	return record, err
}

// Delete implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) Delete(ctx context.Context, name string) error {
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

// Exists implements the TeamsRepository interface.
func (r *MetricsTeamsRepository) Exists(ctx context.Context, name string) (bool, error) {
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
