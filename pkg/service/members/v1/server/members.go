package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/members/repository"
	members "github.com/kleister/kleister-api/pkg/service/members/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the MembersServiceHandler interface.
func (s *MembersServer) List(
	ctx context.Context,
	req *connect.Request[members.ListRequest],
) (*connect.Response[members.ListResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.List(
		ctx,
		req.Msg.Team,
		req.Msg.User,
	)

	if err != nil {
		if err == repository.ErrMemberNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Member, len(records))
	for id, record := range records {
		payload[id] = convertMemberUser(record)
	}

	return connect.NewResponse(&members.ListResponse{
		Members: payload,
	}), nil
}

// Append implements the MembersServiceHandler interface.
func (s *MembersServer) Append(
	ctx context.Context,
	req *connect.Request[members.AppendRequest],
) (*connect.Response[members.AppendResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Append(
		ctx,
		req.Msg.Member.Team,
		req.Msg.Member.User,
	); err != nil {
		if err == repository.ErrMemberNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		if err == repository.ErrIsAssigned {
			return nil, connect.NewError(
				connect.CodeFailedPrecondition,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&members.AppendResponse{
		Message: "successfully appended",
	}), nil
}

// Drop implements the MembersServiceHandler interface.
func (s *MembersServer) Drop(
	ctx context.Context,
	req *connect.Request[members.DropRequest],
) (*connect.Response[members.DropResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Drop(
		ctx,
		req.Msg.Member.Team,
		req.Msg.Member.User,
	); err != nil {
		if err == repository.ErrMemberNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		if err == repository.ErrNotAssigned {
			return nil, connect.NewError(
				connect.CodeFailedPrecondition,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&members.DropResponse{
		Message: "successfully removed user",
	}), nil
}

func convertMemberUser(record *model.Member) *types.Member {
	return &types.Member{
		TeamId:    record.Team.ID,
		TeamSlug:  record.Team.Slug,
		TeamName:  record.Team.Name,
		UserId:    record.User.ID,
		UserSlug:  record.User.Slug,
		UserName:  record.User.Username,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
