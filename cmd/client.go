package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

// Client provides all client related sub-commands.
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

// ClientCreate provides the sub-command to create a client.
func ClientCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#create!")
		},
	}
}

// ClientUpdate provides the sub-command to update a client.
func ClientUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#update!")
		},
	}
}

// ClientDelete provides the sub-command to delete a client.
func ClientDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#delete!")
		},
	}
}

// ClientShow provides the sub-command to show a client.
func ClientShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a client",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute client#show!")
		},
	}
}
