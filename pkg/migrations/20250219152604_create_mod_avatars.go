package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type ModAvatar struct {
			bun.BaseModel `bun:"table:mod_avatars"`

			ID        string    `bun:",pk,type:varchar(20)"`
			ModID     string    `bun:"type:varchar(20)"`
			Slug      string    `bun:"type:varchar(255)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*ModAvatar)(nil)).
			WithForeignKeys().
			ForeignKey(`(mod_id) REFERENCES mods (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type ModAvatar struct {
			bun.BaseModel `bun:"table:mod_avatars"`
		}

		_, err := db.NewDropTable().
			Model((*ModAvatar)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
