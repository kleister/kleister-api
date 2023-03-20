package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	profile "github.com/kleister/kleister-api/pkg/service/profile/v1"
	"github.com/rs/zerolog/log"
)

// ListPacks implements the ProfileServiceHandler interface.
func (s *ProfileServer) ListPacks(
	ctx context.Context,
	req *connect.Request[profile.ListPacksRequest],
) (*connect.Response[profile.ListPacksResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)
	return nil, errors.New("not implemented")
}
