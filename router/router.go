package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/config"
	"github.com/solderapp/solder/controller"
	"github.com/solderapp/solder/router/middleware/error"
	"github.com/solderapp/solder/router/middleware/header"
	"github.com/solderapp/solder/router/middleware/logger"
	"github.com/solderapp/solder/router/middleware/recover"
	"github.com/solderapp/solder/router/middleware/session"
	"github.com/solderapp/solder/static"
	"github.com/solderapp/solder/template"
)

// Load initializes the routing of the application.
func Load(cfg *config.Config, middleware ...gin.HandlerFunc) http.Handler {
	if cfg.Develop {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.SetHTMLTemplate(
		template.Load(),
	)

	e.Use(middleware...)
	e.Use(logger.SetLogger())
	e.Use(recover.SetRecover())
	e.Use(error.SetError())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(session.SetUser())

	r := e.Group(cfg.Server.Root)
	{
		r.StaticFS(
			"/assets",
			static.Load(),
		)

		r.StaticFile(
			"/favicon.ico",
			string(
				static.MustAsset(
					"images/favicon.ico",
				),
			),
		)

		r.GET("", controller.GetIndex)

		api := r.Group("/api")
		{
			api.GET("", controller.GetAPI)

			api.GET("/profile", controller.GetProfile)
			api.PATCH("/profile", controller.PatchProfile)

			api.GET("/minecraft", controller.GetMinecraft)
			api.GET("/minecraft/:filter", controller.CompleteMinecraft)
			api.PATCH("/minecraft", controller.PatchMinecraft)

			api.GET("/forge", controller.GetForge)
			api.GET("/forge/:filter", controller.CompleteForge)
			api.PATCH("/forge", controller.PatchForge)

			packs := api.Group("/packs")
			{
				packs.GET("", controller.GetPacks)
				packs.GET("/:pack", controller.GetPack)
				packs.DELETE("/:pack", controller.DeletePack)
				packs.PATCH("/:pack", controller.PatchPack)
				packs.POST("", controller.PostPack)
			}

			builds := api.Group("/packs/:pack/builds")
			{
				builds.GET("", controller.GetBuilds)
				builds.GET("/:build", controller.GetBuild)
				builds.DELETE("/:build", controller.DeleteBuild)
				builds.PATCH("/:build", controller.PatchBuild)
				builds.POST("", controller.PostBuild)
			}

			mods := api.Group("/mods")
			{
				mods.GET("", controller.GetMods)
				mods.GET("/:mod", controller.GetMod)
				mods.DELETE("/:mod", controller.DeleteMod)
				mods.PATCH("/:mod", controller.PatchMod)
				mods.POST("", controller.PostMod)
			}

			versions := api.Group("/mods/:mod/versions")
			{
				versions.GET("", controller.GetVersions)
				versions.GET("/:version", controller.GetVersion)
				versions.DELETE("/:version", controller.DeleteVersion)
				versions.PATCH("/:version", controller.PatchVersion)
				versions.POST("", controller.PostVersion)
			}

			users := api.Group("/users")
			{
				users.GET("", controller.GetUsers)
				users.GET("/:user", controller.GetUser)
				users.DELETE("/:user", controller.DeleteUser)
				users.PATCH("/:user", controller.PatchUser)
				users.POST("", controller.PostUser)
			}

			keys := api.Group("/keys")
			{
				keys.GET("", controller.GetKeys)
				keys.GET("/:key", controller.GetKey)
				keys.DELETE("/:key", controller.DeleteKey)
				keys.PATCH("/:key", controller.PatchKey)
				keys.POST("", controller.PostKey)
			}

			clients := api.Group("/clients")
			{
				clients.GET("", controller.GetClients)
				clients.GET("/:client", controller.GetClient)
				clients.DELETE("/:client", controller.DeleteClient)
				clients.PATCH("/:client", controller.PatchClient)
				clients.POST("", controller.PostClient)
			}

			solder := api.Group("/")
			{
				solder.GET("/modpack/:pack", controller.GetSolderPack)
				solder.GET("/modpack/:pack/:build", controller.GetSolderBuild)

				solder.GET("/mod/:mod", controller.GetSolderMod)
				solder.GET("/mod/:mod/:version", controller.GetSolderVersion)
			}
		}
	}

	return e
}
