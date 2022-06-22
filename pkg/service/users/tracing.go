package users

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

func (s *tracingService) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.ByBasicAuth")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("username", username)
	defer span.Finish()

	return s.service.ByBasicAuth(ctx, username, password)
}

func (s *tracingService) List(ctx context.Context) ([]*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.List")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.List(ctx)
}

func (s *tracingService) Show(ctx context.Context, id string) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.Show")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	defer span.Finish()

	return s.service.Show(ctx, id)
}

func (s *tracingService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	name := ""

	if user != nil {
		name = user.Username
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.Create")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Create(ctx, user)
}

func (s *tracingService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	id := ""
	name := ""

	if user != nil {
		id = user.ID
		name = user.Username
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.Update")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Update(ctx, user)
}

func (s *tracingService) Delete(ctx context.Context, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.Delete")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Delete(ctx, name)
}

func (s *tracingService) ListTeams(ctx context.Context, name string) ([]*model.TeamUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.ListTeams")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.ListTeams(ctx, name)
}

func (s *tracingService) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.AppendTeam")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.AppendTeam(ctx, userID, teamID, perm)
}

func (s *tracingService) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.PermTeam")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.PermitTeam(ctx, userID, teamID, perm)
}

func (s *tracingService) DropTeam(ctx context.Context, userID, teamID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.Service.DropTeam")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.DropTeam(ctx, userID, teamID)
}
