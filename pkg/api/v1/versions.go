package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	buildversions "github.com/kleister/kleister-api/pkg/service/build_versions"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListVersions implements the v1.ServerInterface.
func (a *API) ListVersions(ctx context.Context, request ListVersionsRequestObject) (ListVersionsResponseObject, error) {
	records, count, err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.VersionParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			ModID: request.ModId,
		},
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ListVersions404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListVersions500JSONResponse{
			Message: ToPtr("Failed to load versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Version, len(records))
	for id, record := range records {
		payload[id] = a.convertVersion(record)
	}

	return ListVersions200JSONResponse{
		Total:    ToPtr(count),
		Versions: ToPtr(payload),
	}, nil
}

// ShowVersion implements the v1.ServerInterface.
func (a *API) ShowVersion(ctx context.Context, request ShowVersionRequestObject) (ShowVersionResponseObject, error) {
	record, err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.VersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
		},
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ShowVersion404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, versions.ErrNotFound) {
			return ShowVersion404JSONResponse{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowVersion500JSONResponse{
			Message: ToPtr("Failed to load version"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowVersion200JSONResponse(
		a.convertVersion(record),
	), nil
}

// CreateVersion implements the v1.ServerInterface.
func (a *API) CreateVersion(ctx context.Context, request CreateVersionRequestObject) (CreateVersionResponseObject, error) {
	record := &model.Version{}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	// TODO: File

	if err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Create(
		ctx,
		model.VersionParams{
			ModID: request.ModId,
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

			return CreateVersion422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate version"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateVersion500JSONResponse{
			Message: ToPtr("Failed to create version"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateVersion200JSONResponse(
		a.convertVersion(record),
	), nil
}

// UpdateVersion implements the v1.ServerInterface.
func (a *API) UpdateVersion(ctx context.Context, request UpdateVersionRequestObject) (UpdateVersionResponseObject, error) {
	record, err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.VersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
		},
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return UpdateVersion404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, versions.ErrNotFound) {
			return UpdateVersion404JSONResponse{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateVersion500JSONResponse{
			Message: ToPtr("Failed to load version"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	// TODO: File

	if err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Update(
		ctx,
		model.VersionParams{
			ModID:     record.ModID,
			VersionID: record.ID,
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

			return UpdateVersion422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate version"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateVersion500JSONResponse{
			Message: ToPtr("Failed to update version"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateVersion200JSONResponse(
		a.convertVersion(record),
	), nil
}

// DeleteVersion implements the v1.ServerInterface.
func (a *API) DeleteVersion(ctx context.Context, request DeleteVersionRequestObject) (DeleteVersionResponseObject, error) {
	record, err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.VersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
		},
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return DeleteVersion404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, versions.ErrNotFound) {
			return DeleteVersion404JSONResponse{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteVersion500JSONResponse{
			Message: ToPtr("Failed to load version"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		model.VersionParams{
			ModID:     record.ModID,
			VersionID: record.ID,
		},
	); err != nil {
		return DeleteVersion400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete version"),
		}, nil
	}

	return DeleteVersion200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted version"),
	}, nil
}

// ListVersionBuilds implements the v1.ServerInterface.
func (a *API) ListVersionBuilds(ctx context.Context, request ListVersionBuildsRequestObject) (ListVersionBuildsResponseObject, error) {
	record, err := a.versions.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		model.VersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
		},
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ListVersionBuilds404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, versions.ErrNotFound) {
			return ListVersionBuilds404JSONResponse{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListVersionBuilds500JSONResponse{
			Message: ToPtr("Failed to load version"),
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
			ModID:     record.ModID,
			VersionID: record.ID,
		},
	)

	if err != nil {
		return ListVersionBuilds500JSONResponse{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]BuildVersion, len(records))
	for id, record := range records {
		payload[id] = a.convertVersionBuild(record)
	}

	return ListVersionBuilds200JSONResponse{
		Total:   ToPtr(count),
		Mod:     ToPtr(a.convertMod(record.Mod)),
		Version: ToPtr(a.convertVersion(record)),
		Builds:  ToPtr(payload),
	}, nil
}

// AttachVersionToBuild implements the v1.ServerInterface.
func (a *API) AttachVersionToBuild(ctx context.Context, request AttachVersionToBuildRequestObject) (AttachVersionToBuildResponseObject, error) {
	if err := a.buildversions.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.BuildVersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
			PackID:    request.Body.Pack,
			BuildID:   request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, buildversions.ErrNotFound) {
			return AttachVersionToBuild404JSONResponse{
				Message: ToPtr("Failed to find version or build"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, buildversions.ErrAlreadyAssigned) {
			return AttachVersionToBuild412JSONResponse{
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return AttachVersionToBuild500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach version to build"),
		}, nil
	}

	return AttachVersionToBuild200JSONResponse{
		Message: ToPtr("Successfully attached version to build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteVersionFromBuild implements the v1.ServerInterface.
func (a *API) DeleteVersionFromBuild(ctx context.Context, request DeleteVersionFromBuildRequestObject) (DeleteVersionFromBuildResponseObject, error) {
	if err := a.buildversions.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.BuildVersionParams{
			ModID:     request.ModId,
			VersionID: request.VersionId,
			PackID:    request.Body.Pack,
			BuildID:   request.Body.Build,
		},
	); err != nil {
		if errors.Is(err, buildversions.ErrNotFound) {
			return DeleteVersionFromBuild404JSONResponse{
				Message: ToPtr("Failed to find version or build"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, buildversions.ErrNotAssigned) {
			return DeleteVersionFromBuild412JSONResponse{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteVersionFromBuild500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop version from build"),
		}, nil
	}

	return DeleteVersionFromBuild200JSONResponse{
		Message: ToPtr("Successfully dropped version from build"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertVersion(record *model.Version) Version {
	result := Version{
		Id:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if record.Mod != nil {
		result.Mod = ToPtr(a.convertMod(record.Mod))
	}

	// TODO: File

	return result
}

func (a *API) convertVersionBuild(record *model.BuildVersion) BuildVersion {
	result := BuildVersion{
		VersionId: record.VersionID,
		BuildId:   record.BuildID,
		Build:     ToPtr(a.convertBuild(record.Build)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
