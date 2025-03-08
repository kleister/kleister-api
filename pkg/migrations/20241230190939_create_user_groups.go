package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserGroup struct {
			bun.BaseModel `bun:"table:user_groups"`

			UserID    string    `bun:",pk,type:varchar(20)"`
			GroupID   string    `bun:",pk,type:varchar(20)"`
			Perm      string    `bun:"type:varchar(32)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*UserGroup)(nil)).
			WithForeignKeys().
			ForeignKey(`(user_id) REFERENCES users (id) ON DELETE CASCADE`).
			ForeignKey(`(group_id) REFERENCES groups (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserGroup struct {
			bun.BaseModel `bun:"table:user_groups"`
		}

		_, err := db.NewDropTable().
			Model((*UserGroup)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
