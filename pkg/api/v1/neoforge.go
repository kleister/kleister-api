package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/neoforge"
)

// ListNeoforges implements the v1.ServerInterface.
func (a *API) ListNeoforges(ctx context.Context, request ListNeoforgesRequestObject) (ListNeoforgesResponseObject, error) {
	params := model.ListParams{}

	if request.Params.Search != nil {
		params.Search = FromPtr(request.Params.Search)
	}

	records, count, err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		params,
	)

	if err != nil {
		return ListNeoforges500JSONResponse{
			Message: ToPtr("Failed to load neoforge versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Neoforge, len(records))
	for id, record := range records {
		payload[id] = a.convertNeoforge(record)
	}

	return ListNeoforges200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// UpdateNeoforge implements the v1.ServerInterface.
func (a *API) UpdateNeoforge(ctx context.Context, _ UpdateNeoforgeRequestObject) (UpdateNeoforgeResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateNeoforge403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	versions, err := neoforge.FetchRemote()

	if err != nil {
		return UpdateNeoforge500JSONResponse{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).Sync(
		ctx,
		versions,
	); err != nil {
		return UpdateNeoforge500JSONResponse{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateNeoforge200JSONResponse{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListNeoforgeBuilds implements the v1.ServerInterface.
func (a *API) ListNeoforgeBuilds(ctx context.Context, request ListNeoforgeBuildsRequestObject) (ListNeoforgeBuildsResponseObject, error) {
	record, err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.NeoforgeId,
	)

	if err != nil {
		if errors.Is(err, neoforge.ErrNotFound) {
			return ListNeoforgeBuilds404JSONResponse{
				Message: ToPtr("Failed to find neoforge"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListNeoforgeBuilds500JSONResponse{
			Message: ToPtr("Failed to load neoforge"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).ListBuilds(
		ctx,
		model.NeoforgeBuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			NeoforgeID: request.NeoforgeId,
		},
	)

	if err != nil {
		return ListNeoforgeBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildWithPack(record)
	}

	return ListNeoforgeBuilds200JSONResponse{
		Total:    ToPtr(count),
		Neoforge: ToPtr(a.convertNeoforge(record)),
		Builds:   ToPtr(payload),
	}, nil
}

// AttachNeoforgeToBuild implements the v1.ServerInterface.
func (a *API) AttachNeoforgeToBuild(ctx context.Context, request AttachNeoforgeToBuildRequestObject) (AttachNeoforgeToBuildResponseObject, error) {
	if err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).AttachBuild(
		ctx,
		model.NeoforgeBuildParams{
			NeoforgeID: request.NeoforgeId,
			PackID:     request.Body.Pack,
			BuildID:    request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, neoforge.ErrNotFound) {
			return AttachNeoforgeToBuild404JSONResponse{
				Message: ToPtr("Failed to find neoforge or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, neoforge.ErrAlreadyAssigned) {
			return AttachNeoforgeToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachNeoforgeToBuild500JSONResponse{
			Message: ToPtr("Failed to attach neoforge to build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return AttachNeoforgeToBuild200JSONResponse{
		Message: ToPtr("Successfully attached neoforge to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteNeoforgeFromBuild implements the v1.ServerInterface.
func (a *API) DeleteNeoforgeFromBuild(ctx context.Context, request DeleteNeoforgeFromBuildRequestObject) (DeleteNeoforgeFromBuildResponseObject, error) {
	if err := a.neoforge.WithPrincipal(
		current.GetUser(ctx),
	).DropBuild(
		ctx,
		model.NeoforgeBuildParams{
			NeoforgeID: request.NeoforgeId,
			PackID:     request.Body.Pack,
			BuildID:    request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, neoforge.ErrNotFound) {
			return DeleteNeoforgeFromBuild404JSONResponse{
				Message: ToPtr("Failed to find neoforge or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, neoforge.ErrNotAssigned) {
			return DeleteNeoforgeFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteNeoforgeFromBuild500JSONResponse{
			Message: ToPtr("Failed to drop neoforge from build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return DeleteNeoforgeFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped neoforge from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertNeoforge(record *model.Neoforge) Neoforge {
	result := Neoforge{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
