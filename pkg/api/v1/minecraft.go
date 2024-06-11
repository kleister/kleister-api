package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
)

// ListMinecrafts implements the v1.ServerInterface.
func (a *API) ListMinecrafts(ctx context.Context, request ListMinecraftsRequestObject) (ListMinecraftsResponseObject, error) {
	params := model.ListParams{}

	if request.Params.Search != nil {
		params.Search = FromPtr(request.Params.Search)
	}

	records, count, err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		params,
	)

	if err != nil {
		return ListMinecrafts500JSONResponse{
			Message: ToPtr("Failed to load minecraft versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Minecraft, len(records))
	for id, record := range records {
		payload[id] = a.convertMinecraft(record)
	}

	return ListMinecrafts200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// UpdateMinecraft implements the v1.ServerInterface.
func (a *API) UpdateMinecraft(ctx context.Context, _ UpdateMinecraftRequestObject) (UpdateMinecraftResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateMinecraft403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	versions, err := minecraft.FetchRemote()

	if err != nil {
		return UpdateMinecraft500JSONResponse{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).Sync(
		ctx,
		versions,
	); err != nil {
		return UpdateMinecraft500JSONResponse{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateMinecraft200JSONResponse{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListMinecraftBuilds implements the v1.ServerInterface.
func (a *API) ListMinecraftBuilds(ctx context.Context, request ListMinecraftBuildsRequestObject) (ListMinecraftBuildsResponseObject, error) {
	record, err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.MinecraftId,
	)

	if err != nil {
		if errors.Is(err, minecraft.ErrNotFound) {
			return ListMinecraftBuilds404JSONResponse{
				Message: ToPtr("Failed to find minecraft"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListMinecraftBuilds500JSONResponse{
			Message: ToPtr("Failed to load minecraft"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).ListBuilds(
		ctx,
		model.MinecraftBuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			MinecraftID: request.MinecraftId,
		},
	)

	if err != nil {
		return ListMinecraftBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildWithPack(record)
	}

	return ListMinecraftBuilds200JSONResponse{
		Total:     ToPtr(count),
		Minecraft: ToPtr(a.convertMinecraft(record)),
		Builds:    ToPtr(payload),
	}, nil
}

// AttachMinecraftToBuild implements the v1.ServerInterface.
func (a *API) AttachMinecraftToBuild(ctx context.Context, request AttachMinecraftToBuildRequestObject) (AttachMinecraftToBuildResponseObject, error) {
	if err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).AttachBuild(
		ctx,
		model.MinecraftBuildParams{
			MinecraftID: request.MinecraftId,
			PackID:      request.Body.Pack,
			BuildID:     request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, minecraft.ErrNotFound) {
			return AttachMinecraftToBuild404JSONResponse{
				Message: ToPtr("Failed to find minecraft or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, minecraft.ErrAlreadyAssigned) {
			return AttachMinecraftToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachMinecraftToBuild500JSONResponse{
			Message: ToPtr("Failed to attach minecraft to build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return AttachMinecraftToBuild200JSONResponse{
		Message: ToPtr("Successfully attached minecraft to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteMinecraftFromBuild implements the v1.ServerInterface.
func (a *API) DeleteMinecraftFromBuild(ctx context.Context, request DeleteMinecraftFromBuildRequestObject) (DeleteMinecraftFromBuildResponseObject, error) {
	if err := a.minecraft.WithPrincipal(
		current.GetUser(ctx),
	).DropBuild(
		ctx,
		model.MinecraftBuildParams{
			MinecraftID: request.MinecraftId,
			PackID:      request.Body.Pack,
			BuildID:     request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, minecraft.ErrNotFound) {
			return DeleteMinecraftFromBuild404JSONResponse{
				Message: ToPtr("Failed to find minecraft or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, minecraft.ErrNotAssigned) {
			return DeleteMinecraftFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteMinecraftFromBuild500JSONResponse{
			Message: ToPtr("Failed to drop minecraft from build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return DeleteMinecraftFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped minecraft from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertMinecraft(record *model.Minecraft) Minecraft {
	result := Minecraft{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Type:      ToPtr(record.Type),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
