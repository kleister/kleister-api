package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/builds/repository"
	builds "github.com/kleister/kleister-api/pkg/service/builds/v1"
	packsRepository "github.com/kleister/kleister-api/pkg/service/packs/repository"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListVersions implements the BuildsServiceHandler interface.
func (s *BuildsServer) ListVersions(
	ctx context.Context,
	req *connect.Request[builds.ListVersionsRequest],
) (*connect.Response[builds.ListVersionsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.packsRepo.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			packsRepository.ErrPackNotFound,
		)
	}

	exists, buildID, err := s.repository.Exists(ctx, packID, req.Msg.Build)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrBuildNotFound,
		)
	}

	records, err := s.repository.ListVersions(ctx, packID, buildID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Version, len(records))
	for id, record := range records {
		payload[id] = convertVersion(record)
	}

	return connect.NewResponse(&builds.ListVersionsResponse{
		Versions: payload,
	}), nil
}

// AttachVersion implements the BuildsServiceHandler interface.
func (s *BuildsServer) AttachVersion(
	ctx context.Context,
	req *connect.Request[builds.AttachVersionRequest],
) (*connect.Response[builds.AttachVersionResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.packsRepo.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			packsRepository.ErrPackNotFound,
		)
	}

	exists, buildID, err := s.repository.Exists(ctx, packID, req.Msg.Build)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrBuildNotFound,
		)
	}

	// TODO: fetch mod and version first
	if err := s.repository.AttachVersion(ctx, packID, buildID, req.Msg.Mod, req.Msg.Version); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&builds.AttachVersionResponse{
		Message: "successfully attached version",
	}), nil
}

// DropVersion implements the BuildsServiceHandler interface.
func (s *BuildsServer) DropVersion(
	ctx context.Context,
	req *connect.Request[builds.DropVersionRequest],
) (*connect.Response[builds.DropVersionResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, packID, err := s.packsRepo.Exists(ctx, req.Msg.Pack)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			packsRepository.ErrPackNotFound,
		)
	}

	exists, buildID, err := s.repository.Exists(ctx, packID, req.Msg.Build)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrBuildNotFound,
		)
	}

	// TODO: fetch mod and version first
	if err := s.repository.DropVersion(ctx, packID, buildID, req.Msg.Mod, req.Msg.Version); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&builds.DropVersionResponse{
		Message: "successfully dropped version",
	}), nil
}

func convertVersion(record *model.Version) *types.Version {
	// TODO: Add missing fields

	return &types.Version{
		Id:        record.ID,
		Slug:      record.Slug,
		Name:      record.Name,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
