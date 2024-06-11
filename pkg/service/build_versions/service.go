package buildversions

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("build or version not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("is already attached")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("is not attached")
)

// Service handles all interactions with buildversions.
type Service interface {
	List(context.Context, model.BuildVersionParams) ([]*model.BuildVersion, int64, error)
	Attach(context.Context, model.BuildVersionParams) error
	Drop(context.Context, model.BuildVersionParams) error
	WithPrincipal(*model.User) Service
}

type service struct {
	buildversions Service
}

// NewService returns a Service that handles all interactions with buildversions.
func NewService(buildversions Service) Service {
	return &service{
		buildversions: buildversions,
	}
}

// WithPrincipal implements the Service interface.
func (s *service) WithPrincipal(principal *model.User) Service {
	return s.buildversions.WithPrincipal(principal)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.BuildVersionParams) ([]*model.BuildVersion, int64, error) {
	return s.buildversions.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.BuildVersionParams) error {
	return s.buildversions.Attach(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.BuildVersionParams) error {
	return s.buildversions.Drop(ctx, params)
}
