package server

import (
	"github.com/solderapp/solder/config"
)

func Load(cfg *config.Config) *Server {
	s := &Server{
		Addr: cfg.Server.Addr,
		Cert: cfg.Server.Cert,
		Key:  cfg.Server.Key,
		Root: cfg.Server.Root,
	}

	return s
}
