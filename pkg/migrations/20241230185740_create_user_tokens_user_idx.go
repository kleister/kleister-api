package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserToken struct {
			bun.BaseModel `bun:"table:user_tokens"`

			ID     string `bun:",pk,type:varchar(20)"`
			UserID string `bun:"type:varchar(20)"`
		}

		_, err := db.NewCreateIndex().
			Model((*UserToken)(nil)).
			Index("user_tokens_user_id_idx").
			Column("user_id").
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserToken struct {
			bun.BaseModel `bun:"table:user_tokens"`
		}

		_, err := db.NewDropIndex().
			Model((*UserToken)(nil)).
			IfExists().
			Index("user_tokens_user_id_idx").
			Exec(ctx)

		return err
	})
}
