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
		Flags:  HealthFlags(cfg),
		Action: HealthAction(cfg),
	}
}

// HealthFlags defines health flags.
func HealthFlags(cfg *config.Config) []cli.Flag {
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

// HealthAction defines health action.
func HealthAction(cfg *config.Config) cli.ActionFunc {
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
				Msg("Failed to request health check")

			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Error().
				Int("code", resp.StatusCode).
				Msg("Health seems to be in bad state")

			return err
		}

		log.Debug().
			Int("code", resp.StatusCode).
			Msg("Health got a good state")

		return nil
	}
}
