package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

// User provides all user related sub-commands.
func User(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "user",
		Usage: "Commands for managing users",
		Subcommands: []cli.Command{
			UserCreate(cfg),
			UserUpdate(cfg),
			UserDelete(cfg),
			UserShow(cfg),
		},
	}
}

// UserCreate provides the sub-command to create a user.
func UserCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#create!")
		},
	}
}

// UserUpdate provides the sub-command to update a user.
func UserUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#update!")
		},
	}
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#delete!")
		},
	}
}

// UserShow provides the sub-command to show a user.
func UserShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#show!")
		},
	}
}
