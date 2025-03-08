package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type PackAvatar struct {
			bun.BaseModel `bun:"table:pack_avatars"`

			ID        string    `bun:",pk,type:varchar(20)"`
			PackID    string    `bun:"type:varchar(20)"`
			Slug      string    `bun:"type:varchar(255)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*PackAvatar)(nil)).
			WithForeignKeys().
			ForeignKey(`(pack_id) REFERENCES packs (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type PackAvatar struct {
			bun.BaseModel `bun:"table:pack_avatars"`
		}

		_, err := db.NewDropTable().
			Model((*PackAvatar)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
