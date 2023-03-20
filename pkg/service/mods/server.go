package mods

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/mods/repository"
	"github.com/kleister/kleister-api/pkg/service/mods/v1/modsv1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/mods/v1/server"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the mods endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.ModsRepository,
	router *chi.Mux,
) {
	mount, handler := modsv1connect.NewModsServiceHandler(
		serverv1.NewModsServer(
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
