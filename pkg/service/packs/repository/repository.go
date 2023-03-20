package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrPackNotFound defines the error if a pack could not be found.
	ErrPackNotFound = errors.New("pack not found")

	// ErrUserNotAssigned defines the error if a teauserm is not assigned.
	ErrUserNotAssigned = errors.New("team is not assigned")

	// ErrUserAlreadyAssigned defines the error if a user is already assigned .
	ErrUserAlreadyAssigned = errors.New("team is already assigned")

	// ErrTeamNotAssigned defines the error if a team is not assigned.
	ErrTeamNotAssigned = errors.New("team is not assigned")

	// ErrTeamAlreadyAssigned defines the error if a team is already assigned.
	ErrTeamAlreadyAssigned = errors.New("team is already assigned")
)

// PacksRepository defines the required functions for the repository.
type PacksRepository interface {
	List(context.Context, string) ([]*model.Pack, error)
	Create(context.Context, *model.Pack) (*model.Pack, error)
	Update(context.Context, *model.Pack) (*model.Pack, error)
	Show(context.Context, string) (*model.Pack, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, string, error)

	ListUsers(context.Context, string, string) ([]*model.UserPack, error)
	AttachUser(context.Context, string, string) error
	DropUser(context.Context, string, string) error

	ListTeams(context.Context, string, string) ([]*model.TeamPack, error)
	AttachTeam(context.Context, string, string) error
	DropTeam(context.Context, string, string) error
}
