package requestid

import (
	"context"

	"github.com/rs/zerolog/hlog"
)

// Get returns the request ID from context.Context.
func Get(ctx context.Context) string {
	id, ok := hlog.IDFromCtx(ctx)

	if ok {
		return id.String()
	}

	return ""
}
