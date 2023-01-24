package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/service/users/repository"
	users "github.com/kleister/kleister-api/pkg/service/users/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the UsersServiceHandler interface.
func (s *UsersServer) List(
	ctx context.Context,
	_ *connect.Request[users.ListRequest],
) (*connect.Response[users.ListResponse], error) {
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

	payload := make([]*types.User, len(records))
	for id, record := range records {
		payload[id] = convertUser(record)
	}

	return connect.NewResponse(&users.ListResponse{
		Users: payload,
	}), nil
}

// Create implements the UsersServiceHandler interface.
func (s *UsersServer) Create(
	ctx context.Context,
	req *connect.Request[users.CreateRequest],
) (*connect.Response[users.CreateResponse], error) {
	// if !current.Admin {
	// 	return nil, connect.NewError(
	// 		connect.CodePermissionDenied,
	// 		fmt.Errorf("only admins can access this resource"),
	// 	)
	// }

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record := &model.User{}

	if req.Msg.User.Slug != nil {
		record.Slug = *req.Msg.User.Slug
	}

	if req.Msg.User.Username != "" {
		record.Username = req.Msg.User.Username
	}

	if req.Msg.User.Password != "" {
		record.Password = req.Msg.User.Password
	}

	if req.Msg.User.Email != "" {
		record.Email = req.Msg.User.Email
	}

	if req.Msg.User.Firstname != nil {
		record.Firstname = *req.Msg.User.Firstname
	}

	if req.Msg.User.Lastname != nil {
		record.Lastname = *req.Msg.User.Lastname
	}

	record.Admin = req.Msg.User.Admin
	record.Active = req.Msg.User.Active

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

	return connect.NewResponse(&users.CreateResponse{
		User: convertUser(created),
	}), nil
}

// Update implements the UsersServiceHandler interface.
func (s *UsersServer) Update(
	ctx context.Context,
	req *connect.Request[users.UpdateRequest],
) (*connect.Response[users.UpdateResponse], error) {
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
		if err == repository.ErrUserNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.User.Slug != nil {
		record.Slug = *req.Msg.User.Slug
	}

	if req.Msg.User.Username != nil {
		record.Username = *req.Msg.User.Username
	}

	if req.Msg.User.Password != nil {
		record.Password = *req.Msg.User.Password
	}

	if req.Msg.User.Email != nil {
		record.Email = *req.Msg.User.Email
	}

	if req.Msg.User.Firstname != nil {
		record.Firstname = *req.Msg.User.Firstname
	}

	if req.Msg.User.Lastname != nil {
		record.Lastname = *req.Msg.User.Lastname
	}

	if req.Msg.User.Admin != nil {
		record.Admin = *req.Msg.User.Admin
	}

	if req.Msg.User.Active != nil {
		record.Active = *req.Msg.User.Active
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

	return connect.NewResponse(&users.UpdateResponse{
		User: convertUser(updated),
	}), nil
}

// Show implements the UsersServiceHandler interface.
func (s *UsersServer) Show(
	ctx context.Context,
	req *connect.Request[users.ShowRequest],
) (*connect.Response[users.ShowResponse], error) {
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
		if err == repository.ErrUserNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&users.ShowResponse{
		User: convertUser(record),
	}), nil
}

// Delete implements the UsersServiceHandler interface.
func (s *UsersServer) Delete(
	ctx context.Context,
	req *connect.Request[users.DeleteRequest],
) (*connect.Response[users.DeleteResponse], error) {
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
		if err == repository.ErrUserNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&users.DeleteResponse{
		Message: "successfully deleted user",
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
		Admin:     record.Admin,
		Active:    record.Active,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
