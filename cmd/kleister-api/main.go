package main

import (
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("KLEISTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "kleister-api",
		Version:  version.Version.String(),
		Usage:    "manage mod packs for minecraft",
		Compiled: time.Now(),

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Value:       "info",
				Usage:       "set logging level",
				EnvVars:     []string{"KLEISTER_LOG_LEVEL"},
				Destination: &config.LogLevel,
			},
			&cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "database driver selection",
				EnvVars:     []string{"KLEISTER_DB_DRIVER"},
				Destination: &config.Database.Driver,
			},
			&cli.StringFlag{
				Name:        "db-name",
				Value:       "umschlag",
				Usage:       "name for database to use",
				EnvVars:     []string{"KLEISTER_DB_NAME"},
				Destination: &config.Database.Name,
			},
			&cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "username for database",
				EnvVars:     []string{"KLEISTER_DB_USERNAME"},
				Destination: &config.Database.Username,
			},
			&cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "password for database",
				EnvVars:     []string{"KLEISTER_DB_PASSWORD"},
				Destination: &config.Database.Password,
			},
			&cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "host for database",
				EnvVars:     []string{"KLEISTER_DB_HOST"},
				Destination: &config.Database.Host,
			},
			&cli.IntFlag{
				Name:        "db-timeout",
				Value:       60,
				Usage:       "timeout for waiting on db",
				EnvVars:     []string{"KLEISTER_DB_TIMEOUT"},
				Destination: &config.Database.Timeout,
			},
		},

		Before: func(c *cli.Context) error {
			return nil
		},

		Commands: []*cli.Command{
			Server(),
			Health(),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
