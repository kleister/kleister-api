package router

import (
	"io"
	"net/http"
	"path"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/middleware/prometheus"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	apiv1 "github.com/kleister/kleister-api/pkg/api/v1"
	restapiv1 "github.com/kleister/kleister-api/pkg/api/v1/restapi"
)

// Server initializes the routing of the server.
func Server(cfg *config.Config, storage store.Store, uploads upload.Upload) http.Handler {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)

	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Route("/api", func(base chi.Router) {
			base.Get("/v1/swagger", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				io.WriteString(w, string(restapiv1.SwaggerJSON))
			})

			if api := apiv1.New(); api != nil {
				base.Mount("/v1", middleware.NoCache(api.Handler))
			}

			if cfg.Server.Pprof {
				base.Mount("/debug", middleware.Profiler())
			}

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

// Metrics initializes the routing of the metrics.
func Metrics(cfg *config.Config, storage store.Store, uploads upload.Upload) http.Handler {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)

	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", prometheus.Handler(cfg.Metrics.Token))

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
