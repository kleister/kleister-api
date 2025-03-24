package handler

import (
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
)

// Handler just provides some simple handlers to serve the frontend.
type Handler struct {
	config *config.Config
}

// Prefix generates the root prefix if a custom path have been defined.
func (h *Handler) Prefix() string {
	if strings.HasSuffix(h.config.Server.Root, "/") {
		return h.config.Server.Root
	}

	return h.config.Server.Root + "/"
}

// New initializes the handler functionaltity to be usable.
func New(cfg *config.Config) *Handler {
	return &Handler{
		config: cfg,
	}
}
