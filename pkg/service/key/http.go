package key

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		keys.Use(session.MustCurrent())
// 		keys.Use(session.MustKeys("display"))

// 		keys.GET("", api.KeyIndex)
// 		keys.GET("/:key", session.SetKey(), api.KeyShow)
// 		keys.DELETE("/:key", session.SetKey(), session.MustKeys("delete"), api.KeyDelete)
// 		keys.PUT("/:key", session.SetKey(), session.MustKeys("change"), api.KeyUpdate)
// 		keys.POST("", session.MustKeys("change"), api.KeyCreate)

// 	return mux
// }
