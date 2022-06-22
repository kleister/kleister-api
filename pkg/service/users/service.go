package users

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a user was not found.
	ErrNotFound = errors.New("user not found")

	// ErrAlreadyAssigned is returned when a user is already assigned.
	ErrAlreadyAssigned = errors.New("user is already assigned")

	// ErrNotAssigned is returned when a user is not assigned.
	ErrNotAssigned = errors.New("user is not assigned")

	// ErrWrongAuth is returned when providing wrong credentials.
	ErrWrongAuth = errors.New("wrong username or password")
)

// Service handles all interactions with users.
type Service interface {
	ByBasicAuth(context.Context, string, string) (*model.User, error)

	List(context.Context) ([]*model.User, error)
	Show(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, string) error

	ListTeams(context.Context, string) ([]*model.TeamUser, error)
	AppendTeam(context.Context, string, string, string) error
	PermitTeam(context.Context, string, string, string) error
	DropTeam(context.Context, string, string) error
}

// Store defines the interface to persist users.
type Store interface {
	ByBasicAuth(context.Context, string, string) (*model.User, error)

	List(context.Context) ([]*model.User, error)
	Show(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, string) error

	ListTeams(context.Context, string) ([]*model.TeamUser, error)
	AppendTeam(context.Context, string, string, string) error
	PermitTeam(context.Context, string, string, string) error
	DropTeam(context.Context, string, string) error
}

type service struct {
	users Store
}

// NewService returns a Service that handles all interactions with users.
func NewService(users Store) Service {
	return &service{
		users: users,
	}
}

func (s *service) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	return s.users.ByBasicAuth(ctx, username, password)
}

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	return s.users.List(ctx)
}

func (s *service) Show(ctx context.Context, id string) (*model.User, error) {
	return s.users.Show(ctx, id)
}

func (s *service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return s.users.Create(ctx, user)
}

func (s *service) Update(ctx context.Context, user *model.User) (*model.User, error) {
	return s.users.Update(ctx, user)
}

func (s *service) Delete(ctx context.Context, name string) error {
	return s.users.Delete(ctx, name)
}

func (s *service) ListTeams(ctx context.Context, name string) ([]*model.TeamUser, error) {
	user, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return s.users.ListTeams(ctx, user.ID)
}

func (s *service) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	return s.users.AppendTeam(ctx, userID, teamID, perm)
}

func (s *service) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	return s.users.PermitTeam(ctx, userID, teamID, perm)
}

func (s *service) DropTeam(ctx context.Context, userID, teamID string) error {
	return s.users.DropTeam(ctx, userID, teamID)
}
