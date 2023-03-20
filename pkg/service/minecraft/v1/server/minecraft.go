package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	minecraft "github.com/kleister/kleister-api/pkg/service/minecraft/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Search implements the MinecraftServiceHandler interface.
func (s *MinecraftServer) Search(
	ctx context.Context,
	req *connect.Request[minecraft.SearchRequest],
) (*connect.Response[minecraft.SearchResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.Search(ctx, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Minecraft, len(records))
	for id, record := range records {
		payload[id] = convertMinecraft(record)
	}

	return connect.NewResponse(&minecraft.SearchResponse{
		Result: payload,
	}), nil
}

// Update implements the MinecraftServiceHandler interface.
func (s *MinecraftServer) Update(
	ctx context.Context,
	_ *connect.Request[minecraft.UpdateRequest],
) (*connect.Response[minecraft.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Update(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&minecraft.UpdateResponse{
		Message: "successfully updated minecraft versions",
	}), nil
}

func convertMinecraft(record *model.Minecraft) *types.Minecraft {
	return &types.Minecraft{
		Id:        record.ID,
		Name:      record.Name,
		Type:      record.Type,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
