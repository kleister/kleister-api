package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/api"
	"github.com/solderapp/solder-api/assets"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/router/middleware/header"
	"github.com/solderapp/solder-api/router/middleware/location"
	"github.com/solderapp/solder-api/router/middleware/logger"
	"github.com/solderapp/solder-api/router/middleware/recovery"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/router/middleware/store"
	"github.com/solderapp/solder-api/template"
	"github.com/solderapp/solder-api/web"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	if config.Debug {
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
	e.Use(recovery.SetRecovery())
	e.Use(location.SetLocation())
	e.Use(store.SetStore())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(header.SetVersion())
	e.Use(session.SetCurrent())

	root := e.Group(config.Server.Root)
	{
		root.StaticFS(
			"/storage",
			gin.Dir(
				config.Server.Storage,
				false,
			),
		)

		root.StaticFS(
			"/assets",
			assets.Load(),
		)

		root.GET("/favicon.ico", web.GetFavicon)
		root.GET("", web.GetIndex)

		base := root.Group("/api")
		{
			base.GET("", api.GetIndex)

			//
			// Profile
			//
			profile := base.Group("/profile")
			{
				profile.Use(session.MustCurrent())

				profile.GET("", api.GetProfile)
				profile.PATCH("", api.PatchProfile)
			}

			//
			// Minecraft
			//
			minecraft := base.Group("/minecraft")
			{
				minecraft.Use(session.MustCurrent())

				minecraft.GET("", api.GetMinecrafts)
				minecraft.GET("/:minecraft", api.GetMinecrafts)
				minecraft.PATCH("", session.MustPacks("change"), api.PatchMinecraft)
			}

			minecraftBuilds := base.Group("/minecraft/:minecraft/builds")
			{
				minecraftBuilds.Use(session.MustPacks("change"))
				minecraftBuilds.Use(session.SetMinecraft())

				minecraftBuilds.GET("", api.GetMinecraftBuilds)
				minecraftBuilds.PATCH("/:build", session.SetBuild(), api.PatchMinecraftBuild)
				minecraftBuilds.DELETE("/:build", session.SetBuild(), api.DeleteMinecraftBuild)
			}

			//
			// Forge
			//
			forge := base.Group("/forge")
			{
				forge.Use(session.MustCurrent())

				forge.GET("", api.GetForges)
				forge.GET("/:forge", api.GetForges)
				forge.PATCH("", session.MustPacks("change"), api.PatchForge)
			}

			forgeBuilds := base.Group("/forge/:forge/builds")
			{
				forgeBuilds.Use(session.MustPacks("change"))
				forgeBuilds.Use(session.SetForge())

				forgeBuilds.GET("", api.GetForgeBuilds)
				forgeBuilds.PATCH("/:build", session.SetBuild(), api.PatchForgeBuild)
				forgeBuilds.DELETE("/:build", session.SetBuild(), api.DeleteForgeBuild)
			}

			//
			// Packs
			//
			packs := base.Group("/packs")
			{
				packs.Use(session.MustPacks("display"))

				packs.GET("", api.GetPacks)
				packs.GET("/:pack", session.SetPack(), api.GetPack)
				packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), api.DeletePack)
				packs.PATCH("/:pack", session.SetPack(), session.MustPacks("change"), api.PatchPack)
				packs.POST("", session.MustPacks("change"), api.PostPack)
			}

			packClients := base.Group("/packs/:pack/clients")
			{
				packClients.Use(session.MustPacks("change"))
				packClients.Use(session.SetPack())

				packClients.GET("", api.GetPackClients)
				packClients.PATCH("/:client", session.SetClient(), api.PatchPackClient)
				packClients.DELETE("/:client", session.SetClient(), api.DeletePackClient)
			}

			//
			// Builds
			//
			builds := base.Group("/packs/:pack/builds")
			{
				builds.Use(session.MustPacks("display"))
				builds.Use(session.SetPack())

				builds.GET("", api.GetBuilds)
				builds.GET("/:build", session.SetBuild(), api.GetBuild)
				builds.DELETE("/:build", session.SetBuild(), session.MustPacks("delete"), api.DeleteBuild)
				builds.PATCH("/:build", session.SetBuild(), session.MustPacks("change"), api.PatchBuild)
				builds.POST("", session.MustPacks("change"), api.PostBuild)
			}

			buildVersions := base.Group("/packs/:pack/builds/:build/versions")
			{
				buildVersions.Use(session.MustPacks("change"))
				buildVersions.Use(session.SetPack())
				buildVersions.Use(session.SetBuild())

				buildVersions.GET("", api.GetBuildVersions)
				buildVersions.PATCH("/:version", session.SetVersion(), api.PatchBuildVersion)
				buildVersions.DELETE("/:version", session.SetVersion(), api.DeleteBuildVersion)
			}

			//
			// Mods
			//
			mods := base.Group("/mods")
			{
				mods.Use(session.MustMods("display"))

				mods.GET("", api.GetMods)
				mods.GET("/:mod", session.SetMod(), api.GetMod)
				mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), api.DeleteMod)
				mods.PATCH("/:mod", session.SetMod(), session.MustMods("change"), api.PatchMod)
				mods.POST("", session.MustMods("change"), api.PostMod)
			}

			modUsers := base.Group("/mods/:mod/users")
			{
				modUsers.Use(session.MustMods("change"))
				modUsers.Use(session.SetMod())

				modUsers.GET("", api.GetModUsers)
				modUsers.PATCH("/:user", session.SetUser(), api.PatchModUser)
				modUsers.DELETE("/:user", session.SetUser(), api.DeleteModUser)
			}

			//
			// Versions
			//
			versions := base.Group("/mods/:mod/versions")
			{
				versions.Use(session.MustMods("display"))
				versions.Use(session.SetMod())

				versions.GET("", api.GetVersions)
				versions.GET("/:version", session.SetVersion(), api.GetVersion)
				versions.DELETE("/:version", session.SetVersion(), session.MustMods("delete"), api.DeleteVersion)
				versions.PATCH("/:version", session.SetVersion(), session.MustMods("change"), api.PatchVersion)
				versions.POST("", session.MustMods("change"), api.PostVersion)
			}

			versionBuilds := base.Group("/mods/:mod/versions/:version/builds")
			{
				versionBuilds.Use(session.MustMods("change"))
				versionBuilds.Use(session.SetMod())
				versionBuilds.Use(session.SetVersion())

				versionBuilds.GET("", api.GetVersionBuilds)
				versionBuilds.PATCH("/:build", session.SetBuild(), api.PatchVersionBuild)
				versionBuilds.DELETE("/:build", session.SetBuild(), api.DeleteVersionBuild)
			}

			//
			// Clients
			//
			clients := base.Group("/clients")
			{
				clients.Use(session.MustClients("display"))

				clients.GET("", api.GetClients)
				clients.GET("/:client", session.SetClient(), api.GetClient)
				clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), api.DeleteClient)
				clients.PATCH("/:client", session.SetClient(), session.MustClients("change"), api.PatchClient)
				clients.POST("", session.MustClients("change"), api.PostClient)
			}

			clientPacks := base.Group("/clients/:client/packs")
			{
				clientPacks.Use(session.MustClients("change"))
				clientPacks.Use(session.SetClient())

				clientPacks.GET("", api.GetClientPacks)
				clientPacks.PATCH("/:pack", session.SetPack(), api.PatchClientPack)
				clientPacks.DELETE("/:pack", session.SetPack(), api.DeleteClientPack)
			}

			//
			// Users
			//
			users := base.Group("/users")
			{
				users.Use(session.MustUsers("display"))

				users.GET("", api.GetUsers)
				users.GET("/:user", session.SetUser(), api.GetUser)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), api.DeleteUser)
				users.PATCH("/:user", session.SetUser(), session.MustUsers("change"), api.PatchUser)
				users.POST("", session.MustUsers("change"), api.PostUser)
			}

			userMods := base.Group("/users/:user/mods")
			{
				userMods.Use(session.MustMods("change"))
				userMods.Use(session.SetUser())

				userMods.GET("", api.GetUserMods)
				userMods.PATCH("/:mod", session.SetMod(), api.PatchUserMod)
				userMods.DELETE("/:mod", session.SetMod(), api.DeleteUserMod)
			}

			//
			// Keys
			//
			keys := base.Group("/keys")
			{
				keys.Use(session.MustKeys("display"))

				keys.GET("", api.GetKeys)
				keys.GET("/:key", session.SetKey(), api.GetKey)
				keys.DELETE("/:key", session.SetKey(), session.MustKeys("delete"), api.DeleteKey)
				keys.PATCH("/:key", session.SetKey(), session.MustKeys("change"), api.PatchKey)
				keys.POST("", session.MustKeys("change"), api.PostKey)
			}

			//
			// Solder
			//
			solder := base.Group("/")
			{
				solder.GET("/modpack/:pack", api.GetSolderPack)
				solder.GET("/modpack/:pack/:build", api.GetSolderBuild)

				solder.GET("/mod/:mod", api.GetSolderMod)
				solder.GET("/mod/:mod/:version", api.GetSolderVersion)
			}
		}
	}

	return e
}
