package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserToken struct {
			bun.BaseModel `bun:"table:user_tokens"`

			ID        string    `bun:",pk,type:varchar(20)"`
			UserID    string    `bun:"type:varchar(20)"`
			Kind      string    `bun:"type:varchar(128)"`
			Token     string    `bun:",unique,type:varchar(128)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*UserToken)(nil)).
			WithForeignKeys().
			ForeignKey(`(user_id) REFERENCES users (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserToken struct {
			bun.BaseModel `bun:"table:user_tokens"`
		}

		_, err := db.NewDropTable().
			Model((*UserToken)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
