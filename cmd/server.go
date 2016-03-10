package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder/config"
	"github.com/solderapp/solder/router"
	"github.com/solderapp/solder/router/middleware/context"
	"github.com/solderapp/solder/server"
	"github.com/solderapp/solder/store"
)

// Server provides the sub-command to start the API server.
func Server(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Solder server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "SOLDER_SERVER_ADDR",
				Destination: &cfg.Server.Addr,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "SOLDER_SERVER_CERT",
				Destination: &cfg.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "SOLDER_SERVER_KEY",
				Destination: &cfg.Server.Key,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "SOLDER_SERVER_ROOT",
				Destination: &cfg.Server.Root,
			},
		},
		Action: func(c *cli.Context) {
			dat := store.Load(
				cfg,
			)

			srv := server.Load(
				cfg,
			)

			srv.Run(
				router.Load(
					cfg,
					context.SetConfig(
						*cfg,
					),
					context.SetStore(
						*dat,
					),
				),
			)
		},
	}
}
