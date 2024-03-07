package router

import (
	"io"
	"net/http"
	"path"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	apiv1 "github.com/kleister/kleister-api/pkg/api/v1"
	restapiv1 "github.com/kleister/kleister-api/pkg/api/v1/restapi"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/middleware/requestid"
	"github.com/kleister/kleister-api/pkg/respond"
	buildVersions "github.com/kleister/kleister-api/pkg/service/build_versions"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/fabric"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/members"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/neoforge"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/quilt"
	teamMods "github.com/kleister/kleister-api/pkg/service/team_mods"
	teamPacks "github.com/kleister/kleister-api/pkg/service/team_packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	userMods "github.com/kleister/kleister-api/pkg/service/user_mods"
	userPacks "github.com/kleister/kleister-api/pkg/service/user_packs"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/session"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	doc "github.com/utahta/swagger-doc"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	registry *metrics.Metrics,
	sess *session.Session,
	uploads upload.Upload,
	storage store.Store,
	minecraftService minecraft.Service,
	forgeService forge.Service,
	neoforgeService neoforge.Service,
	quiltService quilt.Service,
	fabricService fabric.Service,
	teamsService teams.Service,
	usersService users.Service,
	membersService members.Service,
	modsService mods.Service,
	userModsService userMods.Service,
	teamModsService teamMods.Service,
	versionsService versions.Service,
	packsService packs.Service,
	userPacksService userPacks.Service,
	teamPacksService teamPacks.Service,
	buildsService builds.Service,
	buildVersionsService buildVersions.Service,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("request")
	}))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)
	mux.Use(sess.Middleware)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Get("/", func(w http.ResponseWriter, r *http.Request) {

			respond.JSON(
				w,
				r,
				[]string{
					sessionz.Get(
						r.Context(),
						"user",
					),
					sessionz.Get(
						r.Context(),
						"github",
					),
				},
			)

		})

		root.Route("/api/v1", func(r chi.Router) {
			swagger, err := v1.GetSwagger()

			if err != nil {
				log.Error().
					Err(err).
					Str("version", "v1").
					Msg("Failed to load openapi spec")
			}

			swagger.Servers = openapi3.Servers{
				{
					URL: cfg.Server.Host + path.Join(
						cfg.Server.Root,
						"api",
						"v1",
					),
				},
			}

			r.Get("/swagger", func(w http.ResponseWriter, _ *http.Request) {
				respond.JSON(
					w,
					r,
					swagger,
				)
			})

			r.Handle("/docs", oamw.SwaggerUI(oamw.SwaggerUIOpts{
				Path: path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"docs",
				),
				SpecURL: cfg.Server.Host + path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"swagger",
				),
			}, nil))

			r.With(cgmw.OapiRequestValidatorWithOptions(
				swagger,
				&cgmw.Options{
					SilenceServersWarning: true,
					Options: openapi3filter.Options{
						AuthenticationFunc: func(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
							return nil
						},
					},
				},
			)).Mount("/", v1.Handler(
				v1.NewStrictHandler(
					v1.New(
						cfg,
						registry,
						sess,
						uploads,
						storage,
						minecraftService,
						forgeService,
						neoforgeService,
						quiltService,
						fabricService,
						teamsService,
						usersService,
						membersService,
						modsService,
						userModsService,
						teamModsService,
						versionsService,
						packsService,
						userPacksService,
						teamPacksService,
						buildsService,
						buildVersionsService,
					),
					make([]v1.StrictMiddlewareFunc, 0),
				),
			))

			r.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"storage",
				),
			))
		})
	})

	return mux
}

// Metrics initializes the routing of metrics and health.
func Metrics(
	cfg *config.Config,
	registry *metrics.Metrics,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", registry.Handler())

		if cfg.Server.Pprof {
			root.Mount("/debug", middleware.Profiler())
		}

		root.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			_, _ = io.WriteString(w, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			_, _ = io.WriteString(w, http.StatusText(http.StatusOK))
		})
	})

	return mux
}
