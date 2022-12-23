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
			ID: "202206181600001",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        string `gorm:"primaryKey;length:36"`
					Slug      string `gorm:"unique;length:255"`
					Username  string `gorm:"unique;length:255"`
					Password  string `gorm:"length:255"`
					Email     string `gorm:"unique;length:255"`
					Firstname string `gorm:"length:255"`
					Lastname  string `gorm:"length:255"`
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
				type Member struct {
					TeamID    string `gorm:"index:idx_id,unique;length:36"`
					UserID    string `gorm:"index:idx_id,unique;length:36"`
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
			ID: "202206181600004",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type Team struct {
					ID    string    `gorm:"primaryKey"`
					Users []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type Team struct {
					ID    string    `gorm:"primaryKey"`
					Users []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Users")
			},
		},
		{
			ID: "202206181600005",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type User struct {
					ID    string    `gorm:"primaryKey"`
					Teams []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type Member struct {
					TeamID string `gorm:"index:idx_id,unique;length:36"`
					UserID string `gorm:"index:idx_id,unique;length:36"`
				}

				type User struct {
					ID    string    `gorm:"primaryKey"`
					Teams []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Teams")
			},
		},
	}
)
