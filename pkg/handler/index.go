package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/manifest"
	"github.com/kleister/kleister-api/pkg/templates"
	"github.com/rs/zerolog/log"
)

// Index renders the template for embedded frontend.
func (h *Handler) Index() http.HandlerFunc {
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
		render.HTML(w, r, templates.String(
			h.config,
			"index.tmpl",
			struct {
				Prefix      string
				Stylesheets []string
				Javascripts []string
			}{
				Prefix:      h.Prefix(),
				Stylesheets: m.Index().Stylehseets,
				Javascripts: []string{
					m.Index().File,
				},
			},
		))
	}
}
