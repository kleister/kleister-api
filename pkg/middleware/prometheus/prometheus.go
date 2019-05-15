package prometheus

import (
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// ErrInvalidToken is returned when the request token is invalid.
	ErrInvalidToken = errors.New("invalid or missing token")
)

// Handler initializes the prometheus middleware.
func Handler(token string) http.HandlerFunc {
	h := promhttp.Handler()

	return func(w http.ResponseWriter, r *http.Request) {
		if token == "" {
			h.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		if header != "Bearer "+token {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}
}
