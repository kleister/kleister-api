package boltdb

import (
	"context"
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/validate"
)

// Teams implements teams.Store interface.
type Teams struct {
	client *botldbStore
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	err := t.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	for _, record := range records {
		users, err := t.ListUsers(ctx, record.ID)

		if err != nil {
			return records, err
		}

		record.Users = users
	}

	return records, nil
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	if err := t.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, teams.ErrNotFound
		}

		return nil, err
	}

	users, err := t.ListUsers(ctx, record.ID)

	if err != nil {
		return record, err
	}

	record.Users = users
	return record, nil
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", team.Slug),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.ID = uuid.New().String()
	team.UpdatedAt = time.Now().UTC()
	team.CreatedAt = time.Now().UTC()

	if err := t.validateCreate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", team.Slug),
					q.Not(
						q.Eq("ID", team.ID),
					),
				),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.UpdatedAt = time.Now().UTC()

	if err := t.validateUpdate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	record := &model.Team{}
	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return teams.ErrNotFound
		}

		return err
	}

	if err := tx.Select(
		q.Eq("TeamID", record.ID),
	).Delete(new(model.TeamUser)); err != nil {
		return err
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	if err := t.client.handle.Select(
		q.Eq("TeamID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		team := &model.Team{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.TeamID),
		).First(team); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Team = team

		user := &model.User{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.UserID),
		).First(user); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.User = user
	}

	return records, nil
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(new(model.TeamUser)); err == nil {
		return teams.ErrAlreadyAssigned
	}

	record := &model.TeamUser{
		TeamID:    teamID,
		UserID:    userID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListMods implements ListMods from teams.Store interface.
func (t *Teams) ListMods(ctx context.Context, id string) ([]*model.TeamMod, error) {
	records := make([]*model.TeamMod, 0)

	if err := t.client.handle.Select(
		q.Eq("TeamID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		team := &model.Team{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.TeamID),
		).First(team); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Team = team

		mod := &model.Mod{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.ModID),
		).First(mod); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Mod = mod
	}

	return records, nil
}

// AppendMod implements AppendMod from teams.Store interface.
func (t *Teams) AppendMod(ctx context.Context, teamID, modID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("ModID", modID),
		),
	).First(new(model.TeamMod)); err == nil {
		return teams.ErrAlreadyAssigned
	}

	record := &model.TeamMod{
		TeamID:    teamID,
		ModID:     modID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitMod implements PermitMod from teams.Store interface.
func (t *Teams) PermitMod(ctx context.Context, teamID, modID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamMod{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("ModID", modID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropMod implements DropMod from teams.Store interface.
func (t *Teams) DropMod(ctx context.Context, teamID, modID string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamMod{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("ModID", modID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListPacks implements ListPacks from teams.Store interface.
func (t *Teams) ListPacks(ctx context.Context, id string) ([]*model.TeamPack, error) {
	records := make([]*model.TeamPack, 0)

	if err := t.client.handle.Select(
		q.Eq("TeamID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		team := &model.Team{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.TeamID),
		).First(team); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Team = team

		pack := &model.Pack{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.PackID),
		).First(pack); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Pack = pack
	}

	return records, nil
}

// AppendPack implements AppendPack from teams.Store interface.
func (t *Teams) AppendPack(ctx context.Context, teamID, packID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("PackID", packID),
		),
	).First(new(model.TeamPack)); err == nil {
		return teams.ErrAlreadyAssigned
	}

	record := &model.TeamPack{
		TeamID:    teamID,
		PackID:    packID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitPack implements PermitPack from teams.Store interface.
func (t *Teams) PermitPack(ctx context.Context, teamID, packID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamPack{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("PackID", packID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropPack implements DropPack from teams.Store interface.
func (t *Teams) DropPack(ctx context.Context, teamID, packID string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamPack{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("PackID", packID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

func (t *Teams) validateCreate(record *model.Team) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
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

	if t.uniqueValueIsPresent("Name", record.Name, record.ID) {
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

	if t.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
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

	if t.uniqueValueIsPresent("Name", record.Name, record.ID) {
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
	if err := t.client.handle.Select(
		q.And(
			q.Eq(key, val),
			q.Not(
				q.Eq("ID", id),
			),
		),
	).First(new(model.Team)); err == storm.ErrNotFound {
		return false
	}

	return true
}
