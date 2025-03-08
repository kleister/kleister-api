package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserAuth struct {
			bun.BaseModel `bun:"table:user_auths"`

			ID     string `bun:",pk,type:varchar(20)"`
			UserID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*UserAuth)(nil)).
			Index("user_auths_user_id_idx").
			Column("user_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserAuth struct {
			bun.BaseModel `bun:"table:user_auths"`
		}

		_, err := db.NewDropIndex().
			Model((*UserAuth)(nil)).
			IfExists().
			Index("user_auths_user_id_idx").
			Exec(ctx)

		return err
	})
}
