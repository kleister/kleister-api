package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type User struct {
			bun.BaseModel `bun:"table:users"`

			ID        string    `bun:",pk,type:varchar(20)"`
			Scim      string    `bun:"type:varchar(255)"`
			Username  string    `bun:",unique,type:varchar(255)"`
			Hashword  string    `bun:"type:varchar(255)"`
			Email     string    `bun:"type:varchar(255)"`
			Fullname  string    `bun:"type:varchar(255)"`
			Active    bool      `bun:"default:false"`
			Admin     bool      `bun:"default:false"`
			CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*User)(nil)).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type User struct {
			bun.BaseModel `bun:"table:users"`
		}

		_, err := db.NewDropTable().
			Model((*User)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
