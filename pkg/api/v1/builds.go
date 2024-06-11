package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	buildversions "github.com/kleister/kleister-api/pkg/service/build_versions"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/fabric"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/neoforge"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/quilt"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListBuilds implements the v1.ServerInterface.
func (a *API) ListBuilds(ctx context.Context, request ListBuildsRequestObject) (ListBuildsResponseObject, error) {
	records, count, err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.BuildParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			PackID: request.PackId,
		},
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ListBuilds404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	return ListBuilds200JSONResponse{
		Total:  ToPtr(count),
		Builds: ToPtr(payload),
	}, nil
}

// ShowBuild implements the v1.ServerInterface.
func (a *API) ShowBuild(ctx context.Context, request ShowBuildRequestObject) (ShowBuildResponseObject, error) {
	record, err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.BuildParams{
			PackID:  request.PackId,
			BuildID: request.BuildId,
		},
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ShowBuild404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, builds.ErrNotFound) {
			return ShowBuild404JSONResponse{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowBuild500JSONResponse{
			Message: ToPtr("Failed to load build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowBuild200JSONResponse(
		a.convertBuild(record),
	), nil
}

// CreateBuild implements the v1.ServerInterface.
func (a *API) CreateBuild(ctx context.Context, request CreateBuildRequestObject) (CreateBuildResponseObject, error) {
	record := &model.Build{}

	if request.Body.MinecraftId != nil {
		ref, err := a.minecraft.Show(ctx, FromPtr(request.Body.MinecraftId))

		if err != nil {
			if errors.Is(err, minecraft.ErrNotFound) {
				return CreateBuild404JSONResponse{
					Message: ToPtr("Failed to find minecraft"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return CreateBuild500JSONResponse{
				Message: ToPtr("Failed to load minecraft"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.MinecraftID = ToPtr(ref.ID)
	}

	if request.Body.ForgeId != nil {
		ref, err := a.forge.Show(ctx, FromPtr(request.Body.ForgeId))

		if err != nil {
			if errors.Is(err, forge.ErrNotFound) {
				return CreateBuild404JSONResponse{
					Message: ToPtr("Failed to find forge"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return CreateBuild500JSONResponse{
				Message: ToPtr("Failed to load forge"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.ForgeID = ToPtr(ref.ID)
	}

	if request.Body.NeoforgeId != nil {
		ref, err := a.neoforge.Show(ctx, FromPtr(request.Body.NeoforgeId))

		if err != nil {
			if errors.Is(err, neoforge.ErrNotFound) {
				return CreateBuild404JSONResponse{
					Message: ToPtr("Failed to find neoforge"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return CreateBuild500JSONResponse{
				Message: ToPtr("Failed to load neoforge"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.NeoforgeID = ToPtr(ref.ID)
	}

	if request.Body.QuiltId != nil {
		ref, err := a.quilt.Show(ctx, FromPtr(request.Body.QuiltId))

		if err != nil {
			if errors.Is(err, quilt.ErrNotFound) {
				return CreateBuild404JSONResponse{
					Message: ToPtr("Failed to find quilt"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return CreateBuild500JSONResponse{
				Message: ToPtr("Failed to load quilt"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.QuiltID = ToPtr(ref.ID)
	}

	if request.Body.FabricId != nil {
		ref, err := a.fabric.Show(ctx, FromPtr(request.Body.FabricId))

		if err != nil {
			if errors.Is(err, fabric.ErrNotFound) {
				return CreateBuild404JSONResponse{
					Message: ToPtr("Failed to find fabric"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return CreateBuild500JSONResponse{
				Message: ToPtr("Failed to load fabric"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.FabricID = ToPtr(ref.ID)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Java != nil {
		record.Java = FromPtr(request.Body.Java)
	}

	if request.Body.Memory != nil {
		record.Memory = FromPtr(request.Body.Memory)
	}

	if request.Body.Latest != nil {
		record.Latest = FromPtr(request.Body.Latest)
	}

	if request.Body.Recommended != nil {
		record.Recommended = FromPtr(request.Body.Recommended)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	if err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Create(
		ctx,
		model.BuildParams{
			PackID: request.PackId,
		},
		record,
	); err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return CreateBuild404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return CreateBuild422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate build"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateBuild500JSONResponse{
			Message: ToPtr("Failed to create build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateBuild200JSONResponse(
		a.convertBuild(record),
	), nil
}

// UpdateBuild implements the v1.ServerInterface.
func (a *API) UpdateBuild(ctx context.Context, request UpdateBuildRequestObject) (UpdateBuildResponseObject, error) {
	record, err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.BuildParams{
			PackID:  request.PackId,
			BuildID: request.BuildId,
		},
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return UpdateBuild404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, builds.ErrNotFound) {
			return UpdateBuild404JSONResponse{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateBuild500JSONResponse{
			Message: ToPtr("Failed to load build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.MinecraftId != nil {
		ref, err := a.minecraft.Show(ctx, FromPtr(request.Body.MinecraftId))

		if err != nil {
			if errors.Is(err, minecraft.ErrNotFound) {
				return UpdateBuild404JSONResponse{
					Message: ToPtr("Failed to find minecraft"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return UpdateBuild500JSONResponse{
				Message: ToPtr("Failed to load minecraft"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.MinecraftID = ToPtr(ref.ID)
	}

	if request.Body.ForgeId != nil {
		ref, err := a.forge.Show(ctx, FromPtr(request.Body.ForgeId))

		if err != nil {
			if errors.Is(err, forge.ErrNotFound) {
				return UpdateBuild404JSONResponse{
					Message: ToPtr("Failed to find forge"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return UpdateBuild500JSONResponse{
				Message: ToPtr("Failed to load forge"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.ForgeID = ToPtr(ref.ID)
	}

	if request.Body.NeoforgeId != nil {
		ref, err := a.neoforge.Show(ctx, FromPtr(request.Body.NeoforgeId))

		if err != nil {
			if errors.Is(err, neoforge.ErrNotFound) {
				return UpdateBuild404JSONResponse{
					Message: ToPtr("Failed to find neoforge"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return UpdateBuild500JSONResponse{
				Message: ToPtr("Failed to load neoforge"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.NeoforgeID = ToPtr(ref.ID)
	}

	if request.Body.QuiltId != nil {
		ref, err := a.quilt.Show(ctx, FromPtr(request.Body.QuiltId))

		if err != nil {
			if errors.Is(err, quilt.ErrNotFound) {
				return UpdateBuild404JSONResponse{
					Message: ToPtr("Failed to find quilt"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return UpdateBuild500JSONResponse{
				Message: ToPtr("Failed to load quilt"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.QuiltID = ToPtr(ref.ID)
	}

	if request.Body.FabricId != nil {
		ref, err := a.fabric.Show(ctx, FromPtr(request.Body.FabricId))

		if err != nil {
			if errors.Is(err, fabric.ErrNotFound) {
				return UpdateBuild404JSONResponse{
					Message: ToPtr("Failed to find fabric"),
					Status:  ToPtr(http.StatusNotFound),
				}, nil
			}

			return UpdateBuild500JSONResponse{
				Message: ToPtr("Failed to load fabric"),
				Status:  ToPtr(http.StatusInternalServerError),
			}, nil
		}

		record.FabricID = ToPtr(ref.ID)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Java != nil {
		record.Java = FromPtr(request.Body.Java)
	}

	if request.Body.Memory != nil {
		record.Memory = FromPtr(request.Body.Memory)
	}

	if request.Body.Latest != nil {
		record.Latest = FromPtr(request.Body.Latest)
	}

	if request.Body.Recommended != nil {
		record.Recommended = FromPtr(request.Body.Recommended)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	if err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Update(
		ctx,
		model.BuildParams{
			PackID:  record.PackID,
			BuildID: record.ID,
		},
		record,
	); err != nil {
		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return UpdateBuild422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate build"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateBuild500JSONResponse{
			Message: ToPtr("Failed to update build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateBuild200JSONResponse(
		a.convertBuild(record),
	), nil
}

// DeleteBuild implements the v1.ServerInterface.
func (a *API) DeleteBuild(ctx context.Context, request DeleteBuildRequestObject) (DeleteBuildResponseObject, error) {
	record, err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.BuildParams{
			PackID:  request.PackId,
			BuildID: request.BuildId,
		},
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return DeleteBuild404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, builds.ErrNotFound) {
			return DeleteBuild404JSONResponse{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteBuild500JSONResponse{
			Message: ToPtr("Failed to load build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		model.BuildParams{
			PackID:  record.PackID,
			BuildID: record.ID,
		},
	); err != nil {
		return DeleteBuild400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete build"),
		}, nil
	}

	return DeleteBuild200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted build"),
	}, nil
}

// ListBuildVersions implements the v1.ServerInterface.
func (a *API) ListBuildVersions(ctx context.Context, request ListBuildVersionsRequestObject) (ListBuildVersionsResponseObject, error) {
	record, err := a.builds.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.BuildParams{
			PackID:  request.PackId,
			BuildID: request.BuildId,
		},
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ListBuildVersions404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, builds.ErrNotFound) {
			return ListBuildVersions404JSONResponse{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListBuildVersions500JSONResponse{
			Message: ToPtr("Failed to load build"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.buildversions.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.BuildVersionParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			PackID:  record.PackID,
			BuildID: record.ID,
		},
	)

	if err != nil {
		return ListBuildVersions500JSONResponse{
			Message: ToPtr("Failed to load versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]BuildVersion, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildVersion(record)
	}

	return ListBuildVersions200JSONResponse{
		Total:    ToPtr(count),
		Pack:     ToPtr(a.convertPack(record.Pack)),
		Build:    ToPtr(a.convertBuild(record)),
		Versions: ToPtr(payload),
	}, nil
}

// AttachBuildToVersion implements the v1.ServerInterface.
func (a *API) AttachBuildToVersion(ctx context.Context, request AttachBuildToVersionRequestObject) (AttachBuildToVersionResponseObject, error) {
	if err := a.buildversions.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.BuildVersionParams{
			PackID:    request.PackId,
			BuildID:   request.BuildId,
			ModID:     request.Body.Mod,
			VersionID: request.Body.Version,
		},
	); err != nil {
		if errors.Is(err, buildversions.ErrNotFound) {
			return AttachBuildToVersion404JSONResponse{
				Message: ToPtr("Failed to find build or version"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, buildversions.ErrAlreadyAssigned) {
			return AttachBuildToVersion412JSONResponse{
				Message: ToPtr("Version is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachBuildToVersion500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach build to version"),
		}, nil
	}

	return AttachBuildToVersion200JSONResponse{
		Message: ToPtr("Successfully attached build to version"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteBuildFromVersion implements the v1.ServerInterface.
func (a *API) DeleteBuildFromVersion(ctx context.Context, request DeleteBuildFromVersionRequestObject) (DeleteBuildFromVersionResponseObject, error) {
	if err := a.buildversions.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.BuildVersionParams{
			PackID:    request.PackId,
			BuildID:   request.BuildId,
			ModID:     request.Body.Mod,
			VersionID: request.Body.Version,
		},
	); err != nil {
		if errors.Is(err, buildversions.ErrNotFound) {
			return DeleteBuildFromVersion404JSONResponse{
				Message: ToPtr("Failed to find build or version"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, buildversions.ErrNotAssigned) {
			return DeleteBuildFromVersion412JSONResponse{
				Message: ToPtr("Version is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteBuildFromVersion500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop build from version"),
		}, nil
	}

	return DeleteBuildFromVersion200JSONResponse{
		Message: ToPtr("Successfully dropped build from version"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertBuild(record *model.Build) Build {
	result := Build{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Java:      ToPtr(record.Java),
		Memory:    ToPtr(record.Memory),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
		Pack:      ToPtr(a.convertPack(record.Pack)),
	}

	if record.MinecraftID != nil {
		result.MinecraftId = record.MinecraftID
		result.Minecraft = ToPtr(a.convertMinecraft(record.Minecraft))
	}

	if record.ForgeID != nil {
		result.ForgeId = record.ForgeID
		result.Forge = ToPtr(a.convertForge(record.Forge))
	}

	if record.NeoforgeID != nil {
		result.NeoforgeId = record.NeoforgeID
		result.Neoforge = ToPtr(a.convertNeoforge(record.Neoforge))
	}

	if record.QuiltID != nil {
		result.QuiltId = record.QuiltID
		result.Quilt = ToPtr(a.convertQuilt(record.Quilt))
	}

	if record.FabricID != nil {
		result.FabricId = record.FabricID
		result.Fabric = ToPtr(a.convertFabric(record.Fabric))
	}

	return result
}

func (a *API) convertBuildWithPack(record *model.Build) Build {
	result := a.convertBuild(record)
	result.Pack = ToPtr(a.convertPack(record.Pack))

	return result
}

func (a *API) convertBuildVersion(record *model.BuildVersion) BuildVersion {
	result := BuildVersion{
		BuildId:   record.BuildID,
		VersionId: record.VersionID,
		Version:   ToPtr(a.convertVersion(record.Version)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
