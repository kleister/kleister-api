package userteams

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a user team could not be found.
	ErrNotFound = errors.New("team or user not found")

	// ErrAlreadyAssigned defines the error if a user team is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a user team is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with user team.
type Service interface {
	List(context.Context, model.UserTeamParams) ([]*model.UserTeam, int64, error)
	Attach(context.Context, model.UserTeamParams) error
	Permit(context.Context, model.UserTeamParams) error
	Drop(context.Context, model.UserTeamParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	userteams Service
}

// NewService returns a Service that handles all interactions with user teams.
func NewService(userteams Service) Service {
	return &service{
		userteams: userteams,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.userteams.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.UserTeamParams) ([]*model.UserTeam, int64, error) {
	return s.userteams.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.UserTeamParams) error {
	return s.userteams.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.UserTeamParams) error {
	return s.userteams.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.UserTeamParams) error {
	return s.userteams.Drop(ctx, params)
}
