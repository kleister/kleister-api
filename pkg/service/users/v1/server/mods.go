package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	users "github.com/kleister/kleister-api/pkg/service/users/v1"
	"github.com/rs/zerolog/log"
)

// ListMods implements the UsersServiceHandler interface.
func (s *UsersServer) ListMods(
	ctx context.Context,
	req *connect.Request[users.ListModsRequest],
) (*connect.Response[users.ListModsResponse], error) {
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

// AttachMod implements the UsersServiceHandler interface.
func (s *UsersServer) AttachMod(
	ctx context.Context,
	req *connect.Request[users.AttachModRequest],
) (*connect.Response[users.AttachModResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropMod implements the UsersServiceHandler interface.
func (s *UsersServer) DropMod(
	ctx context.Context,
	req *connect.Request[users.DropModRequest],
) (*connect.Response[users.DropModResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
