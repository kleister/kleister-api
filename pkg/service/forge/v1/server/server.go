package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/forge/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewForgeServer initializes the forge server.
func NewForgeServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.ForgeRepository,
) *ForgeServer {
	return &ForgeServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// ForgeServer provides all handlers for forge API.
type ForgeServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.ForgeRepository
}
