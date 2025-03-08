package migrations

import (
	"github.com/uptrace/bun/migrate"
)

var (
	// Migrations provides all available database migrations.
	Migrations = migrate.NewMigrations()
)
