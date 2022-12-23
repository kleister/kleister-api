package teams

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/teams/repository"
	serverv1 "github.com/kleister/kleister-api/pkg/service/teams/v1/server"
	"github.com/kleister/kleister-api/pkg/service/teams/v1/teamsv1connect"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the teams endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.TeamsRepository,
	router *chi.Mux,
) {
	mount, handler := teamsv1connect.NewTeamsServiceHandler(
		serverv1.NewTeamsServer(
			cfg,
			uploads,
			metricz,
			repository,
		),
	)

	router.Mount(
		path.Join(
			cfg.Server.Root,
			mount,
		)+"/",
		http.StripPrefix(
			strings.TrimRight(
				cfg.Server.Root,
				"/",
			),
			handler,
		),
	)
}
