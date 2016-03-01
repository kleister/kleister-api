package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
)

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

func UserCreate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#create!")
		},
	}
}

func UserUpdate(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#update!")
		},
	}
}

func UserDelete(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#delete!")
		},
	}
}

func UserShow(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show an user",
		Action: func(c *cli.Context) {
			logrus.Debug("Execute user#show!")
		},
	}
}
