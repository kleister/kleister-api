package handler

import (
	"io/fs"
	"net/http"
	"path"

	"github.com/kleister/kleister-api/pkg/frontend"
	"github.com/rs/zerolog/log"
)

// Assets provides an handler to read all assets from embedded frontend.
func (h *Handler) Assets() http.Handler {
	content, err := fs.Sub(
		frontend.Load(h.config),
		"assets",
	)

	if err != nil {
		log.Warn().
			Err(err).
			Str("handler", "assets").
			Msg("Failed to load assets")
	}

	return http.StripPrefix(
		path.Join(
			h.config.Server.Root,
			"assets",
		)+"/",
		http.FileServer(
			http.FS(
				content,
			),
		),
	)
}
