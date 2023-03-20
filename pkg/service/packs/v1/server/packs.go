package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs/repository"
	packs "github.com/kleister/kleister-api/pkg/service/packs/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the PacksServiceHandler interface.
func (s *PacksServer) List(
	ctx context.Context,
	req *connect.Request[packs.ListRequest],
) (*connect.Response[packs.ListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.List(ctx, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Pack, len(records))
	for id, record := range records {
		payload[id] = convertPack(record)
	}

	return connect.NewResponse(&packs.ListResponse{
		Packs: payload,
	}), nil
}

// Create implements the PacksServiceHandler interface.
func (s *PacksServer) Create(
	ctx context.Context,
	req *connect.Request[packs.CreateRequest],
) (*connect.Response[packs.CreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record := &model.Pack{}

	if req.Msg.Pack.Slug != nil {
		record.Slug = *req.Msg.Pack.Slug
	}

	if req.Msg.Pack.Name != "" {
		record.Name = req.Msg.Pack.Name
	}

	// TODO: Add back upload

	// TODO: Add icon upload

	// TODO: Add logo upload

	if req.Msg.Pack.Website != nil {
		record.Website = *req.Msg.Pack.Website
	}

	if req.Msg.Pack.Published != nil {
		record.Published = *req.Msg.Pack.Published
	}

	if req.Msg.Pack.Private != nil {
		record.Private = *req.Msg.Pack.Private
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

	return connect.NewResponse(&packs.CreateResponse{
		Pack: convertPack(created),
	}), nil
}

// Update implements the PacksServiceHandler interface.
func (s *PacksServer) Update(
	ctx context.Context,
	req *connect.Request[packs.UpdateRequest],
) (*connect.Response[packs.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrPackNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.Pack.Slug != nil {
		record.Slug = *req.Msg.Pack.Slug
	}

	if req.Msg.Pack.Name != nil {
		record.Name = *req.Msg.Pack.Name
	}

	// TODO: Add back upload

	// TODO: Add icon upload

	// TODO: Add logo upload

	if req.Msg.Pack.Website != nil {
		record.Website = *req.Msg.Pack.Website
	}

	if req.Msg.Pack.Published != nil {
		record.Published = *req.Msg.Pack.Published
	}

	if req.Msg.Pack.Private != nil {
		record.Private = *req.Msg.Pack.Private
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

	return connect.NewResponse(&packs.UpdateResponse{
		Pack: convertPack(updated),
	}), nil
}

// Show implements the PacksServiceHandler interface.
func (s *PacksServer) Show(
	ctx context.Context,
	req *connect.Request[packs.ShowRequest],
) (*connect.Response[packs.ShowResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrPackNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.ShowResponse{
		Pack: convertPack(record),
	}), nil
}

// Delete implements the PacksServiceHandler interface.
func (s *PacksServer) Delete(
	ctx context.Context,
	req *connect.Request[packs.DeleteRequest],
) (*connect.Response[packs.DeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Delete(ctx, req.Msg.Id); err != nil {
		if err == repository.ErrPackNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&packs.DeleteResponse{
		Message: "successfully deleted pack",
	}), nil
}

func convertPack(record *model.Pack) *types.Pack {
	result := &types.Pack{
		Id:        record.ID,
		Slug:      record.Slug,
		Name:      record.Name,
		Website:   record.Website,
		Published: record.Published,
		Private:   record.Private,
		CreatedAt: timestamppb.New(record.CreatedAt),
		UpdatedAt: timestamppb.New(record.UpdatedAt),
	}

	if record.Back != nil {
		result.Back = convertPackBack(record.Back)
	}

	if record.Icon != nil {
		result.Icon = convertPackIcon(record.Icon)
	}

	if record.Logo != nil {
		result.Logo = convertPackLogo(record.Logo)
	}

	return result
}

func convertPackBack(record *model.PackBack) *types.PackBack {
	return &types.PackBack{
		Id:          record.ID,
		PackId:      record.PackID,
		Slug:        record.Slug,
		ContentType: record.ContentType,
		Md5:         record.MD5,
		Path:        record.Path,
		Url:         record.URL,
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   timestamppb.New(record.UpdatedAt),
	}
}

func convertPackIcon(record *model.PackIcon) *types.PackIcon {
	return &types.PackIcon{
		Id:          record.ID,
		PackId:      record.PackID,
		Slug:        record.Slug,
		ContentType: record.ContentType,
		Md5:         record.MD5,
		Path:        record.Path,
		Url:         record.URL,
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   timestamppb.New(record.UpdatedAt),
	}
}

func convertPackLogo(record *model.PackLogo) *types.PackLogo {
	return &types.PackLogo{
		Id:          record.ID,
		PackId:      record.PackID,
		Slug:        record.Slug,
		ContentType: record.ContentType,
		Md5:         record.MD5,
		Path:        record.Path,
		Url:         record.URL,
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   timestamppb.New(record.UpdatedAt),
	}
}
