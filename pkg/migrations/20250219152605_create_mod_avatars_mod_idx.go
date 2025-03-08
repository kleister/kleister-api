package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type ModAvatar struct {
			bun.BaseModel `bun:"table:mod_avatars"`

			ID    string `bun:",pk,type:varchar(20)"`
			ModID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*ModAvatar)(nil)).
			Index("mod_avatars_mod_id_idx").
			Column("mod_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type ModAvatar struct {
			bun.BaseModel `bun:"table:mod_avatars"`
		}

		_, err := db.NewDropIndex().
			Model((*ModAvatar)(nil)).
			IfExists().
			Index("mod_avatars_mod_id_idx").
			Exec(ctx)

		return err
	})
}
