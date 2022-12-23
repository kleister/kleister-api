package router

import (
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/middleware/prometheus"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/service/members/v1/membersv1connect"
	"github.com/kleister/kleister-api/pkg/service/teams/v1/teamsv1connect"
	"github.com/kleister/kleister-api/pkg/service/users/v1/usersv1connect"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	uploads upload.Upload,
) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(requestid.Handler)
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
		root.Handle("/storage/*", uploads.Handler(
			path.Join(
				cfg.Server.Root,
				"storage",
			),
		))

		{
			mount, handler := grpchealth.NewHandler(
				grpchealth.NewStaticChecker(
					usersv1connect.UsersServiceName,
					teamsv1connect.TeamsServiceName,
					membersv1connect.MembersServiceName,
				),
				connect.WithCompressMinBytes(1024),
			)

			root.Mount(
				mount,
				http.StripPrefix(
					strings.TrimRight(
						cfg.Server.Root,
						"/",
					),
					handler,
				),
			)
		}

		{
			mount, handler := grpcreflect.NewHandlerV1(
				grpcreflect.NewStaticReflector(
					usersv1connect.UsersServiceName,
					teamsv1connect.TeamsServiceName,
					membersv1connect.MembersServiceName,
				),
				connect.WithCompressMinBytes(1024),
			)

			root.Mount(
				mount,
				http.StripPrefix(
					strings.TrimRight(
						cfg.Server.Root,
						"/",
					),
					handler,
				),
			)
		}

		{
			mount, handler := grpcreflect.NewHandlerV1Alpha(
				grpcreflect.NewStaticReflector(
					usersv1connect.UsersServiceName,
					teamsv1connect.TeamsServiceName,
					membersv1connect.MembersServiceName,
				),
				connect.WithCompressMinBytes(1024),
			)

			root.Mount(
				mount,
				http.StripPrefix(
					strings.TrimRight(
						cfg.Server.Root,
						"/",
					),
					handler,
				),
			)
		}
	})

	return mux
}

// Metrics initializes the routing of metrics and health.
func Metrics(
	cfg *config.Config,
	metricz *metrics.Metrics,
) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", prometheus.Handler(metricz.Registry, cfg.Metrics.Token))

		if cfg.Metrics.Pprof {
			root.Mount("/debug", middleware.Profiler())
		}

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
