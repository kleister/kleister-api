package v1

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	buildversions "github.com/kleister/kleister-api/pkg/service/build_versions"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/fabric"
	"github.com/kleister/kleister-api/pkg/service/forge"
	"github.com/kleister/kleister-api/pkg/service/minecraft"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/neoforge"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/quilt"
	teammods "github.com/kleister/kleister-api/pkg/service/team_mods"
	teampacks "github.com/kleister/kleister-api/pkg/service/team_packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	usermods "github.com/kleister/kleister-api/pkg/service/user_mods"
	userpacks "github.com/kleister/kleister-api/pkg/service/user_packs"
	userteams "github.com/kleister/kleister-api/pkg/service/user_teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/session"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/upload"
)

var (
	_ StrictServerInterface = (*API)(nil)
)

// New creates a new API that adds the handler implementations.
func New(
	cfg *config.Config,
	registry *metrics.Metrics,
	sess *session.Session,
	uploads upload.Upload,
	storage store.Store,
	minecraftService minecraft.Service,
	forgeService forge.Service,
	neoforgeService neoforge.Service,
	quiltService quilt.Service,
	fabricService fabric.Service,
	teamsService teams.Service,
	usersService users.Service,
	userteamsService userteams.Service,
	modsService mods.Service,
	usermodsService usermods.Service,
	teammodsService teammods.Service,
	versionsService versions.Service,
	packsService packs.Service,
	userpacksService userpacks.Service,
	teampacksService teampacks.Service,
	buildsService builds.Service,
	buildversionsService buildversions.Service,
) *API {
	return &API{
		config:        cfg,
		registry:      registry,
		session:       sess,
		uploads:       uploads,
		storage:       storage,
		minecraft:     minecraftService,
		forge:         forgeService,
		neoforge:      neoforgeService,
		quilt:         quiltService,
		fabric:        fabricService,
		teams:         teamsService,
		users:         usersService,
		userteams:     userteamsService,
		mods:          modsService,
		usermods:      usermodsService,
		teammods:      teammodsService,
		versions:      versionsService,
		packs:         packsService,
		userpacks:     userpacksService,
		teampacks:     teampacksService,
		builds:        buildsService,
		buildversions: buildversionsService,
	}
}

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	config        *config.Config
	registry      *metrics.Metrics
	session       *session.Session
	uploads       upload.Upload
	storage       store.Store
	minecraft     minecraft.Service
	forge         forge.Service
	neoforge      neoforge.Service
	quilt         quilt.Service
	fabric        fabric.Service
	teams         teams.Service
	users         users.Service
	userteams     userteams.Service
	mods          mods.Service
	usermods      usermods.Service
	teammods      teammods.Service
	versions      versions.Service
	packs         packs.Service
	userpacks     userpacks.Service
	teampacks     teampacks.Service
	builds        builds.Service
	buildversions buildversions.Service
}
