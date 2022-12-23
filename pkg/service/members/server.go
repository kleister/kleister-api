package members

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/members/repository"
	"github.com/kleister/kleister-api/pkg/service/members/v1/membersv1connect"
	serverv1 "github.com/kleister/kleister-api/pkg/service/members/v1/server"
	"github.com/kleister/kleister-api/pkg/upload"
)

// RegisterServer is used to register the members endpoints to a router.
func RegisterServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.MembersRepository,
	router *chi.Mux,
) {
	mount, handler := membersv1connect.NewMembersServiceHandler(
		serverv1.NewMembersServer(
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
