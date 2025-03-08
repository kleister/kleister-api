package authn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Machiel/slugify"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

// Provider defines a single authentication provider.
type Provider struct {
	Config   *config.AuthProvider
	OAuth2   *oauth2.Config
	OpenID   *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Logger   zerolog.Logger
}

// Claims tries to map all provider claims properly to the wrapper model.
func (p *Provider) Claims(ctx context.Context, token *oauth2.Token) (user *User, err error) {
	attrs := make(map[string]interface{})

	if p.Config.Driver == "oidc" {
		rawToken, ok := token.Extra("id_token").(string)

		if !ok {
			return nil, ErrMissingToken
		}

		idToken, err := p.Verifier.Verify(
			ctx,
			rawToken,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to verify token: %w", err)
		}

		if err := idToken.Claims(
			&attrs,
		); err != nil {
			return nil, fmt.Errorf("failed to parse claims: %w", err)
		}

		if user, err = p.extractOidcUser(attrs); err != nil {
			return nil, err
		}
	} else {
		resp, err := p.OAuth2.Client(
			ctx,
			token,
		).Get(
			p.Config.Endpoints.Profile,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch userinfo: %w", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("bad status code returned: %d", resp.StatusCode)
		}

		if err := json.NewDecoder(
			resp.Body,
		).Decode(
			&attrs,
		); err != nil {
			return nil, fmt.Errorf("failed to decode userinfo: %w", err)
		}

		switch p.Config.Driver {
		case "entraid":
			if user, err = p.extractEntraidUser(attrs); err != nil {
				return nil, err
			}
		case "google":
			if user, err = p.extractGoogleUser(attrs); err != nil {
				return nil, err
			}
		case "gitea":
			if user, err = p.extractGiteaUser(attrs); err != nil {
				return nil, err
			}
		case "gitlab":
			if user, err = p.extractGitlabUser(attrs); err != nil {
				return nil, err
			}
		case "github":
			if user, err = p.extractGithubUser(attrs); err != nil {
				return nil, err
			}

			if user.Email == "" {
				resp, err := p.OAuth2.Client(
					ctx,
					token,
				).Get(
					p.Config.Endpoints.Email,
				)

				if err != nil {
					return nil, fmt.Errorf("failed to fetch emails: %w", err)
				}

				defer resp.Body.Close()

				var (
					mails []struct {
						Email    string `json:"email"`
						Primary  bool   `json:"primary"`
						Verified bool   `json:"verified"`
					}
				)

				if err := json.NewDecoder(
					resp.Body,
				).Decode(
					&mails,
				); err != nil {
					return nil, fmt.Errorf("failed to decode emails: %w", err)
				}

				for _, mail := range mails {
					if mail.Primary && mail.Verified {
						user.Email = mail.Email
					}
				}
			}
		}
	}

	return user, nil
}

func (p *Provider) extractOidcUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["sub"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Ident = typed
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("mapping", "sub").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Str("mapping", "sub").
			Msg("Failed to fetch attr")
	}

	if p.Config.Mappings.Login != "" {
		if val, ok := user.Raw[p.Config.Mappings.Login]; ok {
			if typed, typeOk := val.(string); typeOk {
				user.Login = typed
			} else {
				p.Logger.Warn().
					Str("attr", "login").
					Str("mapping", p.Config.Mappings.Login).
					Str("type", fmt.Sprintf("%T", val)).
					Msg("Failed to convert attr")
			}
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("mapping", p.Config.Mappings.Login).
				Msg("Failed to fetch attr")
		}
	}

	if p.Config.Mappings.Name != "" {
		if val, ok := user.Raw[p.Config.Mappings.Name]; ok {
			if typed, typeOk := val.(string); typeOk {
				user.Name = typed
			} else {
				p.Logger.Warn().
					Str("attr", "name").
					Str("mapping", p.Config.Mappings.Name).
					Str("type", fmt.Sprintf("%T", val)).
					Msg("Failed to convert attr")
			}
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("mapping", p.Config.Mappings.Name).
				Msg("Failed to fetch attr")
		}
	}

	if p.Config.Mappings.Email != "" {
		if val, ok := user.Raw[p.Config.Mappings.Email]; ok {
			if typed, typeOk := val.(string); typeOk {
				user.Email = typed
			} else {
				p.Logger.Warn().
					Str("attr", "email").
					Str("mapping", p.Config.Mappings.Email).
					Str("type", fmt.Sprintf("%T", val)).
					Msg("Failed to convert attr")
			}
		} else {
			p.Logger.Warn().
				Str("attr", "email").
				Str("mapping", p.Config.Mappings.Email).
				Msg("Failed to fetch attr")
		}
	}

	if p.Config.Mappings.Role != "" {
		if val, ok := user.Raw[p.Config.Mappings.Role]; ok {
			if typed, typeOk := val.([]interface{}); typeOk {
				result := []string{}
				for _, row := range typed {
					result = append(result, row.(string))
				}
				user.Roles = result
			} else {
				p.Logger.Warn().
					Str("attr", "roles").
					Str("mapping", p.Config.Mappings.Role).
					Str("type", fmt.Sprintf("%T", val)).
					Msg("Failed to convert attr")
			}
		} else {
			p.Logger.Warn().
				Str("attr", "roles").
				Str("mapping", p.Config.Mappings.Role).
				Msg("Failed to fetch attr")
		}
	}

	return user, nil
}

func (p *Provider) extractEntraidUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["id"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Ident = typed
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["displayName"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Login = slugify.Slugify(typed)
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "login").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["displayName"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Name = typed
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "name").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["mail"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Email = typed
		} else {
			p.Logger.Warn().
				Str("attr", "email").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "email").
			Msg("Failed to fetch attr")
	}

	return user, nil
}

func (p *Provider) extractGoogleUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["id"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Ident = typed
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["name"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Login = slugify.Slugify(typed)
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "login").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["name"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Name = typed
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "name").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["email"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Email = typed
		} else {
			p.Logger.Warn().
				Str("attr", "email").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "email").
			Msg("Failed to fetch attr")
	}

	return user, nil
}

func (p *Provider) extractGiteaUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["sub"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Ident = typed
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["preferred_username"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Login = typed
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "login").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["name"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Name = typed
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "name").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["email"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Email = typed
		} else {
			p.Logger.Warn().
				Str("attr", "email").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "email").
			Msg("Failed to fetch attr")
	}

	return user, nil
}

func (p *Provider) extractGitlabUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["id"]; ok {
		if typed, typeOk := val.(float64); typeOk {
			user.Ident = fmt.Sprintf("%d", int64(typed))
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["username"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Login = typed
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "login").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["name"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Name = typed
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "name").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["email"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Email = typed
		} else {
			p.Logger.Warn().
				Str("attr", "email").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "email").
			Msg("Failed to fetch attr")
	}

	return user, nil
}

func (p *Provider) extractGithubUser(attrs map[string]interface{}) (*User, error) {
	user := &User{
		Raw: attrs,
	}

	if val, ok := user.Raw["id"]; ok {
		if typed, typeOk := val.(float64); typeOk {
			user.Ident = fmt.Sprintf("%d", int64(typed))
		} else {
			p.Logger.Warn().
				Str("attr", "ident").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "ident").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["login"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Login = typed
		} else {
			p.Logger.Warn().
				Str("attr", "login").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "login").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["name"]; ok {
		if typed, typeOk := val.(string); typeOk {
			user.Name = typed
		} else {
			p.Logger.Warn().
				Str("attr", "name").
				Str("type", fmt.Sprintf("%T", val)).
				Msg("Failed to convert attr")
		}
	} else {
		p.Logger.Warn().
			Str("attr", "name").
			Msg("Failed to fetch attr")
	}

	if val, ok := user.Raw["email"]; ok {
		if val != nil {
			if typed, typeOk := val.(string); typeOk {
				user.Email = typed
			} else {
				p.Logger.Warn().
					Str("attr", "email").
					Str("type", fmt.Sprintf("%T", val)).
					Msg("Failed to convert attr")
			}
		}
	} else {
		if user.Raw["email"] != nil {
			p.Logger.Warn().
				Str("attr", "email").
				Msg("Failed to fetch attr")
		}
	}

	return user, nil
}
