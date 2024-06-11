package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/fabric"
)

// ListFabrics implements the v1.ServerInterface.
func (a *API) ListFabrics(ctx context.Context, request ListFabricsRequestObject) (ListFabricsResponseObject, error) {
	params := model.ListParams{}

	if request.Params.Search != nil {
		params.Search = FromPtr(request.Params.Search)
	}

	records, count, err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		params,
	)

	if err != nil {
		return ListFabrics500JSONResponse{
			Message: ToPtr("Failed to load fabric versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Fabric, len(records))
	for id, record := range records {
		payload[id] = a.convertFabric(record)
	}

	return ListFabrics200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// UpdateFabric implements the v1.ServerInterface.
func (a *API) UpdateFabric(ctx context.Context, _ UpdateFabricRequestObject) (UpdateFabricResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateFabric403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	versions, err := fabric.FetchRemote()

	if err != nil {
		return UpdateFabric500JSONResponse{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).Sync(
		ctx,
		versions,
	); err != nil {
		return UpdateFabric500JSONResponse{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateFabric200JSONResponse{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListFabricBuilds implements the v1.ServerInterface.
func (a *API) ListFabricBuilds(ctx context.Context, request ListFabricBuildsRequestObject) (ListFabricBuildsResponseObject, error) {
	record, err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.FabricId,
	)

	if err != nil {
		if errors.Is(err, fabric.ErrNotFound) {
			return ListFabricBuilds404JSONResponse{
				Message: ToPtr("Failed to find fabric"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListFabricBuilds500JSONResponse{
			Message: ToPtr("Failed to load fabric"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).ListBuilds(
		ctx,
		model.FabricBuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			FabricID: request.FabricId,
		},
	)

	if err != nil {
		return ListFabricBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildWithPack(record)
	}

	return ListFabricBuilds200JSONResponse{
		Total:  ToPtr(count),
		Fabric: ToPtr(a.convertFabric(record)),
		Builds: ToPtr(payload),
	}, nil
}

// AttachFabricToBuild implements the v1.ServerInterface.
func (a *API) AttachFabricToBuild(ctx context.Context, request AttachFabricToBuildRequestObject) (AttachFabricToBuildResponseObject, error) {
	if err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).AttachBuild(
		ctx,
		model.FabricBuildParams{
			FabricID: request.FabricId,
			PackID:   request.Body.Pack,
			BuildID:  request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, fabric.ErrNotFound) {
			return AttachFabricToBuild404JSONResponse{
				Message: ToPtr("Failed to find fabric or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, fabric.ErrAlreadyAssigned) {
			return AttachFabricToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachFabricToBuild500JSONResponse{
			Message: ToPtr("Failed to attach fabric to build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return AttachFabricToBuild200JSONResponse{
		Message: ToPtr("Successfully attached fabric to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteFabricFromBuild implements the v1.ServerInterface.
func (a *API) DeleteFabricFromBuild(ctx context.Context, request DeleteFabricFromBuildRequestObject) (DeleteFabricFromBuildResponseObject, error) {
	if err := a.fabric.WithPrincipal(
		current.GetUser(ctx),
	).DropBuild(
		ctx,
		model.FabricBuildParams{
			FabricID: request.FabricId,
			PackID:   request.Body.Pack,
			BuildID:  request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, fabric.ErrNotFound) {
			return DeleteFabricFromBuild404JSONResponse{
				Message: ToPtr("Failed to find fabric or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, fabric.ErrNotAssigned) {
			return DeleteFabricFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteFabricFromBuild500JSONResponse{
			Message: ToPtr("Failed to drop fabric from build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return DeleteFabricFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped fabric from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertFabric(record *model.Fabric) Fabric {
	result := Fabric{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
