package gormdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// Teams implements teams.Store interface.
type Teams struct {
	client *gormdbStore
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	err := t.client.handle.Order(
		"name ASC",
	).Preload(
		"Users",
	).Preload(
		"Users.Team",
	).Preload(
		"Users.User",
	).Find(
		&records,
	).Error

	return records, err
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	err := t.client.handle.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Preload(
		"Users",
	).Preload(
		"Users.Team",
	).Preload(
		"Users.User",
	).First(
		record,
	).Error

	if err == gorm.ErrRecordNotFound {
		return record, teams.ErrNotFound
	}

	return record, err
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				team.Slug,
			).First(
				&model.Team{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	team.ID = uuid.New().String()

	fmt.Printf("%+v\n", team)

	if err := t.validateCreate(team); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(team).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				team.Slug,
			).Not(
				"id",
				team.ID,
			).First(
				&model.Team{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	if err := t.validateUpdate(team); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(team).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.Team{}

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error; err != nil {
		tx.Rollback()

		if err == gorm.ErrRecordNotFound {
			return teams.ErrNotFound
		}

		return err
	}

	if err := tx.Where(
		"team_id = ?",
		record.ID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	err := t.client.handle.Where(
		"team_id = ?",
		id,
	).Model(
		&model.TeamUser{},
	).Preload(
		"Team",
	).Preload(
		"User",
	).Find(
		&records,
	).Error

	return records, err
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	if t.isAssignedToUser(teamID, userID) {
		return teams.ErrAlreadyAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{
		TeamID: teamID,
		UserID: userID,
		Perm:   perm,
	}

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	if t.isUnassignedFromUser(teamID, userID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{}
	record.Perm = perm

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Model(
		&model.TeamUser{},
	).Updates(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	if t.isUnassignedFromUser(teamID, userID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListMods implements ListMods from teams.Store interface.
func (t *Teams) ListMods(ctx context.Context, id string) ([]*model.TeamMod, error) {
	records := make([]*model.TeamMod, 0)

	err := t.client.handle.Where(
		"team_id = ?",
		id,
	).Model(
		&model.TeamMod{},
	).Preload(
		"Team",
	).Preload(
		"Mod",
	).Find(
		&records,
	).Error

	return records, err
}

// AppendMod implements AppendMod from teams.Store interface.
func (t *Teams) AppendMod(ctx context.Context, teamID, modID, perm string) error {
	if t.isAssignedToMod(teamID, modID) {
		return teams.ErrAlreadyAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamMod{
		TeamID: teamID,
		ModID:  modID,
		Perm:   perm,
	}

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PermitMod implements PermitMod from teams.Store interface.
func (t *Teams) PermitMod(ctx context.Context, teamID, modID, perm string) error {
	if t.isUnassignedFromMod(teamID, modID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamMod{}
	record.Perm = perm

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(
		"team_id = ? AND mod_id = ?",
		teamID,
		modID,
	).Model(
		&model.TeamMod{},
	).Updates(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DropMod implements DropMod from teams.Store interface.
func (t *Teams) DropMod(ctx context.Context, teamID, modID string) error {
	if t.isUnassignedFromMod(teamID, modID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(
		"team_id = ? AND mod_id = ?",
		teamID,
		modID,
	).Delete(
		&model.TeamMod{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListPacks implements ListPacks from teams.Store interface.
func (t *Teams) ListPacks(ctx context.Context, id string) ([]*model.TeamPack, error) {
	records := make([]*model.TeamPack, 0)

	err := t.client.handle.Where(
		"team_id = ?",
		id,
	).Model(
		&model.TeamPack{},
	).Preload(
		"Team",
	).Preload(
		"Pack",
	).Find(
		&records,
	).Error

	return records, err
}

// AppendPack implements AppendPack from teams.Store interface.
func (t *Teams) AppendPack(ctx context.Context, teamID, packID, perm string) error {
	if t.isAssignedToPack(teamID, packID) {
		return teams.ErrAlreadyAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamPack{
		TeamID: teamID,
		PackID: packID,
		Perm:   perm,
	}

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PermitPack implements PermitPack from teams.Store interface.
func (t *Teams) PermitPack(ctx context.Context, teamID, packID, perm string) error {
	if t.isUnassignedFromPack(teamID, packID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamPack{}
	record.Perm = perm

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Model(
		&model.TeamPack{},
	).Updates(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DropPack implements DropPack from teams.Store interface.
func (t *Teams) DropPack(ctx context.Context, teamID, packID string) error {
	if t.isUnassignedFromPack(teamID, packID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Delete(
		&model.TeamPack{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (t *Teams) isAssignedToUser(teamID, userID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected != 0
}

func (t *Teams) isUnassignedFromUser(teamID, userID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected == 0
}

func (t *Teams) isAssignedToMod(teamID, modID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND mod_id = ?",
		teamID,
		modID,
	).Find(
		&model.TeamMod{},
	)

	return res.RowsAffected != 0
}

func (t *Teams) isUnassignedFromMod(teamID, modID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND mod_id = ?",
		teamID,
		modID,
	).Find(
		&model.TeamMod{},
	)

	return res.RowsAffected == 0
}

func (t *Teams) isAssignedToPack(teamID, packID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Find(
		&model.TeamPack{},
	)

	return res.RowsAffected != 0
}

func (t *Teams) isUnassignedFromPack(teamID, packID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND pack_id = ?",
		teamID,
		packID,
	).Find(
		&model.TeamPack{},
	)

	return res.RowsAffected == 0
}

func (t *Teams) validateCreate(record *model.Team) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validateUpdate(record *model.Team) error {
	errs := validate.Errors{}

	if ok := govalidator.IsUUIDv4(record.ID); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "id",
			Error: fmt.Errorf("is not a valid uuid v4"),
		})
	}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validatePerm(perm string) error {
	if ok := govalidator.IsIn(perm, "user", "admin", "owner"); !ok {
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

func (t *Teams) uniqueValueIsPresent(key, val, id string) bool {
	res := t.client.handle.Where(
		fmt.Sprintf("%s = ?", key),
		val,
	).Not(
		"id = ?",
		id,
	).Find(
		&model.Team{},
	)

	return res.RowsAffected != 0
}
