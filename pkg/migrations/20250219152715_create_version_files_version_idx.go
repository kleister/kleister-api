package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type VersionFile struct {
			bun.BaseModel `bun:"table:version_files"`

			ID        string `bun:",pk,type:varchar(20)"`
			VersionID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*VersionFile)(nil)).
			Index("version_files_version_id_idx").
			Column("mod_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type VersionFile struct {
			bun.BaseModel `bun:"table:version_files"`
		}

		_, err := db.NewDropIndex().
			Model((*VersionFile)(nil)).
			IfExists().
			Index("version_files_version_id_idx").
			Exec(ctx)

		return err
	})
}
