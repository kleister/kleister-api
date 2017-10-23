package web

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/storage"
	"github.com/kleister/kleister-api/pkg/template"
)

// Index represents the index page.
func Index(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "web", "index")

	return func(w http.ResponseWriter, r *http.Request) {
		err := template.Load(logger).Execute(
			w,
			map[string]string{
				"Root": config.Server.Root,
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to process index template",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unexpected())
			return
		}
	}
}
