package builds

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

func (s *tracingService) List(ctx context.Context, pack *model.Pack) ([]*model.Build, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "builds.Service.List")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("pack", pack.ID)
	defer span.Finish()

	return s.service.List(ctx, pack)
}

func (s *tracingService) Show(ctx context.Context, pack *model.Pack, id string) (*model.Build, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "builds.Service.Show")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("pack", pack.ID)
	span.SetTag("id", id)
	defer span.Finish()

	return s.service.Show(ctx, pack, id)
}

func (s *tracingService) Create(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	name := ""

	if build != nil {
		name = build.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "builds.Service.Create")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("pack", pack.ID)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Create(ctx, pack, build)
}

func (s *tracingService) Update(ctx context.Context, pack *model.Pack, build *model.Build) (*model.Build, error) {
	id := ""
	name := ""

	if build != nil {
		id = build.ID
		name = build.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "builds.Service.Update")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("pack", pack.ID)
	span.SetTag("id", id)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Update(ctx, pack, build)
}

func (s *tracingService) Delete(ctx context.Context, pack *model.Pack, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "builds.Service.Delete")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("pack", pack.ID)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Delete(ctx, pack, name)
}
