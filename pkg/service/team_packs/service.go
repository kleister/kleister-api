package teamPacks

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

// Service handles all interactions with teamPacks.
type Service interface {
	List(context.Context, string, string) ([]*model.TeamPack, error)
	Attach(context.Context, string, string, string) error
	Permit(context.Context, string, string, string) error
	Drop(context.Context, string, string) error
}

type service struct {
	teamPacks Service
}

// NewService returns a Service that handles all interactions with teamPacks.
func NewService(teamPacks Service) Service {
	return &service{
		teamPacks: teamPacks,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, teamID, packID string) ([]*model.TeamPack, error) {
	return s.teamPacks.List(ctx, teamID, packID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, teamID, packID, perm string) error {
	return s.teamPacks.Attach(ctx, teamID, packID, perm)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, teamID, packID, perm string) error {
	return s.teamPacks.Permit(ctx, teamID, packID, perm)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, teamID, packID string) error {
	return s.teamPacks.Drop(ctx, teamID, packID)
}
