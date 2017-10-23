package web

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/kleister/kleister-api/pkg/storage"
)

// Healthz is a simple health check used by Docker and Kubernetes.
func Healthz(store storage.Store, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plain(w)
	}
}
