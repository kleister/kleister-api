package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrMemberNotFound defines the error if a member could not be found.
	ErrMemberNotFound = errors.New("team or user not found")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("membership is not defined")

	// ErrIsAssigned defines the error if a member is already assigned.
	ErrIsAssigned = errors.New("membership already exists")
)

// MembersRepository defines the required functions for the repository.
type MembersRepository interface {
	List(context.Context, string, string) ([]*model.Member, error)
	Append(context.Context, string, string) error
	Drop(context.Context, string, string) error
}
