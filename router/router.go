package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/controller"
	"github.com/solderapp/solder-api/router/middleware/error"
	"github.com/solderapp/solder-api/router/middleware/header"
	"github.com/solderapp/solder-api/router/middleware/logger"
	"github.com/solderapp/solder-api/router/middleware/recover"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/static"
	"github.com/solderapp/solder-api/template"
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
	e.Use(session.SetCurrent())

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

			api.GET("/profile", session.MustCurrent(), controller.GetProfile)
			api.PATCH("/profile", session.MustCurrent(), controller.PatchProfile)

			api.GET("/minecraft", session.MustCurrent(), controller.GetMinecraft)
			api.GET("/minecraft/:filter", session.MustCurrent(), controller.CompleteMinecraft)
			api.PATCH("/minecraft", session.MustPacks("change"), controller.PatchMinecraft)

			api.GET("/forge", session.MustCurrent(), controller.GetForge)
			api.GET("/forge/:filter", session.MustCurrent(), controller.CompleteForge)
			api.PATCH("/forge", session.MustPacks("change"), controller.PatchForge)

			packs := api.Group("/packs")
			{
				packs.Use(session.MustPacks("display"))

				packs.GET("", controller.GetPacks)
				packs.GET("/:pack", session.SetPack(), controller.GetPack)
				packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), controller.DeletePack)
				packs.PATCH("/:pack", session.SetPack(), session.MustPacks("change"), controller.PatchPack)
				packs.POST("", session.MustPacks("change"), controller.PostPack)
			}

			builds := api.Group("/packs/:pack/builds")
			{
				builds.Use(session.SetPack())
				builds.Use(session.MustPacks("display"))

				builds.GET("", controller.GetBuilds)
				builds.GET("/:build", session.SetBuild(), controller.GetBuild)
				builds.DELETE("/:build", session.SetBuild(), session.MustPacks("delete"), controller.DeleteBuild)
				builds.PATCH("/:build", session.SetBuild(), session.MustPacks("change"), controller.PatchBuild)
				builds.POST("", session.MustPacks("change"), controller.PostBuild)
			}

			mods := api.Group("/mods")
			{
				mods.Use(session.MustMods("display"))

				mods.GET("", controller.GetMods)
				mods.GET("/:mod", session.SetMod(), controller.GetMod)
				mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), controller.DeleteMod)
				mods.PATCH("/:mod", session.SetMod(), session.MustMods("change"), controller.PatchMod)
				mods.POST("", session.MustMods("change"), controller.PostMod)
			}

			versions := api.Group("/mods/:mod/versions")
			{
				versions.Use(session.SetMod())
				versions.Use(session.MustMods("display"))

				versions.GET("", controller.GetVersions)
				versions.GET("/:version", session.SetVersion(), controller.GetVersion)
				versions.DELETE("/:version", session.SetVersion(), session.MustMods("delete"), controller.DeleteVersion)
				versions.PATCH("/:version", session.SetVersion(), session.MustMods("change"), controller.PatchVersion)
				versions.POST("", session.MustMods("change"), controller.PostVersion)
			}

			users := api.Group("/users")
			{
				users.Use(session.MustUsers("display"))

				users.GET("", controller.GetUsers)
				users.GET("/:user", session.SetUser(), controller.GetUser)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), controller.DeleteUser)
				users.PATCH("/:user", session.SetUser(), session.MustUsers("change"), controller.PatchUser)
				users.POST("", session.MustUsers("change"), controller.PostUser)
			}

			keys := api.Group("/keys")
			{
				keys.Use(session.MustKeys("display"))

				keys.GET("", controller.GetKeys)
				keys.GET("/:key", session.SetKey(), controller.GetKey)
				keys.DELETE("/:key", session.SetKey(), session.MustKeys("delete"), controller.DeleteKey)
				keys.PATCH("/:key", session.SetKey(), session.MustKeys("change"), controller.PatchKey)
				keys.POST("", session.MustKeys("change"), controller.PostKey)
			}

			clients := api.Group("/clients")
			{
				clients.Use(session.MustClients("display"))

				clients.GET("", controller.GetClients)
				clients.GET("/:client", session.SetClient(), controller.GetClient)
				clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), controller.DeleteClient)
				clients.PATCH("/:client", session.SetClient(), session.MustClients("change"), controller.PatchClient)
				clients.POST("", session.MustClients("change"), controller.PostClient)
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
