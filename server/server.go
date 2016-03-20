package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/solderapp/solder-api/config"
)

// Server represents the sub config for the server.
type Server struct {
	Addr string
	Cert string
	Key  string
	Root string
}

// Run starts serving the API based on above config.
func (s *Server) Run(handler http.Handler) {
	logrus.Infof("starting server %s", s.Addr)

	if s.Cert != "" && s.Key != "" {
		logrus.Fatal(
			http.ListenAndServeTLS(
				s.Addr,
				s.Cert,
				s.Key,
				handler,
			),
		)
	} else {
		logrus.Fatal(
			http.ListenAndServe(
				s.Addr,
				handler,
			),
		)
	}
}

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
