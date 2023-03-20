package minecraft

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/minecraft/repository"
	"github.com/kleister/kleister-api/pkg/service/minecraft/v1/minecraftv1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/minecraft/v1/server"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the minecraft endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.MinecraftRepository,
	router *chi.Mux,
) {
	mount, handler := minecraftv1connect.NewMinecraftServiceHandler(
		serverv1.NewMinecraftServer(
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
