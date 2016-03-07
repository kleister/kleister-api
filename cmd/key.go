package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

// Key provides all key related sub-commands.
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

// KeyCreate provides the sub-command to create a key.
func KeyCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#create!")
		},
	}
}

// KeyUpdate provides the sub-command to update a key.
func KeyUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#update!")
		},
	}
}

// KeyDelete provides the sub-command to delete a key.
func KeyDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#delete!")
		},
	}
}

// KeyShow provides the sub-command to show a key.
func KeyShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a key",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute key#show!")
		},
	}
}
