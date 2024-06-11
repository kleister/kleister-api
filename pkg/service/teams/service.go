package teams

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a team was not found.
	ErrNotFound = errors.New("team not found")
)

// Service handles all interactions with teams.
type Service interface {
	List(context.Context, model.ListParams) ([]*model.Team, int64, error)
	Show(context.Context, string) (*model.Team, error)
	Create(context.Context, *model.Team) error
	Update(context.Context, *model.Team) error
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
	Column(context.Context, string, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	teams Service
}

// NewService returns a Service that handles all interactions with teams.
func NewService(teams Service) Service {
	return &service{
		teams: teams,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.teams.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.Team, int64, error) {
	return s.teams.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.Team, error) {
	return s.teams.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, team *model.Team) error {
	return s.teams.Create(ctx, team)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, team *model.Team) error {
	return s.teams.Update(ctx, team)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, name string) error {
	return s.teams.Delete(ctx, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, name string) (bool, error) {
	return s.teams.Exists(ctx, name)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, name, col string, val any) error {
	return s.teams.Column(ctx, name, col, val)
}
