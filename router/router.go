package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/api"
	"github.com/kleister/kleister-api/assets"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/router/middleware/header"
	"github.com/kleister/kleister-api/router/middleware/location"
	"github.com/kleister/kleister-api/router/middleware/logger"
	"github.com/kleister/kleister-api/router/middleware/recovery"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/router/middleware/store"
	"github.com/kleister/kleister-api/template"
	"github.com/kleister/kleister-api/web"
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

		root.GET("/favicon.ico", web.Favicon)
		root.GET("", web.Index)

		base := root.Group("/api")
		{
			base.GET("", api.IndexInfo)

			//
			// Auth
			//
			auth := base.Group("/auth")
			{
				auth.GET("/logout", session.MustCurrent(), api.AuthLogout)
				auth.GET("/refresh", session.MustCurrent(), api.AuthRefresh)
				auth.POST("/login", session.MustNobody(), api.AuthLogin)
			}

			//
			// Profile
			//
			profile := base.Group("/profile")
			{
				profile.Use(session.MustCurrent())

				profile.GET("/token", api.ProfileToken)
				profile.GET("/self", api.ProfileShow)
				profile.PATCH("/self", api.ProfileUpdate)
			}

			//
			// Minecraft
			//
			minecraft := base.Group("/minecraft")
			{
				minecraft.Use(session.MustCurrent())

				minecraft.GET("", api.MinecraftIndex)
				minecraft.GET("/:minecraft", api.MinecraftIndex)
				minecraft.PATCH("", session.MustPacks("change"), api.MinecraftUpdate)
			}

			minecraftBuilds := base.Group("/minecraft/:minecraft/builds")
			{
				minecraftBuilds.Use(session.MustPacks("change"))
				minecraftBuilds.Use(session.SetMinecraft())

				minecraftBuilds.GET("", api.MinecraftBuildIndex)
				minecraftBuilds.PATCH("", api.MinecraftBuildAppend)
				minecraftBuilds.DELETE("", api.MinecraftBuildDelete)
			}

			//
			// Forge
			//
			forge := base.Group("/forge")
			{
				forge.Use(session.MustCurrent())

				forge.GET("", api.ForgeIndex)
				forge.GET("/:forge", api.ForgeIndex)
				forge.PATCH("", session.MustPacks("change"), api.ForgeUpdate)
			}

			forgeBuilds := base.Group("/forge/:forge/builds")
			{
				forgeBuilds.Use(session.MustPacks("change"))
				forgeBuilds.Use(session.SetForge())

				forgeBuilds.GET("", api.ForgeBuildIndex)
				forgeBuilds.PATCH("", api.ForgeBuildAppend)
				forgeBuilds.DELETE("", api.ForgeBuildDelete)
			}

			//
			// Packs
			//
			packs := base.Group("/packs")
			{
				packs.Use(session.MustPacks("display"))

				packs.GET("", api.PackIndex)
				packs.GET("/:pack", session.SetPack(), api.PackShow)
				packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), api.PackDelete)
				packs.PATCH("/:pack", session.SetPack(), session.MustPacks("change"), api.PackUpdate)
				packs.POST("", session.MustPacks("change"), api.PackCreate)
			}

			packClients := base.Group("/packs/:pack/clients")
			{
				packClients.Use(session.MustPacks("change"))
				packClients.Use(session.SetPack())

				packClients.GET("", api.PackClientIndex)
				packClients.PATCH("", api.PackClientAppend)
				packClients.DELETE("", api.PackClientDelete)
			}

			packUsers := base.Group("/packs/:pack/users")
			{
				packUsers.Use(session.MustPacks("change"))
				packUsers.Use(session.SetPack())

				packUsers.GET("", api.PackUserIndex)
				packUsers.PATCH("/:user", session.SetUser(), api.PackUserAppend)
				packUsers.DELETE("/:user", session.SetUser(), api.PackUserDelete)
			}

			//
			// Builds
			//
			builds := base.Group("/packs/:pack/builds")
			{
				builds.Use(session.MustPacks("display"))
				builds.Use(session.SetPack())

				builds.GET("", api.BuildIndex)
				builds.GET("/:build", session.SetBuild(), api.BuildShow)
				builds.DELETE("/:build", session.SetBuild(), session.MustPacks("delete"), api.BuildDelete)
				builds.PATCH("/:build", session.SetBuild(), session.MustPacks("change"), api.BuildUpdate)
				builds.POST("", session.MustPacks("change"), api.BuildCreate)
			}

			buildVersions := base.Group("/packs/:pack/builds/:build/versions")
			{
				buildVersions.Use(session.MustPacks("change"))
				buildVersions.Use(session.SetPack())
				buildVersions.Use(session.SetBuild())

				buildVersions.GET("", api.BuildVersionIndex)
				buildVersions.PATCH("", api.BuildVersionAppend)
				buildVersions.DELETE("", api.BuildVersionDelete)
			}

			//
			// Mods
			//
			mods := base.Group("/mods")
			{
				mods.Use(session.MustMods("display"))

				mods.GET("", api.ModIndex)
				mods.GET("/:mod", session.SetMod(), api.ModShow)
				mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), api.ModDelete)
				mods.PATCH("/:mod", session.SetMod(), session.MustMods("change"), api.ModUpdate)
				mods.POST("", session.MustMods("change"), api.ModCreate)
			}

			modUsers := base.Group("/mods/:mod/users")
			{
				modUsers.Use(session.MustMods("change"))
				modUsers.Use(session.SetMod())

				modUsers.GET("", api.ModUserIndex)
				modUsers.PATCH("", api.ModUserAppend)
				modUsers.DELETE("", api.ModUserDelete)
			}

			//
			// Versions
			//
			versions := base.Group("/mods/:mod/versions")
			{
				versions.Use(session.MustMods("display"))
				versions.Use(session.SetMod())

				versions.GET("", api.VersionIndex)
				versions.GET("/:version", session.SetVersion(), api.VersionShow)
				versions.DELETE("/:version", session.SetVersion(), session.MustMods("delete"), api.VersionDelete)
				versions.PATCH("/:version", session.SetVersion(), session.MustMods("change"), api.VersionUpdate)
				versions.POST("", session.MustMods("change"), api.VersionCreate)
			}

			versionBuilds := base.Group("/mods/:mod/versions/:version/builds")
			{
				versionBuilds.Use(session.MustMods("change"))
				versionBuilds.Use(session.SetMod())
				versionBuilds.Use(session.SetVersion())

				versionBuilds.GET("", api.VersionBuildIndex)
				versionBuilds.PATCH("", api.VersionBuildAppend)
				versionBuilds.DELETE("", api.VersionBuildDelete)
			}

			//
			// Clients
			//
			clients := base.Group("/clients")
			{
				clients.Use(session.MustClients("display"))

				clients.GET("", api.ClientIndex)
				clients.GET("/:client", session.SetClient(), api.ClientShow)
				clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), api.ClientDelete)
				clients.PATCH("/:client", session.SetClient(), session.MustClients("change"), api.ClientUpdate)
				clients.POST("", session.MustClients("change"), api.ClientCreate)
			}

			clientPacks := base.Group("/clients/:client/packs")
			{
				clientPacks.Use(session.MustClients("change"))
				clientPacks.Use(session.SetClient())

				clientPacks.GET("", api.ClientPackIndex)
				clientPacks.PATCH("", api.ClientPackAppend)
				clientPacks.DELETE("", api.ClientPackDelete)
			}

			//
			// Users
			//
			users := base.Group("/users")
			{
				users.Use(session.MustUsers("display"))

				users.GET("", api.UserIndex)
				users.GET("/:user", session.SetUser(), api.UserShow)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), api.UserDelete)
				users.PATCH("/:user", session.SetUser(), session.MustUsers("change"), api.UserUpdate)
				users.POST("", session.MustUsers("change"), api.UserCreate)
			}

			userMods := base.Group("/users/:user/mods")
			{
				userMods.Use(session.MustMods("change"))
				userMods.Use(session.SetUser())

				userMods.GET("", api.UserModIndex)
				userMods.PATCH("", api.UserModAppend)
				userMods.DELETE("", api.UserModDelete)
			}

			userPacks := base.Group("/users/:user/packs")
			{
				userPacks.Use(session.MustPacks("change"))
				userPacks.Use(session.SetUser())

				userPacks.GET("", api.UserPackIndex)
				userPacks.PATCH("", api.UserPackAppend)
				userPacks.DELETE("", api.UserPackDelete)
			}

			//
			// Solder
			//
			solder := base.Group("/")
			{
				solder.GET("/modpack", api.SolderPacks)
				solder.GET("/modpack/:pack", api.SolderPack)
				solder.GET("/modpack/:pack/:build", api.SolderBuild)

				solder.GET("/mod", api.SolderMods)
				solder.GET("/mod/:mod", api.SolderMod)
				solder.GET("/mod/:mod/:version", api.SolderVersion)
			}
		}
	}

	return e
}
