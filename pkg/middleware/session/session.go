package session

import (
	"context"
	"encoding/base32"
	"net/http"

	"github.com/codehack/fail"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/storage"
	"github.com/kleister/kleister-api/pkg/token"
)

var (
	// CurrentContextKey defines the key for the user context store.
	CurrentContextKey = &contextKey{"current"}

	// TokenContextKey defines the key for the token context store.
	TokenContextKey = &contextKey{"token"}
)

// Current gets the user from the context.
func Current(c context.Context) *model.User {
	v, ok := c.Value(CurrentContextKey).(*model.User)

	if !ok {
		return nil
	}

	return v
}

// SetCurrent injects the user into the context.
func SetCurrent(store storage.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				record *model.User
			)

			parsed, err := token.Parse(
				r,
				func(t *token.Token) ([]byte, error) {
					var (
						err error
					)

					record, err = store.GetUser(
						t.Text,
					)

					signingKey, _ := base32.StdEncoding.DecodeString(record.Hash)
					return signingKey, err
				},
			)

			ctx := r.Context()

			if err == nil {
				ctx = context.WithValue(ctx, TokenContextKey, parsed)
				ctx = context.WithValue(ctx, CurrentContextKey, record)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustCurrent validates the user access.
func MustCurrent() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := Current(r.Context())

			if user == nil {
				fail.Error(w, fail.Unauthorized("you have to be a authenticated"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// MustNobody validates anonymous users.
func MustNobody() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := Current(r.Context())

			if user != nil {
				fail.Error(w, fail.Unauthorized("you have to be a guest user"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// MustAdmin validates the admin access.
func MustAdmin() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := Current(r.Context())

			if user == nil || !user.Admin || !isAdmin(user.Username) {
				fail.Error(w, fail.Unauthorized("you have to be an admin user"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// isAdmin just checks if the current user is a global admin.
func isAdmin(username string) bool {
	for _, admin := range config.Admin.Users {
		if admin == username {
			return true
		}
	}

	return false
}
