package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kleister/kleister-api/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Health provides the sub-command to perform a health check.
func Health() *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "perform health checks for server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "server-addr",
				Value:       "0.0.0.0:8080",
				Usage:       "address to access the server",
				EnvVars:     []string{"KLEISTER_SERVER_ADDR"},
				Destination: &config.Server.Addr,
			},
		},
		Before: func(c *cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

			switch strings.ToLower(config.LogLevel) {
			case "debug":
				logger = level.NewFilter(logger, level.AllowDebug())
			case "warn":
				logger = level.NewFilter(logger, level.AllowWarn())
			case "error":
				logger = level.NewFilter(logger, level.AllowError())
			default:
				logger = level.NewFilter(logger, level.AllowInfo())
			}

			logger = log.WithPrefix(logger,
				"app", c.App.Name,
				"ts", log.DefaultTimestampUTC,
			)

			resp, err := http.Get(
				fmt.Sprintf(
					"http://%s/healthz",
					config.Server.Addr,
				),
			)

			if err != nil {
				level.Error(logger).Log(
					"msg", "failed to request health check",
					"err", err,
				)

				return err
			}

			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				level.Error(logger).Log(
					"msg", "health seems to be in a bad state",
					"code", resp.StatusCode,
				)

				return err
			}

			return nil
		},
	}
}
