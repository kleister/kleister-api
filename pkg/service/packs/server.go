package packs

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/packs/repository"
	"github.com/kleister/kleister-api/pkg/service/packs/v1/packsv1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/packs/v1/server"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the packs endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.PacksRepository,
	router *chi.Mux,
) {
	mount, handler := packsv1connect.NewPacksServiceHandler(
		serverv1.NewPacksServer(
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
