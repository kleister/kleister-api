package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/members/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewMembersServer initializes the member server.
func NewMembersServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.MembersRepository,
) *MembersServer {
	return &MembersServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// MembersServer provides all handlers for members API.
type MembersServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.MembersRepository
}
