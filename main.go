package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/cmd"
	"github.com/solderapp/solder/config"
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
	}

	app.Commands = []cli.Command{
		cmd.Server(cfg),
		cmd.User(cfg),
		cmd.Client(cfg),
		cmd.Key(cfg),
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
