package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Pack struct {
			bun.BaseModel `bun:"table:packs"`

			ID        string    `bun:",pk,type:varchar(20)"`
			Slug      string    `bun:",unique,type:varchar(255)"`
			Name      string    `bun:",unique,type:varchar(255)"`
			Website   string    `bun:"type:varchar(255)"`
			Public    bool      `bun:"default:true"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Pack)(nil)).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Pack struct {
			bun.BaseModel `bun:"table:packs"`
		}

		_, err := db.NewDropTable().
			Model((*Pack)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
