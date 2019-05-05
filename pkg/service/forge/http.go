package forge

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		forge.Use(session.MustCurrent())

// 		forge.GET("", api.ForgeIndex)
// 		forge.GET("/:forge", api.ForgeIndex)
// 		forge.PUT("", session.MustAdmin(), api.ForgeUpdate)

// 	forgeBuilds := base.Group("/forge/:forge/builds")
// 	{
// 		forgeBuilds.Use(session.MustCurrent())
// 		forgeBuilds.Use(session.SetForge())

// 		forgeBuilds.GET("", session.MustForgeBuilds("display"), api.ForgeBuildIndex)
// 		forgeBuilds.POST("", session.MustForgeBuilds("change"), api.ForgeBuildAppend)
// 		forgeBuilds.DELETE("", session.MustForgeBuilds("change"), api.ForgeBuildDelete)
// 	}

// 	return mux
// }
