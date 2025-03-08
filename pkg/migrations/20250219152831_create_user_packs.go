package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type UserPack struct {
			bun.BaseModel `bun:"table:user_packs"`

			UserID    string    `bun:",pk,type:varchar(20)"`
			PackID    string    `bun:",pk,type:varchar(20)"`
			Perm      string    `bun:"type:varchar(32)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*UserPack)(nil)).
			WithForeignKeys().
			ForeignKey(`(user_id) REFERENCES users (id) ON DELETE CASCADE`).
			ForeignKey(`(pack_id) REFERENCES packs (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type UserPack struct {
			bun.BaseModel `bun:"table:user_packs"`
		}

		_, err := db.NewDropTable().
			Model((*UserPack)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
