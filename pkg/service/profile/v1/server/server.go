package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/profile/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewProfileServer initializes the profile server.
func NewProfileServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.ProfileRepository,
) *ProfileServer {
	return &ProfileServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// ProfileServer provides all handlers for profile API.
type ProfileServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.ProfileRepository
}
