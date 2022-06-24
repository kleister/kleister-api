package versions

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

func (s *tracingService) List(ctx context.Context, mod *model.Mod) ([]*model.Version, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "versions.Service.List")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("mod", mod.ID)
	defer span.Finish()

	return s.service.List(ctx, mod)
}

func (s *tracingService) Show(ctx context.Context, mod *model.Mod, id string) (*model.Version, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "versions.Service.Show")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("mod", mod.ID)
	span.SetTag("id", id)
	defer span.Finish()

	return s.service.Show(ctx, mod, id)
}

func (s *tracingService) Create(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	name := ""

	if version != nil {
		name = version.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "versions.Service.Create")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("mod", mod.ID)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Create(ctx, mod, version)
}

func (s *tracingService) Update(ctx context.Context, mod *model.Mod, version *model.Version) (*model.Version, error) {
	id := ""
	name := ""

	if version != nil {
		id = version.ID
		name = version.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "versions.Service.Update")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("mod", mod.ID)
	span.SetTag("id", id)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Update(ctx, mod, version)
}

func (s *tracingService) Delete(ctx context.Context, mod *model.Mod, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "versions.Service.Delete")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("mod", mod.ID)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Delete(ctx, mod, name)
}
