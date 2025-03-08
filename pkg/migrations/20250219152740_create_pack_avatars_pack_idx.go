package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type PackAvatar struct {
			bun.BaseModel `bun:"table:pack_avatars"`

			ID     string `bun:",pk,type:varchar(20)"`
			PackID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*PackAvatar)(nil)).
			Index("pack_avatars_pack_id_idx").
			Column("pack_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type PackAvatar struct {
			bun.BaseModel `bun:"table:pack_avatars"`
		}

		_, err := db.NewDropIndex().
			Model((*PackAvatar)(nil)).
			IfExists().
			Index("pack_avatars_pack_id_idx").
			Exec(ctx)

		return err
	})
}
