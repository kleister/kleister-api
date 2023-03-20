package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	teams "github.com/kleister/kleister-api/pkg/service/teams/v1"
	"github.com/rs/zerolog/log"
)

// ListUsers implements the TeamsServiceHandler interface.
func (s *TeamsServer) ListUsers(
	ctx context.Context,
	req *connect.Request[teams.ListUsersRequest],
) (*connect.Response[teams.ListUsersResponse], error) {
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

// AttachUser implements the TeamsServiceHandler interface.
func (s *TeamsServer) AttachUser(
	ctx context.Context,
	req *connect.Request[teams.AttachUserRequest],
) (*connect.Response[teams.AttachUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropUser implements the TeamsServiceHandler interface.
func (s *TeamsServer) DropUser(
	ctx context.Context,
	req *connect.Request[teams.DropUserRequest],
) (*connect.Response[teams.DropUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
