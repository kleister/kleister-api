package teams

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
)

var (
	// ErrNotFound is returned when a team was not found.
	ErrNotFound = errors.New("team not found")

	// ErrAlreadyAssigned is returned when a team is already assigned.
	ErrAlreadyAssigned = errors.New("team is already assigned")

	// ErrNotAssigned is returned when a team is not assigned.
	ErrNotAssigned = errors.New("team is not assigned")
)

// Service handles all interactions with teams.
type Service interface {
	List(context.Context) ([]*model.Team, error)
	Show(context.Context, string) (*model.Team, error)
	Create(context.Context, *model.Team) (*model.Team, error)
	Update(context.Context, *model.Team) (*model.Team, error)
	Delete(context.Context, string) error

	ListUsers(context.Context, string) ([]*model.TeamUser, error)
	AppendUser(context.Context, string, string, string) error
	PermitUser(context.Context, string, string, string) error
	DropUser(context.Context, string, string) error

	ListMods(context.Context, string) ([]*model.TeamMod, error)
	AppendMod(context.Context, string, string, string) error
	PermitMod(context.Context, string, string, string) error
	DropMod(context.Context, string, string) error

	ListPacks(context.Context, string) ([]*model.TeamPack, error)
	AppendPack(context.Context, string, string, string) error
	PermitPack(context.Context, string, string, string) error
	DropPack(context.Context, string, string) error
}

// Store defines the interface to persist teams.
type Store interface {
	List(context.Context) ([]*model.Team, error)
	Show(context.Context, string) (*model.Team, error)
	Create(context.Context, *model.Team) (*model.Team, error)
	Update(context.Context, *model.Team) (*model.Team, error)
	Delete(context.Context, string) error

	ListUsers(context.Context, string) ([]*model.TeamUser, error)
	AppendUser(context.Context, string, string, string) error
	PermitUser(context.Context, string, string, string) error
	DropUser(context.Context, string, string) error

	ListMods(context.Context, string) ([]*model.TeamMod, error)
	AppendMod(context.Context, string, string, string) error
	PermitMod(context.Context, string, string, string) error
	DropMod(context.Context, string, string) error

	ListPacks(context.Context, string) ([]*model.TeamPack, error)
	AppendPack(context.Context, string, string, string) error
	PermitPack(context.Context, string, string, string) error
	DropPack(context.Context, string, string) error
}

type service struct {
	teams Store
}

// NewService returns a Service that handles all interactions with teams.
func NewService(teams Store) Service {
	return &service{
		teams: teams,
	}
}

func (s *service) List(ctx context.Context) ([]*model.Team, error) {
	return s.teams.List(ctx)
}

func (s *service) Show(ctx context.Context, id string) (*model.Team, error) {
	return s.teams.Show(ctx, id)
}

func (s *service) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	return s.teams.Create(ctx, team)
}

func (s *service) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	return s.teams.Update(ctx, team)
}

func (s *service) Delete(ctx context.Context, name string) error {
	return s.teams.Delete(ctx, name)
}

func (s *service) ListUsers(ctx context.Context, name string) ([]*model.TeamUser, error) {
	team, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return s.teams.ListUsers(ctx, team.ID)
}

func (s *service) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	return s.teams.AppendUser(ctx, teamID, userID, perm)
}

func (s *service) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	return s.teams.PermitUser(ctx, teamID, userID, perm)
}

func (s *service) DropUser(ctx context.Context, teamID, userID string) error {
	return s.teams.DropUser(ctx, teamID, userID)
}

func (s *service) ListMods(ctx context.Context, name string) ([]*model.TeamMod, error) {
	team, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return s.teams.ListMods(ctx, team.ID)
}

func (s *service) AppendMod(ctx context.Context, teamID, modID, perm string) error {
	return s.teams.AppendMod(ctx, teamID, modID, perm)
}

func (s *service) PermitMod(ctx context.Context, teamID, modID, perm string) error {
	return s.teams.PermitMod(ctx, teamID, modID, perm)
}

func (s *service) DropMod(ctx context.Context, teamID, modID string) error {
	return s.teams.DropMod(ctx, teamID, modID)
}

func (s *service) ListPacks(ctx context.Context, name string) ([]*model.TeamPack, error) {
	team, err := s.Show(ctx, name)

	if err != nil {
		return nil, err
	}

	return s.teams.ListPacks(ctx, team.ID)
}

func (s *service) AppendPack(ctx context.Context, teamID, packID, perm string) error {
	return s.teams.AppendPack(ctx, teamID, packID, perm)
}

func (s *service) PermitPack(ctx context.Context, teamID, packID, perm string) error {
	return s.teams.PermitPack(ctx, teamID, packID, perm)
}

func (s *service) DropPack(ctx context.Context, teamID, packID string) error {
	return s.teams.DropPack(ctx, teamID, packID)
}
