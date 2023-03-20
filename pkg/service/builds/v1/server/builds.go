package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/builds/repository"
	builds "github.com/kleister/kleister-api/pkg/service/builds/v1"
	packsRepository "github.com/kleister/kleister-api/pkg/service/packs/repository"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the BuildsServiceHandler interface.
func (s *BuildsServer) List(
	ctx context.Context,
	req *connect.Request[builds.ListRequest],
) (*connect.Response[builds.ListResponse], error) {
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

	records, err := s.repository.List(ctx, packID)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Build, len(records))
	for id, record := range records {
		payload[id] = convertBuild(record)
	}

	return connect.NewResponse(&builds.ListResponse{
		Builds: payload,
	}), nil
}

// Create implements the BuildsServiceHandler interface.
func (s *BuildsServer) Create(
	ctx context.Context,
	req *connect.Request[builds.CreateRequest],
) (*connect.Response[builds.CreateResponse], error) {
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

	record := &model.Build{
		PackID: packID,
	}

	if req.Msg.Build.Slug != nil {
		record.Slug = *req.Msg.Build.Slug
	}

	if req.Msg.Build.Name != "" {
		record.Name = req.Msg.Build.Name
	}

	// TODO: Add missing fields

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

	return connect.NewResponse(&builds.CreateResponse{
		Build: convertBuild(created),
	}), nil
}

// Update implements the BuildsServiceHandler interface.
func (s *BuildsServer) Update(
	ctx context.Context,
	req *connect.Request[builds.UpdateRequest],
) (*connect.Response[builds.UpdateResponse], error) {
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

	record, err := s.repository.Show(ctx, packID, req.Msg.Id)

	if err != nil {
		if err == repository.ErrBuildNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.Build.Slug != nil {
		record.Slug = *req.Msg.Build.Slug
	}

	if req.Msg.Build.Name != nil {
		record.Name = *req.Msg.Build.Name
	}

	// TODO: Add missing fields

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

	return connect.NewResponse(&builds.UpdateResponse{
		Build: convertBuild(updated),
	}), nil
}

// Show implements the BuildsServiceHandler interface.
func (s *BuildsServer) Show(
	ctx context.Context,
	req *connect.Request[builds.ShowRequest],
) (*connect.Response[builds.ShowResponse], error) {
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

	record, err := s.repository.Show(ctx, packID, req.Msg.Id)

	if err != nil {
		if err == repository.ErrBuildNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&builds.ShowResponse{
		Build: convertBuild(record),
	}), nil
}

// Delete implements the BuildsServiceHandler interface.
func (s *BuildsServer) Delete(
	ctx context.Context,
	req *connect.Request[builds.DeleteRequest],
) (*connect.Response[builds.DeleteResponse], error) {
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

	if err := s.repository.Delete(ctx, packID, req.Msg.Id); err != nil {
		if err == repository.ErrBuildNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&builds.DeleteResponse{
		Message: "successfully deleted build",
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
