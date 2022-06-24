package gormdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// Mods implements mods.Store interface.
type Mods struct {
	client *gormdbStore
}

// List implements List from mods.Store interface.
func (m *Mods) List(ctx context.Context) ([]*model.Mod, error) {
	records := make([]*model.Mod, 0)

	err := m.client.handle.Order(
		"name ASC",
	// ).Preload(
	// 	"Users",
	// ).Preload(
	// 	"Users.Mod",
	// ).Preload(
	// 	"Users.User",
	).Find(
		&records,
	).Error

	return records, err
}

// Show implements Show from mods.Store interface.
func (m *Mods) Show(ctx context.Context, name string) (*model.Mod, error) {
	record := &model.Mod{}

	err := m.client.handle.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	// ).Preload(
	// 	"Users",
	// ).Preload(
	// 	"Users.Mod",
	// ).Preload(
	// 	"Users.User",
	).First(
		record,
	).Error

	if err == gorm.ErrRecordNotFound {
		return record, mods.ErrNotFound
	}

	return record, err
}

// Create implements Create from mods.Store interface.
func (m *Mods) Create(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := m.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if mod.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				mod.Slug = slugify.Slugify(mod.Name)
			} else {
				mod.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", mod.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				mod.Slug,
			).First(
				&model.Mod{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	mod.ID = uuid.New().String()

	fmt.Printf("%+v\n", mod)

	if err := m.validateCreate(mod); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(mod).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return mod, nil
}

// Update implements Update from mods.Store interface.
func (m *Mods) Update(ctx context.Context, mod *model.Mod) (*model.Mod, error) {
	tx := m.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if mod.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				mod.Slug = slugify.Slugify(mod.Name)
			} else {
				mod.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", mod.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				mod.Slug,
			).Not(
				"id",
				mod.ID,
			).First(
				&model.Mod{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	if err := m.validateUpdate(mod); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(mod).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return mod, nil
}

// Delete implements Delete from mods.Store interface.
func (m *Mods) Delete(ctx context.Context, name string) error {
	tx := m.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.Mod{}

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
			return mods.ErrNotFound
		}

		return err
	}

	// if err := tx.Where(
	// 	"mod_id = ?",
	// 	record.ID,
	// ).Delete(
	// 	&model.UserMod{},
	// ).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if err := tx.Delete(
	// 	record,
	// ).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	return tx.Commit().Error
}

func (m *Mods) validateCreate(record *model.Mod) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if m.uniqueValueIsPresent("slug", record.Slug, record.ID) {
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

	if m.uniqueValueIsPresent("name", record.Name, record.ID) {
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

	if m.uniqueValueIsPresent("slug", record.Slug, record.ID) {
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

	if m.uniqueValueIsPresent("name", record.Name, record.ID) {
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
	res := m.client.handle.Where(
		fmt.Sprintf("%s = ?", key),
		val,
	).Not(
		"id = ?",
		id,
	).Find(
		&model.Mod{},
	)

	return res.RowsAffected != 0
}
