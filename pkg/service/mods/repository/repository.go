package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrModNotFound defines the error if a mod could not be found.
	ErrModNotFound = errors.New("mod not found")

	// ErrUserNotAssigned defines the error if a teauserm is not assigned.
	ErrUserNotAssigned = errors.New("team is not assigned")

	// ErrUserAlreadyAssigned defines the error if a user is already assigned .
	ErrUserAlreadyAssigned = errors.New("team is already assigned")

	// ErrTeamNotAssigned defines the error if a team is not assigned.
	ErrTeamNotAssigned = errors.New("team is not assigned")

	// ErrTeamAlreadyAssigned defines the error if a team is already assigned.
	ErrTeamAlreadyAssigned = errors.New("team is already assigned")
)

// ModsRepository defines the required functions for the repository.
type ModsRepository interface {
	List(context.Context, string) ([]*model.Mod, error)
	Create(context.Context, *model.Mod) (*model.Mod, error)
	Update(context.Context, *model.Mod) (*model.Mod, error)
	Show(context.Context, string) (*model.Mod, error)
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, string, error)

	ListUsers(context.Context, string, string) ([]*model.UserMod, error)
	AttachUser(context.Context, string, string) error
	DropUser(context.Context, string, string) error

	ListTeams(context.Context, string, string) ([]*model.TeamMod, error)
	AttachTeam(context.Context, string, string) error
	DropTeam(context.Context, string, string) error
}
