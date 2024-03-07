package store

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	// Migrations define all database migrations.
	Migrations = []*gormigrate.Migration{
		{
			ID: "0001_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        string `gorm:"primaryKey;length:20"`
					Username  string `gorm:"unique;length:255"`
					Hashword  string `gorm:"length:255"`
					Email     string `gorm:"length:255"`
					Fullname  string `gorm:"length:255"`
					Active    bool   `gorm:"default:false"`
					Admin     bool   `gorm:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "0002_create_user_auths_table",
			Migrate: func(tx *gorm.DB) error {
				type UserAuth struct {
					ID        string `gorm:"primaryKey;length:20"`
					UserID    string `gorm:"length:20"`
					Provider  string `gorm:"length:255"`
					Ref       string `gorm:"length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&UserAuth{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_auths")
			},
		},
		{
			ID: "0003_create_user_auths_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserAuth struct {
					UserID string `gorm:"length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Auths []*UserAuth `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Auths")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserAuth struct {
					UserID string `gorm:"length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Auths []*UserAuth `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Auths")
			},
		},
		{
			ID: "0004_create_teams_table",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        string `gorm:"primaryKey;length:20"`
					Slug      string `gorm:"unique;length:255"`
					Name      string `gorm:"unique;length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Team{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("teams")
			},
		},
		{
			ID: "0005_create_members_table",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					TeamID    string `gorm:"index:idx_id,unique;length:20"`
					UserID    string `gorm:"index:idx_id,unique;length:20"`
					Perm      string `gorm:"length:64"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Member{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("members")
			},
		},
		{
			ID: "0006_create_members_teams_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					UserID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID    string    `gorm:"primaryKey"`
					Users []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					UserID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID    string    `gorm:"primaryKey"`
					Users []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Users")
			},
		},
		{
			ID: "0007_create_members_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					UserID string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID    string    `gorm:"primaryKey"`
					Teams []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					UserID string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID    string    `gorm:"primaryKey"`
					Teams []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Teams")
			},
		},
		{
			ID: "0008_create_minecrafts_table",
			Migrate: func(tx *gorm.DB) error {
				type Minecraft struct {
					ID        string `gorm:"primaryKey;length:20"`
					Name      string `gorm:"unique;length:255"`
					Type      string `gorm:"index;length:64"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Minecraft{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("minecrafts")
			},
		},
		{
			ID: "0009_create_forges_table",
			Migrate: func(tx *gorm.DB) error {
				type Forge struct {
					ID        string `gorm:"primaryKey;length:20"`
					Name      string `gorm:"unique;length:255"`
					Minecraft string `gorm:"index;length:64"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Forge{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("forges")
			},
		},
		{
			ID: "0010_create_neoforges_table",
			Migrate: func(tx *gorm.DB) error {
				type Neoforge struct {
					ID        string `gorm:"primaryKey;length:20"`
					Name      string `gorm:"unique;length:255"`
					Minecraft string `gorm:"index;length:64"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Neoforge{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("neoforges")
			},
		},
		{
			ID: "0011_create_quilts_table",
			Migrate: func(tx *gorm.DB) error {
				type Quilt struct {
					ID        string `gorm:"primaryKey;length:20"`
					Name      string `gorm:"unique;length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Quilt{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("quilts")
			},
		},
		{
			ID: "0012_create_fabrics_table",
			Migrate: func(tx *gorm.DB) error {
				type Fabric struct {
					ID        string `gorm:"primaryKey;length:20"`
					Name      string `gorm:"unique;length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Fabric{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("fabrics")
			},
		},
		{
			ID: "0013_create_mods_table",
			Migrate: func(tx *gorm.DB) error {
				type Mod struct {
					ID          string `gorm:"primaryKey;length:20"`
					Slug        string `gorm:"unique;length:255"`
					Name        string `gorm:"unique;length:255"`
					Side        string `gorm:"index;length:36"`
					Description string
					Author      string
					Website     string
					Donate      string
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&Mod{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("mods")
			},
		},

		{
			ID: "0014_create_user_mods_table",
			Migrate: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID    string `gorm:"index:idx_id,unique;length:20"`
					ModID     string `gorm:"index:idx_id,unique;length:20"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&UserMod{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_mods")
			},
		},
		{
			ID: "0015_create_user_mods_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID   string     `gorm:"primaryKey"`
					Mods []*UserMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Mods")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID   string     `gorm:"primaryKey"`
					Mods []*UserMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Mods")
			},
		},
		{
			ID: "0016_create_user_mods_mods_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Mod struct {
					ID    string     `gorm:"primaryKey"`
					Users []*UserMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Mod{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserMod struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Mod struct {
					ID    string     `gorm:"primaryKey"`
					Users []*UserMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Mod{}, "Users")
			},
		},
		{
			ID: "0017_create_team_mods_table",
			Migrate: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID    string `gorm:"index:idx_id,unique;length:20"`
					ModID     string `gorm:"index:idx_id,unique;length:20"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&TeamMod{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("team_mods")
			},
		},
		{
			ID: "0018_create_team_mods_teams_constraint",
			Migrate: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID   string     `gorm:"primaryKey"`
					Mods []*TeamMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Mods")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID   string     `gorm:"primaryKey"`
					Mods []*TeamMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Mods")
			},
		},
		{
			ID: "0019_create_team_mods_mods_constraint",
			Migrate: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Mod struct {
					ID    string     `gorm:"primaryKey"`
					Teams []*TeamMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Mod{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamMod struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					ModID  string `gorm:"index:idx_id,unique;length:20"`
				}

				type Mod struct {
					ID    string     `gorm:"primaryKey"`
					Teams []*TeamMod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Mod{}, "Teams")
			},
		},
		{
			ID: "0020_create_versions_table",
			Migrate: func(tx *gorm.DB) error {
				type Version struct {
					ID        string `gorm:"primaryKey;length:20"`
					ModID     string `gorm:"index;length:20"`
					Slug      string `gorm:"length:255"`
					Name      string `gorm:"length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Version{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("versions")
			},
		},
		{
			ID: "0021_create_versions_mods_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Version struct {
					ID    string `gorm:"primaryKey"`
					ModID string `gorm:"index;length:20"`
				}

				type Mod struct {
					ID       string     `gorm:"primaryKey"`
					Versions []*Version `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Mod{}, "Versions")
			},
			Rollback: func(tx *gorm.DB) error {
				type Version struct {
					ID string `gorm:"primaryKey"`
				}

				type Mod struct {
					ID       string     `gorm:"primaryKey"`
					Versions []*Version `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Mod{}, "Versions")
			},
		},
		{
			ID: "0022_create_version_files_table",
			Migrate: func(tx *gorm.DB) error {
				type VersionFile struct {
					ID          string `gorm:"primaryKey;length:20"`
					VersionID   string `gorm:"index;length:20"`
					Slug        string `gorm:"unique;length:255"`
					ContentType string
					MD5         string `gorm:"column:md5"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&VersionFile{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("version_files")
			},
		},
		{
			ID: "0023_create_version_files_versions_constraint",
			Migrate: func(tx *gorm.DB) error {
				type VersionFile struct {
					ID        string `gorm:"primaryKey"`
					VersionID string `gorm:"index;length:20"`
				}

				type Version struct {
					ID   string       `gorm:"primaryKey"`
					File *VersionFile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Version{}, "File")
			},
			Rollback: func(tx *gorm.DB) error {
				type VersionFile struct {
					ID        string `gorm:"primaryKey"`
					VersionID string `gorm:"index;length:20"`
				}

				type Version struct {
					ID   string       `gorm:"primaryKey"`
					File *VersionFile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Version{}, "File")
			},
		},
		{
			ID: "0024_create_packs_table",
			Migrate: func(tx *gorm.DB) error {
				type Pack struct {
					ID        string `gorm:"primaryKey;length:20"`
					Slug      string `gorm:"unique;length:255"`
					Name      string `gorm:"unique;length:255"`
					Website   string
					Published bool `gorm:"default:true"`
					Private   bool `gorm:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Pack{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("packs")
			},
		},
		{
			ID: "0025_create_pack_backs_table",
			Migrate: func(tx *gorm.DB) error {
				type PackBack struct {
					ID          string `gorm:"primaryKey;length:20"`
					PackID      string `gorm:"index;length:20"`
					Slug        string `gorm:"unique;length:255"`
					ContentType string
					MD5         string `gorm:"column:md5"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&PackBack{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("pack_backs")
			},
		},
		{
			ID: "0026_create_pack_backs_packs_constraint",
			Migrate: func(tx *gorm.DB) error {
				type PackBack struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Back *PackBack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Back")
			},
			Rollback: func(tx *gorm.DB) error {
				type PackBack struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Back *PackBack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Back")
			},
		},
		{
			ID: "0027_create_pack_icons_table",
			Migrate: func(tx *gorm.DB) error {
				type PackIcon struct {
					ID          string `gorm:"primaryKey;length:20"`
					PackID      string `gorm:"index;length:20"`
					Slug        string `gorm:"unique;length:255"`
					ContentType string
					MD5         string `gorm:"column:md5"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&PackIcon{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("pack_icons")
			},
		},
		{
			ID: "0028_create_pack_icons_packs_constraint",
			Migrate: func(tx *gorm.DB) error {
				type PackIcon struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Icon *PackIcon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Icon")
			},
			Rollback: func(tx *gorm.DB) error {
				type PackIcon struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Icon *PackIcon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Icon")
			},
		},
		{
			ID: "0029_create_pack_logos_table",
			Migrate: func(tx *gorm.DB) error {
				type PackLogo struct {
					ID          string `gorm:"primaryKey;length:20"`
					PackID      string `gorm:"index;length:20"`
					Slug        string `gorm:"unique;length:255"`
					ContentType string
					MD5         string `gorm:"column:md5"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&PackLogo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("pack_logos")
			},
		},
		{
			ID: "0030_create_pack_logos_packs_constraint",
			Migrate: func(tx *gorm.DB) error {
				type PackLogo struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Logo *PackLogo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Logo")
			},
			Rollback: func(tx *gorm.DB) error {
				type PackLogo struct {
					ID     string `gorm:"primaryKey"`
					PackID string `gorm:"index;length:20"`
				}

				type Pack struct {
					ID   string    `gorm:"primaryKey"`
					Logo *PackLogo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Logo")
			},
		},
		{
			ID: "0031_create_user_packs_table",
			Migrate: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID    string `gorm:"index:idx_id,unique;length:20"`
					PackID    string `gorm:"index:idx_id,unique;length:20"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&UserPack{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_packs")
			},
		},
		{
			ID: "0032_create_user_packs_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Packs []*UserPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Packs")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Packs []*UserPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Packs")
			},
		},
		{
			ID: "0033_create_user_packs_packs_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Pack struct {
					ID    string      `gorm:"primaryKey"`
					Users []*UserPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserPack struct {
					UserID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Pack struct {
					ID    string      `gorm:"primaryKey"`
					Users []*UserPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Users")
			},
		},
		{
			ID: "0034_create_team_packs_table",
			Migrate: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID    string `gorm:"index:idx_id,unique;length:20"`
					PackID    string `gorm:"index:idx_id,unique;length:20"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&TeamPack{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("team_packs")
			},
		},
		{
			ID: "0035_create_team_packs_teams_constraint",
			Migrate: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Packs []*TeamPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Packs")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Packs []*TeamPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Packs")
			},
		},
		{
			ID: "0036_create_team_packs_packs_constraint",
			Migrate: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Pack struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*TeamPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamPack struct {
					TeamID string `gorm:"index:idx_id,unique;length:20"`
					PackID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Pack struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*TeamPack `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Teams")
			},
		},
		{
			ID: "0037_create_builds_table",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID          string `gorm:"primaryKey;length:20"`
					Slug        string `gorm:"unique;length:255"`
					Name        string `gorm:"unique;length:255"`
					MinecraftID string `gorm:"index;length:20"`
					ForgeID     string `gorm:"index;length:20"`
					NeoforgeID  string `gorm:"index;length:20"`
					QuiltID     string `gorm:"index;length:20"`
					FabricID    string `gorm:"index;length:20"`
					Website     string
					Recommended bool `gorm:"default:false"`
					Published   bool `gorm:"default:true"`
					Private     bool `gorm:"default:false"`
					CreatedAt   time.Time
					UpdatedAt   time.Time
				}

				return tx.Migrator().CreateTable(&Build{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("builds")
			},
		},
		{
			ID: "0038_create_packs_recommended_build_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID string `gorm:"primaryKey;length:20"`
				}

				type Pack struct {
					ID            string `gorm:"primaryKey"`
					RecommendedID string `gorm:"index;length:20"`
					Recommended   *Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Recommended")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID string `gorm:"primaryKey;length:20"`
				}

				type Pack struct {
					ID            string `gorm:"primaryKey"`
					RecommendedID string `gorm:"index;length:20"`
					Recommended   *Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Recommended")
			},
		},
		{
			ID: "0039_create_packs_latest_build_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID string `gorm:"primaryKey;length:20"`
				}

				type Pack struct {
					ID       string  `gorm:"primaryKey"`
					LatestID *string `gorm:"index;length:20"`
					Latest   *Build  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Pack{}, "Latest")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID string `gorm:"primaryKey;length:20"`
				}

				type Pack struct {
					ID       string  `gorm:"primaryKey"`
					LatestID *string `gorm:"index;length:20"`
					Latest   *Build  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Pack{}, "Latest")
			},
		},
		{
			ID: "0040_create_builds_minecrafts_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID          string `gorm:"primaryKey"`
					MinecraftID string `gorm:"index:idx_id;length:20"`
				}

				type Minecraft struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Minecraft{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID          string `gorm:"primaryKey"`
					MinecraftID string `gorm:"index:idx_id;length:20"`
				}

				type Minecraft struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Minecraft{}, "Builds")
			},
		},
		{
			ID: "0041_create_builds_forges_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID      string `gorm:"primaryKey"`
					ForgeID string `gorm:"index:idx_id;length:20"`
				}

				type Forge struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Forge{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID      string `gorm:"primaryKey"`
					ForgeID string `gorm:"index:idx_id;length:20"`
				}

				type Forge struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Forge{}, "Builds")
			},
		},
		{
			ID: "0042_create_builds_neoforges_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID         string `gorm:"primaryKey"`
					NeoforgeID string `gorm:"index:idx_id;length:20"`
				}

				type Neoforge struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Neoforge{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID         string `gorm:"primaryKey"`
					NeoforgeID string `gorm:"index:idx_id;length:20"`
				}

				type Neoforge struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Neoforge{}, "Builds")
			},
		},
		{
			ID: "0043_create_builds_quilts_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID      string `gorm:"primaryKey"`
					QuiltID string `gorm:"index:idx_id;length:20"`
				}

				type Quilt struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Quilt{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID      string `gorm:"primaryKey"`
					QuiltID string `gorm:"index:idx_id;length:20"`
				}

				type Quilt struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Quilt{}, "Builds")
			},
		},
		{
			ID: "0044_create_builds_fabrics_constraint",
			Migrate: func(tx *gorm.DB) error {
				type Build struct {
					ID       string `gorm:"primaryKey"`
					FabricID string `gorm:"index:idx_id;length:20"`
				}

				type Fabric struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Fabric{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type Build struct {
					ID       string `gorm:"primaryKey"`
					FabricID string `gorm:"index:idx_id;length:20"`
				}

				type Fabric struct {
					ID     string   `gorm:"primaryKey;length:20"`
					Builds []*Build `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Fabric{}, "Builds")
			},
		},
		{
			ID: "0045_create_build_versions_table",
			Migrate: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   string `gorm:"index:idx_id,unique;length:20"`
					VersionID string `gorm:"index:idx_id,unique;length:20"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&BuildVersion{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("build_versions")
			},
		},
		{
			ID: "0046_create_build_versions_builds_constraint",
			Migrate: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   string `gorm:"index:idx_id,unique;length:20"`
					VersionID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Build struct {
					ID       string          `gorm:"primaryKey"`
					Versions []*BuildVersion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Build{}, "Versions")
			},
			Rollback: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   string `gorm:"index:idx_id,unique;length:20"`
					VersionID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Build struct {
					ID       string          `gorm:"primaryKey"`
					Versions []*BuildVersion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Build{}, "Versions")
			},
		},
		{
			ID: "0047_create_build_versions_versions_constraint",
			Migrate: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   string `gorm:"index:idx_id,unique;length:20"`
					VersionID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Version struct {
					ID     string          `gorm:"primaryKey"`
					Builds []*BuildVersion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Version{}, "Builds")
			},
			Rollback: func(tx *gorm.DB) error {
				type BuildVersion struct {
					BuildID   string `gorm:"index:idx_id,unique;length:20"`
					VersionID string `gorm:"index:idx_id,unique;length:20"`
				}

				type Version struct {
					ID     string          `gorm:"primaryKey"`
					Builds []*BuildVersion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Version{}, "Builds")
			},
		},
	}
)
