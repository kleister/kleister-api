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

// ListTeams implements the ModsServiceHandler interface.
func (s *ModsServer) ListTeams(
	ctx context.Context,
	req *connect.Request[mods.ListTeamsRequest],
) (*connect.Response[mods.ListTeamsResponse], error) {
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

	records, err := s.repository.ListTeams(ctx, modID, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Team, len(records))
	for id, record := range records {
		payload[id] = convertTeam(record.Team)
	}

	return connect.NewResponse(&mods.ListTeamsResponse{
		Teams: payload,
	}), nil
}

// AttachTeam implements the ModsServiceHandler interface.
func (s *ModsServer) AttachTeam(
	ctx context.Context,
	req *connect.Request[mods.AttachTeamRequest],
) (*connect.Response[mods.AttachTeamResponse], error) {
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

	if err := s.repository.AttachTeam(ctx, modID, req.Msg.Team); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.AttachTeamResponse{
		Message: "successfully attached team",
	}), nil
}

// DropTeam implements the ModsServiceHandler interface.
func (s *ModsServer) DropTeam(
	ctx context.Context,
	req *connect.Request[mods.DropTeamRequest],
) (*connect.Response[mods.DropTeamResponse], error) {
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

	if err := s.repository.DropTeam(ctx, modID, req.Msg.Team); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.DropTeamResponse{
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
