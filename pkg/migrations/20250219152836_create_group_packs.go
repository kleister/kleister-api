package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type GroupPack struct {
			bun.BaseModel `bun:"table:group_packs"`

			GroupID   string    `bun:",pk,type:varchar(20)"`
			PackID    string    `bun:",pk,type:varchar(20)"`
			Perm      string    `bun:"type:varchar(32)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*GroupPack)(nil)).
			WithForeignKeys().
			ForeignKey(`(group_id) REFERENCES groups (id) ON DELETE CASCADE`).
			ForeignKey(`(pack_id) REFERENCES packs (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type GroupPack struct {
			bun.BaseModel `bun:"table:group_packs"`
		}

		_, err := db.NewDropTable().
			Model((*GroupPack)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
