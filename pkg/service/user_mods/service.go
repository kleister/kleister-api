package usermods

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("user or mod not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("user mod already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("user mod is not defined")
)

// Service handles all interactions with usermods.
type Service interface {
	List(context.Context, model.UserModParams) ([]*model.UserMod, int64, error)
	Attach(context.Context, model.UserModParams) error
	Permit(context.Context, model.UserModParams) error
	Drop(context.Context, model.UserModParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	usermods Service
}

// NewService returns a Service that handles all interactions with usermods.
func NewService(usermods Service) Service {
	return &service{
		usermods: usermods,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.usermods.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.UserModParams) ([]*model.UserMod, int64, error) {
	return s.usermods.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.UserModParams) error {
	return s.usermods.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.UserModParams) error {
	return s.usermods.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.UserModParams) error {
	return s.usermods.Drop(ctx, params)
}
