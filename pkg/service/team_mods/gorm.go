package teamMods

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	modsService "github.com/kleister/kleister-api/pkg/service/mods"
	teamsService "github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
	teams  teamsService.Service
	mods   modsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	teams teamsService.Service,
	mods modsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		teams:  teams,
		mods:   mods,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, teamID, modID string) ([]*model.TeamMod, error) {
	q := s.query(ctx)

	switch {
	case teamID != "" && modID == "":
		team, err := s.teamID(ctx, teamID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"team_id = ?",
			team,
		)
	case modID != "" && teamID == "":
		mod, err := s.modID(ctx, modID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"mod_id = ?",
			mod,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.TeamMod, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, teamID, modID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, team, mod) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.TeamMod{
		TeamID: team,
		ModID:  mod,
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
func (s *GormService) Permit(ctx context.Context, teamID, modID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, mod) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.TeamMod{}
	record.Perm = perm

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Where(
		"team_id = ? AND mod_id = ?",
		team,
		mod,
	).Model(
		&model.TeamMod{},
	).Updates(
		record,
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, teamID, modID string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, mod) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"team_id = ? AND mod_id = ?",
		team,
		mod,
	).Delete(
		&model.TeamMod{},
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

func (s *GormService) modID(ctx context.Context, id string) (string, error) {
	record, err := s.mods.Show(ctx, id)

	if err != nil {
		if errors.Is(err, modsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) isAssigned(ctx context.Context, teamID, modID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND user_id = ?",
		teamID,
		modID,
	).Find(
		&model.TeamMod{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, teamID, modID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND user_id = ?",
		teamID,
		modID,
	).Find(
		&model.TeamMod{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.TeamMod{},
	).Preload(
		"Team",
	).Preload(
		"Mod",
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
