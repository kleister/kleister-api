package data

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"gopkg.in/guregu/null.v3"
)

var (
	migrations = []*gormigrate.Migration{
		{
			ID: "201609011300",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Username  string `sql:"unique_index"`
					Email     string `sql:"unique_index"`
					Hash      string `sql:"unique_index"`
					Hashword  string
					Active    bool `sql:"default:false"`
					Admin     bool `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "201609011301",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Team{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("teams").Error
			},
		},
		{
			ID: "201609011302",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID int64 `sql:"index"`
					UserID int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamUser{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_users").Error
			},
		},
		{
			ID: "201609011303",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_users",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011304",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_users",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011305",
			Migrate: func(tx *gorm.DB) error {
				type Forge struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Minecraft string
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Forge{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("forges").Error
			},
		},
		{
			ID: "201609011306",
			Migrate: func(tx *gorm.DB) error {
				type Minecraft struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Type      string
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Minecraft{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("minecrafts").Error
			},
		},
		{
			ID: "201609011307",
			Migrate: func(tx *gorm.DB) error {
				type Pack struct {
					ID            int64    `gorm:"primary_key"`
					RecommendedID null.Int `sql:"index"`
					LatestID      null.Int `sql:"index"`
					Slug          string   `sql:"unique_index"`
					Name          string   `sql:"unique_index"`
					Website       string
					Published     bool `sql:"default:false"`
					Private       bool `sql:"default:false"`
					CreatedAt     time.Time
					UpdatedAt     time.Time
				}

				return tx.CreateTable(&Pack{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("packs").Error
			},
		},
		{
			ID: "201609011308",
			Migrate: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID int64 `sql:"index"`
					PackID int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamPack{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_packs").Error
			},
		},
		{
			ID: "201609011309",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_packs",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011310",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_packs",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011311",
			Migrate: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID int64 `sql:"index"`
					PackID int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&UserPack{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_packs").Error
			},
		},
		{
			ID: "201609011312",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"user_packs",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011313",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"user_packs",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011314",
			Migrate: func(tx *gorm.DB) error {
				type PackBackground struct {
					ID          int64  `gorm:"primary_key"`
					PackID      int64  `sql:"index"`
					Slug        string `sql:"unique_index"`
					ContentType string
					MD5         string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&PackBackground{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("pack_backgrounds").Error
			},
		},
		{
			ID: "201609011315",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"pack_backgrounds",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"CASCADE",
					"CASCADE",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011316",
			Migrate: func(tx *gorm.DB) error {
				type PackIcon struct {
					ID          int64  `gorm:"primary_key"`
					PackID      int64  `sql:"index"`
					Slug        string `sql:"unique_index"`
					ContentType string
					MD5         string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&PackIcon{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("pack_icons").Error
			},
		},
		{
			ID: "201609011317",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"pack_icons",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"CASCADE",
					"CASCADE",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011318",
			Migrate: func(tx *gorm.DB) error {
				type PackLogo struct {
					ID          int64  `gorm:"primary_key"`
					PackID      int64  `sql:"index"`
					Slug        string `sql:"unique_index"`
					ContentType string
					MD5         string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&PackLogo{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("pack_logos").Error
			},
		},
		{
			ID: "201609011319",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"pack_logos",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"CASCADE",
					"CASCADE",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011320",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID          int64    `gorm:"primary_key"`
					PackID      int64    `sql:"index"`
					MinecraftID null.Int `sql:"index"`
					ForgeID     null.Int `sql:"index"`
					Slug        string
					Name        string
					MinJava     string
					MinMemory   string
					Published   bool `sql:"default:false"`
					Private     bool `sql:"default:false"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&Build{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("builds").Error
			},
		},
		{
			ID: "201609011321",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"packs",
				).AddForeignKey(
					"recommended_id",
					"builds(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011322",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"packs",
				).AddForeignKey(
					"latest_id",
					"builds(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011323",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"builds",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011324",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"builds",
				).AddForeignKey(
					"minecraft_id",
					"minecrafts(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011325",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"builds",
				).AddForeignKey(
					"forge_id",
					"forges(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011326",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"builds",
				).AddUniqueIndex(
					"uix_builds_pack_id_slug",
					"pack_id",
					"slug",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"builds",
				).RemoveIndex(
					"uix_builds_pack_id_slug",
				).Error
			},
		},
		{
			ID: "201609011327",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"builds",
				).AddUniqueIndex(
					"uix_builds_pack_id_name",
					"pack_id",
					"name",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"builds",
				).RemoveIndex(
					"uix_builds_pack_id_name",
				).Error
			},
		},
		{
			ID: "201609011328",
			Migrate: func(tx *gorm.DB) error {
				type Mod struct {
					ID          int64  `gorm:"primary_key"`
					Slug        string `sql:"unique_index"`
					Name        string `sql:"unique_index"`
					Description string `sql:"type:text"`
					Author      string
					Website     string
					Donate      string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&Mod{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("mods").Error
			},
		},
		{
			ID: "201609011329",
			Migrate: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID int64 `sql:"index"`
					ModID  int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamMod{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_mods").Error
			},
		},
		{
			ID: "201609011330",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_mods",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011331",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"team_mods",
				).AddForeignKey(
					"mod_id",
					"mods(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011332",
			Migrate: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID int64 `sql:"index"`
					ModID  int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&UserMod{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_mods").Error
			},
		},
		{
			ID: "201609011333",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"user_mods",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011334",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"user_mods",
				).AddForeignKey(
					"mod_id",
					"mods(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011335",
			Migrate: func(tx *gorm.DB) error {
				type Version struct {
					ID        int64 `gorm:"primary_key"`
					ModID     int64 `sql:"index"`
					Slug      string
					Name      string
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Version{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("versions").Error
			},
		},
		{
			ID: "201609011336",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"versions",
				).AddForeignKey(
					"mod_id",
					"mods(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011337",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"versions",
				).AddUniqueIndex(
					"uix_versions_mod_id_slug",
					"mod_id",
					"slug",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"versions",
				).RemoveIndex(
					"uix_versions_mod_id_slug",
				).Error
			},
		},
		{
			ID: "201609011338",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"versions",
				).AddUniqueIndex(
					"uix_versions_mod_id_name",
					"mod_id",
					"name",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"versions",
				).RemoveIndex(
					"uix_versions_mod_id_name",
				).Error
			},
		},
		{
			ID: "201609011339",
			Migrate: func(tx *gorm.DB) error {
				type VersionFile struct {
					ID          int64  `gorm:"primary_key"`
					VersionID   int64  `sql:"index"`
					Slug        string `sql:"unique_index"`
					ContentType string
					MD5         string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.CreateTable(&VersionFile{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("version_files").Error
			},
		},
		{
			ID: "201609011340",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"version_files",
				).AddForeignKey(
					"version_id",
					"versions(id)",
					"CASCADE",
					"CASCADE",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011341",
			Migrate: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   int64 `sql:"index"`
					VersionID int64 `sql:"index"`
				}

				return tx.CreateTable(&BuildVersion{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("build_versions").Error
			},
		},
		{
			ID: "201609011342",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"build_versions",
				).AddForeignKey(
					"build_id",
					"builds(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011343",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"build_versions",
				).AddForeignKey(
					"version_id",
					"versions(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011344",
			Migrate: func(tx *gorm.DB) error {
				type Client struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Value     string `sql:"unique_index"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Client{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("clients").Error
			},
		},
		{
			ID: "201609011345",
			Migrate: func(tx *gorm.DB) error {
				type ClientPack struct {
					ClientID int64 `sql:"index"`
					PackID   int64 `sql:"index"`
				}

				return tx.CreateTable(&ClientPack{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("client_packs").Error
			},
		},
		{
			ID: "201609011346",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"client_packs",
				).AddForeignKey(
					"client_id",
					"clients(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609011347",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return tx.Table(
					"client_packs",
				).AddForeignKey(
					"pack_id",
					"packs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201609091338",
			Migrate: func(tx *gorm.DB) error {
				type Key struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Value     string `sql:"unique_index"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Key{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("keys").Error
			},
		},
		{
			ID: "20160919214501",
			Migrate: func(tx *gorm.DB) error {
				type PackBackground struct {
					MD5 string `gorm:"column:md5"`
				}

				if err := tx.AutoMigrate(&PackBackground{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_backgrounds").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_backgrounds").DropColumn("m_d5").Error
			},
			Rollback: func(tx *gorm.DB) error {
				type PackBackground struct {
					MD5 string `gorm:"column:m_d5"`
				}

				if err := tx.AutoMigrate(&PackBackground{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_backgrounds").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_backgrounds").DropColumn("md5").Error
			},
		},
		{
			ID: "20160919214502",
			Migrate: func(tx *gorm.DB) error {
				type PackIcon struct {
					MD5 string `gorm:"column:md5"`
				}

				if err := tx.AutoMigrate(&PackIcon{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_icons").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_icons").DropColumn("m_d5").Error
			},
			Rollback: func(tx *gorm.DB) error {
				type PackIcon struct {
					MD5 string `gorm:"column:m_d5"`
				}

				if err := tx.AutoMigrate(&PackIcon{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_icons").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_icons").DropColumn("md5").Error
			},
		},
		{
			ID: "20160919214503",
			Migrate: func(tx *gorm.DB) error {
				type PackLogo struct {
					MD5 string `gorm:"column:md5"`
				}

				if err := tx.AutoMigrate(&PackLogo{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_logos").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_logos").DropColumn("m_d5").Error
			},
			Rollback: func(tx *gorm.DB) error {
				type PackLogo struct {
					MD5 string `gorm:"column:m_d5"`
				}

				if err := tx.AutoMigrate(&PackLogo{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("pack_logos").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
					return err
				}

				return tx.Table("pack_logos").DropColumn("md5").Error
			},
		},
		{
			ID: "20160919214504",
			Migrate: func(tx *gorm.DB) error {
				type VersionFile struct {
					MD5 string `gorm:"column:md5"`
				}

				if err := tx.AutoMigrate(&VersionFile{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("version_files").Update("md5", gorm.Expr("m_d5")).Error; err != nil {
					return err
				}

				return tx.Table("version_files").DropColumn("m_d5").Error
			},
			Rollback: func(tx *gorm.DB) error {
				type VersionFile struct {
					MD5 string `gorm:"column:m_d5"`
				}

				if err := tx.AutoMigrate(&VersionFile{}).Error; err != nil {
					return err
				}

				if err := tx.Set("validations:skip_validations", true).Table("version_files").Update("m_d5", gorm.Expr("md5")).Error; err != nil {
					return err
				}

				return tx.Table("version_files").DropColumn("md5").Error
			},
		},
		{
			ID: "20160921222800",
			Migrate: func(tx *gorm.DB) error {
				type Mod struct {
					Side string
				}

				if err := tx.AutoMigrate(&Mod{}).Error; err != nil {
					return err
				}

				return tx.Set("validations:skip_validations", true).Table("mods").Update("side", "both").Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table("mods").DropColumn("side").Error
			},
		},
	}
)
