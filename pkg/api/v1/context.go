package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

const (
	minecraftContext contextKey = "minecraft"
	forgeContext     contextKey = "forge"
	neoforgeContext  contextKey = "neoforge"
	quiltContext     contextKey = "quilt"
	fabricContext    contextKey = "fabric"
	packContext      contextKey = "pack"
	buildContext     contextKey = "build"
	modContext       contextKey = "mod"
	versionContext   contextKey = "version"
	groupContext     contextKey = "group"
	userContext      contextKey = "user"
)

// MinecraftToContext is used to put the requested minecraft into the context.
func (a *API) MinecraftToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "minecraft_id")

		record, err := a.storage.Minecraft.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrMinecraftNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find minecraft"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "MinecraftToContext").
				Str("minecraft", id).
				Msg("Failed to load minecraft")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load minecraft"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			minecraftContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// MinecraftFromContext is used to get the requested minecraft from the context.
func (a *API) MinecraftFromContext(ctx context.Context) *model.Minecraft {
	record, ok := ctx.Value(minecraftContext).(*model.Minecraft)

	if !ok {
		return nil
	}

	return record
}

// ForgeToContext is used to put the requested forge into the context.
func (a *API) ForgeToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "forge_id")

		record, err := a.storage.Forge.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrForgeNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find forge"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "ForgeToContext").
				Str("forge", id).
				Msg("Failed to load forge")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load forge"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			forgeContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ForgeFromContext is used to get the requested forge from the context.
func (a *API) ForgeFromContext(ctx context.Context) *model.Forge {
	record, ok := ctx.Value(forgeContext).(*model.Forge)

	if !ok {
		return nil
	}

	return record
}

// NeoforgeToContext is used to put the requested neoforge into the context.
func (a *API) NeoforgeToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "neoforge_id")

		record, err := a.storage.Neoforge.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrNeoforgeNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find neoforge"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "NeoforgeToContext").
				Str("neoforge", id).
				Msg("Failed to load neoforge")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load neoforge"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			neoforgeContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NeoforgeFromContext is used to get the requested neoforge from the context.
func (a *API) NeoforgeFromContext(ctx context.Context) *model.Neoforge {
	record, ok := ctx.Value(neoforgeContext).(*model.Neoforge)

	if !ok {
		return nil
	}

	return record
}

// QuiltToContext is used to put the requested quilt into the context.
func (a *API) QuiltToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "quilt_id")

		record, err := a.storage.Quilt.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrQuiltNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find quilt"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "QuiltToContext").
				Str("quilt", id).
				Msg("Failed to load quilt")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load quilt"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			quiltContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// QuiltFromContext is used to get the requested quilt from the context.
func (a *API) QuiltFromContext(ctx context.Context) *model.Quilt {
	record, ok := ctx.Value(quiltContext).(*model.Quilt)

	if !ok {
		return nil
	}

	return record
}

// FabricToContext is used to put the requested fabric into the context.
func (a *API) FabricToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "fabric_id")

		record, err := a.storage.Fabric.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrFabricNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find fabric"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "FabricToContext").
				Str("fabric", id).
				Msg("Failed to load fabric")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load fabric"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			fabricContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FabricFromContext is used to get the requested fabric from the context.
func (a *API) FabricFromContext(ctx context.Context) *model.Fabric {
	record, ok := ctx.Value(fabricContext).(*model.Fabric)

	if !ok {
		return nil
	}

	return record
}

// PackToContext is used to put the requested pack into the context.
func (a *API) PackToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "pack_id")

		record, err := a.storage.Packs.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrPackNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find pack"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "PackToContext").
				Str("pack", id).
				Msg("Failed to load pack")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load pack"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			packContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PackFromContext is used to get the requested pack from the context.
func (a *API) PackFromContext(ctx context.Context) *model.Pack {
	record, ok := ctx.Value(packContext).(*model.Pack)

	if !ok {
		return nil
	}

	return record
}

// BuildToContext is used to put the requested build into the context.
func (a *API) BuildToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pack := a.PackFromContext(ctx)

		if pack == nil {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		id := chi.URLParam(r, "build_id")

		record, err := a.storage.Builds.Show(
			ctx,
			pack,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrBuildNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find build"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "BuildToContext").
				Str("pack", pack.ID).
				Str("build", id).
				Msg("Failed to load build")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load build"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			buildContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BuildFromContext is used to get the requested build from the context.
func (a *API) BuildFromContext(ctx context.Context) *model.Build {
	record, ok := ctx.Value(buildContext).(*model.Build)

	if !ok {
		return nil
	}

	return record
}

// ModToContext is used to put the requested mod into the context.
func (a *API) ModToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "mod_id")

		record, err := a.storage.Mods.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrModNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find mod"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "ModToContext").
				Str("mod", id).
				Msg("Failed to load mod")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load mod"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			modContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ModFromContext is used to get the requested mod from the context.
func (a *API) ModFromContext(ctx context.Context) *model.Mod {
	record, ok := ctx.Value(modContext).(*model.Mod)

	if !ok {
		return nil
	}

	return record
}

// VersionToContext is used to put the requested version into the context.
func (a *API) VersionToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		mod := a.ModFromContext(ctx)

		if mod == nil {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		id := chi.URLParam(r, "version_id")

		record, err := a.storage.Versions.Show(
			ctx,
			mod,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrVersionNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find version"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "VersionToContext").
				Str("version", id).
				Msg("Failed to load version")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load version"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			versionContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// VersionFromContext is used to get the requested version from the context.
func (a *API) VersionFromContext(ctx context.Context) *model.Version {
	record, ok := ctx.Value(versionContext).(*model.Version)

	if !ok {
		return nil
	}

	return record
}

// GroupToContext is used to put the requested group into the context.
func (a *API) GroupToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "group_id")

		record, err := a.storage.Groups.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrGroupNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find group"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "GroupToContext").
				Str("group", id).
				Msg("Failed to load group")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load group"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			groupContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GroupFromContext is used to get the requested group from the context.
func (a *API) GroupFromContext(ctx context.Context) *model.Group {
	record, ok := ctx.Value(groupContext).(*model.Group)

	if !ok {
		return nil
	}

	return record
}

// UserToContext is used to put the requested user into the context.
func (a *API) UserToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "user_id")

		record, err := a.storage.Users.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find user"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "UserToContext").
				Str("user", id).
				Msg("Failed to load user")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load user"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			userContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserFromContext is used to get the requested user from the context.
func (a *API) UserFromContext(ctx context.Context) *model.User {
	record, ok := ctx.Value(userContext).(*model.User)

	if !ok {
		return nil
	}

	return record
}
