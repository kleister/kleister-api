package upload

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//go:generate go tool github.com/golang/mock/mockgen -source upload.go -destination mock.go -package upload

var (
	// ErrUnknownDriver defines a named error for unknown upload drivers.
	ErrUnknownDriver = fmt.Errorf("unknown upload driver")
)

// Upload provides the interface for the upload implementations.
type Upload interface {
	Info() map[string]interface{}
	Prepare() (Upload, error)
	Close() error
	Upload(context.Context, string, *bytes.Buffer) error
	Delete(context.Context, string, bool) error
	Handler(string) http.Handler
}
