package web

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/kleister/kleister-api/pkg/storage"
)

// Readyz is a simple ready check used by Docker and Kubernetes.
func Readyz(store storage.Store, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plain(w)
	}
}
