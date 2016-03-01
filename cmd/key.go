package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

func Key(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "key",
		Usage: "Commands for managing keys",
		Subcommands: []cli.Command{
			KeyCreate(cfg),
			KeyUpdate(cfg),
			KeyDelete(cfg),
			KeyShow(cfg),
		},
	}
}

func KeyCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#create!")
		},
	}
}

func KeyUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#update!")
		},
	}
}

func KeyDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#delete!")
		},
	}
}

func KeyShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#show!")
		},
	}
}
