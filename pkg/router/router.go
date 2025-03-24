package router

import (
	"encoding/json"
	"net/http"
	"path"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	oamw "github.com/go-openapi/runtime/middleware"
	v1 "github.com/kleister/kleister-api/pkg/api/v1"
	"github.com/kleister/kleister-api/pkg/authn"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/handler"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/middleware/header"
	"github.com/kleister/kleister-api/pkg/scim"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/upload"
	cgmw "github.com/oapi-codegen/nethttp-middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	registry *metrics.Metrics,
	identity *authn.Authn,
	uploads upload.Upload,
	storage *store.Store,
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
			Msg("Accesslog")
	}))

	mux.Use(render.SetContentType(render.ContentTypeJSON))
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)
	mux.Use(current.Middleware)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		if cfg.Scim.Enabled {
			srv, err := scim.New(
				scim.WithRoot(
					path.Join(
						cfg.Server.Root,
						"api",
						"scim",
						"v2",
					),
				),
				scim.WithStore(
					storage.Handle(),
				),
				scim.WithConfig(
					cfg.Scim,
				),
			).Server()

			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to linitialize scim server")
			}

			root.Mount("/api/scim/v2", srv)
		}

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
					URL: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
					),
				},
			}

			if cfg.Server.Docs {
				r.Get("/spec", func(w http.ResponseWriter, r *http.Request) {
					render.Status(r, http.StatusOK)
					render.JSON(w, r, swagger)
				})

				r.Handle("/docs", oamw.SwaggerUI(oamw.SwaggerUIOpts{
					Path: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
						"docs",
					),
					SpecURL: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
						"spec",
					),
				}, nil))
			}

			apiv1 := v1.New(
				cfg,
				registry,
				identity,
				uploads,
				storage,
			)

			wrapper := v1.ServerInterfaceWrapper{
				Handler: apiv1,
				ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
					apiv1.RenderNotify(w, r, v1.Notification{
						Message: v1.ToPtr(err.Error()),
						Status:  v1.ToPtr(http.StatusBadRequest),
					})
				},
			}

			r.With(cgmw.OapiRequestValidatorWithOptions(
				swagger,
				&cgmw.Options{
					SilenceServersWarning: true,
					Options: openapi3filter.Options{
						AuthenticationFunc: apiv1.Authentication,
					},
					ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(statusCode)

						_ = json.NewEncoder(w).Encode(v1.Notification{
							Message: v1.ToPtr(message),
							Status:  v1.ToPtr(statusCode),
						})
					},
				},
			)).Route("/", func(r chi.Router) {
				r.Route("/auth", func(r chi.Router) {
					r.Group(func(r chi.Router) {
						r.Post("/redirect", wrapper.RedirectAuth)
						r.Post("/login", wrapper.LoginAuth)
						r.Get("/refresh", wrapper.RefreshAuth)
						r.Get("/verify", wrapper.VerifyAuth)
					})

					r.Group(func(r chi.Router) {
						r.Get("/providers", wrapper.ListProviders)

						r.Route("/{provider}", func(r chi.Router) {
							r.Use(render.SetContentType(render.ContentTypeHTML))

							r.Get("/callback", wrapper.CallbackProvider)
							r.Get("/request", wrapper.RequestProvider)
						})
					})
				})

				r.Route("/profile", func(r chi.Router) {
					r.Get("/self", wrapper.ShowProfile)
					r.Put("/self", wrapper.UpdateProfile)
					r.Get("/token", wrapper.TokenProfile)
				})

				r.Route("/minecraft", func(r chi.Router) {
					r.Get("/", wrapper.ListMinecrafts)
					r.With(apiv1.AllowAdminAccessOnly).Put("/", wrapper.UpdateMinecraft)

					r.Route("/{minecraft_id}/builds", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.MinecraftToContext)

						r.Get("/", wrapper.ListMinecraftBuilds)
						r.Delete("/", wrapper.DeleteMinecraftFromBuild)
						r.Post("/", wrapper.AttachMinecraftToBuild)
					})
				})

				r.Route("/forge", func(r chi.Router) {
					r.Get("/", wrapper.ListForges)
					r.With(apiv1.AllowAdminAccessOnly).Put("/", wrapper.UpdateForge)

					r.Route("/{forge_id}/builds", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.ForgeToContext)

						r.Get("/", wrapper.ListForgeBuilds)
						r.Delete("/", wrapper.DeleteForgeFromBuild)
						r.Post("/", wrapper.AttachForgeToBuild)
					})
				})

				r.Route("/neoforge", func(r chi.Router) {
					r.Get("/", wrapper.ListNeoforges)
					r.With(apiv1.AllowAdminAccessOnly).Put("/", wrapper.UpdateNeoforge)

					r.Route("/{neoforge_id}/builds", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.NeoforgeToContext)

						r.Get("/", wrapper.ListNeoforgeBuilds)
						r.Delete("/", wrapper.DeleteNeoforgeFromBuild)
						r.Post("/", wrapper.AttachNeoforgeToBuild)
					})
				})

				r.Route("/quilt", func(r chi.Router) {
					r.Get("/", wrapper.ListQuilts)
					r.With(apiv1.AllowAdminAccessOnly).Put("/", wrapper.UpdateQuilt)

					r.Route("/{quilt_id}/builds", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.QuiltToContext)

						r.Get("/", wrapper.ListQuiltBuilds)
						r.Delete("/", wrapper.DeleteQuiltFromBuild)
						r.Post("/", wrapper.AttachQuiltToBuild)
					})
				})

				r.Route("/fabric", func(r chi.Router) {
					r.Get("/", wrapper.ListFabrics)
					r.With(apiv1.AllowAdminAccessOnly).Put("/", wrapper.UpdateFabric)

					r.Route("/{fabric_id}/builds", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.FabricToContext)

						r.Get("/", wrapper.ListFabricBuilds)
						r.Delete("/", wrapper.DeleteFabricFromBuild)
						r.Post("/", wrapper.AttachFabricToBuild)
					})
				})

				r.Route("/packs", func(r chi.Router) {
					r.Get("/", wrapper.ListPacks)
					r.With(apiv1.AllowCreatePack).Post("/", wrapper.CreatePack)

					r.Route("/{pack_id}", func(r chi.Router) {
						r.Use(apiv1.PackToContext)
						r.Use(apiv1.AllowShowPack)

						r.Get("/", wrapper.ShowPack)
						r.With(apiv1.AllowManagePack).Delete("/", wrapper.DeletePack)
						r.With(apiv1.AllowManagePack).Put("/", wrapper.UpdatePack)

						r.Route("/avatar", func(r chi.Router) {
							r.Use(apiv1.AllowManagePack)

							r.Delete("/", wrapper.DeletePackAvatar)
							r.Post("/", wrapper.CreatePackAvatar)
						})

						r.Route("/users", func(r chi.Router) {
							r.Use(apiv1.AllowManagePack)

							r.Get("/", wrapper.ListPackUsers)
							r.Delete("/", wrapper.DeletePackFromUser)
							r.Post("/", wrapper.AttachPackToUser)
							r.Put("/", wrapper.PermitPackUser)
						})

						r.Route("/groups", func(r chi.Router) {
							r.Use(apiv1.AllowManagePack)

							r.Get("/", wrapper.ListPackGroups)
							r.Delete("/", wrapper.DeletePackFromGroup)
							r.Post("/", wrapper.AttachPackToGroup)
							r.Put("/", wrapper.PermitPackGroup)
						})

						r.Route("/builds", func(r chi.Router) {
							r.Get("/", wrapper.ListBuilds)
							r.With(apiv1.AllowManageBuild).Post("/", wrapper.CreateBuild)

							r.Route("/{build_id}", func(r chi.Router) {
								r.Use(apiv1.BuildToContext)
								r.Use(apiv1.AllowShowBuild)

								r.Get("/", wrapper.ShowBuild)
								r.With(apiv1.AllowManageBuild).Delete("/", wrapper.DeleteBuild)
								r.With(apiv1.AllowManageBuild).Put("/", wrapper.UpdateBuild)

								r.Route("/versions", func(r chi.Router) {
									r.Use(apiv1.AllowManageBuild)

									r.Get("/", wrapper.ListBuildVersions)
									r.Delete("/", wrapper.DeleteBuildFromVersion)
									r.Post("/", wrapper.AttachBuildToVersion)
								})
							})
						})
					})
				})

				r.Route("/mods", func(r chi.Router) {
					r.Get("/", wrapper.ListMods)
					r.With(apiv1.AllowCreateMod).Post("/", wrapper.CreateMod)

					r.Route("/{mod_id}", func(r chi.Router) {
						r.Use(apiv1.ModToContext)
						r.Use(apiv1.AllowShowMod)

						r.Get("/", wrapper.ShowMod)
						r.With(apiv1.AllowManageMod).Delete("/", wrapper.DeleteMod)
						r.With(apiv1.AllowManageMod).Put("/", wrapper.UpdateMod)

						r.Route("/avatar", func(r chi.Router) {
							r.Use(apiv1.AllowManageMod)

							r.Delete("/", wrapper.DeleteModAvatar)
							r.Post("/", wrapper.CreateModAvatar)
						})

						r.Route("/users", func(r chi.Router) {
							r.Use(apiv1.AllowManageMod)

							r.Get("/", wrapper.ListModUsers)
							r.Delete("/", wrapper.DeleteModFromUser)
							r.Post("/", wrapper.AttachModToUser)
							r.Put("/", wrapper.PermitModUser)
						})

						r.Route("/groups", func(r chi.Router) {
							r.Use(apiv1.AllowManageMod)

							r.Get("/", wrapper.ListModGroups)
							r.Delete("/", wrapper.DeleteModFromGroup)
							r.Post("/", wrapper.AttachModToGroup)
							r.Put("/", wrapper.PermitModGroup)
						})

						r.Route("/versions", func(r chi.Router) {
							r.Get("/", wrapper.ListVersions)
							r.With(apiv1.AllowManageVersion).Post("/", wrapper.CreateVersion)

							r.Route("/{version_id}", func(r chi.Router) {
								r.Use(apiv1.VersionToContext)
								r.Use(apiv1.AllowShowVersion)

								r.Get("/", wrapper.ShowVersion)
								r.With(apiv1.AllowManageVersion).Delete("/", wrapper.DeleteVersion)
								r.With(apiv1.AllowManageVersion).Put("/", wrapper.UpdateVersion)

								r.Route("/builds", func(r chi.Router) {
									r.Use(apiv1.AllowManageVersion)

									r.Get("/", wrapper.ListVersionBuilds)
									r.Delete("/", wrapper.DeleteVersionFromBuild)
									r.Post("/", wrapper.AttachVersionToBuild)
								})
							})
						})
					})
				})

				r.Route("/groups", func(r chi.Router) {
					r.Get("/", wrapper.ListGroups)
					r.With(apiv1.AllowAdminAccessOnly).Post("/", wrapper.CreateGroup)

					r.Route("/{group_id}", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.GroupToContext)

						r.Get("/", wrapper.ShowGroup)
						r.Delete("/", wrapper.DeleteGroup)
						r.Put("/", wrapper.UpdateGroup)

						r.Route("/users", func(r chi.Router) {
							r.Get("/", wrapper.ListGroupUsers)
							r.Delete("/", wrapper.DeleteGroupFromUser)
							r.Post("/", wrapper.AttachGroupToUser)
							r.Put("/", wrapper.PermitGroupUser)
						})

						r.Route("/mods", func(r chi.Router) {
							r.Get("/", wrapper.ListGroupMods)
							r.Delete("/", wrapper.DeleteGroupFromMod)
							r.Post("/", wrapper.AttachGroupToMod)
							r.Put("/", wrapper.PermitGroupMod)
						})

						r.Route("/packs", func(r chi.Router) {
							r.Get("/", wrapper.ListGroupPacks)
							r.Delete("/", wrapper.DeleteGroupFromPack)
							r.Post("/", wrapper.AttachGroupToPack)
							r.Put("/", wrapper.PermitGroupPack)
						})
					})
				})

				r.Route("/users", func(r chi.Router) {
					r.Get("/", wrapper.ListUsers)
					r.With(apiv1.AllowAdminAccessOnly).Post("/", wrapper.CreateUser)

					r.Route("/{user_id}", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.UserToContext)

						r.Get("/", wrapper.ShowUser)
						r.Delete("/", wrapper.DeleteUser)
						r.Put("/", wrapper.UpdateUser)

						r.Route("/groups", func(r chi.Router) {
							r.Get("/", wrapper.ListUserGroups)
							r.Delete("/", wrapper.DeleteUserFromGroup)
							r.Post("/", wrapper.AttachUserToGroup)
							r.Put("/", wrapper.PermitUserGroup)
						})

						r.Route("/mods", func(r chi.Router) {
							r.Get("/", wrapper.ListUserMods)
							r.Delete("/", wrapper.DeleteUserFromMod)
							r.Post("/", wrapper.AttachUserToMod)
							r.Put("/", wrapper.PermitUserMod)
						})

						r.Route("/packs", func(r chi.Router) {
							r.Get("/", wrapper.ListUserPacks)
							r.Delete("/", wrapper.DeleteUserFromPack)
							r.Post("/", wrapper.AttachUserToPack)
							r.Put("/", wrapper.PermitUserPack)
						})
					})
				})
			})

			r.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"storage",
				),
			))
		})

		handlers := handler.New(cfg)
		root.Get("/", handlers.Index())
		root.Get("/favicon.svg", handlers.Favicon())
		root.Get("/config.json", handlers.Config())
		root.Get("/manifest.json", handlers.Manifest())
		root.Handle("/assets/*", handlers.Assets())
		root.NotFound(handlers.Index())
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

		if cfg.Metrics.Pprof {
			root.Mount("/debug", middleware.Profiler())
		}

		root.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
		})
	})

	return mux
}
