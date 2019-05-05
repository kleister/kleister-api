package router

import (
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	// "github.com/kleister/kleister-api/pkg/service/auth"
	// "github.com/kleister/kleister-api/pkg/service/client"
	// "github.com/kleister/kleister-api/pkg/service/forge"
	// "github.com/kleister/kleister-api/pkg/service/general"
	// "github.com/kleister/kleister-api/pkg/service/key"
	// "github.com/kleister/kleister-api/pkg/service/minecraft"
	// "github.com/kleister/kleister-api/pkg/service/mod"
	// "github.com/kleister/kleister-api/pkg/service/pack"
	// "github.com/kleister/kleister-api/pkg/service/profile"
	// "github.com/kleister/kleister-api/pkg/service/team"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/middleware/prometheus"
	// "github.com/kleister/kleister-api/pkg/middleware/session"
	// "github.com/kleister/kleister-api/pkg/service/user"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/swagger"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/webhippie/fail"
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

	// userService := user.NewService(user.ServiceOptions{
	// 	Store: store,
	// })

	// userService = user.NewLogging(user.LoggingOptions{
	// 	Service: userService,
	// 	Logger: log.WithPrefix(logger, "service", "user"),
	// })

	// userService = user.NewMetrics(user.MetricsOptions{
	// 	Service: userService,
	// 	Metrics: "user",
	// })

	// mux.Use(session.SetCurrent(store))

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Route("/api", func(base chi.Router) {
			base.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
				content, err := swagger.ReadFile("swagger.json")

				if err != nil {
					log.Error().
						Err(err).
						Msg("failed to read swagger.json")

					fail.ErrorJSON(w, fail.Unexpected())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				io.WriteString(w, string(content))
			})

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

			// base.Get("/", general.Index(store, logger))

			// base.Mount("/auth", auth.NewHandler(store, logger))
			// base.Mount("/profile", profile.NewHandler(store, logger))
			// base.Mount("/keys", keys.NewHandler(store, logger))
			// base.Mount("/minecraft", minecraft.NewHandler(store, logger))
			// base.Mount("/forge", forge.NewHandler(store, logger))
			// base.Mount("/packs", packs.NewHandler(store, logger))
			// base.Mount("/mods", mods.NewHandler(store, logger))
			// base.Mount("/clients", clients.NewHandler(store, logger))
			// base.Mount("/teams", teams.NewHandler(store, logger))

			// base.Mount("/users", user.NewHandler(userService))

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
