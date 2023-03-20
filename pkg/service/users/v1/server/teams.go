package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	users "github.com/kleister/kleister-api/pkg/service/users/v1"
	"github.com/rs/zerolog/log"
)

// ListTeams implements the UsersServiceHandler interface.
func (s *UsersServer) ListTeams(
	ctx context.Context,
	req *connect.Request[users.ListTeamsRequest],
) (*connect.Response[users.ListTeamsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	// records, err := s.repository.List(ctx)

	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, err)
	// }

	// payload := make([]*types.User, len(records))
	// for id, record := range records {
	// 	payload[id] = convertUser(record)
	// }

	// return connect.NewResponse(&users.ListResponse{
	// 	Users: payload,
	// }), nil

	return nil, errors.New("not implemented")
}

// AttachTeam implements the UsersServiceHandler interface.
func (s *UsersServer) AttachTeam(
	ctx context.Context,
	req *connect.Request[users.AttachTeamRequest],
) (*connect.Response[users.AttachTeamResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropTeam implements the UsersServiceHandler interface.
func (s *UsersServer) DropTeam(
	ctx context.Context,
	req *connect.Request[users.DropTeamRequest],
) (*connect.Response[users.DropTeamResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
