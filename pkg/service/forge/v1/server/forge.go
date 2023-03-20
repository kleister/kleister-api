package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	forge "github.com/kleister/kleister-api/pkg/service/forge/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Search implements the ForgeServiceHandler interface.
func (s *ForgeServer) Search(
	ctx context.Context,
	req *connect.Request[forge.SearchRequest],
) (*connect.Response[forge.SearchResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.Search(ctx, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Forge, len(records))
	for id, record := range records {
		payload[id] = convertForge(record)
	}

	return connect.NewResponse(&forge.SearchResponse{
		Result: payload,
	}), nil
}

// Update implements the ForgeServiceHandler interface.
func (s *ForgeServer) Update(
	ctx context.Context,
	_ *connect.Request[forge.UpdateRequest],
) (*connect.Response[forge.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Update(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&forge.UpdateResponse{
		Message: "successfully updated forge versions",
	}), nil
}

func convertForge(record *model.Forge) *types.Forge {
	return &types.Forge{
		Id:        record.ID,
		Name:      record.Name,
		Minecraft: record.Minecraft,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
