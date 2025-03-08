package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserAuth struct {
			bun.BaseModel `bun:"table:user_auths"`

			ID        string    `bun:",pk,type:varchar(20)"`
			UserID    string    `bun:"type:varchar(20)"`
			Provider  string    `bun:"type:varchar(255)"`
			Ref       string    `bun:"type:varchar(255)"`
			Login     string    `bun:"type:varchar(255)"`
			Email     string    `bun:"type:varchar(255)"`
			Name      string    `bun:"type:varchar(255)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*UserAuth)(nil)).
			WithForeignKeys().
			ForeignKey(`(user_id) REFERENCES users (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserAuth struct {
			bun.BaseModel `bun:"table:user_auths"`
		}

		_, err := db.NewDropTable().
			Model((*UserAuth)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
