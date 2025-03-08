package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Fabric struct {
			bun.BaseModel `bun:"table:fabrics"`

			ID        string    `bun:",pk,type:varchar(20)"`
			Name      string    `bun:",unique,type:varchar(64)"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Fabric)(nil)).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Fabric struct {
			bun.BaseModel `bun:"table:fabrics"`
		}

		_, err := db.NewDropTable().
			Model((*Fabric)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
