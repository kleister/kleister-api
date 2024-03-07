package teamPacks

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	packsService "github.com/kleister/kleister-api/pkg/service/packs"
	teamsService "github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
	teams  teamsService.Service
	packs  packsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	teams teamsService.Service,
	packs packsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		teams:  teams,
		packs:  packs,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, teamID, packID string) ([]*model.TeamPack, error) {
	q := s.query(ctx)

	switch {
	case teamID != "" && packID == "":
		team, err := s.teamID(ctx, teamID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"team_id = ?",
			team,
		)
	case packID != "" && teamID == "":
		pack, err := s.packID(ctx, packID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"pack_id = ?",
			pack,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.TeamPack, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, teamID, packID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, team, pack) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.TeamPack{
		TeamID: team,
		PackID: pack,
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
func (s *GormService) Permit(ctx context.Context, teamID, packID, perm string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, pack) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.TeamPack{}
	record.Perm = perm

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Where(
		"team_id = ? AND pack_id = ?",
		team,
		pack,
	).Model(
		&model.TeamPack{},
	).Updates(
		record,
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, teamID, packID string) error {
	team, err := s.teamID(ctx, teamID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, team, pack) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"team_id = ? AND pack_id = ?",
		team,
		pack,
	).Delete(
		&model.TeamPack{},
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

func (s *GormService) packID(ctx context.Context, id string) (string, error) {
	record, err := s.packs.Show(ctx, id)

	if err != nil {
		if errors.Is(err, packsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) isAssigned(ctx context.Context, teamID, packID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Find(
		&model.TeamPack{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, teamID, packID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Find(
		&model.TeamPack{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.TeamPack{},
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
