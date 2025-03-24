package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// Config renders the config for the embedded frontend.
func (h *Handler) Config() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Root string `json:"root"`
		}{
			Root: h.config.Server.Root,
		})
	}
}
