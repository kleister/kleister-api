package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	modsRepository "github.com/kleister/kleister-api/pkg/service/mods/repository"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/service/versions/repository"
	versions "github.com/kleister/kleister-api/pkg/service/versions/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the VersionsServiceHandler interface.
func (s *VersionsServer) List(
	ctx context.Context,
	req *connect.Request[versions.ListRequest],
) (*connect.Response[versions.ListResponse], error) {
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

	records, err := s.repository.List(ctx, modID)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Version, len(records))
	for id, record := range records {
		payload[id] = convertVersion(record)
	}

	return connect.NewResponse(&versions.ListResponse{
		Versions: payload,
	}), nil
}

// Create implements the VersionsServiceHandler interface.
func (s *VersionsServer) Create(
	ctx context.Context,
	req *connect.Request[versions.CreateRequest],
) (*connect.Response[versions.CreateResponse], error) {
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

	record := &model.Version{
		ModID: modID,
	}

	if req.Msg.Version.Slug != nil {
		record.Slug = *req.Msg.Version.Slug
	}

	if req.Msg.Version.Name != "" {
		record.Name = req.Msg.Version.Name
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

	return connect.NewResponse(&versions.CreateResponse{
		Version: convertVersion(created),
	}), nil
}

// Update implements the VersionsServiceHandler interface.
func (s *VersionsServer) Update(
	ctx context.Context,
	req *connect.Request[versions.UpdateRequest],
) (*connect.Response[versions.UpdateResponse], error) {
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

	record, err := s.repository.Show(ctx, modID, req.Msg.Id)

	if err != nil {
		if err == repository.ErrVersionNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.Version.Slug != nil {
		record.Slug = *req.Msg.Version.Slug
	}

	if req.Msg.Version.Name != nil {
		record.Name = *req.Msg.Version.Name
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

	return connect.NewResponse(&versions.UpdateResponse{
		Version: convertVersion(updated),
	}), nil
}

// Show implements the VersionsServiceHandler interface.
func (s *VersionsServer) Show(
	ctx context.Context,
	req *connect.Request[versions.ShowRequest],
) (*connect.Response[versions.ShowResponse], error) {
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

	record, err := s.repository.Show(ctx, modID, req.Msg.Id)

	if err != nil {
		if err == repository.ErrVersionNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&versions.ShowResponse{
		Version: convertVersion(record),
	}), nil
}

// Delete implements the VersionsServiceHandler interface.
func (s *VersionsServer) Delete(
	ctx context.Context,
	req *connect.Request[versions.DeleteRequest],
) (*connect.Response[versions.DeleteResponse], error) {
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

	if err := s.repository.Delete(ctx, modID, req.Msg.Id); err != nil {
		if err == repository.ErrVersionNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&versions.DeleteResponse{
		Message: "successfully deleted version",
	}), nil
}

func convertVersion(record *model.Version) *types.Version {
	// TODO: Add missing fields

	return &types.Version{
		Id:        record.ID,
		Slug:      record.Slug,
		Name:      record.Name,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}
}
