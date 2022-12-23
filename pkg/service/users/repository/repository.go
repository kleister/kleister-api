package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrUserNotFound defines the error if a user could not be found.
	ErrUserNotFound = errors.New("user not found")

	// ErrUserOrTeamNotFound defines the error if a user or team could not be found.
	ErrUserOrTeamNotFound = errors.New("user or team not found")

	// ErrTeamNotAssigned defines the error if a team is not assigned to a user.
	ErrTeamNotAssigned = errors.New("team is not assigned")

	// ErrTeamAlreadyAssigned defines the error if a team is already assigned to a user.
	ErrTeamAlreadyAssigned = errors.New("team is already assigned")
)

// UsersRepository defines the required functions for the repository.
type UsersRepository interface {
	List(context.Context) ([]*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Show(context.Context, string) (*model.User, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
}
