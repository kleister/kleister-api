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

//go:generate oapi-codegen --config=config.yaml ../../../openapi/v1.yaml

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

// AllowAuthenticatedOnly defines a middleware that only requires a
// principal to be set, used for create endpoints where the target object
// doesn't exist yet to check a tier against.
func (a *API) AllowAuthenticatedOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if current.GetUser(ctx) == nil {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Authentication required"),
				Status:  ToPtr(http.StatusForbidden),
			})

			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// allowObjectAccess builds a middleware that checks whether the current
// principal holds the given relation on the object identified by kind and
// id(ctx), bypassing the check entirely for global admins.
func (a *API) allowObjectAccess(kind, relation string, id func(ctx context.Context) string) func(http.Handler) http.Handler {
	return a.allowObject(relation, func(ctx context.Context) string {
		return kind + ":" + id(ctx)
	}, kind)
}

// allowObject builds a middleware that checks whether the current principal
// holds the given relation on the full object reference returned by
// object(ctx) (e.g. "pack:01ABC"), bypassing the check entirely for global
// admins. subject names the kind of resource for error messages/logging.
func (a *API) allowObject(relation string, object func(ctx context.Context) string, subject string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			principal := current.GetUser(ctx)

			if principal == nil {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Authentication required"),
					Status:  ToPtr(http.StatusForbidden),
				})

				return
			}

			if principal.Admin {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			allowed, err := a.storage.Authz().Check(
				ctx,
				"user:"+principal.ID,
				relation,
				object(ctx),
			)

			if err != nil {
				log.Error().
					Err(err).
					Str("subject", subject).
					Str("relation", relation).
					Msg("Failed to check authorization")

				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to check authorization"),
					Status:  ToPtr(http.StatusInternalServerError),
				})

				return
			}

			if !allowed {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr(fmt.Sprintf("Not allowed to access this %s", subject)),
					Status:  ToPtr(http.StatusForbidden),
				})

				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
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
