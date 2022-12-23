package serverv1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/service/users/repository"
	"github.com/kleister/kleister-api/pkg/upload"
)

// NewUsersServer initializes the user server.
func NewUsersServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.UsersRepository,
) *UsersServer {
	return &UsersServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// UsersServer provides all handlers for users API.
type UsersServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.UsersRepository
}
