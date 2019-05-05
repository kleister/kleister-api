package mod

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		mods.Use(session.MustCurrent())
// 		mods.Use(session.MustMods("display"))

// 		mods.GET("", api.ModIndex)
// 		mods.GET("/:mod", session.SetMod(), api.ModShow)
// 		mods.DELETE("/:mod", session.SetMod(), session.MustMods("delete"), api.ModDelete)
// 		mods.PUT("/:mod", session.SetMod(), session.MustMods("change"), api.ModUpdate)
// 		mods.POST("", session.MustMods("change"), api.ModCreate)

// 	modUsers := base.Group("/mods/:mod/users")
// 	{
// 		modUsers.Use(session.MustCurrent())
// 		modUsers.Use(session.SetMod())

// 		modUsers.GET("", session.MustModUsers("display"), api.ModUserIndex)
// 		modUsers.POST("", session.MustModUsers("change"), api.ModUserAppend)
// 		modUsers.PUT("", session.MustModUsers("change"), api.ModUserPerm)
// 		modUsers.DELETE("", session.MustModUsers("change"), api.ModUserDelete)
// 	}

// 	modTeams := base.Group("/mods/:mod/teams")
// 	{
// 		modTeams.Use(session.MustCurrent())
// 		modTeams.Use(session.SetMod())

// 		modTeams.GET("", session.MustModTeams("display"), api.ModTeamIndex)
// 		modTeams.POST("", session.MustTeams("change"), api.ModTeamAppend)
// 		modTeams.PUT("", session.MustTeams("change"), api.ModTeamPerm)
// 		modTeams.DELETE("", session.MustModTeams("change"), api.ModTeamDelete)
// 	}

// 	versions := base.Group("/mods/:mod/versions")
// 	{
// 		versions.Use(session.MustCurrent())
// 		versions.Use(session.SetMod())
// 		versions.Use(session.MustVersions("display"))

// 		versions.GET("", api.VersionIndex)
// 		versions.GET("/:version", session.SetVersion(), api.VersionShow)
// 		versions.DELETE("/:version", session.SetVersion(), session.MustVersions("delete"), api.VersionDelete)
// 		versions.PUT("/:version", session.SetVersion(), session.MustVersions("change"), api.VersionUpdate)
// 		versions.POST("", session.MustVersions("change"), api.VersionCreate)
// 	}

// 	versionBuilds := base.Group("/mods/:mod/versions/:version/builds")
// 	{
// 		versionBuilds.Use(session.MustCurrent())
// 		versionBuilds.Use(session.SetMod())
// 		versionBuilds.Use(session.SetVersion())

// 		versionBuilds.GET("", session.MustVersionBuilds("display"), api.VersionBuildIndex)
// 		versionBuilds.POST("", session.MustVersionBuilds("change"), api.VersionBuildAppend)
// 		versionBuilds.DELETE("", session.MustVersionBuilds("change"), api.VersionBuildDelete)
// 	}

// 	return mux
// }
