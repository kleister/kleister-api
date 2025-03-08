package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type VersionFile struct {
			bun.BaseModel `bun:"table:version_files"`

			ID          string    `bun:",pk,type:varchar(20)"`
			VersionID   string    `bun:"type:varchar(20)"`
			Slug        string    `bun:"type:varchar(255)"`
			ContentType string    `bun:"type:varchar(255)"`
			MD5         string    `bun:"column:md5,type:varchar(255)"`
			CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*VersionFile)(nil)).
			WithForeignKeys().
			ForeignKey(`(version_id) REFERENCES versions (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type VersionFile struct {
			bun.BaseModel `bun:"table:version_files"`
		}

		_, err := db.NewDropTable().
			Model((*VersionFile)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
