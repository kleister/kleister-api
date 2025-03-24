package handler

import (
	"io"
	"net/http"

	"github.com/kleister/kleister-api/pkg/frontend"
	"github.com/rs/zerolog/log"
)

// Favicon returns the favicon for embedded frontend.
func (h *Handler) Favicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := frontend.Load(h.config).Open("favicon.svg")

		if err != nil {
			log.Warn().
				Err(err).
				Str("handler", "favicon").
				Msg("Failed to load favicon")

			http.Error(
				w,
				"Failed to load favicon",
				http.StatusInternalServerError,
			)

			return
		}

		defer file.Close()
		stat, err := file.Stat()

		if err != nil {
			log.Warn().
				Err(err).
				Str("handler", "favicon").
				Msg("Failed to stat favicon")

			http.Error(
				w,
				"Failed to stat favicon",
				http.StatusInternalServerError,
			)

			return
		}

		http.ServeContent(
			w,
			r,
			"favicon.svg",
			stat.ModTime(),
			file.(io.ReadSeeker),
		)
	}
}
