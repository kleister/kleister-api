package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs/repository"
	packs "github.com/kleister/kleister-api/pkg/service/packs/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListUsers implements the PacksServiceHandler interface.
func (s *PacksServer) ListUsers(
	ctx context.Context,
	req *connect.Request[packs.ListUsersRequest],
) (*connect.Response[packs.ListUsersResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.repository.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrPackNotFound,
		)
	}

	records, err := s.repository.ListUsers(ctx, packID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.User, len(records))
	for id, record := range records {
		payload[id] = convertUser(record.User)
	}

	return connect.NewResponse(&packs.ListUsersResponse{
		Users: payload,
	}), nil
}

// AttachUser implements the PacksServiceHandler interface.
func (s *PacksServer) AttachUser(
	ctx context.Context,
	req *connect.Request[packs.AttachUserRequest],
) (*connect.Response[packs.AttachUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.repository.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrPackNotFound,
		)
	}

	if err := s.repository.AttachUser(ctx, packID, req.Msg.User); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.AttachUserResponse{
		Message: "successfully attached user",
	}), nil
}

// DropUser implements the PacksServiceHandler interface.
func (s *PacksServer) DropUser(
	ctx context.Context,
	req *connect.Request[packs.DropUserRequest],
) (*connect.Response[packs.DropUserResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.repository.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrPackNotFound,
		)
	}

	if err := s.repository.DropUser(ctx, packID, req.Msg.User); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.DropUserResponse{
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
