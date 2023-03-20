package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	modsRepository "github.com/kleister/kleister-api/pkg/service/mods/repository"
	"github.com/kleister/kleister-api/pkg/service/versions/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewVersionsServer initializes the version server.
func NewVersionsServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.VersionsRepository,
	modsRepo modsRepository.ModsRepository,
) *VersionsServer {
	return &VersionsServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
		modsRepo:   modsRepo,
	}
}

// VersionsServer provides all handlers for versions API.
type VersionsServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.VersionsRepository
	modsRepo   modsRepository.ModsRepository
}
