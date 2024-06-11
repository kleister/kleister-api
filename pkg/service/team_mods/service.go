package teammods

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("team or mod not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("team mod already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("team mod is not defined")
)

// Service handles all interactions with teammods.
type Service interface {
	List(context.Context, model.TeamModParams) ([]*model.TeamMod, int64, error)
	Attach(context.Context, model.TeamModParams) error
	Permit(context.Context, model.TeamModParams) error
	Drop(context.Context, model.TeamModParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	teammods Service
}

// NewService returns a Service that handles all interactions with teammods.
func NewService(teammods Service) Service {
	return &service{
		teammods: teammods,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.teammods.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.TeamModParams) ([]*model.TeamMod, int64, error) {
	return s.teammods.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.TeamModParams) error {
	return s.teammods.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.TeamModParams) error {
	return s.teammods.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.TeamModParams) error {
	return s.teammods.Drop(ctx, params)
}
