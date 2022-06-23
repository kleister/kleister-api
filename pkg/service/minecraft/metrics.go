package minecraft

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
					Subsystem: "minecraft_service",
					Name:      "request_latency_microseconds",
					Help:      "Histogram of latencies for requests to the minecraft service.",
					Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
				},
				[]string{"method"},
			),
		),
		requestCount: m.RegisterCounter(
			prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: m.Namespace,
					Subsystem: "minecraft_service",
					Name:      "request_count",
					Help:      "Total number of requests to the minecraft service.",
				},
				[]string{"method"},
			),
		),
	}
}

func (s *metricsService) List(ctx context.Context) ([]*model.Minecraft, error) {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("list").Add(1)
		s.requestLatency.WithLabelValues("list").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.List(ctx)
}

func (s *metricsService) Update(ctx context.Context) error {
	defer func(start time.Time) {
		s.requestCount.WithLabelValues("update").Add(1)
		s.requestLatency.WithLabelValues("update").Observe(time.Since(start).Seconds())
	}(time.Now())

	return s.service.Update(ctx)
}
