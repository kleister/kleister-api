package serverv1

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	users "github.com/kleister/kleister-api/pkg/service/users/v1"
	"github.com/rs/zerolog/log"
)

// ListPacks implements the UsersServiceHandler interface.
func (s *UsersServer) ListPacks(
	ctx context.Context,
	req *connect.Request[users.ListPacksRequest],
) (*connect.Response[users.ListPacksResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

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

// AttachPack implements the UsersServiceHandler interface.
func (s *UsersServer) AttachPack(
	ctx context.Context,
	req *connect.Request[users.AttachPackRequest],
) (*connect.Response[users.AttachPackResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}

// DropPack implements the UsersServiceHandler interface.
func (s *UsersServer) DropPack(
	ctx context.Context,
	req *connect.Request[users.DropPackRequest],
) (*connect.Response[users.DropPackResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	log.Debug().Msgf("%+v", req)

	return nil, errors.New("not implemented")
}
