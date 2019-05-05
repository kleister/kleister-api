package minecraft

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		minecraft.Use(session.MustCurrent())

// 		minecraft.GET("", api.MinecraftIndex)
// 		minecraft.GET("/:minecraft", api.MinecraftIndex)
// 		minecraft.PUT("", session.MustAdmin(), api.MinecraftUpdate)

// 	minecraftBuilds := base.Group("/minecraft/:minecraft/builds")
// 	{
// 		minecraftBuilds.Use(session.MustCurrent())
// 		minecraftBuilds.Use(session.SetMinecraft())

// 		minecraftBuilds.GET("", session.MustMinecraftBuilds("display"), api.MinecraftBuildIndex)
// 		minecraftBuilds.POST("", session.MustMinecraftBuilds("change"), api.MinecraftBuildAppend)
// 		minecraftBuilds.DELETE("", session.MustMinecraftBuilds("change"), api.MinecraftBuildDelete)
// 	}

// 	return mux
// }
