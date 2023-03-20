package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	teams "github.com/kleister/kleister-api/pkg/service/teams/v1"
	"github.com/rs/zerolog/log"
)

// ListPacks implements the TeamsServiceHandler interface.
func (s *TeamsServer) ListPacks(
	ctx context.Context,
	req *connect.Request[teams.ListPacksRequest],
) (*connect.Response[teams.ListPacksResponse], error) {
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

// AttachPack implements the TeamsServiceHandler interface.
func (s *TeamsServer) AttachPack(
	ctx context.Context,
	req *connect.Request[teams.AttachPackRequest],
) (*connect.Response[teams.AttachPackResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropPack implements the TeamsServiceHandler interface.
func (s *TeamsServer) DropPack(
	ctx context.Context,
	req *connect.Request[teams.DropPackRequest],
) (*connect.Response[teams.DropPackResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
