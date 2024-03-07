package members

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	teamsService "github.com/kleister/kleister-api/pkg/service/teams"
	usersService "github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
	teams  teamsService.Service
	users  usersService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	teams teamsService.Service,
	users usersService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		teams:  teams,
		users:  users,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	q := s.query(ctx)

	switch {
	case teamID != "" && userID == "":
		team, err := s.teamID(ctx, teamID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"team_id = ?",
			team,
		)
	case userID != "" && teamID == "":
		user, err := s.userID(ctx, userID)
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

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, teamID, userID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, team, user) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.Member{
		TeamID: team,
		UserID: user,
		Perm:   perm,
	}

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Permit implements the Service interface for database persistence.
func (s *GormService) Permit(ctx context.Context, teamID, userID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, user) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.Member{}
	record.Perm = perm

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		team,
		user,
	).Model(
		&model.Member{},
	).Updates(
		record,
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, teamID, userID string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, user) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
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

func (s *GormService) teamID(ctx context.Context, id string) (string, error) {
	record, err := s.teams.Show(ctx, id)

	if err != nil {
		if errors.Is(err, teamsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) userID(ctx context.Context, id string) (string, error) {
	record, err := s.users.Show(ctx, id)

	if err != nil {
		if errors.Is(err, usersService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) isAssigned(ctx context.Context, teamID, userID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.Member{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, teamID, userID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.Member{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.Member{},
	).Preload(
		"Team",
	).Preload(
		"User",
	)
}

func (s *GormService) validatePerm(perm string) error {
	if err := validation.Validate(
		perm,
		validation.In("user", "admin", "owner"),
	); err != nil {
		return validate.Errors{
			Errors: []validate.Error{
				{
					Field: "perm",
					Error: fmt.Errorf("invalid permission value"),
				},
			},
		}
	}

	return nil
}
