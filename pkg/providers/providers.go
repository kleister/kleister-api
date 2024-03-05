package providers

import (
	"fmt"

	"github.com/gopad/gopad-api/pkg/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/azureadv2"
	"github.com/markbates/goth/providers/gitea"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/spf13/viper"
)

var (
	// ErrMissingDiscoveryEndpoint defines the error of an discovery endpoint is missing.
	ErrMissingDiscoveryEndpoint = fmt.Errorf("missing discovery endpoint")
)

// Providers simple stores the registered providers for goth.
type Providers struct {
	config    *config.AuthConfig
	providers []goth.Provider
}

// Register does the registering of providers for goth.
func Register(opts ...Option) error {
	options := newOptions(opts...)

	client := &Providers{
		config:    &config.AuthConfig{},
		providers: make([]goth.Provider, 0),
	}

	if options.Config != "" {
		cfg := viper.New()
		cfg.SetConfigFile(options.Config)

		if err := cfg.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				return fmt.Errorf("failed to find auth config: %w", err)
			case viper.UnsupportedConfigError:
				return fmt.Errorf("unsupport type for auth config: %w", err)
			default:
				return fmt.Errorf("failed to read auth config: %w", err)
			}
		}

		if err := cfg.Unmarshal(client.config); err != nil {
			return fmt.Errorf("failed to parse auth config: %w", err)
		}

		for _, provider := range client.config.Providers {
			switch provider.Driver {
			case "github":
				p, err := githubProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			case "gitea":
				p, err := giteaProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			case "gitlab":
				p, err := gitlabProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			case "google":
				p, err := googleProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			case "azuread":
				p, err := azureadProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			case "oidc":
				p, err := oidcProvider(provider)

				if err != nil {
					return err
				}

				client.providers = append(
					client.providers,
					p,
				)
			default:
				return fmt.Errorf("unsupport auth provider: %s", provider.Driver)
			}
		}

		goth.UseProviders(client.providers...)
	}

	return nil
}

func githubProvider(cfg config.AuthProvider) (*github.Provider, error) {
	authEndpoint := "https://github.com/login/oauth/authorize"
	if cfg.Endpoints.Auth != "" {
		authEndpoint = cfg.Endpoints.Auth
	}

	tokenEndpoint := "https://github.com/login/oauth/access_token"
	if cfg.Endpoints.Auth != "" {
		tokenEndpoint = cfg.Endpoints.Auth
	}

	profileEndpoint := "https://api.github.com/user"
	if cfg.Endpoints.Auth != "" {
		profileEndpoint = cfg.Endpoints.Auth
	}

	emailEndpoint := "https://api.github.com/user/emails"
	if cfg.Endpoints.Auth != "" {
		emailEndpoint = cfg.Endpoints.Auth
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider := github.NewCustomisedURL(
		clientID,
		clientSecret,
		cfg.Callback,
		authEndpoint,
		tokenEndpoint,
		profileEndpoint,
		emailEndpoint,
		cfg.Scopes...,
	)

	provider.SetName(cfg.Name)
	return provider, nil
}

func giteaProvider(cfg config.AuthProvider) (*gitea.Provider, error) {
	authEndpoint := "https://gitea.com/login/oauth/authorize"
	if cfg.Endpoints.Auth != "" {
		authEndpoint = cfg.Endpoints.Auth
	}

	tokenEndpoint := "https://gitea.com/login/oauth/access_token"
	if cfg.Endpoints.Auth != "" {
		tokenEndpoint = cfg.Endpoints.Auth
	}

	profileEndpoint := "https://gitea.com/api/v1/user"
	if cfg.Endpoints.Auth != "" {
		profileEndpoint = cfg.Endpoints.Auth
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider := gitea.NewCustomisedURL(
		clientID,
		clientSecret,
		cfg.Callback,
		authEndpoint,
		tokenEndpoint,
		profileEndpoint,
		cfg.Scopes...,
	)

	provider.SetName(cfg.Name)
	return provider, nil
}

func gitlabProvider(cfg config.AuthProvider) (*gitlab.Provider, error) {
	authEndpoint := "https://gitlab.com/oauth/authorize"
	if cfg.Endpoints.Auth != "" {
		authEndpoint = cfg.Endpoints.Auth
	}

	tokenEndpoint := "https://gitlab.com/oauth/token"
	if cfg.Endpoints.Auth != "" {
		tokenEndpoint = cfg.Endpoints.Auth
	}

	profileEndpoint := "https://gitlab.com/api/v3/user"
	if cfg.Endpoints.Auth != "" {
		profileEndpoint = cfg.Endpoints.Auth
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider := gitlab.NewCustomisedURL(
		clientID,
		clientSecret,
		cfg.Callback,
		authEndpoint,
		tokenEndpoint,
		profileEndpoint,
		cfg.Scopes...,
	)

	provider.SetName(cfg.Name)
	return provider, nil
}

func googleProvider(cfg config.AuthProvider) (*google.Provider, error) {
	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider := google.New(
		clientID,
		clientSecret,
		cfg.Callback,
		cfg.Scopes...,
	)

	provider.SetName(cfg.Name)
	return provider, nil
}

func azureadProvider(cfg config.AuthProvider) (*azureadv2.Provider, error) {
	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	azureScopes := make([]azureadv2.ScopeType, 0)

	for _, scope := range cfg.Scopes {
		azureScopes = append(
			azureScopes,
			azureadv2.ScopeType(scope),
		)
	}

	provider := azureadv2.New(
		clientID,
		clientSecret,
		cfg.Callback,
		azureadv2.ProviderOptions{
			Tenant: azureadv2.CommonTenant,
			Scopes: azureScopes,
		},
	)

	provider.SetName(cfg.Name)
	return provider, nil
}

func oidcProvider(cfg config.AuthProvider) (*openidConnect.Provider, error) {
	if cfg.Endpoints.Discovery == "" {
		return nil, ErrMissingDiscoveryEndpoint
	}

	clientID, err := config.Value(cfg.ClientID)
	if err != nil {
		return nil, err
	}

	clientSecret, err := config.Value(cfg.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider, err := openidConnect.NewNamed(
		cfg.Name,
		clientID,
		clientSecret,
		cfg.Callback,
		cfg.Endpoints.Discovery,
		cfg.Scopes...,
	)

	if err != nil {
		return nil, err
	}

	if cfg.Mappings.Login != "" {
		provider.NickNameClaims = []string{
			cfg.Mappings.Login,
		}
	}

	if cfg.Mappings.Name != "" {
		provider.NameClaims = []string{
			cfg.Mappings.Name,
		}
	}

	if cfg.Mappings.Email != "" {
		provider.EmailClaims = []string{
			cfg.Mappings.Email,
		}
	}

	if cfg.Mappings.Avatar != "" {
		provider.AvatarURLClaims = []string{
			cfg.Mappings.Avatar,
		}
	}

	return provider, nil
}
