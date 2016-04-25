package cmd

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/router"
)

// Server provides the sub-command to start the API server.
func Server() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Solder server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "SOLDER_DB_DRIVER",
				Destination: &config.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "solder",
				Usage:       "Name for database connection",
				EnvVar:      "SOLDER_DB_NAME",
				Destination: &config.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "SOLDER_DB_USERNAME",
				Destination: &config.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "SOLDER_DB_PASSWORD",
				Destination: &config.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "SOLDER_DB_HOST",
				Destination: &config.Database.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "SOLDER_SERVER_ADDR",
				Destination: &config.Server.Addr,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "SOLDER_SERVER_CERT",
				Destination: &config.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "SOLDER_SERVER_KEY",
				Destination: &config.Server.Key,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "SOLDER_SERVER_ROOT",
				Destination: &config.Server.Root,
			},
			cli.StringFlag{
				Name:        "storage",
				Value:       "storage/",
				Usage:       "Folder for storing uploads",
				EnvVar:      "SOLDER_SERVER_STORAGE",
				Destination: &config.Server.Storage,
			},
		},
		Action: func(c *cli.Context) {
			logrus.Infof("starting server on %s", config.Server.Addr)

			if config.Server.Cert != "" && config.Server.Key != "" {
				logrus.Fatal(
					http.ListenAndServeTLS(
						config.Server.Addr,
						config.Server.Cert,
						config.Server.Key,
						router.Load(),
					),
				)
			} else {
				logrus.Fatal(
					http.ListenAndServe(
						config.Server.Addr,
						router.Load(),
					),
				)
			}
		},
	}
}
