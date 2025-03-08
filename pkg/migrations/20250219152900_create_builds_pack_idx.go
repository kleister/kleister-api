package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Build struct {
			bun.BaseModel `bun:"table:builds"`

			ID     string `bun:",pk,type:varchar(20)"`
			PackID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*Build)(nil)).
			Index("builds_pack_id_idx").
			Column("pack_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Build struct {
			bun.BaseModel `bun:"table:builds"`
		}

		_, err := db.NewDropIndex().
			Model((*Build)(nil)).
			IfExists().
			Index("builds_pack_id_idx").
			Exec(ctx)

		return err
	})
}
