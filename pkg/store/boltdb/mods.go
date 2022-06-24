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
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/validate"
)

// Mods implements mods.Store interface.
type Mods struct {
	client *botldbStore
}

// List implements List from mods.Store interface.
func (m *Mods) List(ctx context.Context) ([]*model.Mod, error) {
	records := make([]*model.Mod, 0)

	err := m.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	// for _, record := range records {
	// 	users, err := t.ListUsers(ctx, record.ID)

	// 	if err != nil {
	// 		return records, err
	// 	}

	// 	record.Users = users
	// }

	return records, nil
}

// Show implements Show from mods.Store interface.
func (m *Mods) Show(ctx context.Context, name string) (*model.Mod, error) {
	record := &model.Mod{}

	if err := m.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, mods.ErrNotFound
		}

		return nil, err
	}

	// users, err := t.ListUsers(ctx, record.ID)

	// if err != nil {
	// 	return record, err
	// }

	// record.Users = users
	return record, nil
}

// Create implements Create from mods.Store interface.
func (m *Mods) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx, err := m.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if mod.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				mod.Slug = slugify.Slugify(mod.Name)
			} else {
				mod.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", mod.Name, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", mod.Slug),
			).First(new(model.Mod)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	mod.ID = uuid.New().String()
	mod.UpdatedAt = time.Now().UTC()
	mod.CreatedAt = time.Now().UTC()

	if err := m.validateCreate(mod); err != nil {
		return nil, err
	}

	if err := tx.Save(mod); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mod, nil
}

// Update implements Update from mods.Store interface.
func (m *Mods) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx, err := m.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if mod.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				mod.Slug = slugify.Slugify(mod.Name)
			} else {
				mod.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", mod.Name, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", mod.Slug),
					q.Not(
						q.Eq("ID", mod.ID),
					),
				),
			).First(new(model.Mod)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	mod.UpdatedAt = time.Now().UTC()

	if err := m.validateUpdate(mod); err != nil {
		return nil, err
	}

	if err := tx.Save(mod); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mod, nil
}

// Delete implements Delete from mods.Store interface.
func (m *Mods) Delete(ctx context.Context, name string) error {
	tx, err := m.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	record := &model.Mod{}
	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return mods.ErrNotFound
		}

		return err
	}

	// if err := tx.Select(
	// 	q.Eq("ModID", record.ID),
	// ).Delete(new(model.UserMod)); err != nil {
	// 	return err
	// }

	// if err := tx.DeleteStruct(record); err != nil {
	// 	return err
	// }

	return tx.Commit()
}

func (m *Mods) validateCreate(record *model.Mod) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if m.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
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

	if m.uniqueValueIsPresent("Name", record.Name, record.ID) {
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

func (m *Mods) validateUpdate(record *model.Mod) error {
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

	if m.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
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

	if m.uniqueValueIsPresent("Name", record.Name, record.ID) {
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

func (m *Mods) uniqueValueIsPresent(key, val, id string) bool {
	if err := m.client.handle.Select(
		q.And(
			q.Eq(key, val),
			q.Not(
				q.Eq("ID", id),
			),
		),
	).First(new(model.Mod)); err == storm.ErrNotFound {
		return false
	}

	return true
}
