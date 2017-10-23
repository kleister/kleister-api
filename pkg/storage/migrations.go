package storage

import (
	// "time"

	// "github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
	// "gopkg.in/guregu/null.v3"
)

var (
	migrations = []*migrate.Migration{
	// {
	// 	ID: "201609011300",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type User struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Username  string `sql:"unique_index"`
	// 			Email     string `sql:"unique_index"`
	// 			Hash      string `sql:"unique_index"`
	// 			Hashword  string
	// 			Active    bool `sql:"default:false"`
	// 			Admin     bool `sql:"default:false"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&User{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("user")
	// 	},
	// },
	// {
	// 	ID: "201609011301",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Team struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Team{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team")
	// 	},
	// },
	// {
	// 	ID: "201609011302",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type TeamUser struct {
	// 			TeamID int64 `sql:"index"`
	// 			UserID int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&TeamUser{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team_user")
	// 	},
	// },
	// {
	// 	ID: "201609011303",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_user",
	// 		).AddForeignKey(
	// 			"team_id",
	// 			"team(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011304",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_user",
	// 		).AddForeignKey(
	// 			"user_id",
	// 			"user(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011305",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Forge struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			Minecraft string
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Forge{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("forge")
	// 	},
	// },
	// {
	// 	ID: "201609011306",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Minecraft struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			Type      string
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Minecraft{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("minecraft")
	// 	},
	// },
	// {
	// 	ID: "201609011307",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Pack struct {
	// 			ID            int64    `gorm:"primary_key"`
	// 			RecommendedID null.Int `sql:"index"`
	// 			LatestID      null.Int `sql:"index"`
	// 			Slug          string   `sql:"unique_index"`
	// 			Name          string   `sql:"unique_index"`
	// 			Website       string
	// 			Published     bool `sql:"default:false"`
	// 			Private       bool `sql:"default:false"`
	// 			CreatedAt     time.Time
	// 			UpdatedAt     time.Time
	// 		}

	// 		return engine.CreateTables(&Pack{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("pack")
	// 	},
	// },
	// {
	// 	ID: "201609011308",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type TeamPack struct {
	// 			TeamID int64 `sql:"index"`
	// 			PackID int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&TeamPack{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team_pack")
	// 	},
	// },
	// {
	// 	ID: "201609011309",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_pack",
	// 		).AddForeignKey(
	// 			"team_id",
	// 			"team(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011310",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_pack",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011311",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type UserPack struct {
	// 			UserID int64 `sql:"index"`
	// 			PackID int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&UserPack{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("user_pack")
	// 	},
	// },
	// {
	// 	ID: "201609011312",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_pack",
	// 		).AddForeignKey(
	// 			"user_id",
	// 			"user(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011313",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_pack",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011314",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackBackground struct {
	// 			ID          int64  `gorm:"primary_key"`
	// 			PackID      int64  `sql:"index"`
	// 			Slug        string `sql:"unique_index"`
	// 			ContentType string
	// 			MD5         string
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&PackBackground{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("pack_background")
	// 	},
	// },
	// {
	// 	ID: "201609011315",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"pack_background",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"CASCADE",
	// 			"CASCADE",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011316",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackIcon struct {
	// 			ID          int64  `gorm:"primary_key"`
	// 			PackID      int64  `sql:"index"`
	// 			Slug        string `sql:"unique_index"`
	// 			ContentType string
	// 			MD5         string
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&PackIcon{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("pack_icon")
	// 	},
	// },
	// {
	// 	ID: "201609011317",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"pack_icon",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"CASCADE",
	// 			"CASCADE",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011318",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackLogo struct {
	// 			ID          int64  `gorm:"primary_key"`
	// 			PackID      int64  `sql:"index"`
	// 			Slug        string `sql:"unique_index"`
	// 			ContentType string
	// 			MD5         string
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&PackLogo{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("pack_logo")
	// 	},
	// },
	// {
	// 	ID: "201609011319",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"pack_logo",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"CASCADE",
	// 			"CASCADE",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011320",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Build struct {
	// 			ID          int64    `gorm:"primary_key"`
	// 			PackID      int64    `sql:"index"`
	// 			MinecraftID null.Int `sql:"index"`
	// 			ForgeID     null.Int `sql:"index"`
	// 			Slug        string
	// 			Name        string
	// 			MinJava     string
	// 			MinMemory   string
	// 			Published   bool `sql:"default:false"`
	// 			Private     bool `sql:"default:false"`
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&Build{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("build")
	// 	},
	// },
	// {
	// 	ID: "201609011321",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"pack",
	// 		).AddForeignKey(
	// 			"recommended_id",
	// 			"build(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011322",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"pack",
	// 		).AddForeignKey(
	// 			"latest_id",
	// 			"build(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011323",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"build",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011324",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"build",
	// 		).AddForeignKey(
	// 			"minecraft_id",
	// 			"minecraft(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011325",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"build",
	// 		).AddForeignKey(
	// 			"forge_id",
	// 			"forge(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011326",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"build",
	// 		).AddUniqueIndex(
	// 			"uix_build_pack_id_slug",
	// 			"pack_id",
	// 			"slug",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"build",
	// 		).RemoveIndex(
	// 			"uix_build_pack_id_slug",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201609011327",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"build",
	// 		).AddUniqueIndex(
	// 			"uix_build_pack_id_name",
	// 			"pack_id",
	// 			"name",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"build",
	// 		).RemoveIndex(
	// 			"uix_build_pack_id_name",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201609011328",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Mod struct {
	// 			ID          int64  `gorm:"primary_key"`
	// 			Slug        string `sql:"unique_index"`
	// 			Name        string `sql:"unique_index"`
	// 			Description string `sql:"type:text"`
	// 			Author      string
	// 			Website     string
	// 			Donate      string
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&Mod{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("mod")
	// 	},
	// },
	// {
	// 	ID: "201609011329",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type TeamMod struct {
	// 			TeamID int64 `sql:"index"`
	// 			ModID  int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&TeamMod{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team_mod")
	// 	},
	// },
	// {
	// 	ID: "201609011330",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_mod",
	// 		).AddForeignKey(
	// 			"team_id",
	// 			"team(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011331",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_mod",
	// 		).AddForeignKey(
	// 			"mod_id",
	// 			"mod(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011332",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type UserMod struct {
	// 			UserID int64 `sql:"index"`
	// 			ModID  int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&UserMod{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("user_mod")
	// 	},
	// },
	// {
	// 	ID: "201609011333",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_mod",
	// 		).AddForeignKey(
	// 			"user_id",
	// 			"user(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011334",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_mod",
	// 		).AddForeignKey(
	// 			"mod_id",
	// 			"mod(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011335",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Version struct {
	// 			ID        int64 `gorm:"primary_key"`
	// 			ModID     int64 `sql:"index"`
	// 			Slug      string
	// 			Name      string
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Version{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("version")
	// 	},
	// },
	// {
	// 	ID: "201609011336",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"version",
	// 		).AddForeignKey(
	// 			"mod_id",
	// 			"mod(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011337",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"version",
	// 		).AddUniqueIndex(
	// 			"uix_version_mod_id_slug",
	// 			"mod_id",
	// 			"slug",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"version",
	// 		).RemoveIndex(
	// 			"uix_version_mod_id_slug",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201609011338",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"version",
	// 		).AddUniqueIndex(
	// 			"uix_version_mod_id_name",
	// 			"mod_id",
	// 			"name",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"version",
	// 		).RemoveIndex(
	// 			"uix_version_mod_id_name",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201609011339",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type VersionFile struct {
	// 			ID          int64  `gorm:"primary_key"`
	// 			VersionID   int64  `sql:"index"`
	// 			Slug        string `sql:"unique_index"`
	// 			ContentType string
	// 			MD5         string
	// 			CreatedAt   time.Time
	// 			UpdatedAt   time.Time
	// 		}

	// 		return engine.CreateTables(&VersionFile{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("version_file")
	// 	},
	// },
	// {
	// 	ID: "201609011340",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"version_file",
	// 		).AddForeignKey(
	// 			"version_id",
	// 			"version(id)",
	// 			"CASCADE",
	// 			"CASCADE",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011341",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type BuildVersion struct {
	// 			BuildID   int64 `sql:"index"`
	// 			VersionID int64 `sql:"index"`
	// 		}

	// 		return engine.CreateTables(&BuildVersion{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("build_version")
	// 	},
	// },
	// {
	// 	ID: "201609011342",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"build_version",
	// 		).AddForeignKey(
	// 			"build_id",
	// 			"build(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011343",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"build_version",
	// 		).AddForeignKey(
	// 			"version_id",
	// 			"version(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011344",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Client struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			Value     string `sql:"unique_index"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Client{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("client")
	// 	},
	// },
	// {
	// 	ID: "201609011345",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type ClientPack struct {
	// 			ClientID int64 `sql:"index"`
	// 			PackID   int64 `sql:"index"`
	// 		}

	// 		return engine.CreateTables(&ClientPack{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("client_pack")
	// 	},
	// },
	// {
	// 	ID: "201609011346",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"client_pack",
	// 		).AddForeignKey(
	// 			"client_id",
	// 			"client(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609011347",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"client_pack",
	// 		).AddForeignKey(
	// 			"pack_id",
	// 			"pack(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return migrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201609091338",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Key struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			Value     string `sql:"unique_index"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Key{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("key")
	// 	},
	// },
	// {
	// 	ID: "20160919214501",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackBackground struct {
	// 			MD5 string `gorm:"column:md5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackBackground{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_background").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_background").DropColumn("m_d5").Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		type PackBackground struct {
	// 			MD5 string `gorm:"column:m_d5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackBackground{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_background").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_background").DropColumn("md5").Error
	// 	},
	// },
	// {
	// 	ID: "20160919214502",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackIcon struct {
	// 			MD5 string `gorm:"column:md5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackIcon{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_icon").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_icon").DropColumn("m_d5").Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		type PackIcon struct {
	// 			MD5 string `gorm:"column:m_d5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackIcon{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_icon").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_icon").DropColumn("md5").Error
	// 	},
	// },
	// {
	// 	ID: "20160919214503",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type PackLogo struct {
	// 			MD5 string `gorm:"column:md5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackLogo{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_logo").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_logo").DropColumn("m_d5").Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		type PackLogo struct {
	// 			MD5 string `gorm:"column:m_d5"`
	// 		}

	// 		if err := tx.AutoMigrate(&PackLogo{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("pack_logo").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("pack_logo").DropColumn("md5").Error
	// 	},
	// },
	// {
	// 	ID: "20160919214504",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type VersionFile struct {
	// 			MD5 string `gorm:"column:md5"`
	// 		}

	// 		if err := tx.AutoMigrate(&VersionFile{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("version_file").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("version_file").DropColumn("m_d5").Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		type VersionFile struct {
	// 			MD5 string `gorm:"column:m_d5"`
	// 		}

	// 		if err := tx.AutoMigrate(&VersionFile{}).Error; err != nil {
	// 			return err
	// 		}

	// 		if err := tx.Set("validations:skip_validations", true).Table("version_file").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Table("version_file").DropColumn("md5").Error
	// 	},
	// },
	// {
	// 	ID: "20160921222800",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Mod struct {
	// 			Side string
	// 		}

	// 		if err := tx.AutoMigrate(&Mod{}).Error; err != nil {
	// 			return err
	// 		}

	// 		return tx.Set("validations:skip_validations", true).Table("mod").Update("side", "both").Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table("mod").DropColumn("side").Error
	// 	},
	// },
	}
)
