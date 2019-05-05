package profile

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 		profile.Use(session.MustCurrent())

// 		profile.GET("/token", api.ProfileToken)
// 		profile.GET("/self", api.ProfileShow)
// 		profile.PUT("/self", api.ProfileUpdate)

// 	return mux
// }
