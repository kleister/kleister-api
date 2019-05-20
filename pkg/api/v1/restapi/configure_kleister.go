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

	if api.PackAppendBuildToVersionHandler == nil {
		api.PackAppendBuildToVersionHandler = pack.AppendBuildToVersionHandlerFunc(func(params pack.AppendBuildToVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.AppendBuildToVersion has not yet been implemented")
		})
	}
	if api.ForgeAppendForgeToBuildHandler == nil {
		api.ForgeAppendForgeToBuildHandler = forge.AppendForgeToBuildHandlerFunc(func(params forge.AppendForgeToBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.AppendForgeToBuild has not yet been implemented")
		})
	}
	if api.MinecraftAppendMinecraftToBuildHandler == nil {
		api.MinecraftAppendMinecraftToBuildHandler = minecraft.AppendMinecraftToBuildHandlerFunc(func(params minecraft.AppendMinecraftToBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.AppendMinecraftToBuild has not yet been implemented")
		})
	}
	if api.ModAppendModToTeamHandler == nil {
		api.ModAppendModToTeamHandler = mod.AppendModToTeamHandlerFunc(func(params mod.AppendModToTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.AppendModToTeam has not yet been implemented")
		})
	}
	if api.ModAppendModToUserHandler == nil {
		api.ModAppendModToUserHandler = mod.AppendModToUserHandlerFunc(func(params mod.AppendModToUserParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.AppendModToUser has not yet been implemented")
		})
	}
	if api.PackAppendPackToTeamHandler == nil {
		api.PackAppendPackToTeamHandler = pack.AppendPackToTeamHandlerFunc(func(params pack.AppendPackToTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.AppendPackToTeam has not yet been implemented")
		})
	}
	if api.PackAppendPackToUserHandler == nil {
		api.PackAppendPackToUserHandler = pack.AppendPackToUserHandlerFunc(func(params pack.AppendPackToUserParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.AppendPackToUser has not yet been implemented")
		})
	}
	if api.TeamAppendTeamToModHandler == nil {
		api.TeamAppendTeamToModHandler = team.AppendTeamToModHandlerFunc(func(params team.AppendTeamToModParams) middleware.Responder {
			return middleware.NotImplemented("operation team.AppendTeamToMod has not yet been implemented")
		})
	}
	if api.TeamAppendTeamToPackHandler == nil {
		api.TeamAppendTeamToPackHandler = team.AppendTeamToPackHandlerFunc(func(params team.AppendTeamToPackParams) middleware.Responder {
			return middleware.NotImplemented("operation team.AppendTeamToPack has not yet been implemented")
		})
	}
	if api.TeamAppendTeamToUserHandler == nil {
		api.TeamAppendTeamToUserHandler = team.AppendTeamToUserHandlerFunc(func(params team.AppendTeamToUserParams) middleware.Responder {
			return middleware.NotImplemented("operation team.AppendTeamToUser has not yet been implemented")
		})
	}
	if api.UserAppendUserToModHandler == nil {
		api.UserAppendUserToModHandler = user.AppendUserToModHandlerFunc(func(params user.AppendUserToModParams) middleware.Responder {
			return middleware.NotImplemented("operation user.AppendUserToMod has not yet been implemented")
		})
	}
	if api.UserAppendUserToPackHandler == nil {
		api.UserAppendUserToPackHandler = user.AppendUserToPackHandlerFunc(func(params user.AppendUserToPackParams) middleware.Responder {
			return middleware.NotImplemented("operation user.AppendUserToPack has not yet been implemented")
		})
	}
	if api.UserAppendUserToTeamHandler == nil {
		api.UserAppendUserToTeamHandler = user.AppendUserToTeamHandlerFunc(func(params user.AppendUserToTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation user.AppendUserToTeam has not yet been implemented")
		})
	}
	if api.ModAppendVersionToBuildHandler == nil {
		api.ModAppendVersionToBuildHandler = mod.AppendVersionToBuildHandlerFunc(func(params mod.AppendVersionToBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.AppendVersionToBuild has not yet been implemented")
		})
	}
	if api.PackCreateBuildHandler == nil {
		api.PackCreateBuildHandler = pack.CreateBuildHandlerFunc(func(params pack.CreateBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.CreateBuild has not yet been implemented")
		})
	}
	if api.ModCreateModHandler == nil {
		api.ModCreateModHandler = mod.CreateModHandlerFunc(func(params mod.CreateModParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.CreateMod has not yet been implemented")
		})
	}
	if api.PackCreatePackHandler == nil {
		api.PackCreatePackHandler = pack.CreatePackHandlerFunc(func(params pack.CreatePackParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.CreatePack has not yet been implemented")
		})
	}
	if api.TeamCreateTeamHandler == nil {
		api.TeamCreateTeamHandler = team.CreateTeamHandlerFunc(func(params team.CreateTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation team.CreateTeam has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.ModCreateVersionHandler == nil {
		api.ModCreateVersionHandler = mod.CreateVersionHandlerFunc(func(params mod.CreateVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.CreateVersion has not yet been implemented")
		})
	}
	if api.PackDeleteBuildHandler == nil {
		api.PackDeleteBuildHandler = pack.DeleteBuildHandlerFunc(func(params pack.DeleteBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.DeleteBuild has not yet been implemented")
		})
	}
	if api.PackDeleteBuildFromVersionHandler == nil {
		api.PackDeleteBuildFromVersionHandler = pack.DeleteBuildFromVersionHandlerFunc(func(params pack.DeleteBuildFromVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.DeleteBuildFromVersion has not yet been implemented")
		})
	}
	if api.ForgeDeleteForgeFromBuildHandler == nil {
		api.ForgeDeleteForgeFromBuildHandler = forge.DeleteForgeFromBuildHandlerFunc(func(params forge.DeleteForgeFromBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.DeleteForgeFromBuild has not yet been implemented")
		})
	}
	if api.MinecraftDeleteMinecraftFromBuildHandler == nil {
		api.MinecraftDeleteMinecraftFromBuildHandler = minecraft.DeleteMinecraftFromBuildHandlerFunc(func(params minecraft.DeleteMinecraftFromBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.DeleteMinecraftFromBuild has not yet been implemented")
		})
	}
	if api.ModDeleteModHandler == nil {
		api.ModDeleteModHandler = mod.DeleteModHandlerFunc(func(params mod.DeleteModParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.DeleteMod has not yet been implemented")
		})
	}
	if api.ModDeleteModFromTeamHandler == nil {
		api.ModDeleteModFromTeamHandler = mod.DeleteModFromTeamHandlerFunc(func(params mod.DeleteModFromTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.DeleteModFromTeam has not yet been implemented")
		})
	}
	if api.ModDeleteModFromUserHandler == nil {
		api.ModDeleteModFromUserHandler = mod.DeleteModFromUserHandlerFunc(func(params mod.DeleteModFromUserParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.DeleteModFromUser has not yet been implemented")
		})
	}
	if api.PackDeletePackHandler == nil {
		api.PackDeletePackHandler = pack.DeletePackHandlerFunc(func(params pack.DeletePackParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.DeletePack has not yet been implemented")
		})
	}
	if api.PackDeletePackFromTeamHandler == nil {
		api.PackDeletePackFromTeamHandler = pack.DeletePackFromTeamHandlerFunc(func(params pack.DeletePackFromTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.DeletePackFromTeam has not yet been implemented")
		})
	}
	if api.PackDeletePackFromUserHandler == nil {
		api.PackDeletePackFromUserHandler = pack.DeletePackFromUserHandlerFunc(func(params pack.DeletePackFromUserParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.DeletePackFromUser has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamHandler == nil {
		api.TeamDeleteTeamHandler = team.DeleteTeamHandlerFunc(func(params team.DeleteTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeam has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamFromModHandler == nil {
		api.TeamDeleteTeamFromModHandler = team.DeleteTeamFromModHandlerFunc(func(params team.DeleteTeamFromModParams) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeamFromMod has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamFromPackHandler == nil {
		api.TeamDeleteTeamFromPackHandler = team.DeleteTeamFromPackHandlerFunc(func(params team.DeleteTeamFromPackParams) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeamFromPack has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamFromUserHandler == nil {
		api.TeamDeleteTeamFromUserHandler = team.DeleteTeamFromUserHandlerFunc(func(params team.DeleteTeamFromUserParams) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeamFromUser has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.UserDeleteUserFromModHandler == nil {
		api.UserDeleteUserFromModHandler = user.DeleteUserFromModHandlerFunc(func(params user.DeleteUserFromModParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserFromMod has not yet been implemented")
		})
	}
	if api.UserDeleteUserFromPackHandler == nil {
		api.UserDeleteUserFromPackHandler = user.DeleteUserFromPackHandlerFunc(func(params user.DeleteUserFromPackParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserFromPack has not yet been implemented")
		})
	}
	if api.UserDeleteUserFromTeamHandler == nil {
		api.UserDeleteUserFromTeamHandler = user.DeleteUserFromTeamHandlerFunc(func(params user.DeleteUserFromTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserFromTeam has not yet been implemented")
		})
	}
	if api.ModDeleteVersionHandler == nil {
		api.ModDeleteVersionHandler = mod.DeleteVersionHandlerFunc(func(params mod.DeleteVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.DeleteVersion has not yet been implemented")
		})
	}
	if api.ModDeleteVersionFromBuildHandler == nil {
		api.ModDeleteVersionFromBuildHandler = mod.DeleteVersionFromBuildHandlerFunc(func(params mod.DeleteVersionFromBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.DeleteVersionFromBuild has not yet been implemented")
		})
	}
	if api.PackListBuildVersionsHandler == nil {
		api.PackListBuildVersionsHandler = pack.ListBuildVersionsHandlerFunc(func(params pack.ListBuildVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ListBuildVersions has not yet been implemented")
		})
	}
	if api.PackListBuildsHandler == nil {
		api.PackListBuildsHandler = pack.ListBuildsHandlerFunc(func(params pack.ListBuildsParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ListBuilds has not yet been implemented")
		})
	}
	if api.ForgeListForgeBuildsHandler == nil {
		api.ForgeListForgeBuildsHandler = forge.ListForgeBuildsHandlerFunc(func(params forge.ListForgeBuildsParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ListForgeBuilds has not yet been implemented")
		})
	}
	if api.ForgeListForgesHandler == nil {
		api.ForgeListForgesHandler = forge.ListForgesHandlerFunc(func(params forge.ListForgesParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.ListForges has not yet been implemented")
		})
	}
	if api.MinecraftListMinecraftBuildsHandler == nil {
		api.MinecraftListMinecraftBuildsHandler = minecraft.ListMinecraftBuildsHandlerFunc(func(params minecraft.ListMinecraftBuildsParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.ListMinecraftBuilds has not yet been implemented")
		})
	}
	if api.MinecraftListMinecraftsHandler == nil {
		api.MinecraftListMinecraftsHandler = minecraft.ListMinecraftsHandlerFunc(func(params minecraft.ListMinecraftsParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.ListMinecrafts has not yet been implemented")
		})
	}
	if api.ModListModTeamsHandler == nil {
		api.ModListModTeamsHandler = mod.ListModTeamsHandlerFunc(func(params mod.ListModTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListModTeams has not yet been implemented")
		})
	}
	if api.ModListModUsersHandler == nil {
		api.ModListModUsersHandler = mod.ListModUsersHandlerFunc(func(params mod.ListModUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListModUsers has not yet been implemented")
		})
	}
	if api.ModListModsHandler == nil {
		api.ModListModsHandler = mod.ListModsHandlerFunc(func(params mod.ListModsParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListMods has not yet been implemented")
		})
	}
	if api.PackListPackTeamsHandler == nil {
		api.PackListPackTeamsHandler = pack.ListPackTeamsHandlerFunc(func(params pack.ListPackTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ListPackTeams has not yet been implemented")
		})
	}
	if api.PackListPackUsersHandler == nil {
		api.PackListPackUsersHandler = pack.ListPackUsersHandlerFunc(func(params pack.ListPackUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ListPackUsers has not yet been implemented")
		})
	}
	if api.PackListPacksHandler == nil {
		api.PackListPacksHandler = pack.ListPacksHandlerFunc(func(params pack.ListPacksParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ListPacks has not yet been implemented")
		})
	}
	if api.TeamListTeamModsHandler == nil {
		api.TeamListTeamModsHandler = team.ListTeamModsHandlerFunc(func(params team.ListTeamModsParams) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeamMods has not yet been implemented")
		})
	}
	if api.TeamListTeamPacksHandler == nil {
		api.TeamListTeamPacksHandler = team.ListTeamPacksHandlerFunc(func(params team.ListTeamPacksParams) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeamPacks has not yet been implemented")
		})
	}
	if api.TeamListTeamUsersHandler == nil {
		api.TeamListTeamUsersHandler = team.ListTeamUsersHandlerFunc(func(params team.ListTeamUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeamUsers has not yet been implemented")
		})
	}
	if api.TeamListTeamsHandler == nil {
		api.TeamListTeamsHandler = team.ListTeamsHandlerFunc(func(params team.ListTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeams has not yet been implemented")
		})
	}
	if api.UserListUserModsHandler == nil {
		api.UserListUserModsHandler = user.ListUserModsHandlerFunc(func(params user.ListUserModsParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserMods has not yet been implemented")
		})
	}
	if api.UserListUserPacksHandler == nil {
		api.UserListUserPacksHandler = user.ListUserPacksHandlerFunc(func(params user.ListUserPacksParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserPacks has not yet been implemented")
		})
	}
	if api.UserListUserTeamsHandler == nil {
		api.UserListUserTeamsHandler = user.ListUserTeamsHandlerFunc(func(params user.ListUserTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserTeams has not yet been implemented")
		})
	}
	if api.UserListUsersHandler == nil {
		api.UserListUsersHandler = user.ListUsersHandlerFunc(func(params user.ListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUsers has not yet been implemented")
		})
	}
	if api.ModListVersionBuildsHandler == nil {
		api.ModListVersionBuildsHandler = mod.ListVersionBuildsHandlerFunc(func(params mod.ListVersionBuildsParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListVersionBuilds has not yet been implemented")
		})
	}
	if api.ModListVersionsHandler == nil {
		api.ModListVersionsHandler = mod.ListVersionsHandlerFunc(func(params mod.ListVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListVersions has not yet been implemented")
		})
	}
	if api.AuthLoginUserHandler == nil {
		api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginUser has not yet been implemented")
		})
	}
	if api.ModPermitModTeamHandler == nil {
		api.ModPermitModTeamHandler = mod.PermitModTeamHandlerFunc(func(params mod.PermitModTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.PermitModTeam has not yet been implemented")
		})
	}
	if api.ModPermitModUserHandler == nil {
		api.ModPermitModUserHandler = mod.PermitModUserHandlerFunc(func(params mod.PermitModUserParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.PermitModUser has not yet been implemented")
		})
	}
	if api.PackPermitPackTeamHandler == nil {
		api.PackPermitPackTeamHandler = pack.PermitPackTeamHandlerFunc(func(params pack.PermitPackTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PermitPackTeam has not yet been implemented")
		})
	}
	if api.PackPermitPackUserHandler == nil {
		api.PackPermitPackUserHandler = pack.PermitPackUserHandlerFunc(func(params pack.PermitPackUserParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.PermitPackUser has not yet been implemented")
		})
	}
	if api.TeamPermitTeamModHandler == nil {
		api.TeamPermitTeamModHandler = team.PermitTeamModHandlerFunc(func(params team.PermitTeamModParams) middleware.Responder {
			return middleware.NotImplemented("operation team.PermitTeamMod has not yet been implemented")
		})
	}
	if api.TeamPermitTeamPackHandler == nil {
		api.TeamPermitTeamPackHandler = team.PermitTeamPackHandlerFunc(func(params team.PermitTeamPackParams) middleware.Responder {
			return middleware.NotImplemented("operation team.PermitTeamPack has not yet been implemented")
		})
	}
	if api.TeamPermitTeamUserHandler == nil {
		api.TeamPermitTeamUserHandler = team.PermitTeamUserHandlerFunc(func(params team.PermitTeamUserParams) middleware.Responder {
			return middleware.NotImplemented("operation team.PermitTeamUser has not yet been implemented")
		})
	}
	if api.UserPermitUserModHandler == nil {
		api.UserPermitUserModHandler = user.PermitUserModHandlerFunc(func(params user.PermitUserModParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PermitUserMod has not yet been implemented")
		})
	}
	if api.UserPermitUserPackHandler == nil {
		api.UserPermitUserPackHandler = user.PermitUserPackHandlerFunc(func(params user.PermitUserPackParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PermitUserPack has not yet been implemented")
		})
	}
	if api.UserPermitUserTeamHandler == nil {
		api.UserPermitUserTeamHandler = user.PermitUserTeamHandlerFunc(func(params user.PermitUserTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PermitUserTeam has not yet been implemented")
		})
	}
	if api.AuthRefreshAuthHandler == nil {
		api.AuthRefreshAuthHandler = auth.RefreshAuthHandlerFunc(func(params auth.RefreshAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RefreshAuth has not yet been implemented")
		})
	}
	if api.ForgeSearchForgesHandler == nil {
		api.ForgeSearchForgesHandler = forge.SearchForgesHandlerFunc(func(params forge.SearchForgesParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.SearchForges has not yet been implemented")
		})
	}
	if api.MinecraftSearchMinecraftsHandler == nil {
		api.MinecraftSearchMinecraftsHandler = minecraft.SearchMinecraftsHandlerFunc(func(params minecraft.SearchMinecraftsParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.SearchMinecrafts has not yet been implemented")
		})
	}
	if api.PackShowBuildHandler == nil {
		api.PackShowBuildHandler = pack.ShowBuildHandlerFunc(func(params pack.ShowBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ShowBuild has not yet been implemented")
		})
	}
	if api.ModShowModHandler == nil {
		api.ModShowModHandler = mod.ShowModHandlerFunc(func(params mod.ShowModParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ShowMod has not yet been implemented")
		})
	}
	if api.PackShowPackHandler == nil {
		api.PackShowPackHandler = pack.ShowPackHandlerFunc(func(params pack.ShowPackParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.ShowPack has not yet been implemented")
		})
	}
	if api.ProfileShowProfileHandler == nil {
		api.ProfileShowProfileHandler = profile.ShowProfileHandlerFunc(func(params profile.ShowProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ShowProfile has not yet been implemented")
		})
	}
	if api.TeamShowTeamHandler == nil {
		api.TeamShowTeamHandler = team.ShowTeamHandlerFunc(func(params team.ShowTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation team.ShowTeam has not yet been implemented")
		})
	}
	if api.UserShowUserHandler == nil {
		api.UserShowUserHandler = user.ShowUserHandlerFunc(func(params user.ShowUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ShowUser has not yet been implemented")
		})
	}
	if api.ModShowVersionHandler == nil {
		api.ModShowVersionHandler = mod.ShowVersionHandlerFunc(func(params mod.ShowVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ShowVersion has not yet been implemented")
		})
	}
	if api.ProfileTokenProfileHandler == nil {
		api.ProfileTokenProfileHandler = profile.TokenProfileHandlerFunc(func(params profile.TokenProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.TokenProfile has not yet been implemented")
		})
	}
	if api.PackUpdateBuildHandler == nil {
		api.PackUpdateBuildHandler = pack.UpdateBuildHandlerFunc(func(params pack.UpdateBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.UpdateBuild has not yet been implemented")
		})
	}
	if api.ForgeUpdateForgeHandler == nil {
		api.ForgeUpdateForgeHandler = forge.UpdateForgeHandlerFunc(func(params forge.UpdateForgeParams) middleware.Responder {
			return middleware.NotImplemented("operation forge.UpdateForge has not yet been implemented")
		})
	}
	if api.MinecraftUpdateMinecraftHandler == nil {
		api.MinecraftUpdateMinecraftHandler = minecraft.UpdateMinecraftHandlerFunc(func(params minecraft.UpdateMinecraftParams) middleware.Responder {
			return middleware.NotImplemented("operation minecraft.UpdateMinecraft has not yet been implemented")
		})
	}
	if api.ModUpdateModHandler == nil {
		api.ModUpdateModHandler = mod.UpdateModHandlerFunc(func(params mod.UpdateModParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.UpdateMod has not yet been implemented")
		})
	}
	if api.PackUpdatePackHandler == nil {
		api.PackUpdatePackHandler = pack.UpdatePackHandlerFunc(func(params pack.UpdatePackParams) middleware.Responder {
			return middleware.NotImplemented("operation pack.UpdatePack has not yet been implemented")
		})
	}
	if api.ProfileUpdateProfileHandler == nil {
		api.ProfileUpdateProfileHandler = profile.UpdateProfileHandlerFunc(func(params profile.UpdateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.UpdateProfile has not yet been implemented")
		})
	}
	if api.TeamUpdateTeamHandler == nil {
		api.TeamUpdateTeamHandler = team.UpdateTeamHandlerFunc(func(params team.UpdateTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation team.UpdateTeam has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}
	if api.ModUpdateVersionHandler == nil {
		api.ModUpdateVersionHandler = mod.UpdateVersionHandlerFunc(func(params mod.UpdateVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.UpdateVersion has not yet been implemented")
		})
	}
	if api.AuthVerifyAuthHandler == nil {
		api.AuthVerifyAuthHandler = auth.VerifyAuthHandlerFunc(func(params auth.VerifyAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.VerifyAuth has not yet been implemented")
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
