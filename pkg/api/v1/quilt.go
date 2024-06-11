package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/quilt"
)

// ListQuilts implements the v1.ServerInterface.
func (a *API) ListQuilts(ctx context.Context, request ListQuiltsRequestObject) (ListQuiltsResponseObject, error) {
	params := model.ListParams{}

	if request.Params.Search != nil {
		params.Search = FromPtr(request.Params.Search)
	}

	records, count, err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		params,
	)

	if err != nil {
		return ListQuilts500JSONResponse{
			Message: ToPtr("Failed to load quilt versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Quilt, len(records))
	for id, record := range records {
		payload[id] = a.convertQuilt(record)
	}

	return ListQuilts200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// UpdateQuilt implements the v1.ServerInterface.
func (a *API) UpdateQuilt(ctx context.Context, _ UpdateQuiltRequestObject) (UpdateQuiltResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateQuilt403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	versions, err := quilt.FetchRemote()

	if err != nil {
		return UpdateQuilt500JSONResponse{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).Sync(
		ctx,
		versions,
	); err != nil {
		return UpdateQuilt500JSONResponse{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateQuilt200JSONResponse{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListQuiltBuilds implements the v1.ServerInterface.
func (a *API) ListQuiltBuilds(ctx context.Context, request ListQuiltBuildsRequestObject) (ListQuiltBuildsResponseObject, error) {
	record, err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.QuiltId,
	)

	if err != nil {
		if errors.Is(err, quilt.ErrNotFound) {
			return ListQuiltBuilds404JSONResponse{
				Message: ToPtr("Failed to find quilt"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListQuiltBuilds500JSONResponse{
			Message: ToPtr("Failed to load quilt"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).ListBuilds(
		ctx,
		model.QuiltBuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			QuiltID: request.QuiltId,
		},
	)

	if err != nil {
		return ListQuiltBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildWithPack(record)
	}

	return ListQuiltBuilds200JSONResponse{
		Total:  ToPtr(count),
		Quilt:  ToPtr(a.convertQuilt(record)),
		Builds: ToPtr(payload),
	}, nil
}

// AttachQuiltToBuild implements the v1.ServerInterface.
func (a *API) AttachQuiltToBuild(ctx context.Context, request AttachQuiltToBuildRequestObject) (AttachQuiltToBuildResponseObject, error) {
	if err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).AttachBuild(
		ctx,
		model.QuiltBuildParams{
			QuiltID: request.QuiltId,
			PackID:  request.Body.Pack,
			BuildID: request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, quilt.ErrNotFound) {
			return AttachQuiltToBuild404JSONResponse{
				Message: ToPtr("Failed to find quilt or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, quilt.ErrAlreadyAssigned) {
			return AttachQuiltToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachQuiltToBuild500JSONResponse{
			Message: ToPtr("Failed to attach quilt to build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return AttachQuiltToBuild200JSONResponse{
		Message: ToPtr("Successfully attached quilt to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteQuiltFromBuild implements the v1.ServerInterface.
func (a *API) DeleteQuiltFromBuild(ctx context.Context, request DeleteQuiltFromBuildRequestObject) (DeleteQuiltFromBuildResponseObject, error) {
	if err := a.quilt.WithPrincipal(
		current.GetUser(ctx),
	).DropBuild(
		ctx,
		model.QuiltBuildParams{
			QuiltID: request.QuiltId,
			PackID:  request.Body.Pack,
			BuildID: request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, quilt.ErrNotFound) {
			return DeleteQuiltFromBuild404JSONResponse{
				Message: ToPtr("Failed to find quilt or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, quilt.ErrNotAssigned) {
			return DeleteQuiltFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteQuiltFromBuild500JSONResponse{
			Message: ToPtr("Failed to drop quilt from build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return DeleteQuiltFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped quilt from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertQuilt(record *model.Quilt) Quilt {
	result := Quilt{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
