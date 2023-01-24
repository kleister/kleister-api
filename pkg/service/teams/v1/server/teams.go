package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/teams/repository"
	teams "github.com/kleister/kleister-api/pkg/service/teams/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the TeamsServiceHandler interface.
func (s *TeamsServer) List(
	ctx context.Context,
	_ *connect.Request[teams.ListRequest],
) (*connect.Response[teams.ListResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.List(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Team, len(records))
	for id, record := range records {
		payload[id] = convertTeam(record)
	}

	return connect.NewResponse(&teams.ListResponse{
		Teams: payload,
	}), nil
}

// Create implements the TeamsServiceHandler interface.
func (s *TeamsServer) Create(
	ctx context.Context,
	req *connect.Request[teams.CreateRequest],
) (*connect.Response[teams.CreateResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record := &model.Team{}

	if req.Msg.Team.Slug != nil {
		record.Slug = *req.Msg.Team.Slug
	}

	if req.Msg.Team.Name != "" {
		record.Name = req.Msg.Team.Name
	}

	created, err := s.repository.Create(ctx, record)

	if err != nil {
		if v, ok := err.(validate.Errors); ok {

			log.Debug().Err(err).Msgf("%+v", v.Errors)
			// for _, verr := range v.Errors {
			// 	payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
			// 		Field:   verr.Field,
			// 		Message: verr.Error.Error(),
			// 	})
			// }

			return nil, connect.NewError(connect.CodeFailedPrecondition, err)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&teams.CreateResponse{
		Team: convertTeam(created),
	}), nil
}

// Update implements the TeamsServiceHandler interface.
func (s *TeamsServer) Update(
	ctx context.Context,
	req *connect.Request[teams.UpdateRequest],
) (*connect.Response[teams.UpdateResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrTeamNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.Team.Slug != nil {
		record.Slug = *req.Msg.Team.Slug
	}

	if req.Msg.Team.Name != nil {
		record.Name = *req.Msg.Team.Name
	}

	updated, err := s.repository.Update(ctx, record)

	if err != nil {
		if v, ok := err.(validate.Errors); ok {

			log.Debug().Err(err).Msgf("%+v", v.Errors)
			// for _, verr := range v.Errors {
			// 	payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
			// 		Field:   verr.Field,
			// 		Message: verr.Error.Error(),
			// 	})
			// }

			return nil, connect.NewError(connect.CodeFailedPrecondition, err)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&teams.UpdateResponse{
		Team: convertTeam(updated),
	}), nil
}

// Show implements the TeamsServiceHandler interface.
func (s *TeamsServer) Show(
	ctx context.Context,
	req *connect.Request[teams.ShowRequest],
) (*connect.Response[teams.ShowResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrTeamNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&teams.ShowResponse{
		Team: convertTeam(record),
	}), nil
}

// Delete implements the TeamsServiceHandler interface.
func (s *TeamsServer) Delete(
	ctx context.Context,
	req *connect.Request[teams.DeleteRequest],
) (*connect.Response[teams.DeleteResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Delete(ctx, req.Msg.Id); err != nil {
		if err == repository.ErrTeamNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&teams.DeleteResponse{
		Message: "successfully deleted team",
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
