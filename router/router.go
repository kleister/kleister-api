package router

import (
	"net/http"
	"path"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/api"
	"github.com/kleister/kleister-api/assets"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/router/middleware/header"
	"github.com/kleister/kleister-api/router/middleware/logger"
	"github.com/kleister/kleister-api/router/middleware/prometheus"
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
	e.Use(store.SetStore())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(header.SetVersion())
	e.Use(session.SetCurrent())

	if config.Server.Pprof {
		pprof.Register(
			e,
			&pprof.Options{
				RoutePrefix: path.Join(config.Server.Root, "debug", "pprof"),
			},
		)
	}

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
		root.GET("/metrics", prometheus.Handler())

		base := root.Group("/api")
		{
			base.GET("", api.IndexInfo)

			//
			// Auth
			//
			auth := base.Group("/auth")
			{
				auth.GET("/verify/:token", api.AuthVerify)
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
				profile.PUT("/self", api.ProfileUpdate)
			}

			//
			// Keys
			//
			keys := base.Group("/keys")
			{
				keys.Use(session.MustCurrent())
				keys.Use(session.MustKeys("display"))

				keys.GET("", api.KeyIndex)
				keys.GET("/:key", session.SetKey(), api.KeyShow)
				keys.DELETE("/:key", session.SetKey(), session.MustKeys("delete"), api.KeyDelete)
				keys.PUT("/:key", session.SetKey(), session.MustKeys("change"), api.KeyUpdate)
				keys.POST("", session.MustKeys("change"), api.KeyCreate)
			}

			//
			// Minecraft
			//
			minecraft := base.Group("/minecraft")
			{
				minecraft.Use(session.MustCurrent())

				minecraft.GET("", api.MinecraftIndex)
				minecraft.GET("/:minecraft", api.MinecraftIndex)
				minecraft.PUT("", session.MustAdmin(), api.MinecraftUpdate)
			}

			minecraftBuilds := base.Group("/minecraft/:minecraft/builds")
			{
				minecraftBuilds.Use(session.MustCurrent())
				minecraftBuilds.Use(session.SetMinecraft())

				minecraftBuilds.GET("", session.MustMinecraftBuilds("display"), api.MinecraftBuildIndex)
				minecraftBuilds.POST("", session.MustMinecraftBuilds("change"), api.MinecraftBuildAppend)
				minecraftBuilds.DELETE("", session.MustMinecraftBuilds("change"), api.MinecraftBuildDelete)
			}

			//
			// Forge
			//
			forge := base.Group("/forge")
			{
				forge.Use(session.MustCurrent())

				forge.GET("", api.ForgeIndex)
				forge.GET("/:forge", api.ForgeIndex)
				forge.PUT("", session.MustAdmin(), api.ForgeUpdate)
			}

			forgeBuilds := base.Group("/forge/:forge/builds")
			{
				forgeBuilds.Use(session.MustCurrent())
				forgeBuilds.Use(session.SetForge())

				forgeBuilds.GET("", session.MustForgeBuilds("display"), api.ForgeBuildIndex)
				forgeBuilds.POST("", session.MustForgeBuilds("change"), api.ForgeBuildAppend)
				forgeBuilds.DELETE("", session.MustForgeBuilds("change"), api.ForgeBuildDelete)
			}

			//
			// Packs
			//
			packs := base.Group("/packs")
			{
				packs.Use(session.MustCurrent())
				packs.Use(session.MustPacks("display"))

				packs.GET("", api.PackIndex)
				packs.GET("/:pack", session.SetPack(), api.PackShow)
				packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), api.PackDelete)
				packs.PUT("/:pack", session.SetPack(), session.MustPacks("change"), api.PackUpdate)
				packs.POST("", session.MustPacks("change"), api.PackCreate)
			}

			packClients := base.Group("/packs/:pack/clients")
			{
				packClients.Use(session.MustCurrent())
				packClients.Use(session.SetPack())

				packClients.GET("", session.MustPackClients("display"), api.PackClientIndex)
				packClients.POST("", session.MustPackClients("change"), api.PackClientAppend)
				packClients.DELETE("", session.MustPackClients("change"), api.PackClientDelete)
			}

			packUsers := base.Group("/packs/:pack/users")
			{
				packUsers.Use(session.MustCurrent())
				packUsers.Use(session.SetPack())

				packUsers.GET("", session.MustPackUsers("display"), api.PackUserIndex)
				packUsers.POST("", session.MustPackUsers("change"), api.PackUserAppend)
				packUsers.PUT("", session.MustPackUsers("change"), api.PackUserPerm)
				packUsers.DELETE("", session.MustPackUsers("change"), api.PackUserDelete)
			}

			packTeams := base.Group("/packs/:pack/teams")
			{
				packTeams.Use(session.MustCurrent())
				packTeams.Use(session.SetPack())

				packTeams.GET("", session.MustPackTeams("display"), api.PackTeamIndex)
				packTeams.POST("", session.MustPackTeams("change"), api.PackTeamAppend)
				packTeams.PUT("", session.MustPackTeams("change"), api.PackTeamPerm)
				packTeams.DELETE("", session.MustPackTeams("change"), api.PackTeamDelete)
			}

			//
			// Builds
			//
			builds := base.Group("/packs/:pack/builds")
			{
				builds.Use(session.MustCurrent())
				builds.Use(session.SetPack())
				builds.Use(session.MustBuilds("display"))

				builds.GET("", api.BuildIndex)
				builds.GET("/:build", session.SetBuild(), api.BuildShow)
				builds.DELETE("/:build", session.SetBuild(), session.MustBuilds("delete"), api.BuildDelete)
				builds.PUT("/:build", session.SetBuild(), session.MustBuilds("change"), api.BuildUpdate)
				builds.POST("", session.MustBuilds("change"), api.BuildCreate)
			}

			buildVersions := base.Group("/packs/:pack/builds/:build/versions")
			{
				buildVersions.Use(session.MustCurrent())
				buildVersions.Use(session.SetPack())
				buildVersions.Use(session.SetBuild())

				buildVersions.GET("", session.MustBuildVersions("display"), api.BuildVersionIndex)
				buildVersions.POST("", session.MustBuildVersions("change"), api.BuildVersionAppend)
				buildVersions.DELETE("", session.MustBuildVersions("change"), api.BuildVersionDelete)
			}

			//
			// Mods
			//
			mods := base.Group("/mods")
			{
				mods.Use(session.MustCurrent())
				mods.Use(session.MustMods("display"))

				mods.GET("", api.ModIndex)
				mods.GET("/:mod", session.SetMod(), api.ModShow)
				mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), api.ModDelete)
				mods.PUT("/:mod", session.SetMod(), session.MustMods("change"), api.ModUpdate)
				mods.POST("", session.MustMods("change"), api.ModCreate)
			}

			modUsers := base.Group("/mods/:mod/users")
			{
				modUsers.Use(session.MustCurrent())
				modUsers.Use(session.SetMod())

				modUsers.GET("", session.MustModUsers("display"), api.ModUserIndex)
				modUsers.POST("", session.MustModUsers("change"), api.ModUserAppend)
				modUsers.PUT("", session.MustModUsers("change"), api.ModUserPerm)
				modUsers.DELETE("", session.MustModUsers("change"), api.ModUserDelete)
			}

			modTeams := base.Group("/mods/:mod/teams")
			{
				modTeams.Use(session.MustCurrent())
				modTeams.Use(session.SetMod())

				modTeams.GET("", session.MustModTeams("display"), api.ModTeamIndex)
				modTeams.POST("", session.MustTeams("change"), api.ModTeamAppend)
				modTeams.PUT("", session.MustTeams("change"), api.ModTeamPerm)
				modTeams.DELETE("", session.MustModTeams("change"), api.ModTeamDelete)
			}

			//
			// Versions
			//
			versions := base.Group("/mods/:mod/versions")
			{
				versions.Use(session.MustCurrent())
				versions.Use(session.SetMod())
				versions.Use(session.MustVersions("display"))

				versions.GET("", api.VersionIndex)
				versions.GET("/:version", session.SetVersion(), api.VersionShow)
				versions.DELETE("/:version", session.SetVersion(), session.MustVersions("delete"), api.VersionDelete)
				versions.PUT("/:version", session.SetVersion(), session.MustVersions("change"), api.VersionUpdate)
				versions.POST("", session.MustVersions("change"), api.VersionCreate)
			}

			versionBuilds := base.Group("/mods/:mod/versions/:version/builds")
			{
				versionBuilds.Use(session.MustCurrent())
				versionBuilds.Use(session.SetMod())
				versionBuilds.Use(session.SetVersion())

				versionBuilds.GET("", session.MustVersionBuilds("display"), api.VersionBuildIndex)
				versionBuilds.POST("", session.MustVersionBuilds("change"), api.VersionBuildAppend)
				versionBuilds.DELETE("", session.MustVersionBuilds("change"), api.VersionBuildDelete)
			}

			//
			// Clients
			//
			clients := base.Group("/clients")
			{
				clients.Use(session.MustCurrent())
				clients.Use(session.MustClients("display"))

				clients.GET("", api.ClientIndex)
				clients.GET("/:client", session.SetClient(), api.ClientShow)
				clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), api.ClientDelete)
				clients.PUT("/:client", session.SetClient(), session.MustClients("change"), api.ClientUpdate)
				clients.POST("", session.MustClients("change"), api.ClientCreate)
			}

			clientPacks := base.Group("/clients/:client/packs")
			{
				clientPacks.Use(session.MustCurrent())
				clientPacks.Use(session.SetClient())

				clientPacks.GET("", session.MustClientPacks("display"), api.ClientPackIndex)
				clientPacks.POST("", session.MustClientPacks("change"), api.ClientPackAppend)
				clientPacks.DELETE("", session.MustClientPacks("change"), api.ClientPackDelete)
			}

			//
			// Users
			//
			users := base.Group("/users")
			{
				users.Use(session.MustCurrent())
				users.Use(session.MustUsers("display"))

				users.GET("", api.UserIndex)
				users.GET("/:user", session.SetUser(), api.UserShow)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), api.UserDelete)
				users.PUT("/:user", session.SetUser(), session.MustUsers("change"), api.UserUpdate)
				users.POST("", session.MustUsers("change"), api.UserCreate)
			}

			userTeams := base.Group("/users/:user/teams")
			{
				userTeams.Use(session.MustCurrent())
				userTeams.Use(session.SetUser())

				userTeams.GET("", session.MustUserTeams("display"), api.UserTeamIndex)
				userTeams.POST("", session.MustUserTeams("change"), api.UserTeamAppend)
				userTeams.PUT("", session.MustUserTeams("change"), api.UserTeamPerm)
				userTeams.DELETE("", session.MustUserTeams("change"), api.UserTeamDelete)
			}

			userMods := base.Group("/users/:user/mods")
			{
				userMods.Use(session.MustCurrent())
				userMods.Use(session.SetUser())

				userMods.GET("", session.MustUserMods("display"), api.UserModIndex)
				userMods.POST("", session.MustUserMods("change"), api.UserModAppend)
				userMods.PUT("", session.MustUserMods("change"), api.UserModPerm)
				userMods.DELETE("", session.MustUserMods("change"), api.UserModDelete)
			}

			userPacks := base.Group("/users/:user/packs")
			{
				userPacks.Use(session.MustCurrent())
				userPacks.Use(session.SetUser())

				userPacks.GET("", session.MustUserPacks("display"), api.UserPackIndex)
				userPacks.POST("", session.MustUserPacks("change"), api.UserPackAppend)
				userPacks.PUT("", session.MustUserPacks("change"), api.UserPackPerm)
				userPacks.DELETE("", session.MustUserPacks("change"), api.UserPackDelete)
			}

			//
			// Teams
			//
			teams := base.Group("/teams")
			{
				teams.Use(session.MustCurrent())
				teams.Use(session.MustTeams("display"))

				teams.GET("", api.TeamIndex)
				teams.GET("/:team", session.SetTeam(), api.TeamShow)
				teams.DELETE("/:team", session.SetTeam(), session.MustTeams("delete"), api.TeamDelete)
				teams.PUT("/:team", session.SetTeam(), session.MustTeams("change"), api.TeamUpdate)
				teams.POST("", session.MustTeams("change"), api.TeamCreate)
			}

			teamUsers := base.Group("/teams/:team/users")
			{
				teamUsers.Use(session.MustCurrent())
				teamUsers.Use(session.SetTeam())

				teamUsers.GET("", session.MustTeamUsers("display"), api.TeamUserIndex)
				teamUsers.POST("", session.MustTeamUsers("change"), api.TeamUserAppend)
				teamUsers.PUT("", session.MustTeamUsers("change"), api.TeamUserPerm)
				teamUsers.DELETE("", session.MustTeamUsers("change"), api.TeamUserDelete)
			}

			teamPacks := base.Group("/teams/:team/packs")
			{
				teamPacks.Use(session.MustCurrent())
				teamPacks.Use(session.SetTeam())

				teamPacks.GET("", session.MustTeamPacks("display"), api.TeamPackIndex)
				teamPacks.POST("", session.MustTeamPacks("change"), api.TeamPackAppend)
				teamPacks.PUT("", session.MustTeamPacks("change"), api.TeamPackPerm)
				teamPacks.DELETE("", session.MustTeamPacks("change"), api.TeamPackDelete)
			}

			teamMods := base.Group("/teams/:team/mods")
			{
				teamMods.Use(session.MustCurrent())
				teamMods.Use(session.SetTeam())

				teamMods.GET("", session.MustTeamMods("display"), api.TeamModIndex)
				teamMods.POST("", session.MustTeamMods("change"), api.TeamModAppend)
				teamMods.PUT("", session.MustTeamMods("change"), api.TeamModPerm)
				teamMods.DELETE("", session.MustTeamMods("change"), api.TeamModDelete)
			}

			//
			// Solder
			//
			solder := base.Group("/")
			{
				solder.Use(session.SetSolder())

				solder.GET("/modpack", api.SolderPacks)
				solder.GET("/modpack/:pack", api.SolderPack)
				solder.GET("/modpack/:pack/:build", api.SolderBuild)
			}

			//
			// Verify
			//
			base.GET("/verify/:key", api.KeyVerify)
		}
	}

	return e
}
