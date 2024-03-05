package session

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const (
	// CookieName defines the name of the stored cookie.
	CookieName = "kleister"
)

// Session is a simple wrapper around a session manager.
type Session struct {
	Manager *scs.SessionManager
}

// New simply initializes a new session storage used for OAuth2/OIDC.
func New(opts ...Option) *Session {
	options := newOptions(opts...)

	manager := scs.New()
	manager.Store = options.Store
	manager.Lifetime = options.Lifetime
	manager.Cookie.Name = CookieName
	manager.Cookie.Path = options.Path
	manager.Cookie.Secure = options.Secure

	return &Session{
		Manager: manager,
	}
}

// Middleware defines the middleware to store sessions.
func (s *Session) Middleware(next http.Handler) http.Handler {
	return s.Manager.LoadAndSave(next)
}

// Put simply puts a value into the session cookie.
func (s *Session) Put(ctx context.Context, key, val string) {
	s.Manager.Put(ctx, key, val)
}

// Get simply gets a value from the session cookie.
func (s *Session) Get(ctx context.Context, key string) string {
	return s.Manager.GetString(ctx, key)
}

// Destroy simply wipes the whole session, used for logout.
func (s *Session) Destroy(ctx context.Context) error {
	return s.Manager.Destroy(ctx)
}

// User simply gets the current user ID from the session.
func (s *Session) User(key string) (string, error) {
	encoded, found, err := s.Manager.Store.Find(
		key,
	)

	if err != nil {
		return "", fmt.Errorf("failed to find session: %w", err)
	}

	if !found {
		return "", fmt.Errorf("failed to find session data")
	}

	deadline, values, err := s.Manager.Codec.Decode(
		encoded,
	)

	if err != nil {
		return "", fmt.Errorf("failed to decode session: %w", err)
	}

	user, ok := values["user"]

	if !ok {
		return "", fmt.Errorf("failed to extract user from session")
	}

	if time.Now().UTC().After(deadline) {
		return "", fmt.Errorf("deadline have already been reached")
	}

	return user.(string), nil
}
