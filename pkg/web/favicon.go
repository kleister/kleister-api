package web

import (
	"bytes"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kleister/kleister-api/pkg/assets"
	"github.com/kleister/kleister-api/pkg/storage"
)

// Favicon represents the favicon.
func Favicon(store storage.Store, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _ := assets.ReadFile(
			"images/favicon.ico",
		)

		http.ServeContent(
			w,
			r,
			"favicon.ico",
			time.Now(),
			bytes.NewReader(file),
		)
	}
}
