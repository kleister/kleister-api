package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/forge"
)

// ListForges implements the v1.ServerInterface.
func (a *API) ListForges(ctx context.Context, request ListForgesRequestObject) (ListForgesResponseObject, error) {
	params := model.ListParams{}

	if request.Params.Search != nil {
		params.Search = FromPtr(request.Params.Search)
	}

	records, count, err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		params,
	)

	if err != nil {
		return ListForges500JSONResponse{
			Message: ToPtr("Failed to load forge versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Forge, len(records))
	for id, record := range records {
		payload[id] = a.convertForge(record)
	}

	return ListForges200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// UpdateForge implements the v1.ServerInterface.
func (a *API) UpdateForge(ctx context.Context, _ UpdateForgeRequestObject) (UpdateForgeResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateForge403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	versions, err := forge.FetchRemote()

	if err != nil {
		return UpdateForge500JSONResponse{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).Sync(
		ctx,
		versions,
	); err != nil {
		return UpdateForge500JSONResponse{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateForge200JSONResponse{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListForgeBuilds implements the v1.ServerInterface.
func (a *API) ListForgeBuilds(ctx context.Context, request ListForgeBuildsRequestObject) (ListForgeBuildsResponseObject, error) {
	record, err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ForgeId,
	)

	if err != nil {
		if errors.Is(err, forge.ErrNotFound) {
			return ListForgeBuilds404JSONResponse{
				Message: ToPtr("Failed to find forge"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListForgeBuilds500JSONResponse{
			Message: ToPtr("Failed to load forge"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).ListBuilds(
		ctx,
		model.ForgeBuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			ForgeID: request.ForgeId,
		},
	)

	if err != nil {
		return ListForgeBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildWithPack(record)
	}

	return ListForgeBuilds200JSONResponse{
		Total:  ToPtr(count),
		Forge:  ToPtr(a.convertForge(record)),
		Builds: ToPtr(payload),
	}, nil
}

// AttachForgeToBuild implements the v1.ServerInterface.
func (a *API) AttachForgeToBuild(ctx context.Context, request AttachForgeToBuildRequestObject) (AttachForgeToBuildResponseObject, error) {
	if err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).AttachBuild(
		ctx,
		model.ForgeBuildParams{
			ForgeID: request.ForgeId,
			PackID:  request.Body.Pack,
			BuildID: request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, forge.ErrNotFound) {
			return AttachForgeToBuild404JSONResponse{
				Message: ToPtr("Failed to find forge or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, forge.ErrAlreadyAssigned) {
			return AttachForgeToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachForgeToBuild500JSONResponse{
			Message: ToPtr("Failed to attach forge to build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return AttachForgeToBuild200JSONResponse{
		Message: ToPtr("Successfully attached forge to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteForgeFromBuild implements the v1.ServerInterface.
func (a *API) DeleteForgeFromBuild(ctx context.Context, request DeleteForgeFromBuildRequestObject) (DeleteForgeFromBuildResponseObject, error) {
	if err := a.forge.WithPrincipal(
		current.GetUser(ctx),
	).DropBuild(
		ctx,
		model.ForgeBuildParams{
			ForgeID: request.ForgeId,
			PackID:  request.Body.Pack,
			BuildID: request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, forge.ErrNotFound) {
			return DeleteForgeFromBuild404JSONResponse{
				Message: ToPtr("Failed to find forge or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, forge.ErrNotAssigned) {
			return DeleteForgeFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteForgeFromBuild500JSONResponse{
			Message: ToPtr("Failed to drop forge from build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return DeleteForgeFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped forge from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertForge(record *model.Forge) Forge {
	result := Forge{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Minecraft: ToPtr(record.Minecraft),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
