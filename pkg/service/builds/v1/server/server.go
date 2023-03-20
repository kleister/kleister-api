package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/builds/repository"
	packsRepository "github.com/kleister/kleister-api/pkg/service/packs/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewBuildsServer initializes the build server.
func NewBuildsServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.BuildsRepository,
	packsRepo packsRepository.PacksRepository,
) *BuildsServer {
	return &BuildsServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
		packsRepo:  packsRepo,
	}
}

// BuildsServer provides all handlers for builds API.
type BuildsServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.BuildsRepository
	packsRepo  packsRepository.PacksRepository
}
