package versions

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
					Subsystem: "versions_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the versions service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "versions_service",
					Name:      "request_count",
					Help:      "Total number of requests to the versions service.",
				},
				[]string{"method"},
			),
		),
	}
}

func (s *metricsService) List(ctx context.Context, mod *model.Mod) ([]*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.List(ctx, mod)
}

func (s *metricsService) Show(ctx context.Context, mod *model.Mod, id string) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("show").Add(1)
		s.requestLatency.WithLabelValues("show").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Show(ctx, mod, id)
}

func (s *metricsService) Create(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("create").Add(1)
		s.requestLatency.WithLabelValues("create").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Create(ctx, mod, version)
}

func (s *metricsService) Update(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Update(ctx, mod, version)
}

func (s *metricsService) Delete(ctx context.Context, mod *model.Mod, name string) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("delete").Add(1)
		s.requestLatency.WithLabelValues("delete").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Delete(ctx, mod, name)
}