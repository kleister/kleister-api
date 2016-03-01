package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

func Client(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "client",
		Usage: "Commands for managing clients",
		Subcommands: []cli.Command{
			ClientCreate(cfg),
			ClientUpdate(cfg),
			ClientDelete(cfg),
			ClientShow(cfg),
		},
	}
}

func ClientCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#create!")
		},
	}
}

func ClientUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#update!")
		},
	}
}

func ClientDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#delete!")
		},
	}
}

func ClientShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#show!")
		},
	}
}
