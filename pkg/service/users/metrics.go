package users

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
					Subsystem: "users_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the users service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
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

func (s *metricsService) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("byBasicAuth").Add(1)
		s.requestLatency.WithLabelValues("byBasicAuth").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.ByBasicAuth(ctx, username, password)
}

func (s *metricsService) List(ctx context.Context) ([]*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.List(ctx)
}

func (s *metricsService) Show(ctx context.Context, id string) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show").Add(1)
		s.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Show(ctx, id)
}

func (s *metricsService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create").Add(1)
		s.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Create(ctx, user)
}

func (s *metricsService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Update(ctx, user)
}

func (s *metricsService) Delete(ctx context.Context, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete").Add(1)
		s.requestLatency.WithLabelValues("delete").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Delete(ctx, name)
}

func (s *metricsService) ListTeams(ctx context.Context, name string) ([]*model.TeamUser, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("listTeams").Add(1)
		s.requestLatency.WithLabelValues("listTeams").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.ListTeams(ctx, name)
}

func (s *metricsService) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("appendTeam").Add(1)
		s.requestLatency.WithLabelValues("appendTeam").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.AppendTeam(ctx, userID, teamID, perm)
}

func (s *metricsService) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("permitTeam").Add(1)
		s.requestLatency.WithLabelValues("permitTeam").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.PermitTeam(ctx, userID, teamID, perm)
}

func (s *metricsService) DropTeam(ctx context.Context, userID, teamID string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("dropTeam").Add(1)
		s.requestLatency.WithLabelValues("dropTeam").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.DropTeam(ctx, userID, teamID)
}
