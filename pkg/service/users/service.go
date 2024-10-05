package users

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a user was not found.
	ErrNotFound = errors.New("user not found")

	// ErrWrongCredentials is returned when credentials are wrong.
	ErrWrongCredentials = errors.New("wrong credentials provided")
)

// Service handles all interactions with users.
type Service interface {
	External(context.Context, string, string, string, string, string, bool) (*model.User, error)
	AuthByID(context.Context, string) (*model.User, error)
	AuthByCreds(context.Context, string, string) (*model.User, error)
	List(context.Context, model.ListParams) ([]*model.User, int64, error)
	Show(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) error
	Update(context.Context, *model.User) error
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
	Column(context.Context, string, string, any) error
	WithPrincipal(*model.User) Service
}

type service struct {
	users Service
}

// NewService returns a Service that handles all interactions with users.
func NewService(users Service) Service {
	return &service{
		users: users,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.users.WithPrincipal(principal)
}

// External implements the Service interface.
func (s *service) External(ctx context.Context, provider, ref, username, email, fullname string, admin bool) (*model.User, error) {
	return s.users.External(ctx, provider, ref, username, email, fullname, admin)
}

// AuthByID implements the Service interface.
func (s *service) AuthByID(ctx context.Context, userID string) (*model.User, error) {
	return s.users.AuthByID(ctx, userID)
}

// AuthByCreds implements the Service interface.
func (s *service) AuthByCreds(ctx context.Context, username, password string) (*model.User, error) {
	return s.users.AuthByCreds(ctx, username, password)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.User, int64, error) {
	return s.users.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.User, error) {
	return s.users.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, user *model.User) error {
	return s.users.Create(ctx, user)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, user *model.User) error {
	return s.users.Update(ctx, user)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, name string) error {
	return s.users.Delete(ctx, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, name string) (bool, error) {
	return s.users.Exists(ctx, name)
}

// Column implements the Service interface.
func (s *service) Column(ctx context.Context, name, col string, val any) error {
	return s.users.Column(ctx, name, col, val)
}
