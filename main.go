package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/solderapp/solder-api.v0/cmd"
	"gopkg.in/solderapp/solder-api.v0/config"
)

var (
	version string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg := &config.Config{
		Version: version,
	}

	app := cli.NewApp()
	app.Name = "solder"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Manage mod packs for the Technic launcher"

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if cfg.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Activate debug information",
			EnvVar:      "SOLDER_DEBUG_ENABLED",
			Destination: &cfg.Debug,
		},
		cli.BoolFlag{
			Name:        "develop",
			Usage:       "Activate development mode",
			EnvVar:      "SOLDER_DEVELOP_ENABLED",
			Destination: &cfg.Develop,
		},
		cli.StringFlag{
			Name:        "db-driver",
			Value:       "mysql",
			Usage:       "Database driver selection",
			EnvVar:      "SOLDER_DB_DRIVER",
			Destination: &cfg.Database.Driver,
		},
		cli.StringFlag{
			Name:        "db-username",
			Value:       "root",
			Usage:       "Username for database connection",
			EnvVar:      "SOLDER_DB_USERNAME",
			Destination: &cfg.Database.Username,
		},
		cli.StringFlag{
			Name:        "db-password",
			Value:       "root",
			Usage:       "Password for database connection",
			EnvVar:      "SOLDER_DB_PASSWORD",
			Destination: &cfg.Database.Password,
		},
		cli.StringFlag{
			Name:        "db-name",
			Value:       "solder",
			Usage:       "Name for database connection",
			EnvVar:      "SOLDER_DB_NAME",
			Destination: &cfg.Database.Name,
		},
		cli.StringFlag{
			Name:        "db-host",
			Value:       "localhost:3306",
			Usage:       "Host for database connection",
			EnvVar:      "SOLDER_DB_HOST",
			Destination: &cfg.Database.Host,
		},
	}

	app.Commands = []cli.Command{
		cmd.Server(cfg),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
