package command

import (
	"fmt"
	"net/http"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// Health provides the sub-command to perform a health check.
func Health(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "health",
		Usage:  "Perform health checks",
		Flags:  healthFlags(cfg),
		Before: healthBefore(cfg),
		Action: healthAction(cfg),
	}
}

func healthFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "metrics-addr",
			Value:       defaultMetricsAddr,
			Usage:       "Address to bind the metrics",
			EnvVars:     []string{"KLEISTER_API_METRICS_ADDR"},
			Destination: &cfg.Metrics.Addr,
		},
	}
}

func healthBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		setupLogger(cfg)
		return setupConfig(cfg)
	}
}

func healthAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		resp, err := http.Get(
			fmt.Sprintf(
				"http://%s/healthz",
				cfg.Metrics.Addr,
			),
		)

		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to request health check")

			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Error().
				Int("code", resp.StatusCode).
				Msg("health seems to be in bad state")

			return err
		}

		log.Debug().
			Int("code", resp.StatusCode).
			Msg("health got a good state")

		return nil
	}
}
