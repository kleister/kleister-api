package store

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/gabriel-vasile/mimetype"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/uptrace/bun"
)

// Versions provides all database operations related to versions.
type Versions struct {
	client *Store
}

// List implements the listing of all versions.
func (s *Versions) List(ctx context.Context, mod *model.Mod, params model.ListParams) ([]*model.Version, int64, error) {
	records := make([]*model.Version, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("File").
		Relation("File.Version").
		Where("version.mod_id = ?", mod.ID)

	if val, ok := s.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// Show implements the details for a specific version.
func (s *Versions) Show(ctx context.Context, mod *model.Mod, name string) (*model.Version, error) {
	record := &model.Version{}

	q := s.client.handle.NewSelect().
		Model(record).
		Relation("File").
		Relation("File.Version").
		Where("version.mod_id = ?", mod.ID).
		Where("version.id = ? OR version.name = ?", name, name)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrVersionNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new version.
func (s *Versions) Create(ctx context.Context, mod *model.Mod, record *model.Version) error {
	if record.Name != "" {
		record.Name = slugify.New(
			slugify.Configuration{
				ReplacementMap: slugReplacements,
			},
		).Slugify(
			record.Name,
		)
	}

	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().
			Model(record).
			Exec(ctx); err != nil {
			return err
		}

		if record.FileUpload != nil {
			reader := bytes.NewReader(
				record.FileUpload.Data,
			)

			mtype, err := mimetype.DetectReader(
				reader,
			)

			if err != nil {
				return err
			}

			record.File = &model.VersionFile{
				Version:   record,
				VersionID: record.ID,
				Slug: strings.Join(
					[]string{
						secret.Generate(64),
						mtype.Extension(),
					},
					"",
				),
				ContentType: mtype.String(),
				MD5:         fileChecksum(record.FileUpload.Data),
			}

			if _, err := tx.NewInsert().
				Model(record.File).
				Exec(ctx); err != nil {
				return err
			}

			switch record.FileUpload.Encoding {
			case "base64":
				if err := s.client.upload.Upload(
					ctx,
					filepath.Join(
						"versions",
						mod.ID,
						record.File.Slug,
					),
					bytes.NewBuffer(record.FileUpload.Data),
				); err != nil {
					return err
				}

			default:
				return ErrInvalidUploadEncoding
			}

			record.FileUpload = nil
		}

		return nil
	})
}

// Update implements the update of an existing credential.
func (s *Versions) Update(ctx context.Context, mod *model.Mod, record *model.Version) error {
	if record.Name != "" {
		record.Name = slugify.New(
			slugify.Configuration{
				ReplacementMap: slugReplacements,
			},
		).Slugify(
			record.Name,
		)
	}

	if err := s.validate(ctx, record, true); err != nil {
		return err
	}

	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		q := tx.NewUpdate().
			Model(record).
			Where("mod_id = ?", mod.ID).
			Where("id = ?", record.ID)

		if _, err := q.Exec(ctx); err != nil {
			return err
		}

		if record.FileUpload != nil {
			if record.File != nil {
				if err := s.client.upload.Delete(
					ctx,
					filepath.Join(
						"versions",
						mod.ID,
						record.File.Slug,
					),
					false,
				); err != nil {
					return err
				}

				if _, err := tx.NewDelete().
					Model((*model.VersionFile)(nil)).
					Where("version_id = ?", record.ID).
					Exec(ctx); err != nil {
					return err
				}
			}

			reader := bytes.NewReader(
				record.FileUpload.Data,
			)

			mtype, err := mimetype.DetectReader(
				reader,
			)

			if err != nil {
				return err
			}

			record.File = &model.VersionFile{
				Version:   record,
				VersionID: record.ID,
				Slug: strings.Join(
					[]string{
						secret.Generate(64),
						mtype.Extension(),
					},
					"",
				),
				ContentType: mtype.String(),
				MD5:         fileChecksum(record.FileUpload.Data),
			}

			if _, err := tx.NewInsert().
				Model(record.File).
				Exec(ctx); err != nil {
				return err
			}

			switch record.FileUpload.Encoding {
			case "base64":
				if err := s.client.upload.Upload(
					ctx,
					filepath.Join(
						"versions",
						mod.ID,
						record.File.Slug,
					),
					bytes.NewBuffer(record.FileUpload.Data),
				); err != nil {
					return err
				}

			default:
				return ErrInvalidUploadEncoding
			}

			record.FileUpload = nil
		}

		return nil
	})
}

// Delete implements the deletion of a credential.
func (s *Versions) Delete(ctx context.Context, mod *model.Mod, name string) error {
	record, err := s.Show(ctx, mod, name)

	if err != nil {
		return err
	}

	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if record.File != nil {
			if err := s.client.upload.Delete(
				ctx,
				filepath.Join(
					"versions",
					mod.ID,
					record.File.Slug,
				),
				false,
			); err != nil {
				return err
			}
		}

		q := tx.NewDelete().
			Model((*model.Version)(nil)).
			Where("mod_id = ?", mod.ID).
			Where("id = ?", record.ID)

		if _, err := q.Exec(ctx); err != nil {
			return err
		}

		return nil
	})
}

// ListBuilds implements the listing of all builds for a version.
func (s *Versions) ListBuilds(ctx context.Context, _ *model.Mod, version *model.Version, params model.ListParams) ([]*model.BuildVersion, int64, error) {
	records := make([]*model.BuildVersion, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Version").
		Relation("Build").
		Where("version_id = ?", version.ID)

	if val, ok := s.client.Builds.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// AttachBuild implements the attachment of an version to a build.
func (s *Versions) AttachBuild(ctx context.Context, _ *model.Mod, version *model.Version, params model.BuildVersionParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	assigned, err := s.isBuildAssigned(ctx, build.ID, version.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	record := &model.BuildVersion{
		VersionID: version.ID,
		BuildID:   build.ID,
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropBuild implements the removal of an version from a build.
func (s *Versions) DropBuild(ctx context.Context, _ *model.Mod, version *model.Version, params model.BuildVersionParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	unassigned, err := s.isBuildUnassigned(ctx, build.ID, version.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.BuildVersion)(nil)).
		Where("version_id = ? AND build_id = ?", version.ID, build.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Versions) isBuildAssigned(ctx context.Context, buildID, versionID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.BuildVersion)(nil)).
		Where("build_id = ? AND version_id = ?", buildID, versionID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Versions) isBuildUnassigned(ctx context.Context, buildID, versionID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.BuildVersion)(nil)).
		Where("build_id = ? AND version_id = ?", buildID, versionID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}

func (s *Versions) validate(ctx context.Context, record *model.Version, existing bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID, record.ModID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if !existing {
		if err := validation.Validate(
			record.FileUpload,
			validation.Required,
		); err != nil {
			errs.Errors = append(errs.Errors, validate.Error{
				Field: "upload",
				Error: err,
			})
		}
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (s *Versions) uniqueValueIsPresent(ctx context.Context, key, id, modID string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.Version)(nil)).
			Where("mod_id = ?", modID).
			Where("? = ?", bun.Ident(key), val)

		if id != "" {
			q = q.Where(
				"id != ?",
				id,
			)
		}

		exists, err := q.Exists(ctx)

		if err != nil {
			return err
		}

		if exists {
			return errors.New("is already taken")
		}

		return nil
	}
}

// ValidSort validates the given sorting column.
func (s *Versions) ValidSort(val string) (string, bool) {
	if val == "" {
		return "version.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"name":    "version.name",
		"public":  "version.public",
		"created": "version.created_at",
		"updated": "version.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "version.name", true
}
