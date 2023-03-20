package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/packs/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewPacksServer initializes the pack server.
func NewPacksServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.PacksRepository,
) *PacksServer {
	return &PacksServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// PacksServer provides all handlers for packs API.
type PacksServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.PacksRepository
}
