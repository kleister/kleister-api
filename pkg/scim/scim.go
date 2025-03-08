package scim

import (
	"errors"
	"net/http"

	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/optional"
	"github.com/elimity-com/scim/schema"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

var (
	// ErrInvalidToken is returned when the request token is invalid.
	ErrInvalidToken = errors.New("invalid or missing token")
)

// New initializes the SCIM v2 server handlers.
func New(opts ...Option) *Scim {
	options := newOptions(opts...)

	return &Scim{
		root:   options.Root,
		config: options.Config,
		store:  options.Store,
		logger: log.With().Str("service", "scim").Logger(),
	}
}

// Scim defines the handlers for the SCIM v2 server.
type Scim struct {
	root   string
	config config.Scim
	store  *bun.DB
	logger zerolog.Logger
}

// Server returns the server which can be mounted into an existing mux.
func (s *Scim) Server() (http.HandlerFunc, error) {
	server, err := scim.NewServer(
		&scim.ServerArgs{
			ServiceProviderConfig: &scim.ServiceProviderConfig{
				DocumentationURI: optional.NewString("https://kleister.eu/usage/scim"),
				SupportPatch:     true,
			},
			ResourceTypes: []scim.ResourceType{
				{
					ID:          optional.NewString("User"),
					Name:        "User Account",
					Endpoint:    "/Users",
					Description: optional.NewString("User"),
					Schema: schema.Schema{
						ID:          "urn:ietf:params:scim:schemas:core:2.0:User",
						Name:        optional.NewString("User"),
						Description: optional.NewString("User Account"),
						Attributes: []schema.CoreAttribute{
							schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
								Name:       "userName",
								Required:   true,
								Uniqueness: schema.AttributeUniquenessServer(),
							})),
							schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
								Name: "displayName",
							})),
							schema.SimpleCoreAttribute(schema.SimpleBooleanParams(schema.BooleanParams{
								Name: "active",
							})),
							schema.ComplexCoreAttribute(schema.ComplexParams{
								MultiValued: true,
								Name:        "emails",
								SubAttributes: []schema.SimpleParams{
									schema.SimpleStringParams(schema.StringParams{
										Name: "value",
									}),
									schema.SimpleStringParams(schema.StringParams{
										Name: "display",
									}),
									schema.SimpleStringParams(schema.StringParams{
										CanonicalValues: []string{"work", "home", "other"},
										Name:            "type",
									}),
									schema.SimpleBooleanParams(schema.BooleanParams{
										Name: "primary",
									}),
								},
							}),
						},
					},
					Handler: s.usersHandler(),
				},
				{
					ID:          optional.NewString("Group"),
					Name:        "Group",
					Endpoint:    "/Groups",
					Description: optional.NewString("Group"),
					Schema: schema.Schema{
						ID:          "urn:ietf:params:scim:schemas:core:2.0:Group",
						Name:        optional.NewString("Group"),
						Description: optional.NewString("Group"),
						Attributes: []schema.CoreAttribute{
							schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
								Name:     "displayName",
								Required: true,
							})),
							schema.ComplexCoreAttribute(schema.ComplexParams{
								MultiValued: true,
								Name:        "members",
								SubAttributes: []schema.SimpleParams{
									schema.SimpleStringParams(schema.StringParams{
										Name:       "display",
										Mutability: schema.AttributeMutabilityImmutable(),
									}),
									schema.SimpleStringParams(schema.StringParams{
										Name:       "value",
										Mutability: schema.AttributeMutabilityImmutable(),
									}),
									schema.SimpleReferenceParams(schema.ReferenceParams{
										Name:       "$ref",
										Mutability: schema.AttributeMutabilityImmutable(),
										ReferenceTypes: []schema.AttributeReferenceType{
											"User",
											"Group",
										},
									}),
									schema.SimpleStringParams(schema.StringParams{
										Name:       "type",
										Mutability: schema.AttributeMutabilityImmutable(),
										CanonicalValues: []string{
											"User",
											"Group",
										},
									}),
								},
							}),
						},
					},
					Handler: s.groupsHandler(),
				},
			},
		},
		scim.WithLogger(nil),
	)

	if err != nil {
		return nil, err
	}

	h := http.StripPrefix(
		s.root,
		server,
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if s.config.Token == "" {
			h.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		if header != "Bearer "+s.config.Token {
			http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}, nil
}

func (s *Scim) usersHandler() *userHandlers {
	return &userHandlers{
		config: s.config,
		store:  s.store,
		logger: s.logger.With().Str("type", "users").Logger(),
	}
}

func (s *Scim) groupsHandler() *groupHandlers {
	return &groupHandlers{
		config: s.config,
		store:  s.store,
		logger: s.logger.With().Str("type", "groups").Logger(),
	}
}
