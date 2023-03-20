package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods/repository"
	mods "github.com/kleister/kleister-api/pkg/service/mods/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListUsers implements the ModsServiceHandler interface.
func (s *ModsServer) ListUsers(
	ctx context.Context,
	req *connect.Request[mods.ListUsersRequest],
) (*connect.Response[mods.ListUsersResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.repository.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrModNotFound,
		)
	}

	records, err := s.repository.ListUsers(ctx, modID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.User, len(records))
	for id, record := range records {
		payload[id] = convertUser(record.User)
	}

	return connect.NewResponse(&mods.ListUsersResponse{
		Users: payload,
	}), nil
}

// AttachUser implements the ModsServiceHandler interface.
func (s *ModsServer) AttachUser(
	ctx context.Context,
	req *connect.Request[mods.AttachUserRequest],
) (*connect.Response[mods.AttachUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.repository.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrModNotFound,
		)
	}

	if err := s.repository.AttachUser(ctx, modID, req.Msg.User); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.AttachUserResponse{
		Message: "successfully attached user",
	}), nil
}

// DropUser implements the ModsServiceHandler interface.
func (s *ModsServer) DropUser(
	ctx context.Context,
	req *connect.Request[mods.DropUserRequest],
) (*connect.Response[mods.DropUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.repository.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrModNotFound,
		)
	}

	if err := s.repository.DropUser(ctx, modID, req.Msg.User); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.DropUserResponse{
		Message: "successfully dropped user",
	}), nil
}

func convertUser(record *model.User) *types.User {
	return &types.User{
		Id:        record.ID,
		Slug:      record.Slug,
		Username:  record.Username,
		Email:     record.Email,
		Firstname: record.Firstname,
		Lastname:  record.Lastname,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
