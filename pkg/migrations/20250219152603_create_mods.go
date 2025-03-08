package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Mod struct {
			bun.BaseModel `bun:"table:mods"`

			ID          string    `bun:",pk,type:varchar(20)"`
			Slug        string    `bun:",unique,type:varchar(255)"`
			Name        string    `bun:",unique,type:varchar(255)"`
			Side        string    `bun:"type:varchar(64)"`
			Description string    `bun:"type:text"`
			Author      string    `bun:"type:varchar(255)"`
			Website     string    `bun:"type:varchar(255)"`
			Donate      string    `bun:"type:varchar(255)"`
			Public      bool      `bun:"default:true"`
			CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Mod)(nil)).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Mod struct {
			bun.BaseModel `bun:"table:mods"`
		}

		_, err := db.NewDropTable().
			Model((*Mod)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
