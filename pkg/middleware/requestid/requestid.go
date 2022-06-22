package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	key contextKey = "request_id"
)

type contextKey string

// Get returns the request ID from context.Context.
func Get(ctx context.Context) string {
	value := ctx.Value(key)

	if id, ok := value.(string); ok {
		return id
	}

	return ""
}

// Handler writes the request ID to context if not already exists.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")

		if id == "" {
			id = uuid.New().String()
		}

		r = r.WithContext(
			context.WithValue(
				r.Context(),
				key,
				id,
			),
		)

		next.ServeHTTP(w, r)
	})
}
