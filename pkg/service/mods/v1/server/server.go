package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/mods/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewModsServer initializes the mod server.
func NewModsServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.ModsRepository,
) *ModsServer {
	return &ModsServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// ModsServer provides all handlers for mods API.
type ModsServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.ModsRepository
}
