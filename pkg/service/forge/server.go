package forge

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/forge/repository"
	"github.com/kleister/kleister-api/pkg/service/forge/v1/forgev1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/forge/v1/server"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the forge endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.ForgeRepository,
	router *chi.Mux,
) {
	mount, handler := forgev1connect.NewForgeServiceHandler(
		serverv1.NewForgeServer(
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
