package auth

// import (
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/storage"
// )

// func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
// 	mux := chi.NewRouter()

// 	mux.Get("/verify/{token}", Verify(store, logger))
// 	mux.With(session.MustCurrent()).Get("/logout", Logout(store, logger))
// 	mux.With(session.MustCurrent()).Get("/refresh", Refresh(store, logger))
// 	mux.With(session.MustNobody()).Post("/login", Login(store, logger))

// 	return mux
// }
