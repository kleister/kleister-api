package serverv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods/repository"
	mods "github.com/kleister/kleister-api/pkg/service/mods/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// List implements the ModsServiceHandler interface.
func (s *ModsServer) List(
	ctx context.Context,
	req *connect.Request[mods.ListRequest],
) (*connect.Response[mods.ListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	records, err := s.repository.List(ctx, req.Msg.Query)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	payload := make([]*types.Mod, len(records))
	for id, record := range records {
		payload[id] = convertMod(record)
	}

	return connect.NewResponse(&mods.ListResponse{
		Mods: payload,
	}), nil
}

// Create implements the ModsServiceHandler interface.
func (s *ModsServer) Create(
	ctx context.Context,
	req *connect.Request[mods.CreateRequest],
) (*connect.Response[mods.CreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record := &model.Mod{}

	if req.Msg.Mod.Slug != nil {
		record.Slug = *req.Msg.Mod.Slug
	}

	if req.Msg.Mod.Name != "" {
		record.Name = req.Msg.Mod.Name
	}

	if req.Msg.Mod.Side != nil {
		record.Side = *req.Msg.Mod.Side
	}

	if req.Msg.Mod.Description != nil {
		record.Description = *req.Msg.Mod.Description
	}

	if req.Msg.Mod.Author != nil {
		record.Author = *req.Msg.Mod.Author
	}

	if req.Msg.Mod.Website != nil {
		record.Website = *req.Msg.Mod.Website
	}

	if req.Msg.Mod.Donate != nil {
		record.Donate = *req.Msg.Mod.Donate
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

	return connect.NewResponse(&mods.CreateResponse{
		Mod: convertMod(created),
	}), nil
}

// Update implements the ModsServiceHandler interface.
func (s *ModsServer) Update(
	ctx context.Context,
	req *connect.Request[mods.UpdateRequest],
) (*connect.Response[mods.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrModNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.Mod.Slug != nil {
		record.Slug = *req.Msg.Mod.Slug
	}

	if req.Msg.Mod.Name != nil {
		record.Name = *req.Msg.Mod.Name
	}

	if req.Msg.Mod.Side != nil {
		record.Side = *req.Msg.Mod.Side
	}

	if req.Msg.Mod.Description != nil {
		record.Description = *req.Msg.Mod.Description
	}

	if req.Msg.Mod.Author != nil {
		record.Author = *req.Msg.Mod.Author
	}

	if req.Msg.Mod.Website != nil {
		record.Website = *req.Msg.Mod.Website
	}

	if req.Msg.Mod.Donate != nil {
		record.Donate = *req.Msg.Mod.Donate
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

	return connect.NewResponse(&mods.UpdateResponse{
		Mod: convertMod(updated),
	}), nil
}

// Show implements the ModsServiceHandler interface.
func (s *ModsServer) Show(
	ctx context.Context,
	req *connect.Request[mods.ShowRequest],
) (*connect.Response[mods.ShowResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	record, err := s.repository.Show(ctx, req.Msg.Id)

	if err != nil {
		if err == repository.ErrModNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.ShowResponse{
		Mod: convertMod(record),
	}), nil
}

// Delete implements the ModsServiceHandler interface.
func (s *ModsServer) Delete(
	ctx context.Context,
	req *connect.Request[mods.DeleteRequest],
) (*connect.Response[mods.DeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := s.repository.Delete(ctx, req.Msg.Id); err != nil {
		if err == repository.ErrModNotFound {
			return nil, connect.NewError(
				connect.CodeNotFound,
				err,
			)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&mods.DeleteResponse{
		Message: "successfully deleted mod",
	}), nil
}

func convertMod(record *model.Mod) *types.Mod {
	return &types.Mod{
		Id:          record.ID,
		Slug:        record.Slug,
		Name:        record.Name,
		Side:        record.Side,
		Description: record.Description,
		Author:      record.Author,
		Website:     record.Website,
		Donate:      record.Donate,
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   timestamppb.New(record.UpdatedAt),
	}
}
