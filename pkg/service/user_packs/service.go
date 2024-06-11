package userpacks

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("user or pack not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("user pack already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("user pack is not defined")
)

// Service handles all interactions with userpacks.
type Service interface {
	List(context.Context, model.UserPackParams) ([]*model.UserPack, int64, error)
	Attach(context.Context, model.UserPackParams) error
	Permit(context.Context, model.UserPackParams) error
	Drop(context.Context, model.UserPackParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	userpacks Service
}

// NewService returns a Service that handles all interactions with userpacks.
func NewService(userpacks Service) Service {
	return &service{
		userpacks: userpacks,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.userpacks.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.UserPackParams) ([]*model.UserPack, int64, error) {
	return s.userpacks.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.UserPackParams) error {
	return s.userpacks.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.UserPackParams) error {
	return s.userpacks.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.UserPackParams) error {
	return s.userpacks.Drop(ctx, params)
}
