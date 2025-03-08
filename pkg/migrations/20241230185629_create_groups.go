package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Group struct {
			bun.BaseModel `bun:"table:groups"`

			ID        string    `bun:",pk,type:varchar(20)"`
			Scim      string    `bun:"type:varchar(255)"`
			Slug      string    `bun:",unique,type:varchar(255)"`
			Name      string    `bun:"type:varchar(255)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Group)(nil)).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Group struct {
			bun.BaseModel `bun:"table:groups"`
		}

		_, err := db.NewDropTable().
			Model((*Group)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
