package gormdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// Packs implements packs.Store interface.
type Packs struct {
	client *gormdbStore
}

// List implements List from packs.Store interface.
func (p *Packs) List(ctx context.Context) ([]*model.Pack, error) {
	records := make([]*model.Pack, 0)

	err := p.client.handle.Order(
		"name ASC",
	// ).Preload(
	// 	"Users",
	// ).Preload(
	// 	"Users.Pack",
	// ).Preload(
	// 	"Users.User",
	).Find(
		&records,
	).Error

	return records, err
}

// Show implements Show from packs.Store interface.
func (p *Packs) Show(ctx context.Context, name string) (*model.Pack, error) {
	record := &model.Pack{}

	err := p.client.handle.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	// ).Preload(
	// 	"Users",
	// ).Preload(
	// 	"Users.Pack",
	// ).Preload(
	// 	"Users.User",
	).First(
		record,
	).Error

	if err == gorm.ErrRecordNotFound {
		return record, packs.ErrNotFound
	}

	return record, err
}

// Create implements Create from packs.Store interface.
func (p *Packs) Create(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	tx := p.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if pack.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				pack.Slug = slugify.Slugify(pack.Name)
			} else {
				pack.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", pack.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				pack.Slug,
			).First(
				&model.Pack{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	pack.ID = uuid.New().String()

	fmt.Printf("%+v\n", pack)

	if err := p.validateCreate(pack); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(pack).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pack, nil
}

// Update implements Update from packs.Store interface.
func (p *Packs) Update(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	tx := p.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if pack.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				pack.Slug = slugify.Slugify(pack.Name)
			} else {
				pack.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", pack.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				pack.Slug,
			).Not(
				"id",
				pack.ID,
			).First(
				&model.Pack{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	if err := p.validateUpdate(pack); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(pack).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pack, nil
}

// Delete implements Delete from packs.Store interface.
func (p *Packs) Delete(ctx context.Context, name string) error {
	tx := p.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.Pack{}

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
			return packs.ErrNotFound
		}

		return err
	}

	// if err := tx.Where(
	// 	"pack_id = ?",
	// 	record.ID,
	// ).Delete(
	// 	&model.UserPack{},
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

func (p *Packs) validateCreate(record *model.Pack) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if p.uniqueValueIsPresent("slug", record.Slug, record.ID) {
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

	if p.uniqueValueIsPresent("name", record.Name, record.ID) {
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

func (p *Packs) validateUpdate(record *model.Pack) error {
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

	if p.uniqueValueIsPresent("slug", record.Slug, record.ID) {
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

	if p.uniqueValueIsPresent("name", record.Name, record.ID) {
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

func (p *Packs) uniqueValueIsPresent(key, val, id string) bool {
	res := p.client.handle.Where(
		fmt.Sprintf("%s = ?", key),
		val,
	).Not(
		"id = ?",
		id,
	).Find(
		&model.Pack{},
	)

	return res.RowsAffected != 0
}
