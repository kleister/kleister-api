package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Version struct {
			bun.BaseModel `bun:"table:versions"`

			ID    string `bun:",pk,type:varchar(20)"`
			ModID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*Version)(nil)).
			Index("versions_mod_id_idx").
			Column("mod_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Version struct {
			bun.BaseModel `bun:"table:versions"`
		}

		_, err := db.NewDropIndex().
			Model((*Version)(nil)).
			IfExists().
			Index("versions_mod_id_idx").
			Exec(ctx)

		return err
	})
}
