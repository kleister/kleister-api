package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	profile "github.com/kleister/kleister-api/pkg/service/profile/v1"
	"github.com/rs/zerolog/log"
)

// Show implements the ProfileServiceHandler interface.
func (s *ProfileServer) Show(
	ctx context.Context,
	req *connect.Request[profile.ShowRequest],
) (*connect.Response[profile.ShowResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)
	return nil, errors.New("not implemented")
}

// Update implements the ProfileServiceHandler interface.
func (s *ProfileServer) Update(
	ctx context.Context,
	req *connect.Request[profile.UpdateRequest],
) (*connect.Response[profile.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)
	return nil, errors.New("not implemented")
}
