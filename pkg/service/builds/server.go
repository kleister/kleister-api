package builds

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/builds/repository"
	"github.com/kleister/kleister-api/pkg/service/builds/v1/buildsv1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/builds/v1/server"
	packsRepository "github.com/kleister/kleister-api/pkg/service/packs/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the builds endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.BuildsRepository,
	packsRepo packsRepository.PacksRepository,
	router *chi.Mux,
) {
	mount, handler := buildsv1connect.NewBuildsServiceHandler(
		serverv1.NewBuildsServer(
			cfg,
			uploads,
			metricz,
			repository,
			packsRepo,
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
