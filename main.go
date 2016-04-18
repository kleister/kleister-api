package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-api/cmd"
	"github.com/solderapp/solder-api/config"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "solder"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Manage mod packs for the Technic launcher"

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Server(),
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
