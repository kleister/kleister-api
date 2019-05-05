package user

import (
	"github.com/go-chi/chi"
)

func NewHandler(service Service) *chi.Mux {
	mux := chi.NewRouter()

	// 	users.Use(session.MustCurrent())
	// 	users.Use(session.MustUsers("display"))

	// 	users.GET("", api.UserIndex)
	// 	users.GET("/:user", session.SetUser(), api.UserShow)
	// 	users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), api.UserDelete)
	// 	users.PUT("/:user", session.SetUser(), session.MustUsers("change"), api.UserUpdate)
	// 	users.POST("", session.MustUsers("change"), api.UserCreate)

	// userTeams := base.Group("/users/:user/teams")
	// {
	// 	userTeams.Use(session.MustCurrent())
	// 	userTeams.Use(session.SetUser())

	// 	userTeams.GET("", session.MustUserTeams("display"), api.UserTeamIndex)
	// 	userTeams.POST("", session.MustUserTeams("change"), api.UserTeamAppend)
	// 	userTeams.PUT("", session.MustUserTeams("change"), api.UserTeamPerm)
	// 	userTeams.DELETE("", session.MustUserTeams("change"), api.UserTeamDelete)
	// }

	// userMods := base.Group("/users/:user/mods")
	// {
	// 	userMods.Use(session.MustCurrent())
	// 	userMods.Use(session.SetUser())

	// 	userMods.GET("", session.MustUserMods("display"), api.UserModIndex)
	// 	userMods.POST("", session.MustUserMods("change"), api.UserModAppend)
	// 	userMods.PUT("", session.MustUserMods("change"), api.UserModPerm)
	// 	userMods.DELETE("", session.MustUserMods("change"), api.UserModDelete)
	// }

	// userPacks := base.Group("/users/:user/packs")
	// {
	// 	userPacks.Use(session.MustCurrent())
	// 	userPacks.Use(session.SetUser())

	// 	userPacks.GET("", session.MustUserPacks("display"), api.UserPackIndex)
	// 	userPacks.POST("", session.MustUserPacks("change"), api.UserPackAppend)
	// 	userPacks.PUT("", session.MustUserPacks("change"), api.UserPackPerm)
	// 	userPacks.DELETE("", session.MustUserPacks("change"), api.UserPackDelete)
	// }

	return mux
}
