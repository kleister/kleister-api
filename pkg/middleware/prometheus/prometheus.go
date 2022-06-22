package prometheus

import (
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// ErrInvalidToken is returned when the request token is invalid.
	ErrInvalidToken = errors.New("invalid or missing token")
)

// Handler initializes the prometheus middleware.
func Handler(registry *prometheus.Registry, token string) http.HandlerFunc {
	h := promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			ErrorLog: metrics.Logger{},
		},
	)

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
