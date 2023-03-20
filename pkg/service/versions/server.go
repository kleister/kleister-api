package versions

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	modsRepository "github.com/kleister/kleister-api/pkg/service/mods/repository"
	"github.com/kleister/kleister-api/pkg/service/versions/repository"
	serverv1 "github.com/kleister/kleister-api/pkg/service/versions/v1/server"
	"github.com/kleister/kleister-api/pkg/service/versions/v1/versionsv1connect"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the versions endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.VersionsRepository,
	modsRepo modsRepository.ModsRepository,
	router *chi.Mux,
) {
	mount, handler := versionsv1connect.NewVersionsServiceHandler(
		serverv1.NewVersionsServer(
			cfg,
			uploads,
			metricz,
			repository,
			modsRepo,
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
