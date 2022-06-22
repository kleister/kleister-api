package router

import (
	"io"
	"net/http"
	"path"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/middleware/prometheus"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	doc "github.com/utahta/swagger-doc"

	apiv1 "github.com/kleister/kleister-api/pkg/api/v1"
	restapiv1 "github.com/kleister/kleister-api/pkg/api/v1/restapi"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	uploads upload.Upload,
	usersService users.Service,
	teamsService teams.Service,
	minecraftService minecraft.Service,
	forgeService forge.Service,
) http.Handler {
	mux := chi.NewRouter()

	mux.Use(requestid.Handler)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))

	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Str("request", requestid.Get(r.Context())).
			Str("method", r.Method).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Route("/api", func(base chi.Router) {
			if cfg.Server.Pprof {
				base.Mount("/debug", middleware.Profiler())
			}

			base.Route("/v1", func(v1 chi.Router) {
				if cfg.Server.Docs {
					v1.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)

						io.WriteString(w, string(restapiv1.SwaggerJSON))
					})

					v1.Handle("/docs", doc.NewRedocHandler(
						path.Join(
							cfg.Server.Root,
							"api",
							"v1",
							"swagger",
						),
						path.Join(
							cfg.Server.Root,
							"/api/v1/docs",
						),
					))
				}

				if api := apiv1.New(
					cfg,
					uploads,
					usersService,
					teamsService,
					minecraftService,
					forgeService,
				); api != nil {
					v1.Mount("/", middleware.NoCache(api.Handler))
				}
			})

			base.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"storage",
				),
			))
		})
	})

	return mux
}

// Metrics initializes the routing of metrics and health.
func Metrics(cfg *config.Config, metrics *metrics.Metrics) http.Handler {
	mux := chi.NewRouter()

	mux.Use(requestid.Handler)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", prometheus.Handler(metrics.Registry, cfg.Metrics.Token))

		root.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			io.WriteString(w, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			io.WriteString(w, http.StatusText(http.StatusOK))
		})
	})

	return mux
}
