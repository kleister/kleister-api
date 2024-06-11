package teampacks

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("team or pack not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("team pack already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("team pack is not defined")
)

// Service handles all interactions with teampacks.
type Service interface {
	List(context.Context, model.TeamPackParams) ([]*model.TeamPack, int64, error)
	Attach(context.Context, model.TeamPackParams) error
	Permit(context.Context, model.TeamPackParams) error
	Drop(context.Context, model.TeamPackParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	teampacks Service
}

// NewService returns a Service that handles all interactions with teampacks.
func NewService(teampacks Service) Service {
	return &service{
		teampacks: teampacks,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.teampacks.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.TeamPackParams) ([]*model.TeamPack, int64, error) {
	return s.teampacks.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.TeamPackParams) error {
	return s.teampacks.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.TeamPackParams) error {
	return s.teampacks.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.TeamPackParams) error {
	return s.teampacks.Drop(ctx, params)
}
