package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrTeamNotFound defines the error if a team could not be found.
	ErrTeamNotFound = errors.New("user not found")

	// ErrTeamOrUserNotFound defines the error if a team or user could not be found.
	ErrTeamOrUserNotFound = errors.New("team or user not found")

	// ErrUserNotAssigned defines the error if a team is not assigned to a user.
	ErrUserNotAssigned = errors.New("team is not assigned")

	// ErrUserAlreadyAssigned defines the error if a team is already assigned to a user.
	ErrUserAlreadyAssigned = errors.New("team is already assigned")
)

// TeamsRepository defines the required functions for the repository.
type TeamsRepository interface {
	List(context.Context) ([]*model.Team, error)
	Create(context.Context, *model.Team) (*model.Team, error)
	Update(context.Context, *model.Team) (*model.Team, error)
	Show(context.Context, string) (*model.Team, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
}
