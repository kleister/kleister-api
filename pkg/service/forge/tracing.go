package forge

import (
	"context"

	"github.com/kleister/kleister-api/pkg/model"
	"github.com/opentracing/opentracing-go"
)

// TracingRequestID returns the request ID as string for tracing
type TracingRequestID func(context.Context) string

type tracingService struct {
	service   Service
	requestID TracingRequestID
}

// NewTracingService wraps the Service and provides tracing for its methods.
func NewTracingService(s Service, requestID TracingRequestID) Service {
	return &tracingService{
		service:   s,
		requestID: requestID,
	}
}

func (s *tracingService) List(ctx context.Context) ([]*model.Forge, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "forge.Service.List")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.List(ctx)
}
