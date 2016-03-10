package server

import (
	"gopkg.in/solderapp/solder-api.v0/config"
)

// Load initializes the server of the application.
func Load(cfg *config.Config) *Server {
	s := &Server{
		Addr: cfg.Server.Addr,
		Cert: cfg.Server.Cert,
		Key:  cfg.Server.Key,
		Root: cfg.Server.Root,
	}

	return s
}
