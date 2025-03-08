package authn

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"golang.org/x/oauth2/microsoft"
)

var (
	// ErrMissingIssuerEndpoint defines the error if an issuer endpoint is missing.
	ErrMissingIssuerEndpoint = fmt.Errorf("missing issuer endpoint")

	// ErrMissingToken defines the error if a token is missing while exchange.
	ErrMissingToken = fmt.Errorf("missing token after exchange")
)

// Authn initializes the authn provider handling.
type Authn struct {
	config    *config.AuthConfig
	Providers map[string]*Provider
}

// New initializes the authn provider handling.
func New(opts ...Option) (*Authn, error) {
	options := newOptions(opts...)

	client := &Authn{
		config:    &config.AuthConfig{},
		Providers: make(map[string]*Provider, 0),
	}

	if options.Config != "" {
		cfg := viper.New()
		cfg.SetConfigFile(options.Config)

		if err := cfg.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				return nil, fmt.Errorf("failed to find auth config: %w", err)
			case viper.UnsupportedConfigError:
				return nil, fmt.Errorf("unsupport type for auth config: %w", err)
			default:
				return nil, fmt.Errorf("failed to read auth config: %w", err)
			}
		}

		if err := cfg.Unmarshal(client.config); err != nil {
			return nil, fmt.Errorf("failed to parse auth config: %w", err)
		}

		for _, provider := range client.config.Providers {
			switch provider.Driver {
			case "entraid":
				p, err := entraidProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			case "google":
				p, err := googleProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			case "github":
				p, err := githubProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			case "gitea":
				p, err := giteaProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			case "gitlab":
				p, err := gitlabProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			case "oidc":
				p, err := oidcProvider(provider)

				if err != nil {
					return nil, err
				}

				client.Providers[provider.Name] = p
			default:
				return nil, fmt.Errorf("unsupport auth provider: %s", provider.Driver)
			}
		}
	}

	return client, nil
}

func entraidProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "entraid").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	cfg.Endpoints.Profile = "https://graph.microsoft.com/v1.0/me"

	if cfg.Tenant == "" {
		cfg.Tenant = "common"
	}

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"openid",
			"profile",
			"email",
			"User.Read",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     microsoft.AzureADEndpoint(cfg.Tenant),
		RedirectURL:  cfg.Callback,
		Scopes:       cfg.Scopes,
	}

	return p, nil
}

func googleProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "google").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	cfg.Endpoints.Profile = "https://www.googleapis.com/oauth2/v2/userinfo"

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoints.Google,
		RedirectURL:  cfg.Callback,
		Scopes:       cfg.Scopes,
	}

	return p, nil
}

func githubProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "github").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	if cfg.Endpoints.Auth == "" {
		cfg.Endpoints.Auth = "https://github.com/login/oauth/authorize"
	}

	if cfg.Endpoints.Token == "" {
		cfg.Endpoints.Token = "https://github.com/login/oauth/access_token"
	}

	if cfg.Endpoints.Profile == "" {
		cfg.Endpoints.Profile = "https://api.github.com/user"
	}

	if cfg.Endpoints.Email == "" {
		cfg.Endpoints.Email = "https://api.github.com/user/emails"
	}

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"read:user",
			"user:email",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.Endpoints.Auth,
			TokenURL: cfg.Endpoints.Token,
		},
		RedirectURL: cfg.Callback,
		Scopes:      cfg.Scopes,
	}

	return p, nil
}

func giteaProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "gitea").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	if cfg.Endpoints.Auth == "" {
		cfg.Endpoints.Auth = "https://gitea.com/login/oauth/authorize"
	}

	if cfg.Endpoints.Token == "" {
		cfg.Endpoints.Token = "https://gitea.com/login/oauth/access_token"
	}

	if cfg.Endpoints.Profile == "" {
		cfg.Endpoints.Profile = "https://gitea.com/login/oauth/userinfo"
	}

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"read:user",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.Endpoints.Auth,
			TokenURL: cfg.Endpoints.Token,
		},
		RedirectURL: cfg.Callback,
		Scopes:      cfg.Scopes,
	}

	return p, nil
}

func gitlabProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "gitlab").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	if cfg.Endpoints.Auth == "" {
		cfg.Endpoints.Auth = "https://gitlab.com/oauth/authorize"
	}

	if cfg.Endpoints.Token == "" {
		cfg.Endpoints.Token = "https://gitlab.com/oauth/token"
	}

	if cfg.Endpoints.Profile == "" {
		cfg.Endpoints.Profile = "https://gitlab.com/api/v3/user"
	}

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"openid",
			"profile",
			"email",
			"read_user",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.Endpoints.Auth,
			TokenURL: cfg.Endpoints.Token,
		},
		RedirectURL: cfg.Callback,
		Scopes:      cfg.Scopes,
	}

	return p, nil
}

func oidcProvider(cfg config.AuthProvider) (*Provider, error) {
	logger := log.With().
		Str("service", "provider").
		Str("provider", "oidc").
		Str("name", cfg.Name).
		Logger()

	logger.Info().
		Msg("Registering auth provider")

	p := &Provider{
		Config: &cfg,
		Logger: logger,
	}

	if cfg.Endpoints.Issuer == "" {
		return nil, ErrMissingIssuerEndpoint
	}

	if len(cfg.Scopes) == 0 {
		cfg.Scopes = []string{
			"openid",
			"profile",
			"email",
		}
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}
	p.Config.ClientID = clientID

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}
	p.Config.ClientSecret = clientSecret

	verifier, err := config.Value(cfg.Verifier)
	if err != nil {
		return nil, err
	}
	p.Config.Verifier = verifier

	if p.Config.Verifier == "" {
		p.Config.Verifier = secret.Generate(32)
	}

	oidcProvider, err := oidc.NewProvider(
		context.Background(),
		cfg.Endpoints.Issuer,
	)

	if err != nil {
		return nil, err
	}

	p.OpenID = oidcProvider
	p.Config.Endpoints.Profile = oidcProvider.UserInfoEndpoint()

	p.Verifier = oidcProvider.Verifier(&oidc.Config{
		ClientID: clientID,
	})

	p.OAuth2 = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oidcProvider.Endpoint(),
		RedirectURL:  cfg.Callback,
		Scopes:       cfg.Scopes,
	}

	return p, nil
}
