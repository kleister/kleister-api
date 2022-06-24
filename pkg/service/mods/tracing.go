package mods

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

func (s *tracingService) List(ctx context.Context) ([]*model.Mod, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mods.Service.List")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.List(ctx)
}

func (s *tracingService) Show(ctx context.Context, id string) (*model.Mod, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mods.Service.Show")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	defer span.Finish()

	return s.service.Show(ctx, id)
}

func (s *tracingService) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	name := ""

	if mod != nil {
		name = mod.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "mods.Service.Create")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Create(ctx, mod)
}

func (s *tracingService) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	id := ""
	name := ""

	if mod != nil {
		id = mod.ID
		name = mod.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "mods.Service.Update")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Update(ctx, mod)
}

func (s *tracingService) Delete(ctx context.Context, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mods.Service.Delete")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Delete(ctx, name)
}
