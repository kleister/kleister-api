package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	minecraft "github.com/kleister/kleister-api/pkg/service/minecraft/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListBuilds implements the MinecraftServiceHandler interface.
func (s *MinecraftServer) ListBuilds(
	ctx context.Context,
	req *connect.Request[minecraft.ListBuildsRequest],
) (*connect.Response[minecraft.ListBuildsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.ListBuilds(
		ctx,
		req.Msg.Minecraft,
		req.Msg.Query,
	)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Build, len(records))
	for id, record := range records {
		payload[id] = convertBuild(record)
	}

	return connect.NewResponse(&minecraft.ListBuildsResponse{
		Builds: payload,
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
