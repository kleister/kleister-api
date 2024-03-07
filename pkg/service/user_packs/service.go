package userPacks

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

// Service handles all interactions with userPacks.
type Service interface {
	List(context.Context, string, string) ([]*model.UserPack, error)
	Attach(context.Context, string, string, string) error
	Permit(context.Context, string, string, string) error
	Drop(context.Context, string, string) error
}

type service struct {
	userPacks Service
}

// NewService returns a Service that handles all interactions with userPacks.
func NewService(userPacks Service) Service {
	return &service{
		userPacks: userPacks,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, userID, packID string) ([]*model.UserPack, error) {
	return s.userPacks.List(ctx, userID, packID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, userID, packID, perm string) error {
	return s.userPacks.Attach(ctx, userID, packID, perm)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, userID, packID, perm string) error {
	return s.userPacks.Permit(ctx, userID, packID, perm)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, userID, packID string) error {
	return s.userPacks.Drop(ctx, userID, packID)
}
