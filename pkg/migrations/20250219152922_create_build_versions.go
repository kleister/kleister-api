package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type BuildVersion struct {
			bun.BaseModel `bun:"table:build_versions"`

			BuildID   string    `bun:",pk,type:varchar(20)"`
			VersionID string    `bun:",pk,type:varchar(20)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*BuildVersion)(nil)).
			WithForeignKeys().
			ForeignKey(`(build_id) REFERENCES builds (id) ON DELETE CASCADE`).
			ForeignKey(`(version_id) REFERENCES versions (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type BuildVersion struct {
			bun.BaseModel `bun:"table:build_versions"`
		}

		_, err := db.NewDropTable().
			Model((*BuildVersion)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
