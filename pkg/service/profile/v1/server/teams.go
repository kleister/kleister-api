package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	profile "github.com/kleister/kleister-api/pkg/service/profile/v1"
	"github.com/rs/zerolog/log"
)

// ListTeams implements the ProfileServiceHandler interface.
func (s *ProfileServer) ListTeams(
	ctx context.Context,
	req *connect.Request[profile.ListTeamsRequest],
) (*connect.Response[profile.ListTeamsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)
	return nil, errors.New("not implemented")
}
