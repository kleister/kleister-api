package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
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
