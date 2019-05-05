package team

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		teams.Use(session.MustCurrent())
// 		teams.Use(session.MustTeams("display"))

// 		teams.GET("", api.TeamIndex)
// 		teams.GET("/:team", session.SetTeam(), api.TeamShow)
// 		teams.DELETE("/:team", session.SetTeam(), session.MustTeams("delete"), api.TeamDelete)
// 		teams.PUT("/:team", session.SetTeam(), session.MustTeams("change"), api.TeamUpdate)
// 		teams.POST("", session.MustTeams("change"), api.TeamCreate)

// 	teamUsers := base.Group("/teams/:team/users")
// 	{
// 		teamUsers.Use(session.MustCurrent())
// 		teamUsers.Use(session.SetTeam())

// 		teamUsers.GET("", session.MustTeamUsers("display"), api.TeamUserIndex)
// 		teamUsers.POST("", session.MustTeamUsers("change"), api.TeamUserAppend)
// 		teamUsers.PUT("", session.MustTeamUsers("change"), api.TeamUserPerm)
// 		teamUsers.DELETE("", session.MustTeamUsers("change"), api.TeamUserDelete)
// 	}

// 	teamPacks := base.Group("/teams/:team/packs")
// 	{
// 		teamPacks.Use(session.MustCurrent())
// 		teamPacks.Use(session.SetTeam())

// 		teamPacks.GET("", session.MustTeamPacks("display"), api.TeamPackIndex)
// 		teamPacks.POST("", session.MustTeamPacks("change"), api.TeamPackAppend)
// 		teamPacks.PUT("", session.MustTeamPacks("change"), api.TeamPackPerm)
// 		teamPacks.DELETE("", session.MustTeamPacks("change"), api.TeamPackDelete)
// 	}

// 	teamMods := base.Group("/teams/:team/mods")
// 	{
// 		teamMods.Use(session.MustCurrent())
// 		teamMods.Use(session.SetTeam())

// 		teamMods.GET("", session.MustTeamMods("display"), api.TeamModIndex)
// 		teamMods.POST("", session.MustTeamMods("change"), api.TeamModAppend)
// 		teamMods.PUT("", session.MustTeamMods("change"), api.TeamModPerm)
// 		teamMods.DELETE("", session.MustTeamMods("change"), api.TeamModDelete)
// 	}

// 	return mux
// }
