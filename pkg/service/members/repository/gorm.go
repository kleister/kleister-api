package repository

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
	teamsRepository "github.com/kleister/kleister-api/pkg/service/teams/repository"
	usersRepository "github.com/kleister/kleister-api/pkg/service/users/repository"
	"gorm.io/gorm"
)

// NewGormRepository initializes a new repository for GormDB.
func NewGormRepository(
	handle *gorm.DB,
	teams teamsRepository.TeamsRepository,
	users usersRepository.UsersRepository,
) *GormRepository {
	return &GormRepository{
		handle: handle,
		teams:  teams,
		users:  users,
	}
}

// GormRepository implements the MembersRepository interface.
type GormRepository struct {
	handle *gorm.DB
	teams  teamsRepository.TeamsRepository
	users  usersRepository.UsersRepository
}

// List implements the MembersRepository interface.
func (r *GormRepository) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	q := r.query()

	switch {
	case teamID != "" && userID == "":
		team, err := r.teamID(ctx, teamID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"team_id = ?",
			team,
		)
	case userID != "" && teamID == "":
		user, err := r.userID(ctx, userID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"user_id = ?",
			user,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.Member, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Append implements the MembersRepository interface.
func (r *GormRepository) Append(ctx context.Context, teamID, userID string) error {
	team, err := r.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	user, err := r.userID(ctx, userID)
	if err != nil {
		return err
	}

	if r.isAssigned(team, user) {
		return ErrIsAssigned
	}

	tx := r.handle.Begin()
	defer tx.Rollback()

	record := &model.Member{
		TeamID: team,
		UserID: user,
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the MembersRepository interface.
func (r *GormRepository) Drop(ctx context.Context, teamID, userID string) error {
	team, err := r.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	user, err := r.userID(ctx, userID)
	if err != nil {
		return err
	}

	if r.isUnassigned(team, user) {
		return ErrNotAssigned
	}

	tx := r.handle.Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		team,
		user,
	).Delete(
		&model.Member{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (r *GormRepository) teamID(ctx context.Context, id string) (string, error) {
	record, err := r.teams.Show(ctx, id)

	if err != nil {
		if errors.Is(err, teamsRepository.ErrTeamNotFound) {
			return "", ErrMemberNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (r *GormRepository) userID(ctx context.Context, id string) (string, error) {
	record, err := r.users.Show(ctx, id)

	if err != nil {
		if errors.Is(err, usersRepository.ErrUserNotFound) {
			return "", ErrMemberNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (r *GormRepository) isAssigned(teamID, userID string) bool {
	res := r.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.Member{},
	)

	return res.RowsAffected != 0
}

func (r *GormRepository) isUnassigned(teamID, userID string) bool {
	res := r.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.Member{},
	)

	return res.RowsAffected == 0
}

func (r *GormRepository) query() *gorm.DB {
	return r.handle.Model(
		&model.Member{},
	).Preload(
		"Team",
	).Preload(
		"User",
	)
}
