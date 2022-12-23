package repository

import (
	"context"
	"time"

	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

// MetricsMembersRepository implements MembersRepository interface.
type MetricsMembersRepository struct {
	upstream       MembersRepository
	requestLatency *prometheus.HistogramVec
	errorsCount    *prometheus.CounterVec
	requestCount   *prometheus.CounterVec
}

// NewMetricsRepository wraps the MembersRepository and provides metrics for its methods.
func NewMetricsRepository(repository MembersRepository, metricz *metrics.Metrics) MembersRepository {
	return &MetricsMembersRepository{
		upstream: repository,
		requestLatency: metricz.RegisterHistogram(
			prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: metricz.Namespace,
					Subsystem: "members_repository",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the members repository.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		errorsCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "members_repository",
					Name:      "errors_count",
					Help:      "Total number of errors within the members repository.",
				},
				[]string{"method"},
			),
		),
		requestCount: metricz.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: metricz.Namespace,
					Subsystem: "members_repository",
					Name:      "request_count",
					Help:      "Total number of requests to the members repository.",
				},
				[]string{"method"},
			),
		),
	}
}

// List implements the MembersRepository interface.
func (r *MetricsMembersRepository) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("list").Add(1)
		r.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	records, err := r.upstream.List(ctx, teamID, userID)

	if err != nil {
		r.errorsCount.WithLabelValues("list").Add(1)
	}

	return records, err
}

// Append implements the MembersRepository interface.
func (r *MetricsMembersRepository) Append(ctx context.Context, teamID, userID string) error {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("append").Add(1)
		r.requestLatency.WithLabelValues("append").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := r.upstream.Append(ctx, teamID, userID)

	if err != nil {
		r.errorsCount.WithLabelValues("append").Add(1)
	}

	return err
}

// Drop implements the MembersRepository interface.
func (r *MetricsMembersRepository) Drop(ctx context.Context, teamID, userID string) error {
	defer func(start time.Time) {
		r.requestCount.WithLabelValues("drop").Add(1)
		r.requestLatency.WithLabelValues("drop").Observe(time.Since(start).Seconds())
	}(time.Now())

	err := r.upstream.Drop(ctx, teamID, userID)

	if err != nil {
		r.errorsCount.WithLabelValues("drop").Add(1)
	}

	return err
}
