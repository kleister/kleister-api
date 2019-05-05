package auth

// import (
// 	"encoding/base32"
// 	"net/http"
// 	"time"

// 	"github.com/codehack/fail"
// 	"github.com/go-chi/chi"
// 	"github.com/go-kit/kit/log"
// 	"github.com/go-kit/kit/log/level"
// 	"github.com/json-iterator/go"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/storage"
// 	"github.com/kleister/kleister-api/pkg/token"
// )

// type Service interface {
// 	// VerifyToken
// 	VerifyToken()

// 	// RefreshToken
// 	RefreshToken()

// 	// LogoutUser
// 	LogoutUser()

// 	// LoginUser
// 	LoginUser()
// }

// type ServiceOptions struct {
// 	Store storage.Store
// }

// func NewService(opts ServiceOptions) Service {
// 	return &service{
// 		store: opts.Store,
// 	}
// }

// type service struct {
// 	store storage.Store
// }

// // VerifyToken
// func (s *service) VerifyToken() http.HandlerFunc {
// 	logger = log.WithPrefix(logger, "auth", "verify")

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var (
// 			record *model.User
// 		)

// 		_, err := token.Direct(
// 			chi.URLParam(r, "token"),
// 			func(t *token.Token) ([]byte, error) {
// 				var (
// 					err error
// 				)

// 				record, err = store.GetUser(
// 					t.Text,
// 				)

// 				signingKey, _ := base32.StdEncoding.DecodeString(record.Hash)
// 				return signingKey, err
// 			},
// 		)

// 		if err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "invalid token provided",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).Unauthorized("invalid token provided"))
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)

// 		jsoniter.NewEncoder(w).Encode(struct {
// 			Username  string    `json:"username"`
// 			CreatedAt time.Time `json:"created_at"`
// 		}{
// 			Username:  record.Username,
// 			CreatedAt: record.CreatedAt,
// 		})
// 	}
// }

// // RefreshToken
// func (s *service) RefreshToken() http.HandlerFunc {
// 	logger = log.WithPrefix(logger, "auth", "refresh")

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		record := session.Current(r.Context())

// 		token := token.New(token.SessToken, record.Username)
// 		result, err := token.SignExpiring(record.Hash, config.Session.Expire)

// 		if err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "failed to refresh token",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).Unauthorized("failed to refresh token"))
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)

// 		jsoniter.NewEncoder(w).Encode(result)
// 	}
// }

// // LogoutUser
// func (s *service) LogoutUser() http.HandlerFunc {
// 	logger = log.WithPrefix(logger, "auth", "logout")

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// // LoginUser
// func (s *service) LoginUser() http.HandlerFunc {
// 	logger = log.WithPrefix(logger, "auth", "login")

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		auth := &model.Auth{}

// 		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, loginBodyLimit)).Decode(auth); err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "failed to bind login",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).BadRequest("failed to bind login"))
// 			return
// 		}

// 		user, err := store.GetUser(
// 			auth.Username,
// 		)

// 		if err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "failed to fetch user",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
// 			return
// 		}

// 		if err := user.MatchPassword(auth.Password); err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "failed to match password",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
// 			return
// 		}

// 		token := token.New(token.SessToken, user.Username)
// 		result, err := token.SignExpiring(user.Hash, config.Session.Expire)

// 		if err != nil {
// 			level.Warn(logger).Log(
// 				"msg", "failed to generate token",
// 				"err", err,
// 			)

// 			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)

// 		jsoniter.NewEncoder(w).Encode(result)
// 	}
// }
