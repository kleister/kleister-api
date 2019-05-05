package pack

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		packs.Use(session.MustCurrent())
// 		packs.Use(session.MustPacks("display"))

// 		packs.GET("", api.PackIndex)
// 		packs.GET("/:pack", session.SetPack(), api.PackShow)
// 		packs.DELETE("/:pack", session.SetPack(), session.MustPacks("delete"), api.PackDelete)
// 		packs.PUT("/:pack", session.SetPack(), session.MustPacks("change"), api.PackUpdate)
// 		packs.POST("", session.MustPacks("change"), api.PackCreate)

// 	packClients := base.Group("/packs/:pack/clients")
// 	{
// 		packClients.Use(session.MustCurrent())
// 		packClients.Use(session.SetPack())

// 		packClients.GET("", session.MustPackClients("display"), api.PackClientIndex)
// 		packClients.POST("", session.MustPackClients("change"), api.PackClientAppend)
// 		packClients.DELETE("", session.MustPackClients("change"), api.PackClientDelete)
// 	}

// 	packUsers := base.Group("/packs/:pack/users")
// 	{
// 		packUsers.Use(session.MustCurrent())
// 		packUsers.Use(session.SetPack())

// 		packUsers.GET("", session.MustPackUsers("display"), api.PackUserIndex)
// 		packUsers.POST("", session.MustPackUsers("change"), api.PackUserAppend)
// 		packUsers.PUT("", session.MustPackUsers("change"), api.PackUserPerm)
// 		packUsers.DELETE("", session.MustPackUsers("change"), api.PackUserDelete)
// 	}

// 	packTeams := base.Group("/packs/:pack/teams")
// 	{
// 		packTeams.Use(session.MustCurrent())
// 		packTeams.Use(session.SetPack())

// 		packTeams.GET("", session.MustPackTeams("display"), api.PackTeamIndex)
// 		packTeams.POST("", session.MustPackTeams("change"), api.PackTeamAppend)
// 		packTeams.PUT("", session.MustPackTeams("change"), api.PackTeamPerm)
// 		packTeams.DELETE("", session.MustPackTeams("change"), api.PackTeamDelete)
// 	}

// 	builds := base.Group("/packs/:pack/builds")
// 	{
// 		builds.Use(session.MustCurrent())
// 		builds.Use(session.SetPack())
// 		builds.Use(session.MustBuilds("display"))

// 		builds.GET("", api.BuildIndex)
// 		builds.GET("/:build", session.SetBuild(), api.BuildShow)
// 		builds.DELETE("/:build", session.SetBuild(), session.MustBuilds("delete"), api.BuildDelete)
// 		builds.PUT("/:build", session.SetBuild(), session.MustBuilds("change"), api.BuildUpdate)
// 		builds.POST("", session.MustBuilds("change"), api.BuildCreate)
// 	}

// 	buildVersions := base.Group("/packs/:pack/builds/:build/versions")
// 	{
// 		buildVersions.Use(session.MustCurrent())
// 		buildVersions.Use(session.SetPack())
// 		buildVersions.Use(session.SetBuild())

// 		buildVersions.GET("", session.MustBuildVersions("display"), api.BuildVersionIndex)
// 		buildVersions.POST("", session.MustBuildVersions("change"), api.BuildVersionAppend)
// 		buildVersions.DELETE("", session.MustBuildVersions("change"), api.BuildVersionDelete)
// 	}

// 	return mux
// }
