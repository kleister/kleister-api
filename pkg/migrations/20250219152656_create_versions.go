package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Version struct {
			bun.BaseModel `bun:"table:versions"`

			ID        string    `bun:",pk,type:varchar(20)"`
			ModID     string    `bun:"type:varchar(20)"`
			Name      string    `bun:",unique,type:varchar(255)"`
			Public    bool      `bun:"default:true"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Version)(nil)).
			WithForeignKeys().
			ForeignKey(`(mod_id) REFERENCES mods (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Version struct {
			bun.BaseModel `bun:"table:versions"`
		}

		_, err := db.NewDropTable().
			Model((*Version)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
