package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/assets"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/controller"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/header"
	"github.com/solderapp/solder-api/router/middleware/logger"
	"github.com/solderapp/solder-api/router/middleware/recovery"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/template"
	"github.com/solderapp/solder-api/web"
)

// Load initializes the routing of the application.
func Load(cfg *config.Config, middleware ...gin.HandlerFunc) http.Handler {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.SetHTMLTemplate(
		template.Load(),
	)

	e.Use(middleware...)
	e.Use(context.SetRoot())
	e.Use(logger.SetLogger())
	e.Use(recovery.SetRecovery())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(header.SetVersion())
	e.Use(session.SetCurrent())

	r := e.Group(cfg.Server.Root)
	{
		r.StaticFS(
			"/storage",
			gin.Dir(
				cfg.Server.Storage,
				false,
			),
		)

		r.StaticFS(
			"/assets",
			assets.Load(),
		)

		r.GET("/favicon.ico", web.GetFavicon)
		r.GET("", web.GetIndex)

		api := r.Group("/api")
		{
			api.GET("", controller.GetIndex)

			//
			// Profile
			//
			profile := api.Group("/profile")
			{
				profile.Use(session.MustCurrent())

				profile.GET("", controller.GetProfile)
				profile.PATCH("", controller.PatchProfile)
			}

			//
			// Minecraft
			//
			minecraft := api.Group("/minecraft")
			{
				minecraft.Use(session.MustCurrent())

				minecraft.GET("", controller.GetMinecrafts)
				minecraft.GET("/:minecraft", controller.GetMinecrafts)
				minecraft.PATCH("", session.MustPacks("change"), controller.PatchMinecraft)
			}

			minecraftBuilds := api.Group("/minecraft/:minecraft/builds")
			{
				minecraftBuilds.Use(session.MustPacks("change"))
				minecraftBuilds.Use(session.SetMinecraft())

				minecraftBuilds.GET("", controller.GetMinecraftBuilds)
				minecraftBuilds.PATCH("/:build", session.SetBuild(), controller.PatchMinecraftBuild)
				minecraftBuilds.DELETE("/:build", session.SetBuild(), controller.DeleteMinecraftBuild)
			}

			//
			// Forge
			//
			forge := api.Group("/forge")
			{
				forge.Use(session.MustCurrent())

				forge.GET("", controller.GetForges)
				forge.GET("/:forge", controller.GetForges)
				forge.PATCH("", session.MustPacks("change"), controller.PatchForge)
			}

			forgeBuilds := api.Group("/forge/:forge/builds")
			{
				forgeBuilds.Use(session.MustPacks("change"))
				forgeBuilds.Use(session.SetForge())

				forgeBuilds.GET("", controller.GetForgeBuilds)
				forgeBuilds.PATCH("/:build", session.SetBuild(), controller.PatchForgeBuild)
				forgeBuilds.DELETE("/:build", session.SetBuild(), controller.DeleteForgeBuild)
			}

			//
			// Packs
			//
			packs := api.Group("/packs")
			{
				packs.Use(session.MustPacks("display"))

				packs.GET("", controller.GetPacks)
				packs.GET("/:pack", session.SetPack(), controller.GetPack)
				packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), controller.DeletePack)
				packs.PATCH("/:pack", session.SetPack(), session.MustPacks("change"), controller.PatchPack)
				packs.POST("", session.MustPacks("change"), controller.PostPack)
			}

			packClients := api.Group("/packs/:pack/clients")
			{
				packClients.Use(session.MustPacks("change"))
				packClients.Use(session.SetPack())

				packClients.GET("", controller.GetPackClients)
				packClients.PATCH("/:client", session.SetClient(), controller.PatchPackClient)
				packClients.DELETE("/:client", session.SetClient(), controller.DeletePackClient)
			}

			//
			// Builds
			//
			builds := api.Group("/packs/:pack/builds")
			{
				builds.Use(session.MustPacks("display"))
				builds.Use(session.SetPack())

				builds.GET("", controller.GetBuilds)
				builds.GET("/:build", session.SetBuild(), controller.GetBuild)
				builds.DELETE("/:build", session.SetBuild(), session.MustPacks("delete"), controller.DeleteBuild)
				builds.PATCH("/:build", session.SetBuild(), session.MustPacks("change"), controller.PatchBuild)
				builds.POST("", session.MustPacks("change"), controller.PostBuild)
			}

			buildVersions := api.Group("/packs/:pack/builds/:build/versions")
			{
				buildVersions.Use(session.MustPacks("change"))
				buildVersions.Use(session.SetPack())
				buildVersions.Use(session.SetBuild())

				buildVersions.GET("", controller.GetBuildVersions)
				buildVersions.PATCH("/:version", session.SetVersion(), controller.PatchBuildVersion)
				buildVersions.DELETE("/:version", session.SetVersion(), controller.DeleteBuildVersion)
			}

			//
			// Mods
			//
			mods := api.Group("/mods")
			{
				mods.Use(session.MustMods("display"))

				mods.GET("", controller.GetMods)
				mods.GET("/:mod", session.SetMod(), controller.GetMod)
				mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), controller.DeleteMod)
				mods.PATCH("/:mod", session.SetMod(), session.MustMods("change"), controller.PatchMod)
				mods.POST("", session.MustMods("change"), controller.PostMod)
			}

			modUsers := api.Group("/mods/:mod/users")
			{
				modUsers.Use(session.MustMods("change"))
				modUsers.Use(session.SetMod())

				modUsers.GET("", controller.GetModUsers)
				modUsers.PATCH("/:user", session.SetUser(), controller.PatchModUser)
				modUsers.DELETE("/:user", session.SetUser(), controller.DeleteModUser)
			}

			//
			// Versions
			//
			versions := api.Group("/mods/:mod/versions")
			{
				versions.Use(session.MustMods("display"))
				versions.Use(session.SetMod())

				versions.GET("", controller.GetVersions)
				versions.GET("/:version", session.SetVersion(), controller.GetVersion)
				versions.DELETE("/:version", session.SetVersion(), session.MustMods("delete"), controller.DeleteVersion)
				versions.PATCH("/:version", session.SetVersion(), session.MustMods("change"), controller.PatchVersion)
				versions.POST("", session.MustMods("change"), controller.PostVersion)
			}

			versionBuilds := api.Group("/mods/:mod/versions/:version/builds")
			{
				versionBuilds.Use(session.MustMods("change"))
				versionBuilds.Use(session.SetMod())
				versionBuilds.Use(session.SetVersion())

				versionBuilds.GET("", controller.GetVersionBuilds)
				versionBuilds.PATCH("/:build", session.SetBuild(), controller.PatchVersionBuild)
				versionBuilds.DELETE("/:build", session.SetBuild(), controller.DeleteVersionBuild)
			}

			//
			// Clients
			//
			clients := api.Group("/clients")
			{
				clients.Use(session.MustClients("display"))

				clients.GET("", controller.GetClients)
				clients.GET("/:client", session.SetClient(), controller.GetClient)
				clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), controller.DeleteClient)
				clients.PATCH("/:client", session.SetClient(), session.MustClients("change"), controller.PatchClient)
				clients.POST("", session.MustClients("change"), controller.PostClient)
			}

			clientPacks := api.Group("/clients/:client/packs")
			{
				clientPacks.Use(session.MustClients("change"))
				clientPacks.Use(session.SetClient())

				clientPacks.GET("", controller.GetClientPacks)
				clientPacks.PATCH("/:pack", session.SetPack(), controller.PatchClientPack)
				clientPacks.DELETE("/:pack", session.SetPack(), controller.DeleteClientPack)
			}

			//
			// Users
			//
			users := api.Group("/users")
			{
				users.Use(session.MustUsers("display"))

				users.GET("", controller.GetUsers)
				users.GET("/:user", session.SetUser(), controller.GetUser)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), controller.DeleteUser)
				users.PATCH("/:user", session.SetUser(), session.MustUsers("change"), controller.PatchUser)
				users.POST("", session.MustUsers("change"), controller.PostUser)
			}

			userMods := api.Group("/users/:user/mods")
			{
				userMods.Use(session.MustMods("change"))
				userMods.Use(session.SetUser())

				userMods.GET("", controller.GetUserMods)
				userMods.PATCH("/:mod", session.SetMod(), controller.PatchUserMod)
				userMods.DELETE("/:mod", session.SetMod(), controller.DeleteUserMod)
			}

			//
			// Keys
			//
			keys := api.Group("/keys")
			{
				keys.Use(session.MustKeys("display"))

				keys.GET("", controller.GetKeys)
				keys.GET("/:key", session.SetKey(), controller.GetKey)
				keys.DELETE("/:key", session.SetKey(), session.MustKeys("delete"), controller.DeleteKey)
				keys.PATCH("/:key", session.SetKey(), session.MustKeys("change"), controller.PatchKey)
				keys.POST("", session.MustKeys("change"), controller.PostKey)
			}

			//
			// Solder
			//
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
