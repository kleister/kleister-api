package buildVersions

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
	ErrAlreadyAssigned = errors.New("build version already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("build version is not defined")
)

// Service handles all interactions with buildVersions.
type Service interface {
	List(context.Context, string, string, string, string) ([]*model.BuildVersion, error)
	Attach(context.Context, string, string, string, string) error
	Drop(context.Context, string, string, string, string) error
}

type service struct {
	buildVersions Service
}

// NewService returns a Service that handles all interactions with buildVersions.
func NewService(buildVersions Service) Service {
	return &service{
		buildVersions: buildVersions,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, packID, buildID, modID, versionID string) ([]*model.BuildVersion, error) {
	return s.buildVersions.List(ctx, packID, buildID, modID, versionID)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, packID, buildID, modID, versionID string) error {
	return s.buildVersions.Attach(ctx, packID, buildID, modID, versionID)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, packID, buildID, modID, versionID string) error {
	return s.buildVersions.Drop(ctx, packID, buildID, modID, versionID)
}
