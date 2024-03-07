package userMods

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

// Service handles all interactions with userMods.
type Service interface {
	List(context.Context, string, string) ([]*model.UserMod, error)
	Attach(context.Context, string, string, string) error
	Permit(context.Context, string, string, string) error
	Drop(context.Context, string, string) error
}

type service struct {
	userMods Service
}

// NewService returns a Service that handles all interactions with userMods.
func NewService(userMods Service) Service {
	return &service{
		userMods: userMods,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, userID, modID string) ([]*model.UserMod, error) {
	return s.userMods.List(ctx, userID, modID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, userID, modID, perm string) error {
	return s.userMods.Attach(ctx, userID, modID, perm)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, userID, modID, perm string) error {
	return s.userMods.Permit(ctx, userID, modID, perm)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, userID, modID string) error {
	return s.userMods.Drop(ctx, userID, modID)
}
