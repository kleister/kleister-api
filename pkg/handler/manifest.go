package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/manifest"
	"github.com/rs/zerolog/log"
)

// Manifest renders the manifest from embedded frontend.
func (h *Handler) Manifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := manifest.Read(h.config)

		if err != nil {
			log.Warn().
				Err(err).
				Str("handler", "manifest").
				Msg("Failed to load manifest")

			http.Error(
				w,
				"Failed to load manifest",
				http.StatusInternalServerError,
			)

			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, m)
	}
}
