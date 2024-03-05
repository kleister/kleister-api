package members

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("team or user not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("membership already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("membership is not defined")
)

// Service handles all interactions with members.
type Service interface {
	List(context.Context, string, string) ([]*model.Member, error)
	Attach(context.Context, string, string, string) error
	Permit(context.Context, string, string, string) error
	Drop(context.Context, string, string) error
}

type service struct {
	members Service
}

// NewService returns a Service that handles all interactions with members.
func NewService(members Service) Service {
	return &service{
		members: members,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	return s.members.List(ctx, teamID, userID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, teamID, userID, perm string) error {
	return s.members.Attach(ctx, teamID, userID, perm)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, teamID, userID, perm string) error {
	return s.members.Permit(ctx, teamID, userID, perm)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, teamID, userID string) error {
	return s.members.Drop(ctx, teamID, userID)
}
