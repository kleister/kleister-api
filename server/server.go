package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

type Server struct {
	Addr string
	Cert string
	Key  string
	Root string
}

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
