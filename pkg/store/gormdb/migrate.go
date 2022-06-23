package gormdb

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	migrations = []*gormigrate.Migration{
		{
			ID: "202206181600001",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        string `gorm:"primaryKey;length:36"`
					Slug      string `gorm:"unique;length:255"`
					Email     string `gorm:"unique;length:255"`
					Username  string `gorm:"unique;length:255"`
					Password  string `gorm:"length:255"`
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
			ID: "202206181600002",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        string `gorm:"primaryKey;length:36"`
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
			ID: "202206181600003",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID    string `gorm:"index:idx_id,unique;length:36"`
					UserID    string `gorm:"index:idx_id,unique;length:36"`
					Perm      string `gorm:"length:32"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&TeamUser{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("team_users")
			},
		},
		{
			ID: "202206181600004",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Users []*TeamUser `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Users []*TeamUser `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Users")
			},
		},
		{
			ID: "202206181600005",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*TeamUser `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*TeamUser `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Teams")
			},
		},
		{
			ID: "202206221600006",
			Migrate: func(tx *gorm.DB) error {
				type Minecraft struct {
					ID        string `gorm:"primaryKey;length:36"`
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
			ID: "202206221600007",
			Migrate: func(tx *gorm.DB) error {
				type Forge struct {
					ID        string `gorm:"primaryKey;length:36"`
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
	}
)
