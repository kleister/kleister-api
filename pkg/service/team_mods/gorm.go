package teammods

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	modsService "github.com/kleister/kleister-api/pkg/service/mods"
	teamsService "github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
	teams     teamsService.Service
	mods      modsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
	teams teamsService.Service,
	mods modsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		config: cfg,
		teams:  teams,
		mods:   mods,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.TeamModParams) ([]*model.TeamMod, int64, error) {
	counter := int64(0)
	records := make([]*model.TeamMod, 0)
	q := s.query(ctx).Debug()

	switch {
	case params.TeamID != "" && params.ModID == "":
		team, err := s.teamID(ctx, params.TeamID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"team_id = ?",
			team,
		)

		if val, ok := s.validModSort(params.Sort); ok {
			q = q.Order(strings.Join(
				[]string{
					val,
					sortOrder(params.Order),
				},
				" ",
			))
		}
	case params.ModID != "" && params.TeamID == "":
		mod, err := s.modID(ctx, params.ModID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"mod_id = ?",
			mod,
		)

		if val, ok := s.validTeamSort(params.Sort); ok {
			q = q.Order(strings.Join(
				[]string{
					val,
					sortOrder(params.Order),
				},
				" ",
			))
		}
	default:
		return nil, counter, ErrInvalidListParams
	}

	// if params.Search != "" {
	// 	opts := queryparser.Options{
	// 		CutFn: searchCut,
	// 		Allowed: []string{},
	// 	}

	// 	parser := queryparser.New(
	// 		params.Search,
	// 		opts,
	// 	).Parse()

	// 	for _, name := range opts.Allowed {
	// 		if parser.Has(name) {

	// 			q = q.Where(
	// 				fmt.Sprintf(
	// 					"%s LIKE ?",
	// 					name,
	// 				),
	// 				strings.ReplaceAll(
	// 					parser.GetOne(name),
	// 					"*",
	// 					"%",
	// 				),
	// 			)
	// 		}
	// 	}
	// }

	if err := q.Count(
		&counter,
	).Error; err != nil {
		return nil, counter, err
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, counter, err
	}

	return records, counter, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, params model.TeamModParams) error {
	team, err := s.teamID(ctx, params.TeamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, params.ModID)
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
		Perm:   params.Perm,
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
func (s *GormService) Permit(ctx context.Context, params model.TeamModParams) error {
	team, err := s.teamID(ctx, params.TeamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, params.ModID)
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
	record.Perm = params.Perm

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
func (s *GormService) Drop(ctx context.Context, params model.TeamModParams) error {
	team, err := s.teamID(ctx, params.TeamID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, params.ModID)
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
		"team_id = ? AND mod_id = ?",
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
		"team_id = ? AND mod_id = ?",
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
	).Joins(
		"Team",
	).Preload(
		"Mod",
	).Joins(
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

func (s *GormService) validTeamSort(val string) (string, bool) {
	if val == "" {
		return "Team.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug": "Team.slug",
		"name": "Team.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Team.name", true
}

func (s *GormService) validModSort(val string) (string, bool) {
	if val == "" {
		return "Mod.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug": "Mod.slug",
		"name": "Mod.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Mod.name", true
}
