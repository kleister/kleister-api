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

// ListTeams implements the PacksServiceHandler interface.
func (s *PacksServer) ListTeams(
	ctx context.Context,
	req *connect.Request[packs.ListTeamsRequest],
) (*connect.Response[packs.ListTeamsResponse], error) {
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

	records, err := s.repository.ListTeams(ctx, packID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Team, len(records))
	for id, record := range records {
		payload[id] = convertTeam(record.Team)
	}

	return connect.NewResponse(&packs.ListTeamsResponse{
		Teams: payload,
	}), nil
}

// AttachTeam implements the PacksServiceHandler interface.
func (s *PacksServer) AttachTeam(
	ctx context.Context,
	req *connect.Request[packs.AttachTeamRequest],
) (*connect.Response[packs.AttachTeamResponse], error) {
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

	if err := s.repository.AttachTeam(ctx, packID, req.Msg.Team); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.AttachTeamResponse{
		Message: "successfully attached team",
	}), nil
}

// DropTeam implements the PacksServiceHandler interface.
func (s *PacksServer) DropTeam(
	ctx context.Context,
	req *connect.Request[packs.DropTeamRequest],
) (*connect.Response[packs.DropTeamResponse], error) {
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

	if err := s.repository.DropTeam(ctx, packID, req.Msg.Team); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.DropTeamResponse{
		Message: "successfully dropped team",
	}), nil
}

func convertTeam(record *model.Team) *types.Team {
	return &types.Team{
		Id:        record.ID,
		Slug:      record.Slug,
		Name:      record.Name,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
