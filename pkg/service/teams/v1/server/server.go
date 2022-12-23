package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/teams/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewTeamsServer initializes the team server.
func NewTeamsServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.TeamsRepository,
) *TeamsServer {
	return &TeamsServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// TeamsServer provides all handlers for teams API.
type TeamsServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.TeamsRepository
}
