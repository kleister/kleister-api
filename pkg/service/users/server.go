package users

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/users/repository"
	serverv1 "github.com/kleister/kleister-api/pkg/service/users/v1/server"
	"github.com/kleister/kleister-api/pkg/service/users/v1/usersv1connect"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the users endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.UsersRepository,
	router *chi.Mux,
) {
	mount, handler := usersv1connect.NewUsersServiceHandler(
		serverv1.NewUsersServer(
			cfg,
			uploads,
			metricz,
			repository,
		),
	)

	router.Mount(
		path.Join(
			cfg.Server.Root,
			mount,
		)+"/",
		http.StripPrefix(
			strings.TrimRight(
				cfg.Server.Root,
				"/",
			),
			handler,
		),
	)
}
