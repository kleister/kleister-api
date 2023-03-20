package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/minecraft/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewMinecraftServer initializes the minecraft server.
func NewMinecraftServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.MinecraftRepository,
) *MinecraftServer {
	return &MinecraftServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// MinecraftServer provides all handlers for minecraft API.
type MinecraftServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.MinecraftRepository
}
