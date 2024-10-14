package v1

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/gobwas/glob"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/providers"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/token"
	"github.com/markbates/goth"
	"github.com/rs/zerolog/log"
)

// ExternalInitialize implements the v1.ServerInterface.
func (a *API) ExternalInitialize(ctx context.Context, request ExternalInitializeRequestObject) (ExternalInitializeResponseObject, error) {
	provider, err := goth.GetProvider(
		request.Provider,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect provider")

		return ExternalInitialize404JSONResponse{
			Message: ToPtr("Failed to detect provider"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	session, err := provider.BeginAuth(
		setAuthState(request.Params.State),
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to init provider")

		return ExternalInitialize412JSONResponse{
			Message: ToPtr("Failed to init provider"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	url, err := session.GetAuthURL()

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect auth url")

		return ExternalInitialize404JSONResponse{
			Message: ToPtr("Failed to detect auth url"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	a.session.Put(
		ctx,
		request.Provider,
		session.Marshal(),
	)

	return ExternalInitializeRedirectResponse{
		url: url,
	}, nil
}

// ExternalInitializeRedirectResponse defines the response to redirect to a defined URL.
type ExternalInitializeRedirectResponse struct {
	url string
}

// VisitExternalInitializeResponse implements the middleware.Responder interface for redirects.
func (r ExternalInitializeRedirectResponse) VisitExternalInitializeResponse(w http.ResponseWriter) error {
	w.Header().Set(
		"Location",
		r.url,
	)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.WriteHeader(
		http.StatusTemporaryRedirect,
	)

	return nil
}

// ExternalCallback implements the v1.ServerInterface.
func (a *API) ExternalCallback(ctx context.Context, request ExternalCallbackRequestObject) (ExternalCallbackResponseObject, error) {
	provider, err := goth.GetProvider(
		request.Provider,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect provider")

		return ExternalCallback404JSONResponse{
			Message: ToPtr("Failed to detect provider"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	session, err := provider.UnmarshalSession(
		a.session.Get(
			ctx,
			request.Provider,
		),
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to parse session")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to parse session"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	a.session.Pop(
		ctx,
		request.Provider,
	)

	if err := verifyAuthState(
		request.Params.State,
		session,
	); err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to verify state")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to verify state"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	external, err := provider.FetchUser(
		session,
	)

	log.Debug().
		Str("provider", external.Provider).
		Str("email", external.Email).
		Str("name", external.Name).
		Str("nickname", external.NickName).
		Str("user_id", external.UserID).
		Msg("Requested auth")

	if err == nil {
		nickname := slugify.Slugify(external.NickName)

		user, err := a.users.External(
			ctx,
			external.Provider,
			external.UserID,
			nickname,
			external.Email,
			external.Name,
			detectAdminFor(request.Provider, external),
		)

		if err != nil {
			log.Error().
				Err(err).
				Str("provider", request.Provider).
				Str("username", nickname).
				Msg("Failed to create user")

			return ExternalCallback412JSONResponse{
				Message: ToPtr("Failed to create user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		a.session.Put(
			ctx,
			"user",
			user.ID,
		)

		log.Debug().
			Str("provider", request.Provider).
			Str("username", user.Username).
			Str("email", user.Email).
			Str("external", external.UserID).
			Msg("authenticated")

		return ExternalCallbackRedirectResponse{
			url: strings.Join(
				[]string{
					a.config.Server.Host,
					a.config.Server.Root,
				},
				"",
			),
		}, nil
	}

	authValues := url.Values{}

	if request.Params.Code != nil {
		authValues.Set("code", FromPtr(request.Params.Code))
	}

	if request.Params.State != nil {
		authValues.Set("state", FromPtr(request.Params.State))
	}

	if _, err = session.Authorize(
		provider,
		authValues,
	); err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to authorize session")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to authorize session"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	a.session.Put(
		ctx,
		request.Provider,
		session.Marshal(),
	)

	external, err = provider.FetchUser(session)

	log.Debug().
		Str("provider", external.Provider).
		Str("email", external.Email).
		Str("name", external.Name).
		Str("nickname", external.NickName).
		Str("user_id", external.UserID).
		Msg("Requested auth")

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to fetch user")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to fetch user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	nickname := slugify.Slugify(external.NickName)

	user, err := a.users.External(
		ctx,
		external.Provider,
		external.UserID,
		nickname,
		external.Email,
		external.Name,
		detectAdminFor(request.Provider, external),
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Str("username", nickname).
			Msg("Failed to create user")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to create user"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	a.session.Put(
		ctx,
		"user",
		user.ID,
	)

	log.Debug().
		Str("provider", request.Provider).
		Str("username", user.Username).
		Str("email", user.Email).
		Str("external", external.UserID).
		Msg("authenticated")

	return ExternalCallbackRedirectResponse{
		url: strings.Join(
			[]string{
				a.config.Server.Host,
				a.config.Server.Root,
			},
			"",
		),
	}, nil
}

// ExternalCallbackRedirectResponse defines the response to redirect to a defined URL.
type ExternalCallbackRedirectResponse struct {
	url string
}

// VisitExternalCallbackResponse implements the middleware.Responder interface for redirects.
func (r ExternalCallbackRedirectResponse) VisitExternalCallbackResponse(w http.ResponseWriter) error {
	w.Header().Set(
		"Location",
		r.url,
	)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.WriteHeader(
		http.StatusTemporaryRedirect,
	)

	return nil
}

// LoginAuth implements the v1.ServerInterface.
func (a *API) LoginAuth(ctx context.Context, request LoginAuthRequestObject) (LoginAuthResponseObject, error) {
	user, err := a.users.AuthByCreds(
		ctx,
		request.Body.Username,
		request.Body.Password,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return LoginAuth401JSONResponse{
				Message: ToPtr("Wrong username or password"),
				Status:  ToPtr(http.StatusUnauthorized),
			}, nil
		}

		if errors.Is(err, users.ErrBadCredentials) {
			return LoginAuth401JSONResponse{
				Message: ToPtr("Wrong username or password"),
				Status:  ToPtr(http.StatusUnauthorized),
			}, nil
		}

		log.Error().
			Err(err).
			Str("username", request.Body.Username).
			Msg("Failed to authenticate")

		return LoginAuth500JSONResponse{
			Message: ToPtr("Failed to authenticate user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	result, err := token.New(
		user.Username,
	).Expiring(
		a.config.Session.Secret,
		a.config.Session.Expire,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("username", request.Body.Username).
			Msg("Failed to generate a token")

		return LoginAuth500JSONResponse{
			Message: ToPtr("Failed to generate a token"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return LoginAuth200JSONResponse(
		a.convertAuthToken(result),
	), nil
}

// RefreshAuth implements the v1.ServerInterface.
func (a *API) RefreshAuth(ctx context.Context, _ RefreshAuthRequestObject) (RefreshAuthResponseObject, error) {
	principal := current.GetUser(
		ctx,
	)

	result, err := token.New(
		principal.Username,
	).Expiring(
		a.config.Session.Secret,
		a.config.Session.Expire,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("username", principal.Username).
			Msg("Failed to generate a token")

		return RefreshAuth401JSONResponse{
			Message: ToPtr("Failed to generate a token"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return RefreshAuth200JSONResponse(
		a.convertAuthToken(result),
	), nil
}

// VerifyAuth implements the v1.ServerInterface.
func (a *API) VerifyAuth(ctx context.Context, _ VerifyAuthRequestObject) (VerifyAuthResponseObject, error) {
	principal := current.GetUser(
		ctx,
	)

	return VerifyAuth200JSONResponse(
		a.convertAuthVerify(principal),
	), nil
}

func (a *API) convertAuthToken(record *token.Result) AuthToken {
	if record.ExpiresAt.IsZero() {
		return AuthToken{
			Token: ToPtr(record.Token),
		}
	}

	return AuthToken{
		Token:     ToPtr(record.Token),
		ExpiresAt: ToPtr(record.ExpiresAt),
	}
}

func (a *API) convertAuthVerify(record *model.User) AuthVerify {
	return AuthVerify{
		Username:  ToPtr(record.Username),
		CreatedAt: ToPtr(record.CreatedAt),
	}
}

func setAuthState(state *string) string {
	if state != nil && len(FromPtr(state)) > 0 {
		return FromPtr(state)
	}

	nonceBytes := make([]byte, 64)

	if _, err := io.ReadFull(
		rand.Reader,
		nonceBytes,
	); err != nil {
		log.Error().
			Err(err).
			Msg("Source of randomness unavailable")

		panic("source of randomness unavailable")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func verifyAuthState(state *string, sess goth.Session) error {
	rawAuth, err := sess.GetAuthURL()

	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuth)

	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")

	if originalState != "" && (state == nil || originalState != FromPtr(state)) {
		return fmt.Errorf("state token mismatch")
	}

	return nil
}

func detectAdminFor(providerName string, external goth.User) bool {
	if cfg, ok := providers.Config[providerName]; ok {
		for _, user := range cfg.Admins.Users {
			if user == external.NickName {
				return true
			}
		}

		for _, email := range cfg.Admins.Emails {
			g, err := glob.Compile(email)

			if err != nil {
				log.Error().
					Str("provider", providerName).
					Str("glob", email).
					Msg("Failed to compile email globbing")

				continue
			}

			if g.Match(external.Email) {
				return true
			}
		}

		if cfg.Mappings.Role != "" {
			if mappedRoles, ok := external.RawData[cfg.Mappings.Role]; ok {
				for _, role := range cfg.Admins.Roles {
					for _, mappedRole := range mappedRoles.([]interface{}) {
						if role == mappedRole.(string) {
							return true
						}
					}
				}
			}
		}
	} else {
		log.Error().
			Str("provider", providerName).
			Msg("Failed to detect a provider config")
	}

	return false
}
