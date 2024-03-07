package teamMods

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

// Service handles all interactions with teamMods.
type Service interface {
	List(context.Context, string, string) ([]*model.TeamMod, error)
	Attach(context.Context, string, string, string) error
	Permit(context.Context, string, string, string) error
	Drop(context.Context, string, string) error
}

type service struct {
	teamMods Service
}

// NewService returns a Service that handles all interactions with teamMods.
func NewService(teamMods Service) Service {
	return &service{
		teamMods: teamMods,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, teamID, modID string) ([]*model.TeamMod, error) {
	return s.teamMods.List(ctx, teamID, modID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, teamID, modID, perm string) error {
	return s.teamMods.Attach(ctx, teamID, modID, perm)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, teamID, modID, perm string) error {
	return s.teamMods.Permit(ctx, teamID, modID, perm)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, teamID, modID string) error {
	return s.teamMods.Drop(ctx, teamID, modID)
}
