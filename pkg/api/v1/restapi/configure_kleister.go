// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/auth"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/forge"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/minecraft"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/mod"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/pack"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/profile"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/team"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/user"
)

//go:generate gorunpkg github.com/go-swagger/go-swagger/cmd/swagger generate server --target ../../v1 --name Kleister --spec ../../../../assets/apiv1.yml --exclude-main

func configureFlags(api *operations.KleisterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KleisterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AuthAuthLoginHandler == nil {
		api.AuthAuthLoginHandler = auth.AuthLoginHandlerFunc(func(params auth.AuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthLogin has not yet been implemented")
		})
	}
	if api.AuthAuthRefreshHandler == nil {
		api.AuthAuthRefreshHandler = auth.AuthRefreshHandlerFunc(func(params auth.AuthRefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthRefresh has not yet been implemented")
		})
	}
	if api.AuthAuthVerifyHandler == nil {
		api.AuthAuthVerifyHandler = auth.AuthVerifyHandlerFunc(func(params auth.AuthVerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthVerify has not yet been implemented")
		})
	}
	if api.PackBuildCreateHandler == nil {
		api.PackBuildCreateHandler = pack.BuildCreateHandlerFunc(func(params pack.BuildCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildCreate has not yet been implemented")
		})
	}
	if api.PackBuildDeleteHandler == nil {
		api.PackBuildDeleteHandler = pack.BuildDeleteHandlerFunc(func(params pack.BuildDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildDelete has not yet been implemented")
		})
	}
	if api.PackBuildIndexHandler == nil {
		api.PackBuildIndexHandler = pack.BuildIndexHandlerFunc(func(params pack.BuildIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildIndex has not yet been implemented")
		})
	}
	if api.PackBuildShowHandler == nil {
		api.PackBuildShowHandler = pack.BuildShowHandlerFunc(func(params pack.BuildShowParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildShow has not yet been implemented")
		})
	}
	if api.PackBuildUpdateHandler == nil {
		api.PackBuildUpdateHandler = pack.BuildUpdateHandlerFunc(func(params pack.BuildUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildUpdate has not yet been implemented")
		})
	}
	if api.PackBuildVersionAppendHandler == nil {
		api.PackBuildVersionAppendHandler = pack.BuildVersionAppendHandlerFunc(func(params pack.BuildVersionAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildVersionAppend has not yet been implemented")
		})
	}
	if api.PackBuildVersionDeleteHandler == nil {
		api.PackBuildVersionDeleteHandler = pack.BuildVersionDeleteHandlerFunc(func(params pack.BuildVersionDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildVersionDelete has not yet been implemented")
		})
	}
	if api.PackBuildVersionIndexHandler == nil {
		api.PackBuildVersionIndexHandler = pack.BuildVersionIndexHandlerFunc(func(params pack.BuildVersionIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.BuildVersionIndex has not yet been implemented")
		})
	}
	if api.ForgeForgeBuildAppendHandler == nil {
		api.ForgeForgeBuildAppendHandler = forge.ForgeBuildAppendHandlerFunc(func(params forge.ForgeBuildAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeBuildAppend has not yet been implemented")
		})
	}
	if api.ForgeForgeBuildDeleteHandler == nil {
		api.ForgeForgeBuildDeleteHandler = forge.ForgeBuildDeleteHandlerFunc(func(params forge.ForgeBuildDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeBuildDelete has not yet been implemented")
		})
	}
	if api.ForgeForgeBuildIndexHandler == nil {
		api.ForgeForgeBuildIndexHandler = forge.ForgeBuildIndexHandlerFunc(func(params forge.ForgeBuildIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeBuildIndex has not yet been implemented")
		})
	}
	if api.ForgeForgeIndexHandler == nil {
		api.ForgeForgeIndexHandler = forge.ForgeIndexHandlerFunc(func(params forge.ForgeIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeIndex has not yet been implemented")
		})
	}
	if api.ForgeForgeSearchHandler == nil {
		api.ForgeForgeSearchHandler = forge.ForgeSearchHandlerFunc(func(params forge.ForgeSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeSearch has not yet been implemented")
		})
	}
	if api.ForgeForgeUpdateHandler == nil {
		api.ForgeForgeUpdateHandler = forge.ForgeUpdateHandlerFunc(func(params forge.ForgeUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ForgeUpdate has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftBuildAppendHandler == nil {
		api.MinecraftMinecraftBuildAppendHandler = minecraft.MinecraftBuildAppendHandlerFunc(func(params minecraft.MinecraftBuildAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftBuildAppend has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftBuildDeleteHandler == nil {
		api.MinecraftMinecraftBuildDeleteHandler = minecraft.MinecraftBuildDeleteHandlerFunc(func(params minecraft.MinecraftBuildDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftBuildDelete has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftBuildIndexHandler == nil {
		api.MinecraftMinecraftBuildIndexHandler = minecraft.MinecraftBuildIndexHandlerFunc(func(params minecraft.MinecraftBuildIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftBuildIndex has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftIndexHandler == nil {
		api.MinecraftMinecraftIndexHandler = minecraft.MinecraftIndexHandlerFunc(func(params minecraft.MinecraftIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftIndex has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftSearchHandler == nil {
		api.MinecraftMinecraftSearchHandler = minecraft.MinecraftSearchHandlerFunc(func(params minecraft.MinecraftSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftSearch has not yet been implemented")
		})
	}
	if api.MinecraftMinecraftUpdateHandler == nil {
		api.MinecraftMinecraftUpdateHandler = minecraft.MinecraftUpdateHandlerFunc(func(params minecraft.MinecraftUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.MinecraftUpdate has not yet been implemented")
		})
	}
	if api.ModModCreateHandler == nil {
		api.ModModCreateHandler = mod.ModCreateHandlerFunc(func(params mod.ModCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModCreate has not yet been implemented")
		})
	}
	if api.ModModDeleteHandler == nil {
		api.ModModDeleteHandler = mod.ModDeleteHandlerFunc(func(params mod.ModDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModDelete has not yet been implemented")
		})
	}
	if api.ModModIndexHandler == nil {
		api.ModModIndexHandler = mod.ModIndexHandlerFunc(func(params mod.ModIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModIndex has not yet been implemented")
		})
	}
	if api.ModModShowHandler == nil {
		api.ModModShowHandler = mod.ModShowHandlerFunc(func(params mod.ModShowParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModShow has not yet been implemented")
		})
	}
	if api.ModModTeamAppendHandler == nil {
		api.ModModTeamAppendHandler = mod.ModTeamAppendHandlerFunc(func(params mod.ModTeamAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModTeamAppend has not yet been implemented")
		})
	}
	if api.ModModTeamDeleteHandler == nil {
		api.ModModTeamDeleteHandler = mod.ModTeamDeleteHandlerFunc(func(params mod.ModTeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModTeamDelete has not yet been implemented")
		})
	}
	if api.ModModTeamIndexHandler == nil {
		api.ModModTeamIndexHandler = mod.ModTeamIndexHandlerFunc(func(params mod.ModTeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModTeamIndex has not yet been implemented")
		})
	}
	if api.ModModTeamPermHandler == nil {
		api.ModModTeamPermHandler = mod.ModTeamPermHandlerFunc(func(params mod.ModTeamPermParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModTeamPerm has not yet been implemented")
		})
	}
	if api.ModModUpdateHandler == nil {
		api.ModModUpdateHandler = mod.ModUpdateHandlerFunc(func(params mod.ModUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModUpdate has not yet been implemented")
		})
	}
	if api.ModModUserAppendHandler == nil {
		api.ModModUserAppendHandler = mod.ModUserAppendHandlerFunc(func(params mod.ModUserAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModUserAppend has not yet been implemented")
		})
	}
	if api.ModModUserDeleteHandler == nil {
		api.ModModUserDeleteHandler = mod.ModUserDeleteHandlerFunc(func(params mod.ModUserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModUserDelete has not yet been implemented")
		})
	}
	if api.ModModUserIndexHandler == nil {
		api.ModModUserIndexHandler = mod.ModUserIndexHandlerFunc(func(params mod.ModUserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModUserIndex has not yet been implemented")
		})
	}
	if api.ModModUserPermHandler == nil {
		api.ModModUserPermHandler = mod.ModUserPermHandlerFunc(func(params mod.ModUserPermParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ModUserPerm has not yet been implemented")
		})
	}
	if api.PackPackCreateHandler == nil {
		api.PackPackCreateHandler = pack.PackCreateHandlerFunc(func(params pack.PackCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackCreate has not yet been implemented")
		})
	}
	if api.PackPackDeleteHandler == nil {
		api.PackPackDeleteHandler = pack.PackDeleteHandlerFunc(func(params pack.PackDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackDelete has not yet been implemented")
		})
	}
	if api.PackPackIndexHandler == nil {
		api.PackPackIndexHandler = pack.PackIndexHandlerFunc(func(params pack.PackIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackIndex has not yet been implemented")
		})
	}
	if api.PackPackShowHandler == nil {
		api.PackPackShowHandler = pack.PackShowHandlerFunc(func(params pack.PackShowParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackShow has not yet been implemented")
		})
	}
	if api.PackPackTeamAppendHandler == nil {
		api.PackPackTeamAppendHandler = pack.PackTeamAppendHandlerFunc(func(params pack.PackTeamAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackTeamAppend has not yet been implemented")
		})
	}
	if api.PackPackTeamDeleteHandler == nil {
		api.PackPackTeamDeleteHandler = pack.PackTeamDeleteHandlerFunc(func(params pack.PackTeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackTeamDelete has not yet been implemented")
		})
	}
	if api.PackPackTeamIndexHandler == nil {
		api.PackPackTeamIndexHandler = pack.PackTeamIndexHandlerFunc(func(params pack.PackTeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackTeamIndex has not yet been implemented")
		})
	}
	if api.PackPackTeamPermHandler == nil {
		api.PackPackTeamPermHandler = pack.PackTeamPermHandlerFunc(func(params pack.PackTeamPermParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackTeamPerm has not yet been implemented")
		})
	}
	if api.PackPackUpdateHandler == nil {
		api.PackPackUpdateHandler = pack.PackUpdateHandlerFunc(func(params pack.PackUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackUpdate has not yet been implemented")
		})
	}
	if api.PackPackUserAppendHandler == nil {
		api.PackPackUserAppendHandler = pack.PackUserAppendHandlerFunc(func(params pack.PackUserAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackUserAppend has not yet been implemented")
		})
	}
	if api.PackPackUserDeleteHandler == nil {
		api.PackPackUserDeleteHandler = pack.PackUserDeleteHandlerFunc(func(params pack.PackUserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackUserDelete has not yet been implemented")
		})
	}
	if api.PackPackUserIndexHandler == nil {
		api.PackPackUserIndexHandler = pack.PackUserIndexHandlerFunc(func(params pack.PackUserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackUserIndex has not yet been implemented")
		})
	}
	if api.PackPackUserPermHandler == nil {
		api.PackPackUserPermHandler = pack.PackUserPermHandlerFunc(func(params pack.PackUserPermParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PackUserPerm has not yet been implemented")
		})
	}
	if api.ProfileProfileShowHandler == nil {
		api.ProfileProfileShowHandler = profile.ProfileShowHandlerFunc(func(params profile.ProfileShowParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileShow has not yet been implemented")
		})
	}
	if api.ProfileProfileTokenHandler == nil {
		api.ProfileProfileTokenHandler = profile.ProfileTokenHandlerFunc(func(params profile.ProfileTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileToken has not yet been implemented")
		})
	}
	if api.ProfileProfileUpdateHandler == nil {
		api.ProfileProfileUpdateHandler = profile.ProfileUpdateHandlerFunc(func(params profile.ProfileUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileUpdate has not yet been implemented")
		})
	}
	if api.TeamTeamCreateHandler == nil {
		api.TeamTeamCreateHandler = team.TeamCreateHandlerFunc(func(params team.TeamCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamCreate has not yet been implemented")
		})
	}
	if api.TeamTeamDeleteHandler == nil {
		api.TeamTeamDeleteHandler = team.TeamDeleteHandlerFunc(func(params team.TeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamDelete has not yet been implemented")
		})
	}
	if api.TeamTeamIndexHandler == nil {
		api.TeamTeamIndexHandler = team.TeamIndexHandlerFunc(func(params team.TeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamIndex has not yet been implemented")
		})
	}
	if api.TeamTeamModAppendHandler == nil {
		api.TeamTeamModAppendHandler = team.TeamModAppendHandlerFunc(func(params team.TeamModAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamModAppend has not yet been implemented")
		})
	}
	if api.TeamTeamModDeleteHandler == nil {
		api.TeamTeamModDeleteHandler = team.TeamModDeleteHandlerFunc(func(params team.TeamModDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamModDelete has not yet been implemented")
		})
	}
	if api.TeamTeamModIndexHandler == nil {
		api.TeamTeamModIndexHandler = team.TeamModIndexHandlerFunc(func(params team.TeamModIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamModIndex has not yet been implemented")
		})
	}
	if api.TeamTeamModPermHandler == nil {
		api.TeamTeamModPermHandler = team.TeamModPermHandlerFunc(func(params team.TeamModPermParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamModPerm has not yet been implemented")
		})
	}
	if api.TeamTeamPackAppendHandler == nil {
		api.TeamTeamPackAppendHandler = team.TeamPackAppendHandlerFunc(func(params team.TeamPackAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamPackAppend has not yet been implemented")
		})
	}
	if api.TeamTeamPackDeleteHandler == nil {
		api.TeamTeamPackDeleteHandler = team.TeamPackDeleteHandlerFunc(func(params team.TeamPackDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamPackDelete has not yet been implemented")
		})
	}
	if api.TeamTeamPackIndexHandler == nil {
		api.TeamTeamPackIndexHandler = team.TeamPackIndexHandlerFunc(func(params team.TeamPackIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamPackIndex has not yet been implemented")
		})
	}
	if api.TeamTeamPackPermHandler == nil {
		api.TeamTeamPackPermHandler = team.TeamPackPermHandlerFunc(func(params team.TeamPackPermParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamPackPerm has not yet been implemented")
		})
	}
	if api.TeamTeamShowHandler == nil {
		api.TeamTeamShowHandler = team.TeamShowHandlerFunc(func(params team.TeamShowParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamShow has not yet been implemented")
		})
	}
	if api.TeamTeamUpdateHandler == nil {
		api.TeamTeamUpdateHandler = team.TeamUpdateHandlerFunc(func(params team.TeamUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUpdate has not yet been implemented")
		})
	}
	if api.TeamTeamUserAppendHandler == nil {
		api.TeamTeamUserAppendHandler = team.TeamUserAppendHandlerFunc(func(params team.TeamUserAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserAppend has not yet been implemented")
		})
	}
	if api.TeamTeamUserDeleteHandler == nil {
		api.TeamTeamUserDeleteHandler = team.TeamUserDeleteHandlerFunc(func(params team.TeamUserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserDelete has not yet been implemented")
		})
	}
	if api.TeamTeamUserIndexHandler == nil {
		api.TeamTeamUserIndexHandler = team.TeamUserIndexHandlerFunc(func(params team.TeamUserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserIndex has not yet been implemented")
		})
	}
	if api.TeamTeamUserPermHandler == nil {
		api.TeamTeamUserPermHandler = team.TeamUserPermHandlerFunc(func(params team.TeamUserPermParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserPerm has not yet been implemented")
		})
	}
	if api.UserUserCreateHandler == nil {
		api.UserUserCreateHandler = user.UserCreateHandlerFunc(func(params user.UserCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserCreate has not yet been implemented")
		})
	}
	if api.UserUserDeleteHandler == nil {
		api.UserUserDeleteHandler = user.UserDeleteHandlerFunc(func(params user.UserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserDelete has not yet been implemented")
		})
	}
	if api.UserUserIndexHandler == nil {
		api.UserUserIndexHandler = user.UserIndexHandlerFunc(func(params user.UserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserIndex has not yet been implemented")
		})
	}
	if api.UserUserModAppendHandler == nil {
		api.UserUserModAppendHandler = user.UserModAppendHandlerFunc(func(params user.UserModAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserModAppend has not yet been implemented")
		})
	}
	if api.UserUserModDeleteHandler == nil {
		api.UserUserModDeleteHandler = user.UserModDeleteHandlerFunc(func(params user.UserModDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserModDelete has not yet been implemented")
		})
	}
	if api.UserUserModIndexHandler == nil {
		api.UserUserModIndexHandler = user.UserModIndexHandlerFunc(func(params user.UserModIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserModIndex has not yet been implemented")
		})
	}
	if api.UserUserModPermHandler == nil {
		api.UserUserModPermHandler = user.UserModPermHandlerFunc(func(params user.UserModPermParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserModPerm has not yet been implemented")
		})
	}
	if api.UserUserPackAppendHandler == nil {
		api.UserUserPackAppendHandler = user.UserPackAppendHandlerFunc(func(params user.UserPackAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPackAppend has not yet been implemented")
		})
	}
	if api.UserUserPackDeleteHandler == nil {
		api.UserUserPackDeleteHandler = user.UserPackDeleteHandlerFunc(func(params user.UserPackDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPackDelete has not yet been implemented")
		})
	}
	if api.UserUserPackIndexHandler == nil {
		api.UserUserPackIndexHandler = user.UserPackIndexHandlerFunc(func(params user.UserPackIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPackIndex has not yet been implemented")
		})
	}
	if api.UserUserPackPermHandler == nil {
		api.UserUserPackPermHandler = user.UserPackPermHandlerFunc(func(params user.UserPackPermParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPackPerm has not yet been implemented")
		})
	}
	if api.UserUserShowHandler == nil {
		api.UserUserShowHandler = user.UserShowHandlerFunc(func(params user.UserShowParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserShow has not yet been implemented")
		})
	}
	if api.UserUserTeamAppendHandler == nil {
		api.UserUserTeamAppendHandler = user.UserTeamAppendHandlerFunc(func(params user.UserTeamAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamAppend has not yet been implemented")
		})
	}
	if api.UserUserTeamDeleteHandler == nil {
		api.UserUserTeamDeleteHandler = user.UserTeamDeleteHandlerFunc(func(params user.UserTeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamDelete has not yet been implemented")
		})
	}
	if api.UserUserTeamIndexHandler == nil {
		api.UserUserTeamIndexHandler = user.UserTeamIndexHandlerFunc(func(params user.UserTeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamIndex has not yet been implemented")
		})
	}
	if api.UserUserTeamPermHandler == nil {
		api.UserUserTeamPermHandler = user.UserTeamPermHandlerFunc(func(params user.UserTeamPermParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamPerm has not yet been implemented")
		})
	}
	if api.UserUserUpdateHandler == nil {
		api.UserUserUpdateHandler = user.UserUpdateHandlerFunc(func(params user.UserUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserUpdate has not yet been implemented")
		})
	}
	if api.ModVersionBuildAppendHandler == nil {
		api.ModVersionBuildAppendHandler = mod.VersionBuildAppendHandlerFunc(func(params mod.VersionBuildAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionBuildAppend has not yet been implemented")
		})
	}
	if api.ModVersionBuildDeleteHandler == nil {
		api.ModVersionBuildDeleteHandler = mod.VersionBuildDeleteHandlerFunc(func(params mod.VersionBuildDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionBuildDelete has not yet been implemented")
		})
	}
	if api.ModVersionBuildIndexHandler == nil {
		api.ModVersionBuildIndexHandler = mod.VersionBuildIndexHandlerFunc(func(params mod.VersionBuildIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionBuildIndex has not yet been implemented")
		})
	}
	if api.ModVersionCreateHandler == nil {
		api.ModVersionCreateHandler = mod.VersionCreateHandlerFunc(func(params mod.VersionCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionCreate has not yet been implemented")
		})
	}
	if api.ModVersionDeleteHandler == nil {
		api.ModVersionDeleteHandler = mod.VersionDeleteHandlerFunc(func(params mod.VersionDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionDelete has not yet been implemented")
		})
	}
	if api.ModVersionIndexHandler == nil {
		api.ModVersionIndexHandler = mod.VersionIndexHandlerFunc(func(params mod.VersionIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionIndex has not yet been implemented")
		})
	}
	if api.ModVersionShowHandler == nil {
		api.ModVersionShowHandler = mod.VersionShowHandlerFunc(func(params mod.VersionShowParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionShow has not yet been implemented")
		})
	}
	if api.ModVersionUpdateHandler == nil {
		api.ModVersionUpdateHandler = mod.VersionUpdateHandlerFunc(func(params mod.VersionUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.VersionUpdate has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
