package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	modsRepository "github.com/kleister/kleister-api/pkg/service/mods/repository"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/service/versions/repository"
	versions "github.com/kleister/kleister-api/pkg/service/versions/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListBuilds implements the VersionsServiceHandler interface.
func (s *VersionsServer) ListBuilds(
	ctx context.Context,
	req *connect.Request[versions.ListBuildsRequest],
) (*connect.Response[versions.ListBuildsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.modsRepo.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			modsRepository.ErrModNotFound,
		)
	}

	exists, versionID, err := s.repository.Exists(ctx, modID, req.Msg.Version)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrVersionNotFound,
		)
	}

	records, err := s.repository.ListBuilds(ctx, modID, versionID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Build, len(records))
	for id, record := range records {
		payload[id] = convertBuild(record)
	}

	return connect.NewResponse(&versions.ListBuildsResponse{
		Builds: payload,
	}), nil
}

// AttachBuild implements the VersionsServiceHandler interface.
func (s *VersionsServer) AttachBuild(
	ctx context.Context,
	req *connect.Request[versions.AttachBuildRequest],
) (*connect.Response[versions.AttachBuildResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.modsRepo.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			modsRepository.ErrModNotFound,
		)
	}

	exists, versionID, err := s.repository.Exists(ctx, modID, req.Msg.Version)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrVersionNotFound,
		)
	}

	// TODO: fetch pack and build first
	if err := s.repository.AttachBuild(ctx, modID, versionID, req.Msg.Pack, req.Msg.Build); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&versions.AttachBuildResponse{
		Message: "successfully attached build",
	}), nil
}

// DropBuild implements the VersionsServiceHandler interface.
func (s *VersionsServer) DropBuild(
	ctx context.Context,
	req *connect.Request[versions.DropBuildRequest],
) (*connect.Response[versions.DropBuildResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	exists, modID, err := s.modsRepo.Exists(ctx, req.Msg.Mod)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			modsRepository.ErrModNotFound,
		)
	}

	exists, versionID, err := s.repository.Exists(ctx, modID, req.Msg.Version)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !exists {
		return nil, connect.NewError(
			connect.CodeNotFound,
			repository.ErrVersionNotFound,
		)
	}

	// TODO: fetch pack and build first
	if err := s.repository.AttachBuild(ctx, modID, versionID, req.Msg.Pack, req.Msg.Build); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&versions.DropBuildResponse{
		Message: "successfully dropped build",
	}), nil
}

func convertBuild(record *model.Build) *types.Build {
	// TODO: Add missing fields

	return &types.Build{
		Id:        record.ID,
		Slug:      record.Slug,
		Name:      record.Name,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
