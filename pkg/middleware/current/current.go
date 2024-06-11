package current

import (
	"context"
	"net/http"
	"sync"

	"github.com/kleister/kleister-api/pkg/model"
)

type contextKey string

const (
	generalKey contextKey = "current_context"
)

// Context defines a custom writeable context.
type Context struct {
	mu   sync.RWMutex
	Keys map[string]any
}

// Set is used to write to the context.
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Get is used to read from the contenxt.
func (c *Context) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.Keys[key]
	return value, exists
}

// Middleware initializes the custom context.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), generalKey, &Context{}))
		next.ServeHTTP(w, r)
	})
}

// GetUser returns the current user from context.
func GetUser(ctx context.Context) *model.User {
	general := ctx.Value(generalKey).(*Context)

	value, ok := general.Get("current_user")

	if !ok {
		return nil
	}

	if res, ok := value.(*model.User); ok {
		return res
	}

	return nil
}

// SetUser stores the current user within context.
func SetUser(ctx context.Context, user *model.User) {
	general := ctx.Value(generalKey).(*Context)

	general.Set(
		"current_user",
		user,
	)
}
