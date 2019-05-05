package general

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/json-iterator/go"
	"github.com/kleister/kleister-api/pkg/storage"
	"github.com/kleister/kleister-api/pkg/version"
)

// Index represents the API index.
func Index(store storage.Store, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsoniter.NewEncoder(w).Encode(struct {
			API     string `json:"api"`
			Version string `json:"version"`
			Stream  string `json:"stream"`
		}{
			API:     "Kleister API",
			Version: version.Version.String(),
			Stream:  "master",
		})
	}
}
