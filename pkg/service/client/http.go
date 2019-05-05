package client

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		clients.Use(session.MustCurrent())
// 		clients.Use(session.MustClients("display"))

// 		clients.GET("", api.ClientIndex)
// 		clients.GET("/:client", session.SetClient(), api.ClientShow)
// 		clients.DELETE("/:client", session.SetClient(), session.MustClients("delete"), api.ClientDelete)
// 		clients.PUT("/:client", session.SetClient(), session.MustClients("change"), api.ClientUpdate)
// 		clients.POST("", session.MustClients("change"), api.ClientCreate)

// 	clientPacks := base.Group("/clients/:client/packs")
// 	{
// 		clientPacks.Use(session.MustCurrent())
// 		clientPacks.Use(session.SetClient())

// 		clientPacks.GET("", session.MustClientPacks("display"), api.ClientPackIndex)
// 		clientPacks.POST("", session.MustClientPacks("change"), api.ClientPackAppend)
// 		clientPacks.DELETE("", session.MustClientPacks("change"), api.ClientPackDelete)
// 	}

// 	return mux
// }
