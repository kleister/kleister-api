package v1

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/authn"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/metrics"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/token"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/rs/zerolog/log"
)

//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../../openapi/v1.yaml

var (
	_ ServerInterface = (*API)(nil)

	// ErrUnsupportedImageFormat defines the error for unsupported image formats.
	ErrUnsupportedImageFormat = fmt.Errorf("unsupported avatar file format")
)

func init() {
	openapi3filter.RegisterBodyDecoder("image/jpeg", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/png", openapi3filter.FileBodyDecoder)
}

// New creates a new API that adds the handler implementations.
func New(
	cfg *config.Config,
	registry *metrics.Metrics,
	identity *authn.Authn,
	uploads upload.Upload,
	storage *store.Store,
) *API {
	return &API{
		config:   cfg,
		registry: registry,
		identity: identity,
		uploads:  uploads,
		storage:  storage,
	}
}

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	config   *config.Config
	registry *metrics.Metrics
	identity *authn.Authn
	uploads  upload.Upload
	storage  *store.Store
}

// RenderNotify is a helper to set a correct status for notifications.
func (a *API) RenderNotify(w http.ResponseWriter, r *http.Request, notify Notification) {
	render.Status(
		r,
		FromPtr(notify.Status),
	)

	render.JSON(
		w,
		r,
		notify,
	)
}

// AllowAdminAccessOnly defines a middleware to check permissions.
func (a *API) AllowAdminAccessOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		principal := current.GetUser(ctx)

		if principal == nil {
			render.JSON(w, r, Notification{
				Message: ToPtr("Only admins can access this resource"),
				Status:  ToPtr(http.StatusForbidden),
			})

			return
		}

		if principal.Admin {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		render.JSON(w, r, Notification{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		})
	})
}

// Authentication provides the authentication for the OpenAPI filter.
func (a *API) Authentication(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	authenticating := &model.User{}
	scheme := input.SecuritySchemeName
	operation := input.RequestValidationInput.Route.Operation.OperationID

	logger := log.With().
		Str("scheme", scheme).
		Str("operation", operation).
		Logger()

	switch scheme {
	case "Header":
		header := input.RequestValidationInput.Request.Header.Get(
			input.SecurityScheme.Name,
		)

		if header == "" {
			return fmt.Errorf("missing authorization header")
		}

		t, err := token.Verify(
			a.config.Token.Secret,
			strings.TrimSpace(
				header,
			),
		)

		if err != nil {
			return fmt.Errorf("failed to parse auth token")
		}

		user, err := a.storage.Auth.ByID(
			ctx,
			t.Ident,
		)

		if err != nil {
			logger.Error().
				Err(err).
				Str("user", t.Ident).
				Msg("Failed to find user")

			return fmt.Errorf("failed to find user")
		}

		logger.Trace().
			Str("user", t.Login).
			Msg("Authentication")

		authenticating = user

	case "Bearer":
		header := input.RequestValidationInput.Request.Header.Get(
			"Authorization",
		)

		if header == "" {
			return fmt.Errorf("missing authorization bearer")
		}

		t, err := token.Verify(
			a.config.Token.Secret,
			strings.TrimSpace(
				strings.Replace(
					header,
					"Bearer",
					"",
					1,
				),
			),
		)

		if err != nil {
			return fmt.Errorf("failed to parse auth token")
		}

		user, err := a.storage.Auth.ByID(
			ctx,
			t.Ident,
		)

		if err != nil {
			logger.Error().
				Err(err).
				Str("user", t.Ident).
				Msg("Failed to find user")

			return fmt.Errorf("failed to find user")
		}

		logger.Trace().
			Str("user", t.Login).
			Msg("Authentication")

		authenticating = user

	case "Basic":
		username, password, ok := input.RequestValidationInput.Request.BasicAuth()

		if !ok {
			return fmt.Errorf("missing credentials")
		}

		user, err := a.storage.Auth.ByCreds(
			ctx,
			username,
			password,
		)

		if err != nil {
			logger.Error().
				Err(err).
				Str("user", username).
				Msg("Wrong credentials")

			return fmt.Errorf("wrong credentials")
		}

		logger.Trace().
			Str("user", username).
			Msg("Authentication")

		authenticating = user

	default:
		return fmt.Errorf("unknown security scheme: %s", scheme)
	}

	log.Trace().
		Str("username", authenticating.Username).
		Str("operation", operation).
		Msg("Authenticated")

	current.SetUser(
		input.RequestValidationInput.Request.Context(),
		authenticating,
	)

	return nil
}

func (a *API) transformAvatar(file multipart.File) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})
	src, _, err := image.Decode(file)

	if err != nil {
		return buffer, fmt.Errorf("invalid file format: %w", err)
	}

	resized := imaging.Resize(
		src,
		64,
		64,
		imaging.Lanczos,
	)

	if err := png.Encode(buffer, resized); err != nil {
		return buffer, err
	}

	return buffer, nil
}

func (a *API) resizeAvatar(file multipart.File, meta *multipart.FileHeader) (*bytes.Buffer, error) {
	switch meta.Header.Get("Content-Type") {
	case "image/jpeg":
		return a.transformAvatar(file)
	case "image/png":
		return a.transformAvatar(file)
	}

	return nil, ErrUnsupportedImageFormat
}
