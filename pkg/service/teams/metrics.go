package teams

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
	requestCount   *prometheus.CounterVec
}

// NewMetricsService wraps the Service and provides tracing for its methods.
func NewMetricsService(s Service, m *metrics.Metrics) Service {
	return &metricsService{
		service: s,
		requestLatency: m.RegisterHistogram(
			prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: m.Namespace,
					Subsystem: "teams_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the teams service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "teams_service",
					Name:      "request_count",
					Help:      "Total number of requests to the teams service.",
				},
				[]string{"method"},
			),
		),
	}
}

func (s *metricsService) List(ctx context.Context) ([]*model.Team, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.List(ctx)
}

func (s *metricsService) Show(ctx context.Context, id string) (*model.Team, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show").Add(1)
		s.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Show(ctx, id)
}

func (s *metricsService) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create").Add(1)
		s.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Create(ctx, team)
}

func (s *metricsService) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Update(ctx, team)
}

func (s *metricsService) Delete(ctx context.Context, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete").Add(1)
		s.requestLatency.WithLabelValues("delete").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Delete(ctx, name)
}

func (s *metricsService) ListUsers(ctx context.Context, name string) ([]*model.TeamUser, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("listUsers").Add(1)
		s.requestLatency.WithLabelValues("listUsers").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.ListUsers(ctx, name)
}

func (s *metricsService) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("appendUser").Add(1)
		s.requestLatency.WithLabelValues("appendUser").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.AppendUser(ctx, teamID, userID, perm)
}

func (s *metricsService) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("permitUser").Add(1)
		s.requestLatency.WithLabelValues("permitUser").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.PermitUser(ctx, teamID, userID, perm)
}

func (s *metricsService) DropUser(ctx context.Context, teamID, userID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("dropUser").Add(1)
		s.requestLatency.WithLabelValues("dropUser").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.DropUser(ctx, teamID, userID)
}
