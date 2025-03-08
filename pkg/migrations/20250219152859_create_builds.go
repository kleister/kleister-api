package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		type Build struct {
			bun.BaseModel `bun:"table:builds"`

			ID          string    `bun:",pk,type:varchar(20)"`
			PackID      string    `bun:"type:varchar(20)"`
			Name        string    `bun:"type:varchar(255)"`
			MinecraftID *string   `bun:"type:varchar(20)"`
			ForgeID     *string   `bun:"type:varchar(20)"`
			NeoforgeID  *string   `bun:"type:varchar(20)"`
			QuiltID     *string   `bun:"type:varchar(20)"`
			FabricID    *string   `bun:"type:varchar(20)"`
			Java        string    `bun:"type:varchar(255)"`
			Memory      string    `bun:"type:varchar(255)"`
			Latest      bool      `bun:"default:false"`
			Recommended bool      `bun:"default:false"`
			Public      bool      `bun:"default:true"`
			CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
			UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		}

		_, err := db.NewCreateTable().
			Model((*Build)(nil)).
			WithForeignKeys().
			ForeignKey(`(pack_id) REFERENCES packs (id) ON DELETE CASCADE`).
			ForeignKey(`(minecraft_id) REFERENCES minecrafts (id) ON DELETE CASCADE`).
			ForeignKey(`(forge_id) REFERENCES forges (id) ON DELETE CASCADE`).
			ForeignKey(`(neoforge_id) REFERENCES neoforges (id) ON DELETE CASCADE`).
			ForeignKey(`(quilt_id) REFERENCES quilts (id) ON DELETE CASCADE`).
			ForeignKey(`(fabric_id) REFERENCES fabrics (id) ON DELETE CASCADE`).
			Exec(ctx)

		return err
	}, func(ctx context.Context, db *bun.DB) error {
		type Build struct {
			bun.BaseModel `bun:"table:builds"`
		}

		_, err := db.NewDropTable().
			Model((*Build)(nil)).
			IfExists().
			Exec(ctx)

		return err
	})
}
