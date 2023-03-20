package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	teams "github.com/kleister/kleister-api/pkg/service/teams/v1"
	"github.com/rs/zerolog/log"
)

// ListMods implements the TeamsServiceHandler interface.
func (s *TeamsServer) ListMods(
	ctx context.Context,
	req *connect.Request[teams.ListModsRequest],
) (*connect.Response[teams.ListModsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	// records, err := s.repository.List(ctx)

	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, err)
	// }

	// payload := make([]*types.Team, len(records))
	// for id, record := range records {
	// 	payload[id] = convertTeam(record)
	// }

	// return connect.NewResponse(&teams.ListResponse{
	// 	Teams: payload,
	// }), nil

	return nil, errors.New("not implemented")
}

// AttachMod implements the TeamsServiceHandler interface.
func (s *TeamsServer) AttachMod(
	ctx context.Context,
	req *connect.Request[teams.AttachModRequest],
) (*connect.Response[teams.AttachModResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropMod implements the TeamsServiceHandler interface.
func (s *TeamsServer) DropMod(
	ctx context.Context,
	req *connect.Request[teams.DropModRequest],
) (*connect.Response[teams.DropModResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
