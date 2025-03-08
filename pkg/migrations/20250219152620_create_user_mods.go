package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserMod struct {
			bun.BaseModel `bun:"table:user_mods"`

			UserID    string    `bun:",pk,type:varchar(20)"`
			ModID     string    `bun:",pk,type:varchar(20)"`
			Perm      string    `bun:"type:varchar(32)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*UserMod)(nil)).
			WithForeignKeys().
			ForeignKey(`(user_id) REFERENCES users (id) ON DELETE CASCADE`).
			ForeignKey(`(mod_id) REFERENCES mods (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserMod struct {
			bun.BaseModel `bun:"table:user_mods"`
		}

		_, err := db.NewDropTable().
			Model((*UserMod)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
