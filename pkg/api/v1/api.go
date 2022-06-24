package v1

import (
	"context"
	"net/http"
	"path"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/token"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/log"
)

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	Handler http.Handler
}

// New creates a new API that adds the custom Handler implementations.
func New(
	cfg *config.Config,
	uploads upload.Upload,
	usersService users.Service,
	teamsService teams.Service,
	packsService packs.Service,
	buildsService builds.Service,
	modsService mods.Service,
	versionsService versions.Service,
	minecraftService minecraft.Service,
	forgeService forge.Service,
) *API {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to analyze openapi")

		return nil
	}

	spec.Spec().Host = cfg.Server.Host
	spec.Spec().BasePath = path.Join(
		cfg.Server.Root,
		spec.Spec().BasePath,
	)

	api := operations.NewKleisterAPI(spec)

	api.AuthLoginUserHandler = LoginUserHandler(cfg, usersService)
	api.AuthRefreshAuthHandler = RefreshAuthHandler(cfg)
	api.AuthVerifyAuthHandler = VerifyAuthHandler()

	api.ProfileTokenProfileHandler = TokenProfileHandler(cfg)
	api.ProfileUpdateProfileHandler = UpdateProfileHandler(usersService)
	api.ProfileShowProfileHandler = ShowProfileHandler(usersService)

	api.TeamListTeamsHandler = ListTeamsHandler(teamsService)
	api.TeamShowTeamHandler = ShowTeamHandler(teamsService)
	api.TeamCreateTeamHandler = CreateTeamHandler(teamsService)
	api.TeamUpdateTeamHandler = UpdateTeamHandler(teamsService)
	api.TeamDeleteTeamHandler = DeleteTeamHandler(teamsService)
	api.TeamListTeamUsersHandler = ListTeamUsersHandler(teamsService)
	api.TeamAppendTeamToUserHandler = AppendTeamToUserHandler(teamsService, usersService)
	api.TeamPermitTeamUserHandler = PermitTeamUserHandler(teamsService, usersService)
	api.TeamDeleteTeamFromUserHandler = DeleteTeamFromUserHandler(teamsService, usersService)
	api.TeamListTeamModsHandler = ListTeamModsHandler(teamsService)
	api.TeamAppendTeamToModHandler = AppendTeamToModHandler(teamsService, modsService)
	api.TeamPermitTeamModHandler = PermitTeamModHandler(teamsService, modsService)
	api.TeamDeleteTeamFromModHandler = DeleteTeamFromModHandler(teamsService, modsService)
	api.TeamListTeamPacksHandler = ListTeamPacksHandler(teamsService)
	api.TeamAppendTeamToPackHandler = AppendTeamToPackHandler(teamsService, packsService)
	api.TeamPermitTeamPackHandler = PermitTeamPackHandler(teamsService, packsService)
	api.TeamDeleteTeamFromPackHandler = DeleteTeamFromPackHandler(teamsService, packsService)

	api.UserListUsersHandler = ListUsersHandler(usersService)
	api.UserShowUserHandler = ShowUserHandler(usersService)
	api.UserCreateUserHandler = CreateUserHandler(usersService)
	api.UserUpdateUserHandler = UpdateUserHandler(usersService)
	api.UserDeleteUserHandler = DeleteUserHandler(usersService)
	api.UserListUserTeamsHandler = ListUserTeamsHandler(usersService)
	api.UserAppendUserToTeamHandler = AppendUserToTeamHandler(usersService, teamsService)
	api.UserPermitUserTeamHandler = PermitUserTeamHandler(usersService, teamsService)
	api.UserDeleteUserFromTeamHandler = DeleteUserFromTeamHandler(usersService, teamsService)
	api.UserListUserModsHandler = ListUserModsHandler(usersService)
	api.UserAppendUserToModHandler = AppendUserToModHandler(usersService, modsService)
	api.UserPermitUserModHandler = PermitUserModHandler(usersService, modsService)
	api.UserDeleteUserFromModHandler = DeleteUserFromModHandler(usersService, modsService)
	api.UserListUserPacksHandler = ListUserPacksHandler(usersService)
	api.UserAppendUserToPackHandler = AppendUserToPackHandler(usersService, packsService)
	api.UserPermitUserPackHandler = PermitUserPackHandler(usersService, packsService)
	api.UserDeleteUserFromPackHandler = DeleteUserFromPackHandler(usersService, packsService)

	api.PackListPacksHandler = ListPacksHandler(packsService)
	api.PackShowPackHandler = ShowPackHandler(packsService)
	api.PackCreatePackHandler = CreatePackHandler(packsService)
	api.PackUpdatePackHandler = UpdatePackHandler(packsService)
	api.PackDeletePackHandler = DeletePackHandler(packsService)
	// api.PackListPackTeamsHandler = ListPackTeamsHandler(packsService)
	// api.PackAppendPackToTeamHandler = AppendPackToTeamHandler(packsService, teamsService)
	// api.PackPermitPackTeamHandler = PermitPackTeamHandler(packsService, teamsService)
	// api.PackDeletePackFromTeamHandler = DeletePackFromTeamHandler(packsService, teamsService)
	// api.PackListPackUsersHandler = ListPackUsersHandler(packsService)
	// api.PackAppendPackToUserHandler = AppendPackToUserHandler(packsService, usersService)
	// api.PackPermitPackUserHandler = PermitPackUserHandler(packsService, usersService)
	// api.PackDeletePackFromUserHandler = DeletePackFromUserHandler(packsService, usersService)

	api.PackListBuildsHandler = ListBuildsHandler(packsService, buildsService)
	api.PackShowBuildHandler = ShowBuildHandler(packsService, buildsService)
	api.PackCreateBuildHandler = CreateBuildHandler(packsService, buildsService)
	api.PackUpdateBuildHandler = UpdateBuildHandler(packsService, buildsService)
	api.PackDeleteBuildHandler = DeleteBuildHandler(packsService, buildsService)
	// api.PackListBuildVersionsHandler = ListBuildVersionsHandler(packsService, buildsService)
	// api.PackAppendBuildToVersionHandler = AppendBuildToVersionHandler(packsService, buildsService, versionsService)
	// api.PackDeleteBuildFromVersionHandler = DeleteBuildFromVersionHandler(packsService, buildsService, versionsService)

	api.ModListModsHandler = ListModsHandler(modsService)
	api.ModShowModHandler = ShowModHandler(modsService)
	api.ModCreateModHandler = CreateModHandler(modsService)
	api.ModUpdateModHandler = UpdateModHandler(modsService)
	api.ModDeleteModHandler = DeleteModHandler(modsService)
	// api.ModListModTeamsHandler = ListModTeamsHandler(modsService)
	// api.ModAppendModToTeamHandler = AppendModToTeamHandler(modsService, teamsService)
	// api.ModPermitModTeamHandler = PermitModTeamHandler(modsService, teamsService)
	// api.ModDeleteModFromTeamHandler = DeleteModFromTeamHandler(modsService, teamsService)
	// api.ModListModUsersHandler = ListModUsersHandler(modsService)
	// api.ModAppendModToUserHandler = AppendModToUserHandler(modsService, usersService)
	// api.ModPermitModUserHandler = PermitModUserHandler(modsService, usersService)
	// api.ModDeleteModFromUserHandler = DeleteModFromUserHandler(modsService, usersService)

	api.ModListVersionsHandler = ListVersionsHandler(modsService, versionsService)
	api.ModShowVersionHandler = ShowVersionHandler(modsService, versionsService)
	api.ModCreateVersionHandler = CreateVersionHandler(modsService, versionsService)
	api.ModUpdateVersionHandler = UpdateVersionHandler(modsService, versionsService)
	api.ModDeleteVersionHandler = DeleteVersionHandler(modsService, versionsService)
	// api.ModListVersionBuildsHandler = ListVersionBuildsHandler(modsService, versionsService)
	// api.ModAppendVersionToBuildHandler = AppendVersionToBuildHandler(modsService, versionsService, buildsService)
	// api.ModDeleteVersionFromBuildHandler = DeleteVersionFromBuildHandler(modsService, versionsService, buildsService)

	api.MinecraftListMinecraftsHandler = ListMinecraftsHandler(minecraftService)
	api.MinecraftUpdateMinecraftHandler = UpdateMinecraftHandler(minecraftService)
	// api.MinecraftSearchMinecraftsHandler = SearchMinecraftHandler(minecraftService)
	// api.MinecraftListMinecraftBuildsHandler = ListMinecraftBuilds(minecraftService)
	// api.MinecraftAppendMinecraftToBuildHandler = AppendMinecraftToBuildHandler(minecraftService, buildService)
	// api.MinecraftDeleteMinecraftFromBuildHandler = DeleteMinecraftFromBuildHandler(minecraftService, buildService)

	api.ForgeListForgesHandler = ListForgesHandler(forgeService)
	api.ForgeUpdateForgeHandler = UpdateForgeHandler(forgeService)
	// api.ForgeSearchForgesHandler = SearchForgeHandler(forgeService)
	// api.ForgeListForgeBuildsHandler = ListForgeBuildsHandler(forgeService)
	// api.ForgeAppendForgeToBuildHandler = AppendForgeToBuildHandler(forgeService, buildService)
	// api.ForgeDeleteForgeFromBuildHandler = DeleteForgeFromBuildHandler(forgeService, buildService)

	// TODO: needs context for request id logging?
	api.HeaderAuth = func(val string) (*models.User, error) {
		t, err := token.Parse(val, cfg.Session.Secret)

		if err != nil {
			log.Warn().
				Err(err).
				Str("token", val).
				Msg("failed to parse token")

			return nil, errors.New(401, "incorrect auth")
		}

		user, err := usersService.Show(
			context.Background(),
			t.Text,
		)

		if err != nil {
			log.Warn().
				Err(err).
				Str("token", val).
				Msg("failed to fetch user")

			return nil, errors.New(401, "incorrect auth")
		}

		return convertUser(user), nil
	}

	// TODO: needs context for request id logging?
	api.BasicAuth = func(username, password string) (*models.User, error) {
		user, err := usersService.ByBasicAuth(
			context.Background(),
			username,
			password,
		)

		if err != nil {
			log.Warn().
				Err(err).
				Str("username", username).
				Msg("failed to auth user")

			return nil, errors.New(401, "incorrect auth")
		}

		return convertUser(user), nil
	}

	return &API{
		Handler: api.Serve(nil),
	}
}
