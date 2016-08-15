package cmd

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/router"
	"github.com/urfave/cli"
)

// Server provides the sub-command to start the API server.
func Server() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Kleister API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "KLEISTER_DB_DRIVER",
				Destination: &config.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "kleister",
				Usage:       "Name for database connection",
				EnvVar:      "KLEISTER_DB_NAME",
				Destination: &config.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "KLEISTER_DB_USERNAME",
				Destination: &config.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "KLEISTER_DB_PASSWORD",
				Destination: &config.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "KLEISTER_DB_HOST",
				Destination: &config.Database.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "KLEISTER_SERVER_ADDR",
				Destination: &config.Server.Addr,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "KLEISTER_SERVER_CERT",
				Destination: &config.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "KLEISTER_SERVER_KEY",
				Destination: &config.Server.Key,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "KLEISTER_SERVER_ROOT",
				Destination: &config.Server.Root,
			},
			cli.StringFlag{
				Name:        "storage",
				Value:       "storage/",
				Usage:       "Folder for storing uploads",
				EnvVar:      "KLEISTER_SERVER_STORAGE",
				Destination: &config.Server.Storage,
			},
			cli.DurationFlag{
				Name:        "expire",
				Value:       time.Hour * 24,
				Usage:       "Session expire duration",
				EnvVar:      "KLEISTER_SESSION_EXPIRE",
				Destination: &config.Session.Expire,
			},
		},
		Action: func(c *cli.Context) {
			logrus.Infof("Starting the API on %s", config.Server.Addr)

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
